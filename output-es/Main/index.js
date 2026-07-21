import * as $runtime from "../runtime.js";
import * as Data$dArgonaut$dDecode$dError from "../Data.Argonaut.Decode.Error/index.js";
import * as Data$dArgonaut$dParser from "../Data.Argonaut.Parser/index.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dString$dCommon from "../Data.String.Common/index.js";
import * as Data$dTraversable from "../Data.Traversable/index.js";
import * as Effect$dAff from "../Effect.Aff/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
import * as Gopurs$dCodeGen from "../Gopurs.CodeGen/index.js";
import * as Gopurs$dRuntime from "../Gopurs.Runtime/index.js";
import * as Node$dEncoding from "../Node.Encoding/index.js";
import * as Node$dFS$dAff from "../Node.FS.Aff/index.js";
import * as Node$dFS$dAsync from "../Node.FS.Async/index.js";
import * as Node$dFS$dStats from "../Node.FS.Stats/index.js";
import * as PureScript$dBackend$dOptimizer$dBuilder from "../PureScript.Backend.Optimizer.Builder/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn$dJson from "../PureScript.Backend.Optimizer.CoreFn.Json/index.js";
import * as PureScript$dBackend$dOptimizer$dCoreFn$dSort from "../PureScript.Backend.Optimizer.CoreFn.Sort/index.js";
const filterA = /* #__PURE__ */ Data$dArray.filterA(Effect$dAff.applicativeAff);
const traverse = /* #__PURE__ */ (() => Data$dTraversable.traversableArray.traverse(Effect$dAff.applicativeAff))();
const fromFoldable = /* #__PURE__ */ Data$dFoldable.foldrArray(Data$dList$dTypes.Cons)(Data$dList$dTypes.Nil);
const buildModules = /* #__PURE__ */ PureScript$dBackend$dOptimizer$dBuilder.buildModules(Effect$dAff.monadAff);
const readCoreFnModule = filePath => Effect$dAff._bind(Effect$dAff.try(Node$dFS$dAff.toAff1(Node$dFS$dAsync.stat)(filePath)))(statRes => {
  if (
    (() => {
      if (statRes.tag === "Left") { return false; }
      if (statRes.tag === "Right") { return true; }
      $runtime.fail();
    })()
  ) {
    return Effect$dAff._bind(Node$dFS$dAff.toAff2(Node$dFS$dAsync.readTextFile)(Node$dEncoding.UTF8)(filePath))(contents => {
      const $0 = Data$dArgonaut$dParser._jsonParser(Data$dEither.Left, Data$dEither.Right, contents);
      const v = (() => {
        if ($0.tag === "Left") {
          const $1 = $0._1;
          return v => Data$dEither.$Either("Left", $1);
        }
        if ($0.tag === "Right") {
          const $1 = $0._1;
          return f => f($1);
        }
        $runtime.fail();
      })()(x => {
        const $1 = PureScript$dBackend$dOptimizer$dCoreFn$dJson.decodeModule$p(PureScript$dBackend$dOptimizer$dCoreFn$dJson.decodeAnn)(x);
        if ($1.tag === "Left") { return Data$dEither.$Either("Left", Data$dArgonaut$dDecode$dError.printJsonDecodeError($1._1)); }
        if ($1.tag === "Right") { return Data$dEither.$Either("Right", $1._1); }
        $runtime.fail();
      });
      if (v.tag === "Left") {
        return Effect$dAff._bind(Effect$dAff._liftEffect(Effect$dConsole.error("Failed to decode " + filePath + ": " + v._1)))(() => Effect$dAff._pure(Data$dMaybe.Nothing));
      }
      if (v.tag === "Right") { return Effect$dAff._pure(Data$dMaybe.$Maybe("Just", v._1)); }
      $runtime.fail();
    });
  }
  return Effect$dAff._pure(Data$dMaybe.Nothing);
});
const main = /* #__PURE__ */ (() => {
  const $0 = Effect$dAff._makeFiber(
    Effect$dAff.ffiUtil,
    Effect$dAff._bind(Node$dFS$dAff.toAff1(Node$dFS$dAsync.readdir)("output"))(files => Effect$dAff._bind(filterA(f => Effect$dAff._bind(Node$dFS$dAff.toAff1(Node$dFS$dAsync.stat)("output/" + f))(stat => Effect$dAff._pure(Node$dFS$dStats.isDirectoryImpl(stat))))(files))(validDirs => Effect$dAff._bind(traverse(dir => readCoreFnModule("output/" + dir + "/corefn.json"))(validDirs))(mbModules => {
      const finalModules = PureScript$dBackend$dOptimizer$dCoreFn$dSort.sortModules(Data$dList$dTypes.foldableList)(fromFoldable(Data$dArray.mapMaybe(x => x)(mbModules)));
      return Effect$dAff._bind(Effect$dAff.try(Node$dFS$dAff.toAff1(Node$dFS$dAsync.mkdir)("output/gopurs_runtime")))(() => Effect$dAff._bind(Node$dFS$dAff.toAff3(Node$dFS$dAsync.writeTextFile)(Node$dEncoding.UTF8)("output/gopurs_runtime/runtime.go")(Gopurs$dRuntime.runtimeGoCode))(() => Effect$dAff._bind(Node$dFS$dAff.toAff3(Node$dFS$dAsync.writeTextFile)(Node$dEncoding.UTF8)("output/go.mod")("module gopurs/output\n\ngo 1.22\n"))(() => Effect$dAff._bind(buildModules({
        directives: Data$dMap$dInternal.Leaf,
        analyzeCustom: v => v1 => Data$dMaybe.Nothing,
        foreignSemantics: Data$dMap$dInternal.Leaf,
        traceIdents: Data$dMap$dInternal.Leaf,
        onPrepareModule: v => m => Effect$dAff._pure(m),
        onCodegenModule: v => v1 => backendMod => v2 => Node$dFS$dAff.toAff3(Node$dFS$dAsync.writeTextFile)(Node$dEncoding.UTF8)("output/" + backendMod.name + "/" + Data$dString$dCommon.replaceAll(".")("_")(backendMod.name) + ".go")(Gopurs$dCodeGen.translate(Data$dFunctor.arrayMap(i => Data$dString$dCommon.split(".")(i._2))(v1.imports))(backendMod))
      })(finalModules))(() => Node$dFS$dAff.toAff3(Node$dFS$dAsync.writeTextFile)(Node$dEncoding.UTF8)("output/main.go")("package main\n\nimport (\n\t\"gopurs/output/Test1110\"\n\t\"gopurs/output/gopurs_runtime\"\n)\n\nfunc main() {\n\tgopurs_runtime.Apply(Test1110.Main, gopurs_runtime.Value{})\n}\n")))));
    })))
  );
  return () => {
    const fiber = $0();
    fiber.run();
  };
})();
export {buildModules, filterA, fromFoldable, main, readCoreFnModule, traverse};
