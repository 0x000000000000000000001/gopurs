// | This module provides compatability functions for constructing `Aff`s which
// | are defined via the FFI.
import * as Data_Either from "../Data.Either/index.js";
import * as Effect_Aff from "../Effect.Aff/index.js";
import * as Effect_Uncurried from "../Effect.Uncurried/index.js";
var EffectFnCanceler = function (x) {
    return x;
};
var EffectFnAff = function (x) {
    return x;
};

// | Lift a FFI definition into an `Aff`. `EffectFnAff` makes use of `EffectFn` so
// | `Effect` thunks are unnecessary. A definition might follow this example:
// |
// | ```javascript
// | exports._myAff = function (onError, onSuccess) {
// |   var cancel = doSomethingAsync(function (err, res) {
// |     if (err) {
// |       onError(err);
// |     } else {
// |       onSuccess(res);
// |     }
// |   });
// |   return function (cancelError, onCancelerError, onCancelerSuccess) {
// |     cancel();
// |     onCancelerSuccess();
// |   };
// | };
// | ```
// |
// | ```purescript
// | foreign import _myAff :: EffectFnAff String
// |
// | myAff :: Aff String
// | myAff = fromEffectFnAff _myAff
// | ````
var fromEffectFnAff = function (v) {
    return Effect_Aff.makeAff(function (k) {
        return function __do() {
            var v1 = v(function ($5) {
                return k(Data_Either.Left.create($5))();
            }, function ($6) {
                return k(Data_Either.Right.create($6))();
            });
            return function (e) {
                return Effect_Aff.makeAff(function (k2) {
                    return function __do() {
                        v1(e, function ($7) {
                            return k2(Data_Either.Left.create($7))();
                        }, function ($8) {
                            return k2(Data_Either.Right.create($8))();
                        });
                        return Effect_Aff.nonCanceler;
                    };
                });
            };
        };
    });
};
export {
    EffectFnAff,
    EffectFnCanceler,
    fromEffectFnAff
};
export {
    mkEffectFn1,
    mkEffectFn2,
    mkEffectFn3,
    runEffectFn1,
    runEffectFn2,
    runEffectFn3
} from "../Effect.Uncurried/index.js";
