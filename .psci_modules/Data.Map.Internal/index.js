// | This module defines a type of maps as height-balanced (AVL) binary trees.
// | Efficient set operations are implemented in terms of
// | <https://www.cs.cmu.edu/~guyb/papers/BFS16.pdf>
import * as Control_Applicative from "../Control.Applicative/index.js";
import * as Control_Apply from "../Control.Apply/index.js";
import * as Control_Category from "../Control.Category/index.js";
import * as Data_Boolean from "../Data.Boolean/index.js";
import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Foldable from "../Data.Foldable/index.js";
import * as Data_FoldableWithIndex from "../Data.FoldableWithIndex/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_List_Types from "../Data.List.Types/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Monoid from "../Data.Monoid/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Ordering from "../Data.Ordering/index.js";
import * as Data_Ring from "../Data.Ring/index.js";
import * as Data_Semigroup from "../Data.Semigroup/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Data_Traversable from "../Data.Traversable/index.js";
import * as Data_Tuple from "../Data.Tuple/index.js";
import * as Data_Unfoldable from "../Data.Unfoldable/index.js";
import * as Data_Unit from "../Data.Unit/index.js";
var $runtime_lazy = function (name, moduleName, init) {
    var state = 0;
    var val;
    return function (lineNumber) {
        if (state === 2) return val;
        if (state === 1) throw new ReferenceError(name + " was needed before it finished initializing (module " + moduleName + ", line " + lineNumber + ")", moduleName, lineNumber);
        state = 1;
        val = init();
        state = 2;
        return val;
    };
};
var identity = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity1 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);
var identity2 = /* #__PURE__ */ Control_Category.identity(Control_Category.categoryFn);

// | `Map k v` represents maps from keys of type `k` to values of type `v`.
var Leaf = /* #__PURE__ */ (function () {
    function Leaf() {

    };
    Leaf.value = new Leaf();
    return Leaf;
})();

// | `Map k v` represents maps from keys of type `k` to values of type `v`.
var Node = /* #__PURE__ */ (function () {
    function Node(value0, value1, value2, value3, value4, value5) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
        this.value3 = value3;
        this.value4 = value4;
        this.value5 = value5;
    };
    Node.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return function (value3) {
                    return function (value4) {
                        return function (value5) {
                            return new Node(value0, value1, value2, value3, value4, value5);
                        };
                    };
                };
            };
        };
    };
    return Node;
})();

// | Low-level iteration state for a `Map`. Must be consumed using
// | an appropriate stepper.
var IterLeaf = /* #__PURE__ */ (function () {
    function IterLeaf() {

    };
    IterLeaf.value = new IterLeaf();
    return IterLeaf;
})();

// | Low-level iteration state for a `Map`. Must be consumed using
// | an appropriate stepper.
var IterEmit = /* #__PURE__ */ (function () {
    function IterEmit(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    IterEmit.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new IterEmit(value0, value1, value2);
            };
        };
    };
    return IterEmit;
})();

// | Low-level iteration state for a `Map`. Must be consumed using
// | an appropriate stepper.
var IterNode = /* #__PURE__ */ (function () {
    function IterNode(value0, value1) {
        this.value0 = value0;
        this.value1 = value1;
    };
    IterNode.create = function (value0) {
        return function (value1) {
            return new IterNode(value0, value1);
        };
    };
    return IterNode;
})();
var IterDone = /* #__PURE__ */ (function () {
    function IterDone() {

    };
    IterDone.value = new IterDone();
    return IterDone;
})();
var IterNext = /* #__PURE__ */ (function () {
    function IterNext(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    IterNext.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new IterNext(value0, value1, value2);
            };
        };
    };
    return IterNext;
})();
var Split = /* #__PURE__ */ (function () {
    function Split(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    Split.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new Split(value0, value1, value2);
            };
        };
    };
    return Split;
})();
var SplitLast = /* #__PURE__ */ (function () {
    function SplitLast(value0, value1, value2) {
        this.value0 = value0;
        this.value1 = value1;
        this.value2 = value2;
    };
    SplitLast.create = function (value0) {
        return function (value1) {
            return function (value2) {
                return new SplitLast(value0, value1, value2);
            };
        };
    };
    return SplitLast;
})();

// | Low-level Node constructor which maintains the height and size invariants
// | This is unsafe because it assumes the child Maps are ordered and balanced.
var unsafeNode = function (k, v, l, r) {
    if (l instanceof Leaf) {
        if (r instanceof Leaf) {
            return new Node(1, 1, k, v, l, r);
        };
        if (r instanceof Node) {
            return new Node(1 + r.value0 | 0, 1 + r.value1 | 0, k, v, l, r);
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 702, column 5 - line 706, column 39): " + [ r.constructor.name ]);
    };
    if (l instanceof Node) {
        if (r instanceof Leaf) {
            return new Node(1 + l.value0 | 0, 1 + l.value1 | 0, k, v, l, r);
        };
        if (r instanceof Node) {
            return new Node(1 + (function () {
                var $212 = l.value0 > r.value0;
                if ($212) {
                    return l.value0;
                };
                return r.value0;
            })() | 0, (1 + l.value1 | 0) + r.value1 | 0, k, v, l, r);
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 708, column 5 - line 712, column 68): " + [ r.constructor.name ]);
    };
    throw new Error("Failed pattern match at Data.Map.Internal (line 700, column 32 - line 712, column 68): " + [ l.constructor.name ]);
};

// | Converts a Map to a MapIter for iteration using a MapStepper.
var toMapIter = /* #__PURE__ */ (function () {
    return Data_Function.flip(IterNode.create)(IterLeaf.value);
})();
var stepWith = function (f) {
    return function (next) {
        return function (done) {
            var go = function ($copy_v) {
                var $tco_done = false;
                var $tco_result;
                function $tco_loop(v) {
                    if (v instanceof IterLeaf) {
                        $tco_done = true;
                        return done(Data_Unit.unit);
                    };
                    if (v instanceof IterEmit) {
                        $tco_done = true;
                        return next(v.value0, v.value1, v.value2);
                    };
                    if (v instanceof IterNode) {
                        $copy_v = f(v.value1)(v.value0);
                        return;
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 940, column 8 - line 946, column 20): " + [ v.constructor.name ]);
                };
                while (!$tco_done) {
                    $tco_result = $tco_loop($copy_v);
                };
                return $tco_result;
            };
            return go;
        };
    };
};

// | Calculate the number of key/value pairs in a map
var size = function (v) {
    if (v instanceof Leaf) {
        return 0;
    };
    if (v instanceof Node) {
        return v.value1;
    };
    throw new Error("Failed pattern match at Data.Map.Internal (line 618, column 8 - line 620, column 24): " + [ v.constructor.name ]);
};

// | Create a map with one key/value pair
var singleton = function (k) {
    return function (v) {
        return new Node(1, 1, k, v, Leaf.value, Leaf.value);
    };
};

