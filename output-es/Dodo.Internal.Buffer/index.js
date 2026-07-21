import * as Data$dFunction from "../Data.Function/index.js";
import * as Data$dList from "../Data.List/index.js";
import * as Data$dList$dTypes from "../Data.List.Types/index.js";
const Buffer = x => x;
const revert = v => ({...v, queue: Data$dList.drop(1)(v.queue)});
const $$new = buffer => ({buffer, queue: Data$dList$dTypes.Nil});
const modify = f => v => {
  if (v.queue.tag === "Cons") { return {...v, queue: Data$dList$dTypes.$List("Cons", Data$dList$dTypes.$List("Cons", f, v.queue._1), v.queue._2)}; }
  return {...v, buffer: f(v.buffer)};
};
const isBranching = v => v.queue.tag !== "Nil";
const commit = v => (
  {buffer: Data$dList$dTypes.foldableList.foldr(b => a => Data$dList$dTypes.foldableList.foldr(Data$dFunction.apply)(a)(b))(v.buffer)(v.queue), queue: Data$dList$dTypes.Nil}
);
const $$get = x => commit(x).buffer;
const branch = v => ({...v, queue: Data$dList$dTypes.$List("Cons", Data$dList$dTypes.Nil, v.queue)});
export {Buffer, branch, commit, $$get as get, isBranching, modify, $$new as new, revert};
