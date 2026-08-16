package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Done"))
	})
	return cache_Main_main
}

var cache_Main_foo gopurs_runtime.Value
var once_Main_foo sync.Once

func Get_Main_foo() gopurs_runtime.Value {
	once_Main_foo.Do(func() {
		cache_Main_foo = func() gopurs_runtime.Value {
			arr := func() []int64 {
				arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Bind_arrayBind(), func() gopurs_runtime.Value {
						arr := [][][]int64{[][]int64{[]int64{1, 2, 3}, []int64{4, 5}}, [][]int64{[]int64{6}}}
						boxed := make([]gopurs_runtime.Value, len(arr))
						for i, v := range arr {
							boxed[i] = func() gopurs_runtime.Value {
								arr := v
								boxed := make([]gopurs_runtime.Value, len(arr))
								for i, v := range arr {
									boxed[i] = func() gopurs_runtime.Value {
										arr := v
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = gopurs_runtime.Int(v)
										}
										return gopurs_runtime.Array(boxed)
									}()
								}
								return gopurs_runtime.Array(boxed)
							}()
						}
						return gopurs_runtime.Array(boxed)
					}(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
						return func() gopurs_runtime.Value {
							arr := func() []int64 {
								arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
									arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Bind_arrayBind(), func() gopurs_runtime.Value {
										arr := func() [][]int64 {
											arr := *(*[]gopurs_runtime.Value)(v_0.UnsafePtr)
											unboxed := make([][]int64, len(arr))
											for i, v := range arr {
												unboxed[i] = func() []int64 {
													arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
													unboxed := make([]int64, len(arr))
													for i, v := range arr {
														unboxed[i] = v.IntVal
													}
													return unboxed
												}()
											}
											return unboxed
										}()
										boxed := make([]gopurs_runtime.Value, len(arr))
										for i, v := range arr {
											boxed[i] = func() gopurs_runtime.Value {
												arr := v
												boxed := make([]gopurs_runtime.Value, len(arr))
												for i, v := range arr {
													boxed[i] = gopurs_runtime.Int(v)
												}
												return gopurs_runtime.Array(boxed)
											}()
										}
										return gopurs_runtime.Array(boxed)
									}(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
										return func() gopurs_runtime.Value {
											arr := func() []int64 {
												arr := *(*[]gopurs_runtime.Value)(v1_1.UnsafePtr)
												unboxed := make([]int64, len(arr))
												for i, v := range arr {
													unboxed[i] = v.IntVal
												}
												return unboxed
											}()
											boxed := make([]gopurs_runtime.Value, len(arr))
											for i, v := range arr {
												boxed[i] = gopurs_runtime.Int(v)
											}
											return gopurs_runtime.Array(boxed)
										}()
									})).UnsafePtr)
									unboxed := make([]gopurs_runtime.Value, len(arr))
									for i, v := range arr {
										unboxed[i] = v
									}
									return unboxed
								}()).UnsafePtr)
								unboxed := make([]int64, len(arr))
								for i, v := range arr {
									unboxed[i] = v.IntVal
								}
								return unboxed
							}()
							boxed := make([]gopurs_runtime.Value, len(arr))
							for i, v := range arr {
								boxed[i] = gopurs_runtime.Int(v)
							}
							return gopurs_runtime.Array(boxed)
						}()
					})).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr {
						unboxed[i] = v
					}
					return unboxed
				}()).UnsafePtr)
				unboxed := make([]int64, len(arr))
				for i, v := range arr {
					unboxed[i] = v.IntVal
				}
				return unboxed
			}()
			boxed := make([]gopurs_runtime.Value, len(arr))
			for i, v := range arr {
				boxed[i] = gopurs_runtime.Int(v)
			}
			return gopurs_runtime.Array(boxed)
		}()
	})
	return cache_Main_foo
}