// | Low-level Node constructor which maintains the balance invariants.
// | This is unsafe because it assumes the child Maps are ordered.
var unsafeBalancedNode = /* #__PURE__ */ (function () {
    var height = function (v) {
        if (v instanceof Leaf) {
            return 0;
        };
        if (v instanceof Node) {
            return v.value0;
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 757, column 12 - line 759, column 26): " + [ v.constructor.name ]);
    };
    var rotateLeft = function (k, v, l, rk, rv, rl, rr) {
        if (rl instanceof Node && rl.value0 > height(rr)) {
            return unsafeNode(rl.value2, rl.value3, unsafeNode(k, v, l, rl.value4), unsafeNode(rk, rv, rl.value5, rr));
        };
        return unsafeNode(rk, rv, unsafeNode(k, v, l, rl), rr);
    };
    var rotateRight = function (k, v, lk, lv, ll, lr, r) {
        if (lr instanceof Node && height(ll) <= lr.value0) {
            return unsafeNode(lr.value2, lr.value3, unsafeNode(lk, lv, ll, lr.value4), unsafeNode(k, v, lr.value5, r));
        };
        return unsafeNode(lk, lv, ll, unsafeNode(k, v, lr, r));
    };
    return function (k, v, l, r) {
        if (l instanceof Leaf) {
            if (r instanceof Leaf) {
                return singleton(k)(v);
            };
            if (r instanceof Node && r.value0 > 1) {
                return rotateLeft(k, v, l, r.value2, r.value3, r.value4, r.value5);
            };
            return unsafeNode(k, v, l, r);
        };
        if (l instanceof Node) {
            if (r instanceof Node) {
                if (r.value0 > (l.value0 + 1 | 0)) {
                    return rotateLeft(k, v, l, r.value2, r.value3, r.value4, r.value5);
                };
                if (l.value0 > (r.value0 + 1 | 0)) {
                    return rotateRight(k, v, l.value2, l.value3, l.value4, l.value5, r);
                };
            };
            if (r instanceof Leaf && l.value0 > 1) {
                return rotateRight(k, v, l.value2, l.value3, l.value4, l.value5, r);
            };
            return unsafeNode(k, v, l, r);
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 717, column 40 - line 738, column 34): " + [ l.constructor.name ]);
    };
})();
var $lazy_unsafeSplit = /* #__PURE__ */ $runtime_lazy("unsafeSplit", "Data.Map.Internal", function () {
    return function (comp, k, m) {
        if (m instanceof Leaf) {
            return new Split(Data_Maybe.Nothing.value, Leaf.value, Leaf.value);
        };
        if (m instanceof Node) {
            var v = comp(k)(m.value2);
            if (v instanceof Data_Ordering.LT) {
                var v1 = $lazy_unsafeSplit(793)(comp, k, m.value4);
                return new Split(v1.value0, v1.value1, unsafeBalancedNode(m.value2, m.value3, v1.value2, m.value5));
            };
            if (v instanceof Data_Ordering.GT) {
                var v1 = $lazy_unsafeSplit(796)(comp, k, m.value5);
                return new Split(v1.value0, unsafeBalancedNode(m.value2, m.value3, m.value4, v1.value1), v1.value2);
            };
            if (v instanceof Data_Ordering.EQ) {
                return new Split(new Data_Maybe.Just(m.value3), m.value4, m.value5);
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 791, column 5 - line 799, column 30): " + [ v.constructor.name ]);
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 787, column 34 - line 799, column 30): " + [ m.constructor.name ]);
    };
});
var unsafeSplit = /* #__PURE__ */ $lazy_unsafeSplit(786);
var $lazy_unsafeSplitLast = /* #__PURE__ */ $runtime_lazy("unsafeSplitLast", "Data.Map.Internal", function () {
    return function (k, v, l, r) {
        if (r instanceof Leaf) {
            return new SplitLast(k, v, l);
        };
        if (r instanceof Node) {
            var v1 = $lazy_unsafeSplitLast(779)(r.value2, r.value3, r.value4, r.value5);
            return new SplitLast(v1.value0, v1.value1, unsafeBalancedNode(k, v, l, v1.value2));
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 776, column 37 - line 780, column 57): " + [ r.constructor.name ]);
    };
});
var unsafeSplitLast = /* #__PURE__ */ $lazy_unsafeSplitLast(775);

// | Low-level Node constructor from two Maps.
// | This is unsafe because it assumes the child Maps are ordered.
var unsafeJoinNodes = function (v, v1) {
    if (v instanceof Leaf) {
        return v1;
    };
    if (v instanceof Node) {
        var v2 = unsafeSplitLast(v.value2, v.value3, v.value4, v.value5);
        return unsafeBalancedNode(v2.value0, v2.value1, v2.value2, v1);
    };
    throw new Error("Failed pattern match at Data.Map.Internal (line 764, column 25 - line 768, column 38): " + [ v.constructor.name, v1.constructor.name ]);
};
var $lazy_unsafeDifference = /* #__PURE__ */ $runtime_lazy("unsafeDifference", "Data.Map.Internal", function () {
    return function (comp, l, r) {
        if (l instanceof Leaf) {
            return Leaf.value;
        };
        if (r instanceof Leaf) {
            return l;
        };
        if (r instanceof Node) {
            var v = unsafeSplit(comp, r.value2, l);
            var l$prime = $lazy_unsafeDifference(841)(comp, v.value1, r.value4);
            var r$prime = $lazy_unsafeDifference(842)(comp, v.value2, r.value5);
            return unsafeJoinNodes(l$prime, r$prime);
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 836, column 39 - line 843, column 33): " + [ l.constructor.name, r.constructor.name ]);
    };
});
var unsafeDifference = /* #__PURE__ */ $lazy_unsafeDifference(835);
var $lazy_unsafeIntersectionWith = /* #__PURE__ */ $runtime_lazy("unsafeIntersectionWith", "Data.Map.Internal", function () {
    return function (comp, app, l, r) {
        if (l instanceof Leaf) {
            return Leaf.value;
        };
        if (r instanceof Leaf) {
            return Leaf.value;
        };
        if (r instanceof Node) {
            var v = unsafeSplit(comp, r.value2, l);
            var l$prime = $lazy_unsafeIntersectionWith(825)(comp, app, v.value1, r.value4);
            var r$prime = $lazy_unsafeIntersectionWith(826)(comp, app, v.value2, r.value5);
            if (v.value0 instanceof Data_Maybe.Just) {
                return unsafeBalancedNode(r.value2, app(v.value0.value0)(r.value3), l$prime, r$prime);
            };
            if (v.value0 instanceof Data_Maybe.Nothing) {
                return unsafeJoinNodes(l$prime, r$prime);
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 827, column 5 - line 831, column 37): " + [ v.value0.constructor.name ]);
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 820, column 49 - line 831, column 37): " + [ l.constructor.name, r.constructor.name ]);
    };
});
var unsafeIntersectionWith = /* #__PURE__ */ $lazy_unsafeIntersectionWith(819);
var $lazy_unsafeUnionWith = /* #__PURE__ */ $runtime_lazy("unsafeUnionWith", "Data.Map.Internal", function () {
    return function (comp, app, l, r) {
        if (l instanceof Leaf) {
            return r;
        };
        if (r instanceof Leaf) {
            return l;
        };
        if (r instanceof Node) {
            var v = unsafeSplit(comp, r.value2, l);
            var l$prime = $lazy_unsafeUnionWith(809)(comp, app, v.value1, r.value4);
            var r$prime = $lazy_unsafeUnionWith(810)(comp, app, v.value2, r.value5);
            if (v.value0 instanceof Data_Maybe.Just) {
                return unsafeBalancedNode(r.value2, app(v.value0.value0)(r.value3), l$prime, r$prime);
            };
            if (v.value0 instanceof Data_Maybe.Nothing) {
                return unsafeBalancedNode(r.value2, r.value3, l$prime, r$prime);
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 811, column 5 - line 815, column 46): " + [ v.value0.constructor.name ]);
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 804, column 42 - line 815, column 46): " + [ l.constructor.name, r.constructor.name ]);
    };
});
var unsafeUnionWith = /* #__PURE__ */ $lazy_unsafeUnionWith(803);

// | Compute the union of two maps, using the specified function
// | to combine values for duplicate keys.
var unionWith = function (dictOrd) {
    var compare = Data_Ord.compare(dictOrd);
    return function (app) {
        return function (m1) {
            return function (m2) {
                return unsafeUnionWith(compare, app, m1, m2);
            };
        };
    };
};

// | Compute the union of two maps, preferring values from the first map in the case
// | of duplicate keys
var union = function (dictOrd) {
    return unionWith(dictOrd)(Data_Function["const"]);
};

// | Update or delete the value for a key in a map
var update = function (dictOrd) {
    return function (f) {
        return function (k) {
            var go = function (v) {
                if (v instanceof Leaf) {
                    return Leaf.value;
                };
                if (v instanceof Node) {
                    var v1 = Data_Ord.compare(dictOrd)(k)(v.value2);
                    if (v1 instanceof Data_Ordering.LT) {
                        return unsafeBalancedNode(v.value2, v.value3, go(v.value4), v.value5);
                    };
                    if (v1 instanceof Data_Ordering.GT) {
                        return unsafeBalancedNode(v.value2, v.value3, v.value4, go(v.value5));
                    };
                    if (v1 instanceof Data_Ordering.EQ) {
                        var v2 = f(v.value3);
                        if (v2 instanceof Data_Maybe.Nothing) {
                            return unsafeJoinNodes(v.value4, v.value5);
                        };
                        if (v2 instanceof Data_Maybe.Just) {
                            return new Node(v.value0, v.value1, v.value2, v2.value0, v.value4, v.value5);
                        };
                        throw new Error("Failed pattern match at Data.Map.Internal (line 531, column 11 - line 535, column 38): " + [ v2.constructor.name ]);
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 527, column 7 - line 535, column 38): " + [ v1.constructor.name ]);
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 524, column 8 - line 535, column 38): " + [ v.constructor.name ]);
            };
            return go;
        };
    };
};

