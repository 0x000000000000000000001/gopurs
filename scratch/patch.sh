#!/bin/bash
sed -i.bak -e '374,378c\
    getExprType :: TcoExpr -> ExprType\
    getExprType (TcoExpr _ (Typed ty inner)) = ty\
    getExprType _ = Any -- fallback' src/Gopurs/CodeGen.purs

sed -i.bak -e '447,451c\
                                  paramsWithTypes = case extractExprFuncType (getExprType fn.val) of\
                                    Just { fArgs } -> Array.zip fn.args (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) fArgs <> Array.replicate (Array.length fn.args - Array.length fArgs) TypeValue)\
                                    Nothing -> map (\\p -> Tuple p TypeValue) fn.args\
\
                                  newBound = foldl (\\acc (Tuple idStr goType) -> Map.insert idStr { name: idStr, goType } acc) Map.empty paramsWithTypes' src/Gopurs/CodeGen.purs

sed -i.bak -e '608,612c\
getExprType :: TcoExpr -> ExprType\
getExprType (TcoExpr _ syn) = case syn of\
  Typed t inner -> t' src/Gopurs/CodeGen.purs

sed -i.bak -e '1445,1447c\
          paramsWithTypes = case extractExprFuncType (getExprType tcoExpr) of\
            Just { fArgs } -> Array.zipWith (\\(Tuple mbI lvl) goType -> Tuple (localId mbI lvl) goType) args (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) fArgs <> Array.replicate (Array.length args - Array.length fArgs) TypeValue)\
            Nothing -> map (\\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args' src/Gopurs/CodeGen.purs

sed -i.bak -e '1512,1514c\
          paramsWithTypes = case extractExprFuncType (getExprType tcoExpr) of\
            Just { fArgs } -> Array.zipWith (\\(Tuple mbI lvl) goType -> Tuple (localId mbI lvl) goType) args (map (exprTypeToGoType (unsafePerformEffect (Ref.read helpersRef)).pointerAdtPaths (unsafePerformEffect (Ref.read helpersRef)).enumAdts (unsafePerformEffect (Ref.read helpersRef)).elidedCtors modNameStr) fArgs <> Array.replicate (Array.length args - Array.length fArgs) TypeValue)\
            Nothing -> map (\\(Tuple mbI lvl) -> Tuple (localId mbI lvl) TypeValue) args' src/Gopurs/CodeGen.purs

