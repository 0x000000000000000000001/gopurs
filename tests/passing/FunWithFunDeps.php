<?php
$exports['fnil'] = [];

$exports['fcons'] = function($hd) {
    return function($tl) use ($hd) {
        return \array_merge([$hd], $tl);
    };
};

$exports['fappendImpl'] = function($left) {
    return function($right) use ($left) {
        return \array_merge($left, $right);
    };
};

$exports['fflattenImpl'] = function($v) {
    $accRef = [];
    foreach ($v as $indexRef) {
        $accRef = \array_merge($accRef, $indexRef);
    }
    return $accRef;
};

$exports['ftoArray'] = function($vect) {
    return $vect;
};