// | Render a `Map` as a `String`
var showTree = function (dictShow) {
    return function (dictShow1) {
        var go = function (ind) {
            return function (v) {
                if (v instanceof Leaf) {
                    return ind + "Leaf";
                };
                if (v instanceof Node) {
                    return ind + ("[" + (Data_Show.show(Data_Show.showInt)(v.value0) + ("] " + (Data_Show.show(dictShow)(v.value2) + (" => " + (Data_Show.show(dictShow1)(v.value3) + "\x0a")))))) + (go(ind + "    ")(v.value4) + "\x0a" + go(ind + "    ")(v.value5));
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 233, column 12 - line 238, column 34): " + [ v.constructor.name ]);
            };
        };
        return go("");
    };
};
var semigroupMap = function () {
    return function (dictOrd) {
        return function (dictSemigroup) {
            return {
                append: unionWith(dictOrd)(Data_Semigroup.append(dictSemigroup))
            };
        };
    };
};
var semigroupMap1 = /* #__PURE__ */ semigroupMap();

// | Delete a key and its corresponding value from a map, returning the value
// | as well as the subsequent map.
var pop = function (dictOrd) {
    var compare = Data_Ord.compare(dictOrd);
    return function (k) {
        return function (m) {
            var v = unsafeSplit(compare, k, m);
            return Data_Functor.map(Data_Maybe.functorMaybe)(function (a) {
                return new Data_Tuple.Tuple(a, unsafeJoinNodes(v.value1, v.value2));
            })(v.value0);
        };
    };
};

// | Test if a key is a member of a map
var member = function (dictOrd) {
    return function (k) {
        var go = function ($copy_v) {
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v) {
                if (v instanceof Leaf) {
                    $tco_done = true;
                    return false;
                };
                if (v instanceof Node) {
                    var v1 = Data_Ord.compare(dictOrd)(k)(v.value2);
                    if (v1 instanceof Data_Ordering.LT) {
                        $copy_v = v.value4;
                        return;
                    };
                    if (v1 instanceof Data_Ordering.GT) {
                        $copy_v = v.value5;
                        return;
                    };
                    if (v1 instanceof Data_Ordering.EQ) {
                        $tco_done = true;
                        return true;
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 459, column 7 - line 462, column 19): " + [ v1.constructor.name ]);
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 456, column 8 - line 462, column 19): " + [ v.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($copy_v);
            };
            return $tco_result;
        };
        return go;
    };
};

// | Applies a function to each key/value pair in a map, discarding entries
// | where the function returns `Nothing`.
var mapMaybeWithKey = function (dictOrd) {
    return function (f) {
        var go = function (v) {
            if (v instanceof Leaf) {
                return Leaf.value;
            };
            if (v instanceof Node) {
                var v2 = f(v.value2)(v.value3);
                if (v2 instanceof Data_Maybe.Just) {
                    return unsafeBalancedNode(v.value2, v2.value0, go(v.value4), go(v.value5));
                };
                if (v2 instanceof Data_Maybe.Nothing) {
                    return unsafeJoinNodes(go(v.value4), go(v.value5));
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 661, column 7 - line 665, column 47): " + [ v2.constructor.name ]);
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 658, column 8 - line 665, column 47): " + [ v.constructor.name ]);
        };
        return go;
    };
};

// | Applies a function to each value in a map, discarding entries where the
// | function returns `Nothing`.
var mapMaybe = function (dictOrd) {
    var $712 = mapMaybeWithKey(dictOrd);
    return function ($713) {
        return $712(Data_Function["const"]($713));
    };
};

// | Look up a value for the specified key, or the greatest one less than it
var lookupLE = function (dictOrd) {
    return function (k) {
        var go = function (v) {
            if (v instanceof Leaf) {
                return Data_Maybe.Nothing.value;
            };
            if (v instanceof Node) {
                var v1 = Data_Ord.compare(dictOrd)(k)(v.value2);
                if (v1 instanceof Data_Ordering.LT) {
                    return go(v.value4);
                };
                if (v1 instanceof Data_Ordering.GT) {
                    var v2 = go(v.value5);
                    if (v2 instanceof Data_Maybe.Nothing) {
                        return new Data_Maybe.Just({
                            key: v.value2,
                            value: v.value3
                        });
                    };
                    return v2;
                };
                if (v1 instanceof Data_Ordering.EQ) {
                    return new Data_Maybe.Just({
                        key: v.value2,
                        value: v.value3
                    });
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 295, column 7 - line 302, column 38): " + [ v1.constructor.name ]);
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 292, column 8 - line 302, column 38): " + [ v.constructor.name ]);
        };
        return go;
    };
};

// | Look up a value for the specified key, or the least one greater than it
var lookupGE = function (dictOrd) {
    return function (k) {
        var go = function (v) {
            if (v instanceof Leaf) {
                return Data_Maybe.Nothing.value;
            };
            if (v instanceof Node) {
                var v1 = Data_Ord.compare(dictOrd)(k)(v.value2);
                if (v1 instanceof Data_Ordering.LT) {
                    var v2 = go(v.value4);
                    if (v2 instanceof Data_Maybe.Nothing) {
                        return new Data_Maybe.Just({
                            key: v.value2,
                            value: v.value3
                        });
                    };
                    return v2;
                };
                if (v1 instanceof Data_Ordering.GT) {
                    return go(v.value5);
                };
                if (v1 instanceof Data_Ordering.EQ) {
                    return new Data_Maybe.Just({
                        key: v.value2,
                        value: v.value3
                    });
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 327, column 7 - line 333, column 42): " + [ v1.constructor.name ]);
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 324, column 8 - line 333, column 42): " + [ v.constructor.name ]);
        };
        return go;
    };
};

// | Look up a value for the specified key
var lookup = function (dictOrd) {
    return function (k) {
        var go = function ($copy_v) {
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(v) {
                if (v instanceof Leaf) {
                    $tco_done = true;
                    return Data_Maybe.Nothing.value;
                };
                if (v instanceof Node) {
                    var v1 = Data_Ord.compare(dictOrd)(k)(v.value2);
                    if (v1 instanceof Data_Ordering.LT) {
                        $copy_v = v.value4;
                        return;
                    };
                    if (v1 instanceof Data_Ordering.GT) {
                        $copy_v = v.value5;
                        return;
                    };
                    if (v1 instanceof Data_Ordering.EQ) {
                        $tco_done = true;
                        return new Data_Maybe.Just(v.value3);
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 283, column 7 - line 286, column 22): " + [ v1.constructor.name ]);
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 280, column 8 - line 286, column 22): " + [ v.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($copy_v);
            };
            return $tco_result;
        };
        return go;
    };
};
var iterMapU = function (iter) {
    return function (v) {
        if (v instanceof Leaf) {
            return iter;
        };
        if (v instanceof Node) {
            if (v.value4 instanceof Leaf) {
                if (v.value5 instanceof Leaf) {
                    return new IterEmit(v.value2, v.value3, iter);
                };
                return new IterEmit(v.value2, v.value3, new IterNode(v.value5, iter));
            };
            if (v.value5 instanceof Leaf) {
                return new IterEmit(v.value2, v.value3, new IterNode(v.value4, iter));
            };
            return new IterEmit(v.value2, v.value3, new IterNode(v.value4, new IterNode(v.value5, iter)));
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 973, column 17 - line 988, column 56): " + [ v.constructor.name ]);
    };
};

// | Steps a `MapIter` in arbitrary order with a CPS encoding.
var stepUnorderedCps = /* #__PURE__ */ stepWith(iterMapU);
var stepUnfoldrUnordered = /* #__PURE__ */ (function () {
    var step = function (k, v, next) {
        return new Data_Maybe.Just(new Data_Tuple.Tuple(new Data_Tuple.Tuple(k, v), next));
    };
    return stepUnorderedCps(step)(function (v) {
        return Data_Maybe.Nothing.value;
    });
})();

