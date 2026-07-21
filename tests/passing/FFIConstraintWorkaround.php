<?php
$exports['showImpl'] = function($showFn) {
    return function($val) use ($showFn) {
        return $showFn($val);
    };
};
