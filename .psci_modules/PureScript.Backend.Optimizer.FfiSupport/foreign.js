import fs from 'fs';
import path from 'path';

export const hashString = (str) => {
    let hash = 5381;
    let i = str.length;
    while(i) {
        hash = (hash * 33) ^ str.charCodeAt(--i);
    }
    return (hash >>> 0).toString();
};

let cachedScanDirs = null;

function getScanDirs(mbFfiDir, extraSpagoDirs) {
    if (cachedScanDirs !== null) return cachedScanDirs;
    
    const rootDir = process.cwd();
    const scanDirs = [];
    
    const spagoDirs = [
        path.join(rootDir, '.spago'),
        path.join(rootDir, 'spago.d')
    ];
    for (const d of extraSpagoDirs) {
        spagoDirs.push(path.join(rootDir, d));
    }
    
    for (const spagoDir of spagoDirs) {
        if (fs.existsSync(spagoDir) && fs.statSync(spagoDir).isDirectory()) {
            const packages = fs.readdirSync(spagoDir);
            for (const pkg of packages) {
                const pkgDir = path.join(spagoDir, pkg);
                if (fs.statSync(pkgDir).isDirectory()) {
                    let hasVersion = false;
                    const subdirs = fs.readdirSync(pkgDir);
                    for (const subdir of subdirs) {
                        const versionDir = path.join(pkgDir, subdir);
                        if (subdir.startsWith('v') && fs.statSync(versionDir).isDirectory()) {
                            scanDirs.push(versionDir);
                            hasVersion = true;
                        }
                    }
                    if (!hasVersion) {
                        scanDirs.push(pkgDir);
                    }
                }
            }
        }
    }
    
    if (mbFfiDir) {
        scanDirs.push(path.join(rootDir, mbFfiDir));
    } else {
        // Fallback to searching the local src/ dir if mbFfiDir is not provided
        scanDirs.push(rootDir);
    }
    
    cachedScanDirs = scanDirs;
    return scanDirs;
}

const ffiFileIndexes = {};

function buildFfiFileIndex(scanDirs, extension) {
    if (ffiFileIndexes[extension]) return;
    const index = new Set();
    
    function walk(dir) {
        let entries;
        try {
            entries = fs.readdirSync(dir, { withFileTypes: true });
        } catch (e) {
            return;
        }
        for (const entry of entries) {
            const res = path.join(dir, entry.name);
            if (entry.isDirectory()) {
                walk(res);
            } else if (entry.name.endsWith(extension)) {
                index.add(res);
            }
        }
    }
    
    for (const d of scanDirs) {
        walk(d);
    }
    ffiFileIndexes[extension] = index;
}

export const findFfiFileImpl = function(extension) {
    return function(extraSpagoDirs) {
        return function(mbFfiDir) {
            return function(modNameStr) {
                return function(mbModulePath) {
                    return function() {
                        if (mbModulePath) {
                            const ffiPath = mbModulePath.replace(/\.purs$/, extension);
                            if (fs.existsSync(ffiPath)) {
                                return ffiPath;
                            }
                        }
                        
                        const scanDirs = getScanDirs(mbFfiDir, extraSpagoDirs);
                        buildFfiFileIndex(scanDirs, extension);
                        
                        const index = ffiFileIndexes[extension];
                        
                        for (const dir of scanDirs) {
                            const searchPaths = [
                                path.join(dir, 'src', ...modNameStr.split('.')) + extension,
                                path.join(dir, 'src', modNameStr + extension),
                                path.join(dir, modNameStr + extension)
                            ];
                            for (const p of searchPaths) {
                                if (index.has(p)) {
                                    return p;
                                }
                            }
                        }
                        return null;
                    };
                };
            };
        };
    };
};