// | Convert a map to an unfoldable structure of key/value pairs
// |
// | While this traversal is up to 10% faster in benchmarks than `toUnfoldable`,
// | it leaks the underlying map stucture, making it only suitable for applications
// | where order is irrelevant.
// |
// | If you are unsure, use `toUnfoldable`
var toUnfoldableUnordered = function (dictUnfoldable) {
    var $714 = Data_Unfoldable.unfoldr(dictUnfoldable)(stepUnfoldrUnordered);
    return function ($715) {
        return $714(toMapIter($715));
    };
};

// | Steps a `MapIter` in arbitrary order.
var stepUnordered = /* #__PURE__ */ (function () {
    return stepUnorderedCps(function (k, v, next) {
        return new IterNext(k, v, next);
    })(Data_Function["const"](IterDone.value));
})();
var iterMapR = /* #__PURE__ */ (function () {
    var go = function ($copy_iter) {
        return function ($copy_v) {
            var $tco_var_iter = $copy_iter;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(iter, v) {
                if (v instanceof Leaf) {
                    $tco_done = true;
                    return iter;
                };
                if (v instanceof Node) {
                    if (v.value5 instanceof Leaf) {
                        $tco_var_iter = new IterEmit(v.value2, v.value3, iter);
                        $copy_v = v.value4;
                        return;
                    };
                    $tco_var_iter = new IterEmit(v.value2, v.value3, new IterNode(v.value4, iter));
                    $copy_v = v.value5;
                    return;
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 963, column 13 - line 970, column 48): " + [ v.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_iter, $copy_v);
            };
            return $tco_result;
        };
    };
    return go;
})();

// | Steps a `MapIter` in descending order with a CPS encoding.
var stepDescCps = /* #__PURE__ */ stepWith(iterMapR);

// | Steps a `MapIter` in descending order.
var stepDesc = /* #__PURE__ */ (function () {
    return stepDescCps(function (k, v, next) {
        return new IterNext(k, v, next);
    })(Data_Function["const"](IterDone.value));
})();
var iterMapL = /* #__PURE__ */ (function () {
    var go = function ($copy_iter) {
        return function ($copy_v) {
            var $tco_var_iter = $copy_iter;
            var $tco_done = false;
            var $tco_result;
            function $tco_loop(iter, v) {
                if (v instanceof Leaf) {
                    $tco_done = true;
                    return iter;
                };
                if (v instanceof Node) {
                    if (v.value5 instanceof Leaf) {
                        $tco_var_iter = new IterEmit(v.value2, v.value3, iter);
                        $copy_v = v.value4;
                        return;
                    };
                    $tco_var_iter = new IterEmit(v.value2, v.value3, new IterNode(v.value5, iter));
                    $copy_v = v.value4;
                    return;
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 951, column 13 - line 958, column 48): " + [ v.constructor.name ]);
            };
            while (!$tco_done) {
                $tco_result = $tco_loop($tco_var_iter, $copy_v);
            };
            return $tco_result;
        };
    };
    return go;
})();

// | Steps a `MapIter` in ascending order with a CPS encoding.
var stepAscCps = /* #__PURE__ */ stepWith(iterMapL);

// | Steps a `MapIter` in ascending order.
var stepAsc = /* #__PURE__ */ (function () {
    return stepAscCps(function (k, v, next) {
        return new IterNext(k, v, next);
    })(Data_Function["const"](IterDone.value));
})();
var eqMapIter = function (dictEq) {
    return function (dictEq1) {
        return {
            eq: (function () {
                var go = function ($copy_a) {
                    return function ($copy_b) {
                        var $tco_var_a = $copy_a;
                        var $tco_done = false;
                        var $tco_result;
                        function $tco_loop(a, b) {
                            var v = stepAsc(a);
                            if (v instanceof IterNext) {
                                var v2 = stepAsc(b);
                                if (v2 instanceof IterNext && (Data_Eq.eq(dictEq)(v.value0)(v2.value0) && Data_Eq.eq(dictEq1)(v.value1)(v2.value1))) {
                                    $tco_var_a = v.value2;
                                    $copy_b = v2.value2;
                                    return;
                                };
                                $tco_done = true;
                                return false;
                            };
                            if (v instanceof IterDone) {
                                $tco_done = true;
                                return true;
                            };
                            throw new Error("Failed pattern match at Data.Map.Internal (line 859, column 14 - line 868, column 13): " + [ v.constructor.name ]);
                        };
                        while (!$tco_done) {
                            $tco_result = $tco_loop($tco_var_a, $copy_b);
                        };
                        return $tco_result;
                    };
                };
                return go;
            })()
        };
    };
};
var ordMapIter = function (dictOrd) {
    var eqMapIter1 = eqMapIter(dictOrd.Eq0());
    return function (dictOrd1) {
        var eqMapIter2 = eqMapIter1(dictOrd1.Eq0());
        return {
            compare: (function () {
                var go = function ($copy_a) {
                    return function ($copy_b) {
                        var $tco_var_a = $copy_a;
                        var $tco_done = false;
                        var $tco_result;
                        function $tco_loop(a, b) {
                            var v = stepAsc(b);
                            var v1 = stepAsc(a);
                            if (v1 instanceof IterNext && v instanceof IterNext) {
                                var v3 = Data_Ord.compare(dictOrd)(v1.value0)(v.value0);
                                if (v3 instanceof Data_Ordering.EQ) {
                                    var v4 = Data_Ord.compare(dictOrd1)(v1.value1)(v.value1);
                                    if (v4 instanceof Data_Ordering.EQ) {
                                        $tco_var_a = v1.value2;
                                        $copy_b = v.value2;
                                        return;
                                    };
                                    $tco_done = true;
                                    return v4;
                                };
                                $tco_done = true;
                                return v3;
                            };
                            if (v1 instanceof IterDone) {
                                if (v instanceof IterDone) {
                                    $tco_done = true;
                                    return Data_Ordering.EQ.value;
                                };
                                $tco_done = true;
                                return Data_Ordering.LT.value;
                            };
                            if (v instanceof IterDone) {
                                $tco_done = true;
                                return Data_Ordering.GT.value;
                            };
                            throw new Error("Failed pattern match at Data.Map.Internal (line 873, column 14 - line 891, column 11): " + [ v1.constructor.name, v.constructor.name ]);
                        };
                        while (!$tco_done) {
                            $tco_result = $tco_loop($tco_var_a, $copy_b);
                        };
                        return $tco_result;
                    };
                };
                return go;
            })(),
            Eq0: function () {
                return eqMapIter2;
            }
        };
    };
};
var stepUnfoldr = /* #__PURE__ */ (function () {
    var step = function (k, v, next) {
        return new Data_Maybe.Just(new Data_Tuple.Tuple(new Data_Tuple.Tuple(k, v), next));
    };
    return stepAscCps(step)(function (v) {
        return Data_Maybe.Nothing.value;
    });
})();

// | Convert a map to an unfoldable structure of key/value pairs where the keys are in ascending order
var toUnfoldable = function (dictUnfoldable) {
    var $716 = Data_Unfoldable.unfoldr(dictUnfoldable)(stepUnfoldr);
    return function ($717) {
        return $716(toMapIter($717));
    };
};
var showMap = function (dictShow) {
    var showTuple = Data_Tuple.showTuple(dictShow);
    return function (dictShow1) {
        var showArray = Data_Show.showArray(showTuple(dictShow1));
        return {
            show: function (as) {
                return "(fromFoldable " + (Data_Show.show(showArray)(toUnfoldable(Data_Unfoldable.unfoldableArray)(as)) + ")");
            }
        };
    };
};

// | Test whether one map contains all of the keys and values contained in another map
var isSubmap = function (dictOrd) {
    return function (dictEq) {
        var go = function (m1) {
            return function (m2) {
                if (m1 instanceof Leaf) {
                    return true;
                };
                if (m1 instanceof Node) {
                    var v1 = lookup(dictOrd)(m1.value2)(m2);
                    if (v1 instanceof Data_Maybe.Nothing) {
                        return false;
                    };
                    if (v1 instanceof Data_Maybe.Just) {
                        return Data_Eq.eq(dictEq)(m1.value3)(v1.value0) && (go(m1.value4)(m2) && go(m1.value5)(m2));
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 611, column 7 - line 614, column 40): " + [ v1.constructor.name ]);
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 608, column 14 - line 614, column 40): " + [ m1.constructor.name ]);
            };
        };
        return go;
    };
};

