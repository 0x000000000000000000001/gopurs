import fs from 'fs';
import path from 'path';

let cachedScanDirs = null;

function getScanDirs(mbFfiDir) {
    if (cachedScanDirs !== null) return cachedScanDirs;
    
    const rootDir = process.cwd();
    const scanDirs = [];
    
    const spagoDirs = [
        path.join(rootDir, '.spago'),
        path.join(rootDir, 'spago.d')
    ];
    
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

let goFileIndex = null;

function buildGoFileIndex(scanDirs) {
    if (goFileIndex !== null) return;
    goFileIndex = new Set();
    
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
            } else if (entry.name.endsWith('.go')) {
                goFileIndex.add(res);
            }
        }
    }
    
    for (const d of scanDirs) {
        walk(d);
    }
}

export const findFfiFileImpl = function(mbFfiDir) {
    return function(modNameStr) {
        return function(mbModulePath) {
            return function() {
                if (mbModulePath) {
                    const goPath = mbModulePath.replace(/\.purs$/, '.go');
                    if (fs.existsSync(goPath)) {
                        return goPath;
                    }
                }
                
                const scanDirs = getScanDirs(mbFfiDir);
                buildGoFileIndex(scanDirs);
                
                for (const dir of scanDirs) {
                    // Search in dir/src/Data/Show.go, dir/src/Show.go, etc.
                    const searchPaths = [
                        path.join(dir, 'src', ...modNameStr.split('.')) + '.go',
                        path.join(dir, 'src', modNameStr + '.go'),
                        path.join(dir, modNameStr + '.go')
                    ];
                    for (const p of searchPaths) {
                        if (goFileIndex.has(p)) {
                            return p;
                        }
                    }
                }
                return null;
            };
        };
    };
};
