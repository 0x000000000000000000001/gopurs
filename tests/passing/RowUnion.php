<?php
$exports['mergeImpl'] = function($l) {
    return function($r) use ($l) {
        return (object) \array_merge((array) $r, (array) $l);
    };
};