// | Test if a map is empty
var isEmpty = function (v) {
    if (v instanceof Leaf) {
        return true;
    };
    return false;
};

// | Compute the intersection of two maps, using the specified function
// | to combine values for duplicate keys.
var intersectionWith = function (dictOrd) {
    var compare = Data_Ord.compare(dictOrd);
    return function (app) {
        return function (m1) {
            return function (m2) {
                return unsafeIntersectionWith(compare, app, m1, m2);
            };
        };
    };
};

// | Compute the intersection of two maps, preferring values from the first map in the case
// | of duplicate keys.
var intersection = function (dictOrd) {
    return intersectionWith(dictOrd)(Data_Function["const"]);
};

// | Inserts or updates a value with the given function.
// |
// | The combining function is called with the existing value as the first
// | argument and the new value as the second argument.
var insertWith = function (dictOrd) {
    return function (app) {
        return function (k) {
            return function (v) {
                var go = function (v1) {
                    if (v1 instanceof Leaf) {
                        return singleton(k)(v);
                    };
                    if (v1 instanceof Node) {
                        var v2 = Data_Ord.compare(dictOrd)(k)(v1.value2);
                        if (v2 instanceof Data_Ordering.LT) {
                            return unsafeBalancedNode(v1.value2, v1.value3, go(v1.value4), v1.value5);
                        };
                        if (v2 instanceof Data_Ordering.GT) {
                            return unsafeBalancedNode(v1.value2, v1.value3, v1.value4, go(v1.value5));
                        };
                        if (v2 instanceof Data_Ordering.EQ) {
                            return new Node(v1.value0, v1.value1, k, app(v1.value3)(v), v1.value4, v1.value5);
                        };
                        throw new Error("Failed pattern match at Data.Map.Internal (line 486, column 7 - line 489, column 44): " + [ v2.constructor.name ]);
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 483, column 8 - line 489, column 44): " + [ v1.constructor.name ]);
                };
                return go;
            };
        };
    };
};

// | Insert or replace a key/value pair in a map
var insert = function (dictOrd) {
    return function (k) {
        return function (v) {
            var go = function (v1) {
                if (v1 instanceof Leaf) {
                    return singleton(k)(v);
                };
                if (v1 instanceof Node) {
                    var v2 = Data_Ord.compare(dictOrd)(k)(v1.value2);
                    if (v2 instanceof Data_Ordering.LT) {
                        return unsafeBalancedNode(v1.value2, v1.value3, go(v1.value4), v1.value5);
                    };
                    if (v2 instanceof Data_Ordering.GT) {
                        return unsafeBalancedNode(v1.value2, v1.value3, v1.value4, go(v1.value5));
                    };
                    if (v2 instanceof Data_Ordering.EQ) {
                        return new Node(v1.value0, v1.value1, k, v, v1.value4, v1.value5);
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 471, column 7 - line 474, column 35): " + [ v2.constructor.name ]);
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 468, column 8 - line 474, column 35): " + [ v1.constructor.name ]);
            };
            return go;
        };
    };
};
var functorMap = {
    map: function (f) {
        var go = function (v) {
            if (v instanceof Leaf) {
                return Leaf.value;
            };
            if (v instanceof Node) {
                return new Node(v.value0, v.value1, v.value2, f(v.value3), go(v.value4), go(v.value5));
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 147, column 10 - line 150, column 39): " + [ v.constructor.name ]);
        };
        return go;
    }
};
var functorWithIndexMap = {
    mapWithIndex: function (f) {
        var go = function (v) {
            if (v instanceof Leaf) {
                return Leaf.value;
            };
            if (v instanceof Node) {
                return new Node(v.value0, v.value1, v.value2, f(v.value2)(v.value3), go(v.value4), go(v.value5));
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 155, column 10 - line 158, column 41): " + [ v.constructor.name ]);
        };
        return go;
    },
    Functor0: function () {
        return functorMap;
    }
};
var foldableMap = {
    foldr: function (f) {
        return function (z) {
            var $lazy_go = $runtime_lazy("go", "Data.Map.Internal", function () {
                return function (m$prime, z$prime) {
                    if (m$prime instanceof Leaf) {
                        return z$prime;
                    };
                    if (m$prime instanceof Node) {
                        return $lazy_go(172)(m$prime.value4, f(m$prime.value3)($lazy_go(172)(m$prime.value5, z$prime)));
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 169, column 26 - line 172, column 43): " + [ m$prime.constructor.name ]);
                };
            });
            var go = $lazy_go(169);
            return function (m) {
                return go(m, z);
            };
        };
    },
    foldl: function (f) {
        return function (z) {
            var $lazy_go = $runtime_lazy("go", "Data.Map.Internal", function () {
                return function (z$prime, m$prime) {
                    if (m$prime instanceof Leaf) {
                        return z$prime;
                    };
                    if (m$prime instanceof Node) {
                        return $lazy_go(178)(f($lazy_go(178)(z$prime, m$prime.value4))(m$prime.value3), m$prime.value5);
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 175, column 26 - line 178, column 43): " + [ m$prime.constructor.name ]);
                };
            });
            var go = $lazy_go(175);
            return function (m) {
                return go(z, m);
            };
        };
    },
    foldMap: function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        return function (f) {
            var go = function (v) {
                if (v instanceof Leaf) {
                    return Data_Monoid.mempty(dictMonoid);
                };
                if (v instanceof Node) {
                    return Data_Semigroup.append(Semigroup0)(go(v.value4))(Data_Semigroup.append(Semigroup0)(f(v.value3))(go(v.value5)));
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 181, column 10 - line 184, column 28): " + [ v.constructor.name ]);
            };
            return go;
        };
    }
};
var foldableWithIndexMap = {
    foldrWithIndex: function (f) {
        return function (z) {
            var $lazy_go = $runtime_lazy("go", "Data.Map.Internal", function () {
                return function (m$prime, z$prime) {
                    if (m$prime instanceof Leaf) {
                        return z$prime;
                    };
                    if (m$prime instanceof Node) {
                        return $lazy_go(192)(m$prime.value4, f(m$prime.value2)(m$prime.value3)($lazy_go(192)(m$prime.value5, z$prime)));
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 189, column 26 - line 192, column 45): " + [ m$prime.constructor.name ]);
                };
            });
            var go = $lazy_go(189);
            return function (m) {
                return go(m, z);
            };
        };
    },
    foldlWithIndex: function (f) {
        return function (z) {
            var $lazy_go = $runtime_lazy("go", "Data.Map.Internal", function () {
                return function (z$prime, m$prime) {
                    if (m$prime instanceof Leaf) {
                        return z$prime;
                    };
                    if (m$prime instanceof Node) {
                        return $lazy_go(198)(f(m$prime.value2)($lazy_go(198)(z$prime, m$prime.value4))(m$prime.value3), m$prime.value5);
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 195, column 26 - line 198, column 45): " + [ m$prime.constructor.name ]);
                };
            });
            var go = $lazy_go(195);
            return function (m) {
                return go(z, m);
            };
        };
    },
    foldMapWithIndex: function (dictMonoid) {
        var Semigroup0 = dictMonoid.Semigroup0();
        return function (f) {
            var go = function (v) {
                if (v instanceof Leaf) {
                    return Data_Monoid.mempty(dictMonoid);
                };
                if (v instanceof Node) {
                    return Data_Semigroup.append(Semigroup0)(go(v.value4))(Data_Semigroup.append(Semigroup0)(f(v.value2)(v.value3))(go(v.value5)));
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 201, column 10 - line 204, column 30): " + [ v.constructor.name ]);
            };
            return go;
        };
    },
    Foldable0: function () {
        return foldableMap;
    }
};

