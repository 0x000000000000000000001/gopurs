<?php

$assertThrowsImpl = function($arg) {
    return function($f) use ($arg) {
        return function() use ($f, $arg) {
            set_error_handler(function($errno, $errstr, $errfile, $errline) {
                throw new \ErrorException($errstr, 0, $errno, $errfile, $errline);
            });
            try {
                $f($arg);
                restore_error_handler();
            } catch (\Throwable $e) {
                restore_error_handler();
                return $e->getMessage();
            }
            restore_error_handler();
            throw new \Exception("Assertion failed: An error should have been thrown");
        };
    };
};

$exports['assertThrowsImpl'] = $assertThrowsImpl;

