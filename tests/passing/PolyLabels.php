<?php
$exports['unsafeGet'] = function($s) {
    return function($o) use ($s) {
        if (is_array($o)) return $o[$s];
        if (is_object($o)) return $o->{$s};
        return null;
    };
};

$exports['unsafeSet'] = function($s) {
    return function($a) use ($s) {
        return function($o) use ($s, $a) {
            if (is_array($o)) {
                $o1 = $o;
                $o1[$s] = $a;
                return $o1;
            } else {
                $o1 = clone $o;
                $o1->{$s} = $a;
                return $o1;
            }
        };
    };
};
