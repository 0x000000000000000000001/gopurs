import * as $runtime from "../runtime.js";
import * as Control$dBind from "../Control.Bind/index.js";
import * as Data$dArray from "../Data.Array/index.js";
import * as Data$dEither from "../Data.Either/index.js";
import * as Data$dFoldable from "../Data.Foldable/index.js";
import * as Data$dFunctor from "../Data.Functor/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
import * as Data$dMap$dInternal from "../Data.Map.Internal/index.js";
import * as Data$dMaybe from "../Data.Maybe/index.js";
import * as Data$dOrd from "../Data.Ord/index.js";
import * as Data$dSemiring from "../Data.Semiring/index.js";
import * as Data$dTuple from "../Data.Tuple/index.js";
import * as Data$dUnfoldable from "../Data.Unfoldable/index.js";
const $ModuleSort = (tag, _1) => ({tag, _1});
const toUnfoldable = /* #__PURE__ */ (() => {
  const $0 = Data$dUnfoldable.unfoldableArray.unfoldr(Data$dMap$dInternal.stepUnfoldr);
  return x => $0(Data$dMap$dInternal.$MapIter("IterNode", x, Data$dMap$dInternal.IterLeaf));
})();
const toUnfoldable1 = /* #__PURE__ */ (() => {
  const $0 = Data$dUnfoldable.unfoldableArray.unfoldr(xs => {
    if (xs.tag === "Nil") { return Data$dMaybe.Nothing; }
    if (xs.tag === "Cons") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(xs._1, xs._2)); }
    $runtime.fail();
  });
  return x => $0((() => {
    const go = (m$p, z$p) => {
      if (m$p.tag === "Leaf") { return z$p; }
      if (m$p.tag === "Node") { return go(m$p._5, Data$dList$dTypes.$List("Cons", m$p._3, go(m$p._6, z$p))); }
      $runtime.fail();
    };
    return go(x, Data$dList$dTypes.Nil);
  })());
})();
const fromFoldable = /* #__PURE__ */ Data$dFoldable.foldlArray(m => a => Data$dMap$dInternal.insert(Data$dOrd.ordString)(a)()(m))(Data$dMap$dInternal.Leaf);
const fromFoldable1 = /* #__PURE__ */ Data$dMap$dInternal.fromFoldable(Data$dOrd.ordString)(Data$dFoldable.foldableArray);
const lookup = k => {
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      if (v.tag === "Leaf") {
        go$c = false;
        go$r = Data$dMaybe.Nothing;
        continue;
      }
      if (v.tag === "Node") {
        const v1 = Data$dOrd.ordString.compare(k)(v._3);
        if (v1 === "LT") {
          go$a0 = v._5;
          continue;
        }
        if (v1 === "GT") {
          go$a0 = v._6;
          continue;
        }
        if (v1 === "EQ") {
          go$c = false;
          go$r = Data$dMaybe.$Maybe("Just", v._4);
          continue;
        }
      }
      $runtime.fail();
    }
    return go$r;
  };
  return go;
};
const toUnfoldable2 = /* #__PURE__ */ (() => Data$dUnfoldable.unfoldableArray.unfoldr(xs => {
  if (xs.tag === "Nil") { return Data$dMaybe.Nothing; }
  if (xs.tag === "Cons") { return Data$dMaybe.$Maybe("Just", Data$dTuple.$Tuple(xs._1, xs._2)); }
  $runtime.fail();
}))();
const Sorted = value0 => $ModuleSort("Sorted", value0);
const CycleDetected = value0 => $ModuleSort("CycleDetected", value0);
const topoSort = dictOrd => graph => {
  const importCounts = Data$dMap$dInternal.fromFoldableWith(dictOrd)(Data$dFoldable.foldableArray)(Data$dSemiring.intAdd)(Control$dBind.arrayBind(toUnfoldable(graph))(v => [
    Data$dTuple.$Tuple(v._1, 0),
    ...Data$dFunctor.arrayMap(a => Data$dTuple.$Tuple(a, 1))(toUnfoldable1(v._2))
  ]));
  const depthFirst = v => {
    const $0 = v.curr;
    if (
      (() => {
        const go = go$a0$copy => {
          let go$a0 = go$a0$copy, go$c = true, go$r;
          while (go$c) {
            const v$1 = go$a0;
            if (v$1.tag === "Leaf") {
              go$c = false;
              go$r = false;
              continue;
            }
            if (v$1.tag === "Node") {
              const v1 = dictOrd.compare($0)(v$1._3);
              if (v1 === "LT") {
                go$a0 = v$1._5;
                continue;
              }
              if (v1 === "GT") {
                go$a0 = v$1._6;
                continue;
              }
              if (v1 === "EQ") {
                go$c = false;
                go$r = true;
                continue;
              }
            }
            $runtime.fail();
          }
          return go$r;
        };
        return go(v.visited);
      })()
    ) {
      return Data$dMaybe.$Maybe("Just", Data$dList$dTypes.$List("Cons", $0, v.path));
    }
    if (
      (() => {
        const go = go$a0$copy => {
          let go$a0 = go$a0$copy, go$c = true, go$r;
          while (go$c) {
            const v$1 = go$a0;
            if (v$1.tag === "Leaf") {
              go$c = false;
              go$r = Data$dMaybe.Nothing;
              continue;
            }
            if (v$1.tag === "Node") {
              const v1 = dictOrd.compare($0)(v$1._3);
              if (v1 === "LT") {
                go$a0 = v$1._5;
                continue;
              }
              if (v1 === "GT") {
                go$a0 = v$1._6;
                continue;
              }
              if (v1 === "EQ") {
                go$c = false;
                go$r = Data$dMaybe.$Maybe("Just", v$1._4);
                continue;
              }
            }
            $runtime.fail();
          }
          return go$r;
        };
        const $1 = go(graph);
        if ($1.tag === "Nothing") { return true; }
        if ($1.tag === "Just") { return $1._1.tag === "Leaf"; }
        $runtime.fail();
      })()
    ) {
      return Data$dMaybe.Nothing;
    }
    const go = go$a0$copy => {
      let go$a0 = go$a0$copy, go$c = true, go$r;
      while (go$c) {
        const v$1 = go$a0;
        if (v$1.tag === "Leaf") {
          go$c = false;
          go$r = Data$dMaybe.Nothing;
          continue;
        }
        if (v$1.tag === "Node") {
          const v1 = dictOrd.compare($0)(v$1._3);
          if (v1 === "LT") {
            go$a0 = v$1._5;
            continue;
          }
          if (v1 === "GT") {
            go$a0 = v$1._6;
            continue;
          }
          if (v1 === "EQ") {
            go$c = false;
            go$r = Data$dMaybe.$Maybe("Just", v$1._4);
            continue;
          }
        }
        $runtime.fail();
      }
      return go$r;
    };
    const $1 = go(graph);
    if ($1.tag === "Just") {
      const go$1 = go$1$a0$copy => go$1$a1$copy => {
        let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
        while (go$1$c) {
          const b = go$1$a0, v$1 = go$1$a1;
          if (v$1.tag === "Nil") {
            go$1$c = false;
            go$1$r = b;
            continue;
          }
          if (v$1.tag === "Cons") {
            go$1$a0 = (() => {
              if (b.tag === "Nothing") { return false; }
              if (b.tag === "Just") { return true; }
              $runtime.fail();
            })()
              ? b
              : depthFirst({path: Data$dList$dTypes.$List("Cons", $0, v.path), visited: Data$dMap$dInternal.insert(dictOrd)($0)()(v.visited), curr: v$1._1});
            go$1$a1 = v$1._2;
            continue;
          }
          $runtime.fail();
        }
        return go$1$r;
      };
      return go$1(Data$dMaybe.Nothing)((() => {
        const go$2 = (m$p, z$p) => {
          if (m$p.tag === "Leaf") { return z$p; }
          if (m$p.tag === "Node") { return go$2(m$p._5, Data$dList$dTypes.$List("Cons", m$p._3, go$2(m$p._6, z$p))); }
          $runtime.fail();
        };
        return go$2($1._1, Data$dList$dTypes.Nil);
      })());
    }
    if ($1.tag === "Nothing") { return Data$dMaybe.Nothing; }
    $runtime.fail();
  };
  const go = go$a0$copy => {
    let go$a0 = go$a0$copy, go$c = true, go$r;
    while (go$c) {
      const v = go$a0;
      const $0 = Data$dMap$dInternal.findMin(v.roots);
      const v1 = $0.tag === "Just" ? Data$dMaybe.$Maybe("Just", $0._1.key) : Data$dMaybe.Nothing;
      if (v1.tag === "Nothing") {
        if (
          (() => {
            const go$1 = v$1 => {
              if (v$1.tag === "Leaf") { return true; }
              if (v$1.tag === "Node") { return go$1(v$1._5) && 0 === v$1._4 && go$1(v$1._6); }
              $runtime.fail();
            };
            return go$1(v.usages);
          })()
        ) {
          go$c = false;
          go$r = Data$dEither.$Either("Right", {roots: v.roots, sorted: v.sorted, usages: v.usages});
          continue;
        }
        const go$1 = go$1$a0$copy => go$1$a1$copy => {
          let go$1$a0 = go$1$a0$copy, go$1$a1 = go$1$a1$copy, go$1$c = true, go$1$r;
          while (go$1$c) {
            const b = go$1$a0, v$1 = go$1$a1;
            if (v$1.tag === "Nil") {
              go$1$c = false;
              go$1$r = b;
              continue;
            }
            if (v$1.tag === "Cons") {
              go$1$a0 = (() => {
                if (b.tag === "Nothing") { return false; }
                if (b.tag === "Just") { return true; }
                $runtime.fail();
              })()
                ? b
                : depthFirst({path: Data$dList$dTypes.Nil, visited: Data$dMap$dInternal.Leaf, curr: v$1._1});
              go$1$a1 = v$1._2;
              continue;
            }
            $runtime.fail();
          }
          return go$1$r;
        };
        const go$2 = v$1 => {
          if (v$1.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
          if (v$1.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v$1._1, v$1._2, v$1._3, undefined, go$2(v$1._5), go$2(v$1._6)); }
          $runtime.fail();
        };
        const detectCycles = go$1(Data$dMaybe.Nothing)((() => {
          const go$3 = (m$p, z$p) => {
            if (m$p.tag === "Leaf") { return z$p; }
            if (m$p.tag === "Node") { return go$3(m$p._5, Data$dList$dTypes.$List("Cons", m$p._3, go$3(m$p._6, z$p))); }
            $runtime.fail();
          };
          return go$3(
            go$2(Data$dMap$dInternal.filterWithKey(dictOrd)(a => count => {
              const go$4 = go$4$a0$copy => {
                let go$4$a0 = go$4$a0$copy, go$4$c = true, go$4$r;
                while (go$4$c) {
                  const v$1 = go$4$a0;
                  if (v$1.tag === "Leaf") {
                    go$4$c = false;
                    go$4$r = Data$dMaybe.Nothing;
                    continue;
                  }
                  if (v$1.tag === "Node") {
                    const v1$1 = dictOrd.compare(a)(v$1._3);
                    if (v1$1 === "LT") {
                      go$4$a0 = v$1._5;
                      continue;
                    }
                    if (v1$1 === "GT") {
                      go$4$a0 = v$1._6;
                      continue;
                    }
                    if (v1$1 === "EQ") {
                      go$4$c = false;
                      go$4$r = Data$dMaybe.$Maybe("Just", v$1._4);
                      continue;
                    }
                  }
                  $runtime.fail();
                }
                return go$4$r;
              };
              const $1 = go$4(graph);
              return count > 0 && (() => {
                if ($1.tag === "Nothing") { return false; }
                if ($1.tag === "Just") { return $1._1.tag !== "Leaf"; }
                $runtime.fail();
              })();
            })(v.usages)),
            Data$dList$dTypes.Nil
          );
        })());
        if (detectCycles.tag === "Just") {
          go$c = false;
          go$r = Data$dEither.$Either("Left", detectCycles._1);
          continue;
        }
        if (detectCycles.tag === "Nothing") {
          go$c = false;
          go$r = Data$dEither.$Either("Left", Data$dList$dTypes.Nil);
          continue;
        }
        $runtime.fail();
      }
      if (v1.tag === "Just") {
        const $1 = v1._1;
        const go$1 = go$1$a0$copy => {
          let go$1$a0 = go$1$a0$copy, go$1$c = true, go$1$r;
          while (go$1$c) {
            const v$1 = go$1$a0;
            if (v$1.tag === "Leaf") {
              go$1$c = false;
              go$1$r = Data$dMaybe.Nothing;
              continue;
            }
            if (v$1.tag === "Node") {
              const v1$1 = dictOrd.compare($1)(v$1._3);
              if (v1$1 === "LT") {
                go$1$a0 = v$1._5;
                continue;
              }
              if (v1$1 === "GT") {
                go$1$a0 = v$1._6;
                continue;
              }
              if (v1$1 === "EQ") {
                go$1$c = false;
                go$1$r = Data$dMaybe.$Maybe("Just", v$1._4);
                continue;
              }
            }
            $runtime.fail();
          }
          return go$1$r;
        };
        const $2 = go$1(graph);
        const reachable = (() => {
          if ($2.tag === "Nothing") { return Data$dMap$dInternal.Leaf; }
          if ($2.tag === "Just") { return $2._1; }
          $runtime.fail();
        })();
        const go$2 = go$2$a0$copy => go$2$a1$copy => {
          let go$2$a0 = go$2$a0$copy, go$2$a1 = go$2$a1$copy, go$2$c = true, go$2$r;
          while (go$2$c) {
            const b = go$2$a0, v$1 = go$2$a1;
            if (v$1.tag === "Nil") {
              go$2$c = false;
              go$2$r = b;
              continue;
            }
            if (v$1.tag === "Cons") {
              go$2$a0 = Data$dMap$dInternal.insertWith(dictOrd)(Data$dSemiring.intAdd)(v$1._1)(-1)(b);
              go$2$a1 = v$1._2;
              continue;
            }
            $runtime.fail();
          }
          return go$2$r;
        };
        const usages$p = go$2(v.usages)((() => {
          const go$3 = (m$p, z$p) => {
            if (m$p.tag === "Leaf") { return z$p; }
            if (m$p.tag === "Node") { return go$3(m$p._5, Data$dList$dTypes.$List("Cons", m$p._3, go$3(m$p._6, z$p))); }
            $runtime.fail();
          };
          return go$3(reachable, Data$dList$dTypes.Nil);
        })());
        go$a0 = {
          roots: (() => {
            const go$3 = go$3$a0$copy => go$3$a1$copy => {
              let go$3$a0 = go$3$a0$copy, go$3$a1 = go$3$a1$copy, go$3$c = true, go$3$r;
              while (go$3$c) {
                const b = go$3$a0, v$1 = go$3$a1;
                if (v$1.tag === "Nil") {
                  go$3$c = false;
                  go$3$r = b;
                  continue;
                }
                if (v$1.tag === "Cons") {
                  go$3$a0 = (() => {
                    const $3 = v$1._1;
                    const go$4 = go$4$a0$copy => {
                      let go$4$a0 = go$4$a0$copy, go$4$c = true, go$4$r;
                      while (go$4$c) {
                        const v$2 = go$4$a0;
                        if (v$2.tag === "Leaf") {
                          go$4$c = false;
                          go$4$r = Data$dMaybe.Nothing;
                          continue;
                        }
                        if (v$2.tag === "Node") {
                          const v1$1 = dictOrd.compare($3)(v$2._3);
                          if (v1$1 === "LT") {
                            go$4$a0 = v$2._5;
                            continue;
                          }
                          if (v1$1 === "GT") {
                            go$4$a0 = v$2._6;
                            continue;
                          }
                          if (v1$1 === "EQ") {
                            go$4$c = false;
                            go$4$r = Data$dMaybe.$Maybe("Just", v$2._4);
                            continue;
                          }
                        }
                        $runtime.fail();
                      }
                      return go$4$r;
                    };
                    const $4 = go$4(usages$p);
                    if ($4.tag === "Just") {
                      if ($4._1 === 0) { return Data$dMap$dInternal.insert(dictOrd)($3)()(b); }
                      return b;
                    }
                    if ($4.tag === "Nothing") { return b; }
                    $runtime.fail();
                  })();
                  go$3$a1 = v$1._2;
                  continue;
                }
                $runtime.fail();
              }
              return go$3$r;
            };
            return go$3(Data$dMap$dInternal.delete(dictOrd)($1)(v.roots))((() => {
              const go$4 = (m$p, z$p) => {
                if (m$p.tag === "Leaf") { return z$p; }
                if (m$p.tag === "Node") { return go$4(m$p._5, Data$dList$dTypes.$List("Cons", m$p._3, go$4(m$p._6, z$p))); }
                $runtime.fail();
              };
              return go$4(reachable, Data$dList$dTypes.Nil);
            })());
          })(),
          sorted: Data$dList$dTypes.$List("Cons", $1, v.sorted),
          usages: usages$p
        };
        continue;
      }
      $runtime.fail();
    }
    return go$r;
  };
  const $0 = go({
    roots: (() => {
      const go$1 = v => {
        if (v.tag === "Leaf") { return Data$dMap$dInternal.Leaf; }
        if (v.tag === "Node") { return Data$dMap$dInternal.$$$Map("Node", v._1, v._2, v._3, undefined, go$1(v._5), go$1(v._6)); }
        $runtime.fail();
      };
      return go$1(Data$dMap$dInternal.filterWithKey(dictOrd)(k => v => v === 0)(importCounts));
    })(),
    sorted: Data$dList$dTypes.Nil,
    usages: importCounts
  });
  if ($0.tag === "Left") { return Data$dEither.$Either("Left", $0._1); }
  if ($0.tag === "Right") { return Data$dEither.$Either("Right", $0._1.sorted); }
  $runtime.fail();
};
const moduleGraph = k => {
  const $0 = Data$dFunctor.arrayMap(x => {
    const $0 = k(x);
    return Data$dTuple.$Tuple($0.name.name, fromFoldable(Data$dFunctor.arrayMap(v => v.module.name)($0.imports)));
  });
  return x => fromFoldable1($0(x));
};
const sortModules = k => moduleHeaders => {
  const knownModuleHeaders = fromFoldable1(Data$dFunctor.arrayMap(a => Data$dTuple.$Tuple(k(a).name.name, a))(moduleHeaders));
  const v = topoSort(Data$dOrd.ordString)(moduleGraph(k)(moduleHeaders));
  if (v.tag === "Left") { return $ModuleSort("CycleDetected", Data$dArray.mapMaybe(a => lookup(a)(knownModuleHeaders))(toUnfoldable2(v._1))); }
  if (v.tag === "Right") { return $ModuleSort("Sorted", Data$dArray.mapMaybe(a => lookup(a)(knownModuleHeaders))(toUnfoldable2(v._1))); }
  $runtime.fail();
};
export {$ModuleSort, CycleDetected, Sorted, fromFoldable, fromFoldable1, lookup, moduleGraph, sortModules, toUnfoldable, toUnfoldable1, toUnfoldable2, topoSort};