// | Get a list of the keys contained in a map
var keys = /* #__PURE__ */ (function () {
    return Data_FoldableWithIndex.foldrWithIndex(foldableWithIndexMap)(function (k) {
        return function (v) {
            return function (acc) {
                return new Data_List_Types.Cons(k, acc);
            };
        };
    })(Data_List_Types.Nil.value);
})();
var traversableMap = {
    traverse: function (dictApplicative) {
        var Apply0 = dictApplicative.Apply0();
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            var go = function (v) {
                if (v instanceof Leaf) {
                    return Control_Applicative.pure(dictApplicative)(Leaf.value);
                };
                if (v instanceof Node) {
                    return Control_Apply.apply(Apply0)(Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (l$prime) {
                        return function (v$prime) {
                            return function (r$prime) {
                                return new Node(v.value0, v.value1, v.value2, v$prime, l$prime, r$prime);
                            };
                        };
                    })(go(v.value4)))(f(v.value3)))(go(v.value5));
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 209, column 10 - line 215, column 19): " + [ v.constructor.name ]);
            };
            return go;
        };
    },
    sequence: function (dictApplicative) {
        return Data_Traversable.traverse(traversableMap)(dictApplicative)(identity);
    },
    Functor0: function () {
        return functorMap;
    },
    Foldable1: function () {
        return foldableMap;
    }
};
var traversableWithIndexMap = {
    traverseWithIndex: function (dictApplicative) {
        var Apply0 = dictApplicative.Apply0();
        var Functor0 = (dictApplicative.Apply0()).Functor0();
        return function (f) {
            var go = function (v) {
                if (v instanceof Leaf) {
                    return Control_Applicative.pure(dictApplicative)(Leaf.value);
                };
                if (v instanceof Node) {
                    return Control_Apply.apply(Apply0)(Control_Apply.apply(Apply0)(Data_Functor.map(Functor0)(function (l$prime) {
                        return function (v$prime) {
                            return function (r$prime) {
                                return new Node(v.value0, v.value1, v.value2, v$prime, l$prime, r$prime);
                            };
                        };
                    })(go(v.value4)))(f(v.value2)(v.value3)))(go(v.value5));
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 221, column 10 - line 227, column 19): " + [ v.constructor.name ]);
            };
            return go;
        };
    },
    FunctorWithIndex0: function () {
        return functorWithIndexMap;
    },
    FoldableWithIndex1: function () {
        return foldableWithIndexMap;
    },
    Traversable2: function () {
        return traversableMap;
    }
};

// | Get a list of the values contained in a map
var values = /* #__PURE__ */ (function () {
    return Data_Foldable.foldr(foldableMap)(Data_List_Types.Cons.create)(Data_List_Types.Nil.value);
})();
var foldSubmapBy = function (dictOrd) {
    return function (appendFn) {
        return function (memptyValue) {
            return function (kmin) {
                return function (kmax) {
                    return function (f) {
                        var tooSmall = (function () {
                            if (kmin instanceof Data_Maybe.Just) {
                                return function (k) {
                                    return Data_Ord.lessThan(dictOrd)(k)(kmin.value0);
                                };
                            };
                            if (kmin instanceof Data_Maybe.Nothing) {
                                return Data_Function["const"](false);
                            };
                            throw new Error("Failed pattern match at Data.Map.Internal (line 389, column 7 - line 393, column 22): " + [ kmin.constructor.name ]);
                        })();
                        var tooLarge = (function () {
                            if (kmax instanceof Data_Maybe.Just) {
                                return function (k) {
                                    return Data_Ord.greaterThan(dictOrd)(k)(kmax.value0);
                                };
                            };
                            if (kmax instanceof Data_Maybe.Nothing) {
                                return Data_Function["const"](false);
                            };
                            throw new Error("Failed pattern match at Data.Map.Internal (line 396, column 7 - line 400, column 22): " + [ kmax.constructor.name ]);
                        })();
                        var inBounds = (function () {
                            if (kmin instanceof Data_Maybe.Just && kmax instanceof Data_Maybe.Just) {
                                return function (k) {
                                    return Data_Ord.lessThanOrEq(dictOrd)(kmin.value0)(k) && Data_Ord.lessThanOrEq(dictOrd)(k)(kmax.value0);
                                };
                            };
                            if (kmin instanceof Data_Maybe.Just && kmax instanceof Data_Maybe.Nothing) {
                                return function (k) {
                                    return Data_Ord.lessThanOrEq(dictOrd)(kmin.value0)(k);
                                };
                            };
                            if (kmin instanceof Data_Maybe.Nothing && kmax instanceof Data_Maybe.Just) {
                                return function (k) {
                                    return Data_Ord.lessThanOrEq(dictOrd)(k)(kmax.value0);
                                };
                            };
                            if (kmin instanceof Data_Maybe.Nothing && kmax instanceof Data_Maybe.Nothing) {
                                return Data_Function["const"](true);
                            };
                            throw new Error("Failed pattern match at Data.Map.Internal (line 403, column 7 - line 411, column 21): " + [ kmin.constructor.name, kmax.constructor.name ]);
                        })();
                        var go = function (v) {
                            if (v instanceof Leaf) {
                                return memptyValue;
                            };
                            if (v instanceof Node) {
                                return appendFn(appendFn((function () {
                                    var $575 = tooSmall(v.value2);
                                    if ($575) {
                                        return memptyValue;
                                    };
                                    return go(v.value4);
                                })())((function () {
                                    var $576 = inBounds(v.value2);
                                    if ($576) {
                                        return f(v.value2)(v.value3);
                                    };
                                    return memptyValue;
                                })()))((function () {
                                    var $577 = tooLarge(v.value2);
                                    if ($577) {
                                        return memptyValue;
                                    };
                                    return go(v.value5);
                                })());
                            };
                            throw new Error("Failed pattern match at Data.Map.Internal (line 413, column 10 - line 419, column 66): " + [ v.constructor.name ]);
                        };
                        return go;
                    };
                };
            };
        };
    };
};

// | Fold over the entries of a given map where the key is between a lower and
// | an upper bound. Passing `Nothing` as either the lower or upper bound
// | argument means that the fold has no lower or upper bound, i.e. the fold
// | starts from (or ends with) the smallest (or largest) key in the map.
// |
// | ```purescript
// | foldSubmap (Just 1) (Just 2) (\_ v -> [v])
// |  (fromFoldable [Tuple 0 "zero", Tuple 1 "one", Tuple 2 "two", Tuple 3 "three"])
// |  == ["one", "two"]
// |
// | foldSubmap Nothing (Just 2) (\_ v -> [v])
// |  (fromFoldable [Tuple 0 "zero", Tuple 1 "one", Tuple 2 "two", Tuple 3 "three"])
// |  == ["zero", "one", "two"]
// | ```
var foldSubmap = function (dictOrd) {
    return function (dictMonoid) {
        return foldSubmapBy(dictOrd)(Data_Semigroup.append(dictMonoid.Semigroup0()))(Data_Monoid.mempty(dictMonoid));
    };
};

// | Returns the pair with the least key
var findMin = function ($copy_v) {
    var $tco_done = false;
    var $tco_result;
    function $tco_loop(v) {
        if (v instanceof Leaf) {
            $tco_done = true;
            return Data_Maybe.Nothing.value;
        };
        if (v instanceof Node) {
            if (v.value4 instanceof Leaf) {
                $tco_done = true;
                return new Data_Maybe.Just({
                    key: v.value2,
                    value: v.value3
                });
            };
            $copy_v = v.value4;
            return;
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 361, column 11 - line 366, column 21): " + [ v.constructor.name ]);
    };
    while (!$tco_done) {
        $tco_result = $tco_loop($copy_v);
    };
    return $tco_result;
};

// | Look up a value for the least key greater than the specified key
var lookupGT = function (dictOrd) {
    return function (k) {
        var go = function (v) {
            if (v instanceof Leaf) {
                return Data_Maybe.Nothing.value;
            };
            if (v instanceof Node) {
                var v1 = Data_Ord.compare(dictOrd)(k)(v.value2);
                if (v1 instanceof Data_Ordering.LT) {
                    var v2 = go(v.value4);
                    if (v2 instanceof Data_Maybe.Nothing) {
                        return new Data_Maybe.Just({
                            key: v.value2,
                            value: v.value3
                        });
                    };
                    return v2;
                };
                if (v1 instanceof Data_Ordering.GT) {
                    return go(v.value5);
                };
                if (v1 instanceof Data_Ordering.EQ) {
                    return findMin(v.value5);
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 342, column 7 - line 348, column 25): " + [ v1.constructor.name ]);
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 339, column 8 - line 348, column 25): " + [ v.constructor.name ]);
        };
        return go;
    };
};

