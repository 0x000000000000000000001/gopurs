<?php
$exports['runtimeImportImpl'] = function($nothing) {
    return function($just) use ($nothing) {
        return function($moduleName) use ($nothing, $just) {
            return function($body) use ($nothing, $just, $moduleName) {
                return function() use ($nothing, $just, $moduleName, $body) {
                    try {
                        if (!file_exists(__DIR__ . '/../' . $moduleName . '/index.php')) {
                            throw new \Exception("File not found");
                        }
                        require_once __DIR__ . '/../' . $moduleName . '/index.php';
                        return $body($nothing)();
                    } catch (\Throwable $err) {
                        return $body($just($err->getMessage()))();
                    }
                };
            };
        };
    };
};
