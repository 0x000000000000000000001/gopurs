#!/bin/bash
sed -i.bak -e '396,400c\
                      let ret = if Array.length args < Array.length fArgs then TypeValue else exprTypeToGoType pointerAdtPaths enumAdts elidedCtors modNameStr fRet\
                      in unsafePerformEffect ((if show name == "(Ident \\"contains\\")" then Console.log ("contains fRetGo=" <> goTypeToStr ret <> ", typeSig=" <> show (isJust typeSig) <> ", argsLen=" <> show (Array.length args) <> ", fArgsLen=" <> show (Array.length fArgs)) else pure unit) *> pure ret)' src/Gopurs/CodeGen.purs
