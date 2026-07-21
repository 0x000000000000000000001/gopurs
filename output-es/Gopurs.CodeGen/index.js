import * as $runtime from "../runtime.js";
import * as Control$dBind from "../Control.Bind/index.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dSet from "../Data.Set/index.js";
import * as Data$dString$dCommon from "../Data.String.Common/index.js";
import * as Gopurs$dGoAst from "../Gopurs.GoAst/index.js";
import * as Gopurs$dPrinter from "../Gopurs.Printer/index.js";
const translateExpr = v => {
  if (v.tag === "Var") {
    if (v._1._2 === "bindE") {
      return Gopurs$dGoAst.$GoExpr(
        "GoCall",
        Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime"), "Func"),
        [
          Gopurs$dGoAst.$GoExpr(
            "GoFunc",
            ["a"],
            Gopurs$dGoAst.$GoExpr(
              "GoReturn",
              Gopurs$dGoAst.$GoExpr(
                "GoCall",
                Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime"), "Func"),
                [
                  Gopurs$dGoAst.$GoExpr(
                    "GoFunc",
                    ["f"],
                    Gopurs$dGoAst.$GoExpr(
                      "GoReturn",
                      Gopurs$dGoAst.$GoExpr(
                        "GoCall",
                        Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime"), "Func"),
                        [
                          Gopurs$dGoAst.$GoExpr(
                            "GoFunc",
                            ["_"],
                            Gopurs$dGoAst.$GoExpr(
                              "GoBlock",
                              [
                                Gopurs$dGoAst.$GoExpr("GoVar", "resA := gopurs_runtime.Apply(a, gopurs_runtime.Value{})"),
                                Gopurs$dGoAst.$GoExpr("GoVar", "resB := gopurs_runtime.Apply(f, resA)"),
                                Gopurs$dGoAst.$GoExpr(
                                  "GoReturn",
                                  Gopurs$dGoAst.$GoExpr(
                                    "GoCall",
                                    Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime"), "Apply"),
                                    [Gopurs$dGoAst.$GoExpr("GoVar", "resB"), Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime.Value{}")]
                                  )
                                )
                              ]
                            )
                          )
                        ]
                      )
                    )
                  )
                ]
              )
            )
          )
        ]
      );
    }
    if (v._1._2 === "log") {
      return Gopurs$dGoAst.$GoExpr(
        "GoCall",
        Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime"), "Func"),
        [
          Gopurs$dGoAst.$GoExpr(
            "GoFunc",
            ["x"],
            Gopurs$dGoAst.$GoExpr(
              "GoReturn",
              Gopurs$dGoAst.$GoExpr(
                "GoCall",
                Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime"), "Func"),
                [
                  Gopurs$dGoAst.$GoExpr(
                    "GoFunc",
                    ["_"],
                    Gopurs$dGoAst.$GoExpr(
                      "GoBlock",
                      [
                        Gopurs$dGoAst.$GoExpr(
                          "GoCall",
                          Gopurs$dGoAst.$GoExpr("GoVar", "fmt.Println"),
                          [Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "x"), "StrVal")]
                        ),
                        Gopurs$dGoAst.$GoExpr("GoReturn", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime.Value{}"))
                      ]
                    )
                  )
                ]
              )
            )
          )
        ]
      );
    }
    if (v._1._2 === "showStringImpl") {
      return Gopurs$dGoAst.$GoExpr(
        "GoCall",
        Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime"), "Func"),
        [
          Gopurs$dGoAst.$GoExpr(
            "GoFunc",
            ["s"],
            Gopurs$dGoAst.$GoExpr(
              "GoReturn",
              Gopurs$dGoAst.$GoExpr(
                "GoCall",
                Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime"), "Str"),
                [
                  Gopurs$dGoAst.$GoExpr(
                    "GoCall",
                    Gopurs$dGoAst.$GoExpr("GoVar", "fmt.Sprintf"),
                    [Gopurs$dGoAst.$GoExpr("GoString", "%q"), Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "s"), "StrVal")]
                  )
                ]
              )
            )
          )
        ]
      );
    }
    return Gopurs$dGoAst.$GoExpr("GoVar", Data$dString$dCommon.replaceAll("$")("_")(v._1._2));
  }
  if (v.tag === "Local") {
    if (v._1.tag === "Just") { return Gopurs$dGoAst.$GoExpr("GoVar", Data$dString$dCommon.replaceAll("$")("_")(v._1._1)); }
    if (v._1.tag === "Nothing") { return Gopurs$dGoAst.$GoExpr("GoVar", "_"); }
    return Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime.Value{}");
  }
  if (v.tag === "Lit") {
    if (v._1.tag === "LitString") {
      return Gopurs$dGoAst.$GoExpr(
        "GoCall",
        Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime"), "Str"),
        [Gopurs$dGoAst.$GoExpr("GoString", v._1._1)]
      );
    }
    return Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime.Value{}");
  }
  if (v.tag === "App") {
    return Gopurs$dGoAst.$GoExpr(
      "GoCall",
      Gopurs$dGoAst.$GoExpr("GoSelector", Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime"), "Apply"),
      [
        translateExpr(v._1),
        translateExpr((() => {
          if (0 < v._2.length) { return v._2[0]; }
          $runtime.fail();
        })())
      ]
    );
  }
  return Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime.Value{}");
};
const translateBinding = v => {
  const safeName = Data$dString$dCommon.replaceAll("$")("_")(v._1);
  return Data$dMaybe.$Maybe("Just", {identifier: safeName === "main" ? "Main" : safeName, expression: translateExpr(v._2)});
};
const translateBindingGroup = bg => Data$dArray.mapMaybe(translateBinding)(bg.bindings);
const translate = v => backendMod => Gopurs$dPrinter.printGoFile({
  packageName: Data$dString$dCommon.replaceAll(".")("_")(backendMod.name),
  imports: ["gopurs/output/gopurs_runtime", "fmt"],
  decls: [
    ...Data$dFunctor.arrayMap(f => ({identifier: f, expression: Gopurs$dGoAst.$GoExpr("GoVar", "gopurs_runtime.Value{}")}))(Data$dArray.fromFoldableImpl(
      Data$dSet.foldableSet.foldr,
      backendMod.foreign
    )),
    ...Control$dBind.arrayBind(Data$dArray.fromFoldableImpl(Data$dFoldable.foldrArray, backendMod.bindings))(translateBindingGroup)
  ]
});
export {translate, translateBinding, translateBindingGroup, translateExpr};
