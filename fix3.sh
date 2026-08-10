#!/bin/bash
git checkout src/Gopurs/Runtime.purs

sed -i '' 's/)/)\
\
var EscapeSink any/g' src/Gopurs/Runtime.purs

sed -i '' 's/return Value{Type: TypeFunc, UnsafePtr: \*(\*unsafe.Pointer)(unsafe.Pointer(&f))}/EscapeSink = f; return Value{Type: TypeFunc, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(\&f))}/g' src/Gopurs/Runtime.purs

for i in {2..11}; do
  sed -i '' "s/return Value{Type: TypeFunc${i}, UnsafePtr: \*(\*unsafe.Pointer)(unsafe.Pointer(&f))}/EscapeSink = f; return Value{Type: TypeFunc${i}, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(\&f))}/g" src/Gopurs/Runtime.purs
done