// | Returns the pair with the greatest key
var findMax = function ($copy_v) {
    var $tco_done = false;
    var $tco_result;
    function $tco_loop(v) {
        if (v instanceof Leaf) {
            $tco_done = true;
            return Data_Maybe.Nothing.value;
        };
        if (v instanceof Node) {
            if (v.value5 instanceof Leaf) {
                $tco_done = true;
                return new Data_Maybe.Just({
                    key: v.value2,
                    value: v.value3
                });
            };
            $copy_v = v.value5;
            return;
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 352, column 11 - line 357, column 21): " + [ v.constructor.name ]);
    };
    while (!$tco_done) {
        $tco_result = $tco_loop($copy_v);
    };
    return $tco_result;
};

// | Look up a value for the greatest key less than the specified key
var lookupLT = function (dictOrd) {
    return function (k) {
        var go = function (v) {
            if (v instanceof Leaf) {
                return Data_Maybe.Nothing.value;
            };
            if (v instanceof Node) {
                var v1 = Data_Ord.compare(dictOrd)(k)(v.value2);
                if (v1 instanceof Data_Ordering.LT) {
                    return go(v.value4);
                };
                if (v1 instanceof Data_Ordering.GT) {
                    var v2 = go(v.value5);
                    if (v2 instanceof Data_Maybe.Nothing) {
                        return new Data_Maybe.Just({
                            key: v.value2,
                            value: v.value3
                        });
                    };
                    return v2;
                };
                if (v1 instanceof Data_Ordering.EQ) {
                    return findMax(v.value4);
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 311, column 7 - line 318, column 21): " + [ v1.constructor.name ]);
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 308, column 8 - line 318, column 21): " + [ v.constructor.name ]);
        };
        return go;
    };
};

// | Filter out those key/value pairs of a map for which a predicate
// | fails to hold.
var filterWithKey = function (dictOrd) {
    return function (f) {
        var go = function (v) {
            if (v instanceof Leaf) {
                return Leaf.value;
            };
            if (v instanceof Node) {
                if (f(v.value2)(v.value3)) {
                    return unsafeBalancedNode(v.value2, v.value3, go(v.value4), go(v.value5));
                };
                if (Data_Boolean.otherwise) {
                    return unsafeJoinNodes(go(v.value4), go(v.value5));
                };
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 627, column 8 - line 633, column 47): " + [ v.constructor.name ]);
        };
        return go;
    };
};

// | Filter out those key/value pairs of a map for which a predicate
// | on the key fails to hold.
var filterKeys = function (dictOrd) {
    return function (f) {
        var go = function (v) {
            if (v instanceof Leaf) {
                return Leaf.value;
            };
            if (v instanceof Node) {
                if (f(v.value2)) {
                    return unsafeBalancedNode(v.value2, v.value3, go(v.value4), go(v.value5));
                };
                if (Data_Boolean.otherwise) {
                    return unsafeJoinNodes(go(v.value4), go(v.value5));
                };
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 640, column 8 - line 646, column 47): " + [ v.constructor.name ]);
        };
        return go;
    };
};

// | Filter out those key/value pairs of a map for which a predicate
// | on the value fails to hold.
var filter = function (dictOrd) {
    var $718 = filterWithKey(dictOrd);
    return function ($719) {
        return $718(Data_Function["const"]($719));
    };
};
var eqMap = function (dictEq) {
    var eqMapIter1 = eqMapIter(dictEq);
    return function (dictEq1) {
        var eqMapIter2 = eqMapIter1(dictEq1);
        return {
            eq: function (xs) {
                return function (ys) {
                    if (xs instanceof Leaf) {
                        if (ys instanceof Leaf) {
                            return true;
                        };
                        return false;
                    };
                    if (xs instanceof Node) {
                        if (ys instanceof Node && xs.value1 === ys.value1) {
                            return Data_Eq.eq(eqMapIter2)(toMapIter(xs))(toMapIter(ys));
                        };
                        return false;
                    };
                    throw new Error("Failed pattern match at Data.Map.Internal (line 94, column 14 - line 105, column 16): " + [ xs.constructor.name ]);
                };
            }
        };
    };
};
var ordMap = function (dictOrd) {
    var ordMapIter1 = ordMapIter(dictOrd);
    var eqMap1 = eqMap(dictOrd.Eq0());
    return function (dictOrd1) {
        var ordMapIter2 = ordMapIter1(dictOrd1);
        var eqMap2 = eqMap1(dictOrd1.Eq0());
        return {
            compare: function (xs) {
                return function (ys) {
                    if (xs instanceof Leaf) {
                        if (ys instanceof Leaf) {
                            return Data_Ordering.EQ.value;
                        };
                        return Data_Ordering.LT.value;
                    };
                    if (ys instanceof Leaf) {
                        return Data_Ordering.GT.value;
                    };
                    return Data_Ord.compare(ordMapIter2)(toMapIter(xs))(toMapIter(ys));
                };
            },
            Eq0: function () {
                return eqMap2;
            }
        };
    };
};
var eq1Map = function (dictEq) {
    var eqMap1 = eqMap(dictEq);
    return {
        eq1: function (dictEq1) {
            return Data_Eq.eq(eqMap1(dictEq1));
        }
    };
};
var ord1Map = function (dictOrd) {
    var ordMap1 = ordMap(dictOrd);
    var eq1Map1 = eq1Map(dictOrd.Eq0());
    return {
        compare1: function (dictOrd1) {
            return Data_Ord.compare(ordMap1(dictOrd1));
        },
        Eq10: function () {
            return eq1Map1;
        }
    };
};

// | An empty map
var empty = /* #__PURE__ */ (function () {
    return Leaf.value;
})();

// | Convert any foldable collection of key/value pairs to a map.
// | On key collision, later values take precedence over earlier ones.
var fromFoldable = function (dictOrd) {
    return function (dictFoldable) {
        return Data_Foldable.foldl(dictFoldable)(function (m) {
            return function (v) {
                return insert(dictOrd)(v.value0)(v.value1)(m);
            };
        })(empty);
    };
};

// | Convert any foldable collection of key/value pairs to a map.
// | On key collision, the values are configurably combined.
var fromFoldableWith = function (dictOrd) {
    return function (dictFoldable) {
        return function (f) {
            var f$prime = insertWith(dictOrd)(Data_Function.flip(f));
            return Data_Foldable.foldl(dictFoldable)(function (m) {
                return function (v) {
                    return f$prime(v.value0)(v.value1)(m);
                };
            })(empty);
        };
    };
};

// | Convert any indexed foldable collection into a map.
var fromFoldableWithIndex = function (dictOrd) {
    return function (dictFoldableWithIndex) {
        return Data_FoldableWithIndex.foldlWithIndex(dictFoldableWithIndex)(function (k) {
            return function (m) {
                return function (v) {
                    return insert(dictOrd)(k)(v)(m);
                };
            };
        })(empty);
    };
};
var monoidSemigroupMap = function () {
    return function (dictOrd) {
        var semigroupMap2 = semigroupMap1(dictOrd);
        return function (dictSemigroup) {
            var semigroupMap3 = semigroupMap2(dictSemigroup);
            return {
                mempty: empty,
                Semigroup0: function () {
                    return semigroupMap3;
                }
            };
        };
    };
};

// | Returns a new map containing all entries of the given map which lie
// | between a given lower and upper bound, treating `Nothing` as no bound i.e.
// | including the smallest (or largest) key in the map, no matter how small
// | (or large) it is. For example:
// |
// | ```purescript
// | submap (Just 1) (Just 2)
// |   (fromFoldable [Tuple 0 "zero", Tuple 1 "one", Tuple 2 "two", Tuple 3 "three"])
// |   == fromFoldable [Tuple 1 "one", Tuple 2 "two"]
// |
// | submap Nothing (Just 2)
// |   (fromFoldable [Tuple 0 "zero", Tuple 1 "one", Tuple 2 "two", Tuple 3 "three"])
// |   == fromFoldable [Tuple 0 "zero", Tuple 1 "one", Tuple 2 "two"]
// | ```
// |
// | The function is entirely specified by the following
// | property:
// |
// | ```purescript
// | Given any m :: Map k v, mmin :: Maybe k, mmax :: Maybe k, key :: k,
// |   let m' = submap mmin mmax m in
// |     if (maybe true (\min -> min <= key) mmin &&
// |         maybe true (\max -> max >= key) mmax)
// |       then lookup key m == lookup key m'
// |       else not (member key m')
// | ```
var submap = function (dictOrd) {
    var union1 = union(dictOrd);
    return function (kmin) {
        return function (kmax) {
            return foldSubmapBy(dictOrd)(union1)(empty)(kmin)(kmax)(singleton);
        };
    };
};

