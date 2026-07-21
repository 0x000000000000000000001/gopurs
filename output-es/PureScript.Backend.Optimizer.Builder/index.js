import * as $runtime from "../runtime.js";
import * as Data$dList from "../Data.List/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as PureScript$dBackend$dOptimizer$dConvert from "../PureScript.Backend.Optimizer.Convert/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn from "../PureScript.Backend.Optimizer.CoreFn/index.js";
import * as PureScript$dBackend$dOptimizer$dSemantics from "../PureScript.Backend.Optimizer.Semantics/index.js";
const buildModules = dictMonad => {
  const Bind1 = dictMonad.Bind1();
  const $$void = Bind1.Apply0().Functor0().map(v => {});
  return options => coreFnModules => {
    const go = go$a0$copy => go$a1$copy => {
      let go$a0 = go$a0$copy, go$a1 = go$a1$copy, go$c = true, go$r;
      while (go$c) {
        const b = go$a0, v = go$a1;
        if (v.tag === "Nil") {
          go$c = false;
          go$r = b;
          continue;
        }
        if (v.tag === "Cons") {
          go$a0 = b + 1 | 0;
          go$a1 = v._2;
          continue;
        }
        $runtime.fail();
      }
      return go$r;
    };
    const moduleCount = go(0)(coreFnModules);
    return $$void(Data$dList.foldM(dictMonad)(v => coreFnModule => {
      const $0 = v.directives;
      const $1 = v.implementations;
      const $2 = v.moduleIndex;
      const buildEnv = {implementations: $1, moduleCount, moduleIndex: $2};
      return Bind1.bind(options.onPrepareModule(buildEnv)(coreFnModule))(v1 => {
        const v2 = PureScript$dBackend$dOptimizer$dConvert.toBackendModule(v1)({
          analyzeCustom: options.analyzeCustom,
          currentModule: v1.name,
          currentLevel: 0,
          toLevel: Data$dMap$dInternal.Leaf,
          implementations: $1,
          moduleImplementations: Data$dMap$dInternal.Leaf,
          directives: $0,
          dataTypes: Data$dMap$dInternal.Leaf,
          foreignSemantics: options.foreignSemantics,
          rewriteLimit: 10000,
          traceIdents: options.traceIdents,
          optimizationSteps: []
        });
        const $3 = v2._2;
        const go$1 = (m$p, z$p) => {
          if (m$p.tag === "Leaf") { return z$p; }
          if (m$p.tag === "Node") {
            return go$1(m$p._5, Data$dMap$dInternal.insert(PureScript$dBackend$dOptimizer$dCoreFn.ordQualified(Data$dOrd.ordString))(m$p._3)(m$p._4)(go$1(m$p._6, z$p)));
          }
          $runtime.fail();
        };
        const newImplementations = go$1($3.implementations, $1);
        return Bind1.bind(options.onCodegenModule({...buildEnv, implementations: newImplementations})(v1)($3)(v2._1))(() => dictMonad.Applicative0().pure({
          directives: (() => {
            const go$2 = (m$p, z$p) => {
              if (m$p.tag === "Leaf") { return z$p; }
              if (m$p.tag === "Node") { return go$2(m$p._5, Data$dMap$dInternal.insert(PureScript$dBackend$dOptimizer$dSemantics.ordEvalRef)(m$p._3)(m$p._4)(go$2(m$p._6, z$p))); }
              $runtime.fail();
            };
            return go$2($3.directives, $0);
          })(),
          implementations: newImplementations,
          moduleIndex: $2 + 1 | 0
        }));
      });
    })({directives: options.directives, implementations: Data$dMap$dInternal.Leaf, moduleIndex: 0})(coreFnModules));
  };
};
export {buildModules};