// | Compute the union of a collection of maps
var unions = function (dictOrd) {
    var union1 = union(dictOrd);
    return function (dictFoldable) {
        return Data_Foldable.foldl(dictFoldable)(union1)(empty);
    };
};

// | Difference of two maps. Return elements of the first map where
// | the keys do not exist in the second map.
var difference = function (dictOrd) {
    var compare = Data_Ord.compare(dictOrd);
    return function (m1) {
        return function (m2) {
            return unsafeDifference(compare, m1, m2);
        };
    };
};

// | Delete a key and its corresponding value from a map.
var $$delete = function (dictOrd) {
    return function (k) {
        var go = function (v) {
            if (v instanceof Leaf) {
                return Leaf.value;
            };
            if (v instanceof Node) {
                var v1 = Data_Ord.compare(dictOrd)(k)(v.value2);
                if (v1 instanceof Data_Ordering.LT) {
                    return unsafeBalancedNode(v.value2, v.value3, go(v.value4), v.value5);
                };
                if (v1 instanceof Data_Ordering.GT) {
                    return unsafeBalancedNode(v.value2, v.value3, v.value4, go(v.value5));
                };
                if (v1 instanceof Data_Ordering.EQ) {
                    return unsafeJoinNodes(v.value4, v.value5);
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 498, column 7 - line 501, column 43): " + [ v1.constructor.name ]);
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 495, column 8 - line 501, column 43): " + [ v.constructor.name ]);
        };
        return go;
    };
};

// | Check whether the underlying tree satisfies the height, size, and ordering invariants.
// |
// | This function is provided for internal use.
var checkValid = function (dictOrd) {
    var go = function (v) {
        if (v instanceof Leaf) {
            return true;
        };
        if (v instanceof Node) {
            if (v.value4 instanceof Leaf) {
                if (v.value5 instanceof Leaf) {
                    return true;
                };
                if (v.value5 instanceof Node) {
                    return v.value0 === 2 && (v.value5.value0 === 1 && (v.value1 > v.value5.value1 && (Data_Ord.greaterThan(dictOrd)(v.value5.value2)(v.value2) && go(v.value5))));
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 264, column 11 - line 268, column 60): " + [ v.value5.constructor.name ]);
            };
            if (v.value4 instanceof Node) {
                if (v.value5 instanceof Leaf) {
                    return v.value0 === 2 && (v.value4.value0 === 1 && (v.value1 > v.value4.value1 && (Data_Ord.lessThan(dictOrd)(v.value4.value2)(v.value2) && go(v.value4))));
                };
                if (v.value5 instanceof Node) {
                    return v.value0 > v.value5.value0 && (Data_Ord.greaterThan(dictOrd)(v.value5.value2)(v.value2) && (v.value0 > v.value4.value0 && (Data_Ord.lessThan(dictOrd)(v.value4.value2)(v.value2) && (Data_Ord.abs(Data_Ord.ordInt)(Data_Ring.ringInt)(v.value5.value0 - v.value4.value0 | 0) < 2 && (((v.value5.value1 + v.value4.value1 | 0) + 1 | 0) === v.value1 && (go(v.value4) && go(v.value5)))))));
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 270, column 11 - line 274, column 108): " + [ v.value5.constructor.name ]);
            };
            throw new Error("Failed pattern match at Data.Map.Internal (line 262, column 7 - line 274, column 108): " + [ v.value4.constructor.name ]);
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 259, column 8 - line 274, column 108): " + [ v.constructor.name ]);
    };
    return go;
};

// | Filter a map of optional values, keeping only the key/value pairs which
// | contain a value, creating a new map.
var catMaybes = function (dictOrd) {
    return mapMaybe(dictOrd)(identity1);
};
var applyMap = function (dictOrd) {
    return {
        apply: intersectionWith(dictOrd)(identity2),
        Functor0: function () {
            return functorMap;
        }
    };
};
var bindMap = function (dictOrd) {
    var applyMap1 = applyMap(dictOrd);
    return {
        bind: function (m) {
            return function (f) {
                return mapMaybeWithKey(dictOrd)(function (k) {
                    var $720 = lookup(dictOrd)(k);
                    return function ($721) {
                        return $720(f($721));
                    };
                })(m);
            };
        },
        Apply0: function () {
            return applyMap1;
        }
    };
};

// | Returns true if at least one map element satisfies the given predicate,
// | iterating the map only as necessary and stopping as soon as the predicate
// | yields true.
var anyWithKey = function (predicate) {
    var go = function (v) {
        if (v instanceof Leaf) {
            return false;
        };
        if (v instanceof Node) {
            return predicate(v.value2)(v.value3) || (go(v.value4) || go(v.value5));
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 693, column 8 - line 695, column 62): " + [ v.constructor.name ]);
    };
    return go;
};

// | Returns true if at least one map element satisfies the given predicateon the value,
// | iterating the map only as necessary and stopping as soon as the predicate
// | yields true.
var any = function (predicate) {
    var go = function (v) {
        if (v instanceof Leaf) {
            return false;
        };
        if (v instanceof Node) {
            return predicate(v.value3) || (go(v.value4) || go(v.value5));
        };
        throw new Error("Failed pattern match at Data.Map.Internal (line 683, column 8 - line 685, column 58): " + [ v.constructor.name ]);
    };
    return go;
};

// | Insert the value, delete a value, or update a value for a key in a map
var alter = function (dictOrd) {
    var compare = Data_Ord.compare(dictOrd);
    return function (f) {
        return function (k) {
            return function (m) {
                var v = unsafeSplit(compare, k, m);
                var v2 = f(v.value0);
                if (v2 instanceof Data_Maybe.Nothing) {
                    return unsafeJoinNodes(v.value1, v.value2);
                };
                if (v2 instanceof Data_Maybe.Just) {
                    return unsafeBalancedNode(k, v2.value0, v.value1, v.value2);
                };
                throw new Error("Failed pattern match at Data.Map.Internal (line 514, column 3 - line 518, column 41): " + [ v2.constructor.name ]);
            };
        };
    };
};
var altMap = function (dictOrd) {
    return {
        alt: union(dictOrd),
        Functor0: function () {
            return functorMap;
        }
    };
};
var plusMap = function (dictOrd) {
    var altMap1 = altMap(dictOrd);
    return {
        empty: empty,
        Alt0: function () {
            return altMap1;
        }
    };
};
export {
    Leaf,
    Node,
    showTree,
    empty,
    isEmpty,
    singleton,
    checkValid,
    insert,
    insertWith,
    lookup,
    lookupLE,
    lookupLT,
    lookupGE,
    lookupGT,
    findMin,
    findMax,
    foldSubmap,
    submap,
    fromFoldable,
    fromFoldableWith,
    fromFoldableWithIndex,
    toUnfoldable,
    toUnfoldableUnordered,
    $$delete as delete,
    pop,
    member,
    alter,
    update,
    keys,
    values,
    union,
    unionWith,
    unions,
    intersection,
    intersectionWith,
    difference,
    isSubmap,
    size,
    filterWithKey,
    filterKeys,
    filter,
    mapMaybeWithKey,
    mapMaybe,
    catMaybes,
    any,
    anyWithKey,
    IterDone,
    IterNext,
    toMapIter,
    stepAsc,
    stepAscCps,
    stepDesc,
    stepDescCps,
    stepUnordered,
    stepUnorderedCps,
    unsafeNode,
    unsafeBalancedNode,
    unsafeJoinNodes,
    unsafeSplit,
    Split,
    eq1Map,
    eqMap,
    ord1Map,
    ordMap,
    showMap,
    semigroupMap,
    monoidSemigroupMap,
    altMap,
    plusMap,
    functorMap,
    functorWithIndexMap,
    applyMap,
    bindMap,
    foldableMap,
    foldableWithIndexMap,
    traversableMap,
    traversableWithIndexMap,
    eqMapIter,
    ordMapIter
};
