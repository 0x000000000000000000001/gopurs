package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Main_R gopurs_runtime.Value
var once_Main_R sync.Once

func Get_Main_R() gopurs_runtime.Value {
	once_Main_R.Do(func() {
		cache_Main_R = gopurs_runtime.Value{Type: 9, IntVal: int64(3558538316), UnsafePtr: nil}
	})
	return cache_Main_R
}

var cache_Main_B gopurs_runtime.Value
var once_Main_B sync.Once

func Get_Main_B() gopurs_runtime.Value {
	once_Main_B.Do(func() {
		cache_Main_B = gopurs_runtime.Value{Type: 9, IntVal: int64(4250879068), UnsafePtr: nil}
	})
	return cache_Main_B
}

var cache_Main_E gopurs_runtime.Value
var once_Main_E sync.Once

func Get_Main_E() gopurs_runtime.Value {
	once_Main_E.Do(func() {
		cache_Main_E = gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((*Constructor_Main_T)(nil))}
	})
	return cache_Main_E
}

var cache_Main_T gopurs_runtime.Value
var once_Main_T sync.Once

func Get_Main_T() gopurs_runtime.Value {
	once_Main_T.Do(func() {
		cache_Main_T = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(value3 gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((&Constructor_Main_T{1, uint32(value0.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Main_T](value1), value2.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_T](value3)}))}
					})
				})
			})
		})
	})
	return cache_Main_T
}

var cache_Main_max gopurs_runtime.Value
var once_Main_max sync.Once

func Get_Main_max() gopurs_runtime.Value {
	once_Main_max.Do(func() {
		cache_Main_max = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_max(x_0_box.IntVal, y_1_box.IntVal))
		})
	})
	return cache_Main_max
}

var cache_Main_describe gopurs_runtime.Value
var once_Main_describe sync.Once

func Get_Main_describe() gopurs_runtime.Value {
	once_Main_describe.Do(func() {
		cache_Main_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Red-Black Tree (100k Worst-Case Insertions):"))
	})
	return cache_Main_describe
}

var cache_Main_depth gopurs_runtime.Value
var once_Main_depth sync.Once

func Get_Main_depth() gopurs_runtime.Value {
	once_Main_depth.Do(func() {
		cache_Main_depth = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Main_depth(gopurs_runtime.CoerceToStruct[Constructor_Main_T](v_0_box)))
		})
	})
	return cache_Main_depth
}

var cache_Main_balance gopurs_runtime.Value
var once_Main_balance sync.Once

func Get_Main_balance() gopurs_runtime.Value {
	once_Main_balance.Do(func() {
		cache_Main_balance = gopurs_runtime.Func4(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value, v3_3_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(Call_Main_balance(uint32(v_0_box.IntVal), gopurs_runtime.CoerceToStruct[Constructor_Main_T](v1_1_box), v2_2_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_T](v3_3_box)))}
		})
	})
	return cache_Main_balance
}

var cache_Main_insert gopurs_runtime.Value
var once_Main_insert sync.Once

func Get_Main_insert() gopurs_runtime.Value {
	once_Main_insert.Do(func() {
		cache_Main_insert = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(Call_Main_insert(x_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_T](s_1_box)))}
		})
	})
	return cache_Main_insert
}

var cache_Main_buildTree gopurs_runtime.Value
var once_Main_buildTree sync.Once

func Get_Main_buildTree() gopurs_runtime.Value {
	once_Main_buildTree.Do(func() {
		cache_Main_buildTree = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(Call_Main_buildTree(v_0_box.IntVal, gopurs_runtime.CoerceToStruct[Constructor_Main_T](v1_1_box)))}
		})
	})
	return cache_Main_buildTree
}

var cache_Main_act gopurs_runtime.Value
var once_Main_act sync.Once

func Get_Main_act() gopurs_runtime.Value {
	once_Main_act.Do(func() {
		cache_Main_act = Call_Effect_Console_logShow(Rebox_Main_1636311157_1386611502(Rebox_Main_1386611502_1636311157(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[gopurs_runtime.Value]](Get_Data_Show_showInt()))), gopurs_runtime.Int(Call_Main_depth(Call_Main_buildTree(int64(100000), (*Constructor_Main_T)(nil)))))
	})
	return cache_Main_act
}

var cache_Main_main gopurs_runtime.Value
var once_Main_main sync.Once

func Get_Main_main() gopurs_runtime.Value {
	once_Main_main.Do(func() {
		cache_Main_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_Main_describe(), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return Get_Main_act()
		}))
	})
	return cache_Main_main
}

func Call_Main___gopurs_owned_balance_0_consume(__arg0 uint32, __arg1 *Constructor_Main_T, __arg2 int64, __arg3 *Constructor_Main_T, __donor *Constructor_Main_T) *Constructor_Main_T {
__owned_loop:
	for {
		if false {
			continue __owned_loop
		}
		if (__arg0) == (4250879068) {
			if (__arg1) != (nil) {
				if (__arg1.V0) == (3558538316) {
					if (__arg1.V1) != (nil) {
						if (__arg1.V1.V0) == (3558538316) {
							__let_scalar_40 := int64(__arg1.V1.V2)
							_ = __let_scalar_40
							__let_scalar_41 := int64(__arg1.V2)
							_ = __let_scalar_41
							__let_scalar_42 := int64(__arg2)
							_ = __let_scalar_42
							__scalar_43 := uint32(3558538316)
							_ = __scalar_43
							__scalar_44 := uint32(4250879068)
							_ = __scalar_44
							__read_45 := __arg1.V1.V1
							_ = __read_45
							__scalar_46 := int64(__let_scalar_40)
							_ = __scalar_46
							__read_47 := __arg1.V1.V3
							_ = __read_47
							__scalar_48 := int64(__let_scalar_41)
							_ = __scalar_48
							__scalar_49 := uint32(4250879068)
							_ = __scalar_49
							__read_50 := __arg1.V3
							_ = __read_50
							__scalar_51 := int64(__let_scalar_42)
							_ = __scalar_51
							__read_52 := __arg3
							_ = __read_52
							__donor_slot_55 := __donor
							_ = __donor_slot_55
							__dead_53 := __arg1
							_ = __dead_53
							__dead_54 := __arg1.V1
							_ = __dead_54
							__dead_53.Rc = 1
							__dead_53.V0 = __scalar_44
							__dead_53.V1 = __read_45
							__dead_53.V2 = __scalar_46
							__dead_53.V3 = __read_47
							__dead_54.Rc = 1
							__dead_54.V0 = __scalar_49
							__dead_54.V1 = __read_50
							__dead_54.V2 = __scalar_51
							__dead_54.V3 = __read_52
							var __cell_56 *Constructor_Main_T
							if (__donor_slot_55) != (nil) {
								__cell_56 = __donor_slot_55
								__donor_slot_55 = nil
							} else {

							}
							if (__cell_56) == (nil) {
								__cell_56 = new(Constructor_Main_T)
							} else {

							}
							__cell_56.Rc = 1
							__cell_56.V0 = __scalar_43
							__cell_56.V1 = __dead_53
							__cell_56.V2 = __scalar_48
							__cell_56.V3 = __dead_54
							return __cell_56
						} else {
							if (__arg1.V3) != (nil) {
								if (__arg1.V3.V0) == (3558538316) {
									__let_scalar_65 := int64(__arg1.V2)
									_ = __let_scalar_65
									__let_scalar_66 := int64(__arg1.V3.V2)
									_ = __let_scalar_66
									__let_scalar_67 := int64(__arg2)
									_ = __let_scalar_67
									__scalar_68 := uint32(3558538316)
									_ = __scalar_68
									__scalar_69 := uint32(4250879068)
									_ = __scalar_69
									__read_70 := __arg1.V1
									_ = __read_70
									__scalar_71 := int64(__let_scalar_65)
									_ = __scalar_71
									__read_72 := __arg1.V3.V1
									_ = __read_72
									__scalar_73 := int64(__let_scalar_66)
									_ = __scalar_73
									__scalar_74 := uint32(4250879068)
									_ = __scalar_74
									__read_75 := __arg1.V3.V3
									_ = __read_75
									__scalar_76 := int64(__let_scalar_67)
									_ = __scalar_76
									__read_77 := __arg3
									_ = __read_77
									__donor_slot_80 := __donor
									_ = __donor_slot_80
									__dead_78 := __arg1
									_ = __dead_78
									__dead_79 := __arg1.V3
									_ = __dead_79
									__dead_78.Rc = 1
									__dead_78.V0 = __scalar_69
									__dead_78.V1 = __read_70
									__dead_78.V2 = __scalar_71
									__dead_78.V3 = __read_72
									__dead_79.Rc = 1
									__dead_79.V0 = __scalar_74
									__dead_79.V1 = __read_75
									__dead_79.V2 = __scalar_76
									__dead_79.V3 = __read_77
									var __cell_81 *Constructor_Main_T
									if (__donor_slot_80) != (nil) {
										__cell_81 = __donor_slot_80
										__donor_slot_80 = nil
									} else {

									}
									if (__cell_81) == (nil) {
										__cell_81 = new(Constructor_Main_T)
									} else {

									}
									__cell_81.Rc = 1
									__cell_81.V0 = __scalar_68
									__cell_81.V1 = __dead_78
									__cell_81.V2 = __scalar_73
									__cell_81.V3 = __dead_79
									return __cell_81
								} else {
									if ((__arg3) != (nil)) && ((__arg3.V0) == (3558538316)) {
										if (__arg3.V1) != (nil) {
											if (__arg3.V1.V0) == (3558538316) {
												__let_scalar_98 := int64(__arg2)
												_ = __let_scalar_98
												__let_scalar_99 := int64(__arg3.V1.V2)
												_ = __let_scalar_99
												__let_scalar_100 := int64(__arg3.V2)
												_ = __let_scalar_100
												__scalar_101 := uint32(3558538316)
												_ = __scalar_101
												__scalar_102 := uint32(4250879068)
												_ = __scalar_102
												__read_103 := __arg1
												_ = __read_103
												__scalar_104 := int64(__let_scalar_98)
												_ = __scalar_104
												__read_105 := __arg3.V1.V1
												_ = __read_105
												__scalar_106 := int64(__let_scalar_99)
												_ = __scalar_106
												__scalar_107 := uint32(4250879068)
												_ = __scalar_107
												__read_108 := __arg3.V1.V3
												_ = __read_108
												__scalar_109 := int64(__let_scalar_100)
												_ = __scalar_109
												__read_110 := __arg3.V3
												_ = __read_110
												__donor_slot_113 := __donor
												_ = __donor_slot_113
												__dead_111 := __arg3
												_ = __dead_111
												__dead_112 := __arg3.V1
												_ = __dead_112
												__dead_111.Rc = 1
												__dead_111.V0 = __scalar_102
												__dead_111.V1 = __read_103
												__dead_111.V2 = __scalar_104
												__dead_111.V3 = __read_105
												__dead_112.Rc = 1
												__dead_112.V0 = __scalar_107
												__dead_112.V1 = __read_108
												__dead_112.V2 = __scalar_109
												__dead_112.V3 = __read_110
												var __cell_114 *Constructor_Main_T
												if (__donor_slot_113) != (nil) {
													__cell_114 = __donor_slot_113
													__donor_slot_113 = nil
												} else {

												}
												if (__cell_114) == (nil) {
													__cell_114 = new(Constructor_Main_T)
												} else {

												}
												__cell_114.Rc = 1
												__cell_114.V0 = __scalar_101
												__cell_114.V1 = __dead_111
												__cell_114.V2 = __scalar_106
												__cell_114.V3 = __dead_112
												return __cell_114
											} else {
												if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
													__let_scalar_115 := int64(__arg2)
													_ = __let_scalar_115
													__let_scalar_116 := int64(__arg3.V2)
													_ = __let_scalar_116
													__let_scalar_117 := int64(__arg3.V3.V2)
													_ = __let_scalar_117
													__scalar_118 := uint32(3558538316)
													_ = __scalar_118
													__scalar_119 := uint32(4250879068)
													_ = __scalar_119
													__read_120 := __arg1
													_ = __read_120
													__scalar_121 := int64(__let_scalar_115)
													_ = __scalar_121
													__read_122 := __arg3.V1
													_ = __read_122
													__scalar_123 := int64(__let_scalar_116)
													_ = __scalar_123
													__scalar_124 := uint32(4250879068)
													_ = __scalar_124
													__read_125 := __arg3.V3.V1
													_ = __read_125
													__scalar_126 := int64(__let_scalar_117)
													_ = __scalar_126
													__read_127 := __arg3.V3.V3
													_ = __read_127
													__donor_slot_130 := __donor
													_ = __donor_slot_130
													__dead_128 := __arg3
													_ = __dead_128
													__dead_129 := __arg3.V3
													_ = __dead_129
													__dead_128.Rc = 1
													__dead_128.V0 = __scalar_119
													__dead_128.V1 = __read_120
													__dead_128.V2 = __scalar_121
													__dead_128.V3 = __read_122
													__dead_129.Rc = 1
													__dead_129.V0 = __scalar_124
													__dead_129.V1 = __read_125
													__dead_129.V2 = __scalar_126
													__dead_129.V3 = __read_127
													var __cell_131 *Constructor_Main_T
													if (__donor_slot_130) != (nil) {
														__cell_131 = __donor_slot_130
														__donor_slot_130 = nil
													} else {

													}
													if (__cell_131) == (nil) {
														__cell_131 = new(Constructor_Main_T)
													} else {

													}
													__cell_131.Rc = 1
													__cell_131.V0 = __scalar_118
													__cell_131.V1 = __dead_128
													__cell_131.V2 = __scalar_123
													__cell_131.V3 = __dead_129
													return __cell_131
												} else {
													__let_scalar_90 := uint32(__arg0)
													_ = __let_scalar_90
													__let_scalar_91 := int64(__arg2)
													_ = __let_scalar_91
													__scalar_92 := uint32(__let_scalar_90)
													_ = __scalar_92
													__read_93 := __arg1
													_ = __read_93
													__scalar_94 := int64(__let_scalar_91)
													_ = __scalar_94
													__read_95 := __arg3
													_ = __read_95
													__donor_slot_96 := __donor
													_ = __donor_slot_96
													var __cell_97 *Constructor_Main_T
													if (__donor_slot_96) != (nil) {
														__cell_97 = __donor_slot_96
														__donor_slot_96 = nil
													} else {

													}
													if (__cell_97) == (nil) {
														__cell_97 = new(Constructor_Main_T)
													} else {

													}
													__cell_97.Rc = 1
													__cell_97.V0 = __scalar_92
													__cell_97.V1 = __read_93
													__cell_97.V2 = __scalar_94
													__cell_97.V3 = __read_95
													return __cell_97
												}
											}
										} else {
											if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
												__let_scalar_132 := int64(__arg2)
												_ = __let_scalar_132
												__let_scalar_133 := int64(__arg3.V2)
												_ = __let_scalar_133
												__let_scalar_134 := int64(__arg3.V3.V2)
												_ = __let_scalar_134
												__scalar_135 := uint32(3558538316)
												_ = __scalar_135
												__scalar_136 := uint32(4250879068)
												_ = __scalar_136
												__read_137 := __arg1
												_ = __read_137
												__scalar_138 := int64(__let_scalar_132)
												_ = __scalar_138
												__read_139 := __arg3.V1
												_ = __read_139
												__scalar_140 := int64(__let_scalar_133)
												_ = __scalar_140
												__scalar_141 := uint32(4250879068)
												_ = __scalar_141
												__read_142 := __arg3.V3.V1
												_ = __read_142
												__scalar_143 := int64(__let_scalar_134)
												_ = __scalar_143
												__read_144 := __arg3.V3.V3
												_ = __read_144
												__donor_slot_147 := __donor
												_ = __donor_slot_147
												__dead_145 := __arg3
												_ = __dead_145
												__dead_146 := __arg3.V3
												_ = __dead_146
												__dead_145.Rc = 1
												__dead_145.V0 = __scalar_136
												__dead_145.V1 = __read_137
												__dead_145.V2 = __scalar_138
												__dead_145.V3 = __read_139
												__dead_146.Rc = 1
												__dead_146.V0 = __scalar_141
												__dead_146.V1 = __read_142
												__dead_146.V2 = __scalar_143
												__dead_146.V3 = __read_144
												var __cell_148 *Constructor_Main_T
												if (__donor_slot_147) != (nil) {
													__cell_148 = __donor_slot_147
													__donor_slot_147 = nil
												} else {

												}
												if (__cell_148) == (nil) {
													__cell_148 = new(Constructor_Main_T)
												} else {

												}
												__cell_148.Rc = 1
												__cell_148.V0 = __scalar_135
												__cell_148.V1 = __dead_145
												__cell_148.V2 = __scalar_140
												__cell_148.V3 = __dead_146
												return __cell_148
											} else {
												__let_scalar_82 := uint32(__arg0)
												_ = __let_scalar_82
												__let_scalar_83 := int64(__arg2)
												_ = __let_scalar_83
												__scalar_84 := uint32(__let_scalar_82)
												_ = __scalar_84
												__read_85 := __arg1
												_ = __read_85
												__scalar_86 := int64(__let_scalar_83)
												_ = __scalar_86
												__read_87 := __arg3
												_ = __read_87
												__donor_slot_88 := __donor
												_ = __donor_slot_88
												var __cell_89 *Constructor_Main_T
												if (__donor_slot_88) != (nil) {
													__cell_89 = __donor_slot_88
													__donor_slot_88 = nil
												} else {

												}
												if (__cell_89) == (nil) {
													__cell_89 = new(Constructor_Main_T)
												} else {

												}
												__cell_89.Rc = 1
												__cell_89.V0 = __scalar_84
												__cell_89.V1 = __read_85
												__cell_89.V2 = __scalar_86
												__cell_89.V3 = __read_87
												return __cell_89
											}
										}
									} else {
										__let_scalar_57 := uint32(__arg0)
										_ = __let_scalar_57
										__let_scalar_58 := int64(__arg2)
										_ = __let_scalar_58
										__scalar_59 := uint32(__let_scalar_57)
										_ = __scalar_59
										__read_60 := __arg1
										_ = __read_60
										__scalar_61 := int64(__let_scalar_58)
										_ = __scalar_61
										__read_62 := __arg3
										_ = __read_62
										__donor_slot_63 := __donor
										_ = __donor_slot_63
										var __cell_64 *Constructor_Main_T
										if (__donor_slot_63) != (nil) {
											__cell_64 = __donor_slot_63
											__donor_slot_63 = nil
										} else {

										}
										if (__cell_64) == (nil) {
											__cell_64 = new(Constructor_Main_T)
										} else {

										}
										__cell_64.Rc = 1
										__cell_64.V0 = __scalar_59
										__cell_64.V1 = __read_60
										__cell_64.V2 = __scalar_61
										__cell_64.V3 = __read_62
										return __cell_64
									}
								}
							} else {
								if ((__arg3) != (nil)) && ((__arg3.V0) == (3558538316)) {
									if (__arg3.V1) != (nil) {
										if (__arg3.V1.V0) == (3558538316) {
											__let_scalar_165 := int64(__arg2)
											_ = __let_scalar_165
											__let_scalar_166 := int64(__arg3.V1.V2)
											_ = __let_scalar_166
											__let_scalar_167 := int64(__arg3.V2)
											_ = __let_scalar_167
											__scalar_168 := uint32(3558538316)
											_ = __scalar_168
											__scalar_169 := uint32(4250879068)
											_ = __scalar_169
											__read_170 := __arg1
											_ = __read_170
											__scalar_171 := int64(__let_scalar_165)
											_ = __scalar_171
											__read_172 := __arg3.V1.V1
											_ = __read_172
											__scalar_173 := int64(__let_scalar_166)
											_ = __scalar_173
											__scalar_174 := uint32(4250879068)
											_ = __scalar_174
											__read_175 := __arg3.V1.V3
											_ = __read_175
											__scalar_176 := int64(__let_scalar_167)
											_ = __scalar_176
											__read_177 := __arg3.V3
											_ = __read_177
											__donor_slot_180 := __donor
											_ = __donor_slot_180
											__dead_178 := __arg3
											_ = __dead_178
											__dead_179 := __arg3.V1
											_ = __dead_179
											__dead_178.Rc = 1
											__dead_178.V0 = __scalar_169
											__dead_178.V1 = __read_170
											__dead_178.V2 = __scalar_171
											__dead_178.V3 = __read_172
											__dead_179.Rc = 1
											__dead_179.V0 = __scalar_174
											__dead_179.V1 = __read_175
											__dead_179.V2 = __scalar_176
											__dead_179.V3 = __read_177
											var __cell_181 *Constructor_Main_T
											if (__donor_slot_180) != (nil) {
												__cell_181 = __donor_slot_180
												__donor_slot_180 = nil
											} else {

											}
											if (__cell_181) == (nil) {
												__cell_181 = new(Constructor_Main_T)
											} else {

											}
											__cell_181.Rc = 1
											__cell_181.V0 = __scalar_168
											__cell_181.V1 = __dead_178
											__cell_181.V2 = __scalar_173
											__cell_181.V3 = __dead_179
											return __cell_181
										} else {
											if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
												__let_scalar_182 := int64(__arg2)
												_ = __let_scalar_182
												__let_scalar_183 := int64(__arg3.V2)
												_ = __let_scalar_183
												__let_scalar_184 := int64(__arg3.V3.V2)
												_ = __let_scalar_184
												__scalar_185 := uint32(3558538316)
												_ = __scalar_185
												__scalar_186 := uint32(4250879068)
												_ = __scalar_186
												__read_187 := __arg1
												_ = __read_187
												__scalar_188 := int64(__let_scalar_182)
												_ = __scalar_188
												__read_189 := __arg3.V1
												_ = __read_189
												__scalar_190 := int64(__let_scalar_183)
												_ = __scalar_190
												__scalar_191 := uint32(4250879068)
												_ = __scalar_191
												__read_192 := __arg3.V3.V1
												_ = __read_192
												__scalar_193 := int64(__let_scalar_184)
												_ = __scalar_193
												__read_194 := __arg3.V3.V3
												_ = __read_194
												__donor_slot_197 := __donor
												_ = __donor_slot_197
												__dead_195 := __arg3
												_ = __dead_195
												__dead_196 := __arg3.V3
												_ = __dead_196
												__dead_195.Rc = 1
												__dead_195.V0 = __scalar_186
												__dead_195.V1 = __read_187
												__dead_195.V2 = __scalar_188
												__dead_195.V3 = __read_189
												__dead_196.Rc = 1
												__dead_196.V0 = __scalar_191
												__dead_196.V1 = __read_192
												__dead_196.V2 = __scalar_193
												__dead_196.V3 = __read_194
												var __cell_198 *Constructor_Main_T
												if (__donor_slot_197) != (nil) {
													__cell_198 = __donor_slot_197
													__donor_slot_197 = nil
												} else {

												}
												if (__cell_198) == (nil) {
													__cell_198 = new(Constructor_Main_T)
												} else {

												}
												__cell_198.Rc = 1
												__cell_198.V0 = __scalar_185
												__cell_198.V1 = __dead_195
												__cell_198.V2 = __scalar_190
												__cell_198.V3 = __dead_196
												return __cell_198
											} else {
												__let_scalar_157 := uint32(__arg0)
												_ = __let_scalar_157
												__let_scalar_158 := int64(__arg2)
												_ = __let_scalar_158
												__scalar_159 := uint32(__let_scalar_157)
												_ = __scalar_159
												__read_160 := __arg1
												_ = __read_160
												__scalar_161 := int64(__let_scalar_158)
												_ = __scalar_161
												__read_162 := __arg3
												_ = __read_162
												__donor_slot_163 := __donor
												_ = __donor_slot_163
												var __cell_164 *Constructor_Main_T
												if (__donor_slot_163) != (nil) {
													__cell_164 = __donor_slot_163
													__donor_slot_163 = nil
												} else {

												}
												if (__cell_164) == (nil) {
													__cell_164 = new(Constructor_Main_T)
												} else {

												}
												__cell_164.Rc = 1
												__cell_164.V0 = __scalar_159
												__cell_164.V1 = __read_160
												__cell_164.V2 = __scalar_161
												__cell_164.V3 = __read_162
												return __cell_164
											}
										}
									} else {
										if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
											__let_scalar_199 := int64(__arg2)
											_ = __let_scalar_199
											__let_scalar_200 := int64(__arg3.V2)
											_ = __let_scalar_200
											__let_scalar_201 := int64(__arg3.V3.V2)
											_ = __let_scalar_201
											__scalar_202 := uint32(3558538316)
											_ = __scalar_202
											__scalar_203 := uint32(4250879068)
											_ = __scalar_203
											__read_204 := __arg1
											_ = __read_204
											__scalar_205 := int64(__let_scalar_199)
											_ = __scalar_205
											__read_206 := __arg3.V1
											_ = __read_206
											__scalar_207 := int64(__let_scalar_200)
											_ = __scalar_207
											__scalar_208 := uint32(4250879068)
											_ = __scalar_208
											__read_209 := __arg3.V3.V1
											_ = __read_209
											__scalar_210 := int64(__let_scalar_201)
											_ = __scalar_210
											__read_211 := __arg3.V3.V3
											_ = __read_211
											__donor_slot_214 := __donor
											_ = __donor_slot_214
											__dead_212 := __arg3
											_ = __dead_212
											__dead_213 := __arg3.V3
											_ = __dead_213
											__dead_212.Rc = 1
											__dead_212.V0 = __scalar_203
											__dead_212.V1 = __read_204
											__dead_212.V2 = __scalar_205
											__dead_212.V3 = __read_206
											__dead_213.Rc = 1
											__dead_213.V0 = __scalar_208
											__dead_213.V1 = __read_209
											__dead_213.V2 = __scalar_210
											__dead_213.V3 = __read_211
											var __cell_215 *Constructor_Main_T
											if (__donor_slot_214) != (nil) {
												__cell_215 = __donor_slot_214
												__donor_slot_214 = nil
											} else {

											}
											if (__cell_215) == (nil) {
												__cell_215 = new(Constructor_Main_T)
											} else {

											}
											__cell_215.Rc = 1
											__cell_215.V0 = __scalar_202
											__cell_215.V1 = __dead_212
											__cell_215.V2 = __scalar_207
											__cell_215.V3 = __dead_213
											return __cell_215
										} else {
											__let_scalar_149 := uint32(__arg0)
											_ = __let_scalar_149
											__let_scalar_150 := int64(__arg2)
											_ = __let_scalar_150
											__scalar_151 := uint32(__let_scalar_149)
											_ = __scalar_151
											__read_152 := __arg1
											_ = __read_152
											__scalar_153 := int64(__let_scalar_150)
											_ = __scalar_153
											__read_154 := __arg3
											_ = __read_154
											__donor_slot_155 := __donor
											_ = __donor_slot_155
											var __cell_156 *Constructor_Main_T
											if (__donor_slot_155) != (nil) {
												__cell_156 = __donor_slot_155
												__donor_slot_155 = nil
											} else {

											}
											if (__cell_156) == (nil) {
												__cell_156 = new(Constructor_Main_T)
											} else {

											}
											__cell_156.Rc = 1
											__cell_156.V0 = __scalar_151
											__cell_156.V1 = __read_152
											__cell_156.V2 = __scalar_153
											__cell_156.V3 = __read_154
											return __cell_156
										}
									}
								} else {
									__let_scalar_32 := uint32(__arg0)
									_ = __let_scalar_32
									__let_scalar_33 := int64(__arg2)
									_ = __let_scalar_33
									__scalar_34 := uint32(__let_scalar_32)
									_ = __scalar_34
									__read_35 := __arg1
									_ = __read_35
									__scalar_36 := int64(__let_scalar_33)
									_ = __scalar_36
									__read_37 := __arg3
									_ = __read_37
									__donor_slot_38 := __donor
									_ = __donor_slot_38
									var __cell_39 *Constructor_Main_T
									if (__donor_slot_38) != (nil) {
										__cell_39 = __donor_slot_38
										__donor_slot_38 = nil
									} else {

									}
									if (__cell_39) == (nil) {
										__cell_39 = new(Constructor_Main_T)
									} else {

									}
									__cell_39.Rc = 1
									__cell_39.V0 = __scalar_34
									__cell_39.V1 = __read_35
									__cell_39.V2 = __scalar_36
									__cell_39.V3 = __read_37
									return __cell_39
								}
							}
						}
					} else {
						if (__arg1.V3) != (nil) {
							if (__arg1.V3.V0) == (3558538316) {
								__let_scalar_224 := int64(__arg1.V2)
								_ = __let_scalar_224
								__let_scalar_225 := int64(__arg1.V3.V2)
								_ = __let_scalar_225
								__let_scalar_226 := int64(__arg2)
								_ = __let_scalar_226
								__scalar_227 := uint32(3558538316)
								_ = __scalar_227
								__scalar_228 := uint32(4250879068)
								_ = __scalar_228
								__read_229 := __arg1.V1
								_ = __read_229
								__scalar_230 := int64(__let_scalar_224)
								_ = __scalar_230
								__read_231 := __arg1.V3.V1
								_ = __read_231
								__scalar_232 := int64(__let_scalar_225)
								_ = __scalar_232
								__scalar_233 := uint32(4250879068)
								_ = __scalar_233
								__read_234 := __arg1.V3.V3
								_ = __read_234
								__scalar_235 := int64(__let_scalar_226)
								_ = __scalar_235
								__read_236 := __arg3
								_ = __read_236
								__donor_slot_239 := __donor
								_ = __donor_slot_239
								__dead_237 := __arg1
								_ = __dead_237
								__dead_238 := __arg1.V3
								_ = __dead_238
								__dead_237.Rc = 1
								__dead_237.V0 = __scalar_228
								__dead_237.V1 = __read_229
								__dead_237.V2 = __scalar_230
								__dead_237.V3 = __read_231
								__dead_238.Rc = 1
								__dead_238.V0 = __scalar_233
								__dead_238.V1 = __read_234
								__dead_238.V2 = __scalar_235
								__dead_238.V3 = __read_236
								var __cell_240 *Constructor_Main_T
								if (__donor_slot_239) != (nil) {
									__cell_240 = __donor_slot_239
									__donor_slot_239 = nil
								} else {

								}
								if (__cell_240) == (nil) {
									__cell_240 = new(Constructor_Main_T)
								} else {

								}
								__cell_240.Rc = 1
								__cell_240.V0 = __scalar_227
								__cell_240.V1 = __dead_237
								__cell_240.V2 = __scalar_232
								__cell_240.V3 = __dead_238
								return __cell_240
							} else {
								if ((__arg3) != (nil)) && ((__arg3.V0) == (3558538316)) {
									if (__arg3.V1) != (nil) {
										if (__arg3.V1.V0) == (3558538316) {
											__let_scalar_257 := int64(__arg2)
											_ = __let_scalar_257
											__let_scalar_258 := int64(__arg3.V1.V2)
											_ = __let_scalar_258
											__let_scalar_259 := int64(__arg3.V2)
											_ = __let_scalar_259
											__scalar_260 := uint32(3558538316)
											_ = __scalar_260
											__scalar_261 := uint32(4250879068)
											_ = __scalar_261
											__read_262 := __arg1
											_ = __read_262
											__scalar_263 := int64(__let_scalar_257)
											_ = __scalar_263
											__read_264 := __arg3.V1.V1
											_ = __read_264
											__scalar_265 := int64(__let_scalar_258)
											_ = __scalar_265
											__scalar_266 := uint32(4250879068)
											_ = __scalar_266
											__read_267 := __arg3.V1.V3
											_ = __read_267
											__scalar_268 := int64(__let_scalar_259)
											_ = __scalar_268
											__read_269 := __arg3.V3
											_ = __read_269
											__donor_slot_272 := __donor
											_ = __donor_slot_272
											__dead_270 := __arg3
											_ = __dead_270
											__dead_271 := __arg3.V1
											_ = __dead_271
											__dead_270.Rc = 1
											__dead_270.V0 = __scalar_261
											__dead_270.V1 = __read_262
											__dead_270.V2 = __scalar_263
											__dead_270.V3 = __read_264
											__dead_271.Rc = 1
											__dead_271.V0 = __scalar_266
											__dead_271.V1 = __read_267
											__dead_271.V2 = __scalar_268
											__dead_271.V3 = __read_269
											var __cell_273 *Constructor_Main_T
											if (__donor_slot_272) != (nil) {
												__cell_273 = __donor_slot_272
												__donor_slot_272 = nil
											} else {

											}
											if (__cell_273) == (nil) {
												__cell_273 = new(Constructor_Main_T)
											} else {

											}
											__cell_273.Rc = 1
											__cell_273.V0 = __scalar_260
											__cell_273.V1 = __dead_270
											__cell_273.V2 = __scalar_265
											__cell_273.V3 = __dead_271
											return __cell_273
										} else {
											if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
												__let_scalar_274 := int64(__arg2)
												_ = __let_scalar_274
												__let_scalar_275 := int64(__arg3.V2)
												_ = __let_scalar_275
												__let_scalar_276 := int64(__arg3.V3.V2)
												_ = __let_scalar_276
												__scalar_277 := uint32(3558538316)
												_ = __scalar_277
												__scalar_278 := uint32(4250879068)
												_ = __scalar_278
												__read_279 := __arg1
												_ = __read_279
												__scalar_280 := int64(__let_scalar_274)
												_ = __scalar_280
												__read_281 := __arg3.V1
												_ = __read_281
												__scalar_282 := int64(__let_scalar_275)
												_ = __scalar_282
												__scalar_283 := uint32(4250879068)
												_ = __scalar_283
												__read_284 := __arg3.V3.V1
												_ = __read_284
												__scalar_285 := int64(__let_scalar_276)
												_ = __scalar_285
												__read_286 := __arg3.V3.V3
												_ = __read_286
												__donor_slot_289 := __donor
												_ = __donor_slot_289
												__dead_287 := __arg3
												_ = __dead_287
												__dead_288 := __arg3.V3
												_ = __dead_288
												__dead_287.Rc = 1
												__dead_287.V0 = __scalar_278
												__dead_287.V1 = __read_279
												__dead_287.V2 = __scalar_280
												__dead_287.V3 = __read_281
												__dead_288.Rc = 1
												__dead_288.V0 = __scalar_283
												__dead_288.V1 = __read_284
												__dead_288.V2 = __scalar_285
												__dead_288.V3 = __read_286
												var __cell_290 *Constructor_Main_T
												if (__donor_slot_289) != (nil) {
													__cell_290 = __donor_slot_289
													__donor_slot_289 = nil
												} else {

												}
												if (__cell_290) == (nil) {
													__cell_290 = new(Constructor_Main_T)
												} else {

												}
												__cell_290.Rc = 1
												__cell_290.V0 = __scalar_277
												__cell_290.V1 = __dead_287
												__cell_290.V2 = __scalar_282
												__cell_290.V3 = __dead_288
												return __cell_290
											} else {
												__let_scalar_249 := uint32(__arg0)
												_ = __let_scalar_249
												__let_scalar_250 := int64(__arg2)
												_ = __let_scalar_250
												__scalar_251 := uint32(__let_scalar_249)
												_ = __scalar_251
												__read_252 := __arg1
												_ = __read_252
												__scalar_253 := int64(__let_scalar_250)
												_ = __scalar_253
												__read_254 := __arg3
												_ = __read_254
												__donor_slot_255 := __donor
												_ = __donor_slot_255
												var __cell_256 *Constructor_Main_T
												if (__donor_slot_255) != (nil) {
													__cell_256 = __donor_slot_255
													__donor_slot_255 = nil
												} else {

												}
												if (__cell_256) == (nil) {
													__cell_256 = new(Constructor_Main_T)
												} else {

												}
												__cell_256.Rc = 1
												__cell_256.V0 = __scalar_251
												__cell_256.V1 = __read_252
												__cell_256.V2 = __scalar_253
												__cell_256.V3 = __read_254
												return __cell_256
											}
										}
									} else {
										if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
											__let_scalar_291 := int64(__arg2)
											_ = __let_scalar_291
											__let_scalar_292 := int64(__arg3.V2)
											_ = __let_scalar_292
											__let_scalar_293 := int64(__arg3.V3.V2)
											_ = __let_scalar_293
											__scalar_294 := uint32(3558538316)
											_ = __scalar_294
											__scalar_295 := uint32(4250879068)
											_ = __scalar_295
											__read_296 := __arg1
											_ = __read_296
											__scalar_297 := int64(__let_scalar_291)
											_ = __scalar_297
											__read_298 := __arg3.V1
											_ = __read_298
											__scalar_299 := int64(__let_scalar_292)
											_ = __scalar_299
											__scalar_300 := uint32(4250879068)
											_ = __scalar_300
											__read_301 := __arg3.V3.V1
											_ = __read_301
											__scalar_302 := int64(__let_scalar_293)
											_ = __scalar_302
											__read_303 := __arg3.V3.V3
											_ = __read_303
											__donor_slot_306 := __donor
											_ = __donor_slot_306
											__dead_304 := __arg3
											_ = __dead_304
											__dead_305 := __arg3.V3
											_ = __dead_305
											__dead_304.Rc = 1
											__dead_304.V0 = __scalar_295
											__dead_304.V1 = __read_296
											__dead_304.V2 = __scalar_297
											__dead_304.V3 = __read_298
											__dead_305.Rc = 1
											__dead_305.V0 = __scalar_300
											__dead_305.V1 = __read_301
											__dead_305.V2 = __scalar_302
											__dead_305.V3 = __read_303
											var __cell_307 *Constructor_Main_T
											if (__donor_slot_306) != (nil) {
												__cell_307 = __donor_slot_306
												__donor_slot_306 = nil
											} else {

											}
											if (__cell_307) == (nil) {
												__cell_307 = new(Constructor_Main_T)
											} else {

											}
											__cell_307.Rc = 1
											__cell_307.V0 = __scalar_294
											__cell_307.V1 = __dead_304
											__cell_307.V2 = __scalar_299
											__cell_307.V3 = __dead_305
											return __cell_307
										} else {
											__let_scalar_241 := uint32(__arg0)
											_ = __let_scalar_241
											__let_scalar_242 := int64(__arg2)
											_ = __let_scalar_242
											__scalar_243 := uint32(__let_scalar_241)
											_ = __scalar_243
											__read_244 := __arg1
											_ = __read_244
											__scalar_245 := int64(__let_scalar_242)
											_ = __scalar_245
											__read_246 := __arg3
											_ = __read_246
											__donor_slot_247 := __donor
											_ = __donor_slot_247
											var __cell_248 *Constructor_Main_T
											if (__donor_slot_247) != (nil) {
												__cell_248 = __donor_slot_247
												__donor_slot_247 = nil
											} else {

											}
											if (__cell_248) == (nil) {
												__cell_248 = new(Constructor_Main_T)
											} else {

											}
											__cell_248.Rc = 1
											__cell_248.V0 = __scalar_243
											__cell_248.V1 = __read_244
											__cell_248.V2 = __scalar_245
											__cell_248.V3 = __read_246
											return __cell_248
										}
									}
								} else {
									__let_scalar_216 := uint32(__arg0)
									_ = __let_scalar_216
									__let_scalar_217 := int64(__arg2)
									_ = __let_scalar_217
									__scalar_218 := uint32(__let_scalar_216)
									_ = __scalar_218
									__read_219 := __arg1
									_ = __read_219
									__scalar_220 := int64(__let_scalar_217)
									_ = __scalar_220
									__read_221 := __arg3
									_ = __read_221
									__donor_slot_222 := __donor
									_ = __donor_slot_222
									var __cell_223 *Constructor_Main_T
									if (__donor_slot_222) != (nil) {
										__cell_223 = __donor_slot_222
										__donor_slot_222 = nil
									} else {

									}
									if (__cell_223) == (nil) {
										__cell_223 = new(Constructor_Main_T)
									} else {

									}
									__cell_223.Rc = 1
									__cell_223.V0 = __scalar_218
									__cell_223.V1 = __read_219
									__cell_223.V2 = __scalar_220
									__cell_223.V3 = __read_221
									return __cell_223
								}
							}
						} else {
							if ((__arg3) != (nil)) && ((__arg3.V0) == (3558538316)) {
								if (__arg3.V1) != (nil) {
									if (__arg3.V1.V0) == (3558538316) {
										__let_scalar_324 := int64(__arg2)
										_ = __let_scalar_324
										__let_scalar_325 := int64(__arg3.V1.V2)
										_ = __let_scalar_325
										__let_scalar_326 := int64(__arg3.V2)
										_ = __let_scalar_326
										__scalar_327 := uint32(3558538316)
										_ = __scalar_327
										__scalar_328 := uint32(4250879068)
										_ = __scalar_328
										__read_329 := __arg1
										_ = __read_329
										__scalar_330 := int64(__let_scalar_324)
										_ = __scalar_330
										__read_331 := __arg3.V1.V1
										_ = __read_331
										__scalar_332 := int64(__let_scalar_325)
										_ = __scalar_332
										__scalar_333 := uint32(4250879068)
										_ = __scalar_333
										__read_334 := __arg3.V1.V3
										_ = __read_334
										__scalar_335 := int64(__let_scalar_326)
										_ = __scalar_335
										__read_336 := __arg3.V3
										_ = __read_336
										__donor_slot_339 := __donor
										_ = __donor_slot_339
										__dead_337 := __arg3
										_ = __dead_337
										__dead_338 := __arg3.V1
										_ = __dead_338
										__dead_337.Rc = 1
										__dead_337.V0 = __scalar_328
										__dead_337.V1 = __read_329
										__dead_337.V2 = __scalar_330
										__dead_337.V3 = __read_331
										__dead_338.Rc = 1
										__dead_338.V0 = __scalar_333
										__dead_338.V1 = __read_334
										__dead_338.V2 = __scalar_335
										__dead_338.V3 = __read_336
										var __cell_340 *Constructor_Main_T
										if (__donor_slot_339) != (nil) {
											__cell_340 = __donor_slot_339
											__donor_slot_339 = nil
										} else {

										}
										if (__cell_340) == (nil) {
											__cell_340 = new(Constructor_Main_T)
										} else {

										}
										__cell_340.Rc = 1
										__cell_340.V0 = __scalar_327
										__cell_340.V1 = __dead_337
										__cell_340.V2 = __scalar_332
										__cell_340.V3 = __dead_338
										return __cell_340
									} else {
										if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
											__let_scalar_341 := int64(__arg2)
											_ = __let_scalar_341
											__let_scalar_342 := int64(__arg3.V2)
											_ = __let_scalar_342
											__let_scalar_343 := int64(__arg3.V3.V2)
											_ = __let_scalar_343
											__scalar_344 := uint32(3558538316)
											_ = __scalar_344
											__scalar_345 := uint32(4250879068)
											_ = __scalar_345
											__read_346 := __arg1
											_ = __read_346
											__scalar_347 := int64(__let_scalar_341)
											_ = __scalar_347
											__read_348 := __arg3.V1
											_ = __read_348
											__scalar_349 := int64(__let_scalar_342)
											_ = __scalar_349
											__scalar_350 := uint32(4250879068)
											_ = __scalar_350
											__read_351 := __arg3.V3.V1
											_ = __read_351
											__scalar_352 := int64(__let_scalar_343)
											_ = __scalar_352
											__read_353 := __arg3.V3.V3
											_ = __read_353
											__donor_slot_356 := __donor
											_ = __donor_slot_356
											__dead_354 := __arg3
											_ = __dead_354
											__dead_355 := __arg3.V3
											_ = __dead_355
											__dead_354.Rc = 1
											__dead_354.V0 = __scalar_345
											__dead_354.V1 = __read_346
											__dead_354.V2 = __scalar_347
											__dead_354.V3 = __read_348
											__dead_355.Rc = 1
											__dead_355.V0 = __scalar_350
											__dead_355.V1 = __read_351
											__dead_355.V2 = __scalar_352
											__dead_355.V3 = __read_353
											var __cell_357 *Constructor_Main_T
											if (__donor_slot_356) != (nil) {
												__cell_357 = __donor_slot_356
												__donor_slot_356 = nil
											} else {

											}
											if (__cell_357) == (nil) {
												__cell_357 = new(Constructor_Main_T)
											} else {

											}
											__cell_357.Rc = 1
											__cell_357.V0 = __scalar_344
											__cell_357.V1 = __dead_354
											__cell_357.V2 = __scalar_349
											__cell_357.V3 = __dead_355
											return __cell_357
										} else {
											__let_scalar_316 := uint32(__arg0)
											_ = __let_scalar_316
											__let_scalar_317 := int64(__arg2)
											_ = __let_scalar_317
											__scalar_318 := uint32(__let_scalar_316)
											_ = __scalar_318
											__read_319 := __arg1
											_ = __read_319
											__scalar_320 := int64(__let_scalar_317)
											_ = __scalar_320
											__read_321 := __arg3
											_ = __read_321
											__donor_slot_322 := __donor
											_ = __donor_slot_322
											var __cell_323 *Constructor_Main_T
											if (__donor_slot_322) != (nil) {
												__cell_323 = __donor_slot_322
												__donor_slot_322 = nil
											} else {

											}
											if (__cell_323) == (nil) {
												__cell_323 = new(Constructor_Main_T)
											} else {

											}
											__cell_323.Rc = 1
											__cell_323.V0 = __scalar_318
											__cell_323.V1 = __read_319
											__cell_323.V2 = __scalar_320
											__cell_323.V3 = __read_321
											return __cell_323
										}
									}
								} else {
									if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
										__let_scalar_358 := int64(__arg2)
										_ = __let_scalar_358
										__let_scalar_359 := int64(__arg3.V2)
										_ = __let_scalar_359
										__let_scalar_360 := int64(__arg3.V3.V2)
										_ = __let_scalar_360
										__scalar_361 := uint32(3558538316)
										_ = __scalar_361
										__scalar_362 := uint32(4250879068)
										_ = __scalar_362
										__read_363 := __arg1
										_ = __read_363
										__scalar_364 := int64(__let_scalar_358)
										_ = __scalar_364
										__read_365 := __arg3.V1
										_ = __read_365
										__scalar_366 := int64(__let_scalar_359)
										_ = __scalar_366
										__scalar_367 := uint32(4250879068)
										_ = __scalar_367
										__read_368 := __arg3.V3.V1
										_ = __read_368
										__scalar_369 := int64(__let_scalar_360)
										_ = __scalar_369
										__read_370 := __arg3.V3.V3
										_ = __read_370
										__donor_slot_373 := __donor
										_ = __donor_slot_373
										__dead_371 := __arg3
										_ = __dead_371
										__dead_372 := __arg3.V3
										_ = __dead_372
										__dead_371.Rc = 1
										__dead_371.V0 = __scalar_362
										__dead_371.V1 = __read_363
										__dead_371.V2 = __scalar_364
										__dead_371.V3 = __read_365
										__dead_372.Rc = 1
										__dead_372.V0 = __scalar_367
										__dead_372.V1 = __read_368
										__dead_372.V2 = __scalar_369
										__dead_372.V3 = __read_370
										var __cell_374 *Constructor_Main_T
										if (__donor_slot_373) != (nil) {
											__cell_374 = __donor_slot_373
											__donor_slot_373 = nil
										} else {

										}
										if (__cell_374) == (nil) {
											__cell_374 = new(Constructor_Main_T)
										} else {

										}
										__cell_374.Rc = 1
										__cell_374.V0 = __scalar_361
										__cell_374.V1 = __dead_371
										__cell_374.V2 = __scalar_366
										__cell_374.V3 = __dead_372
										return __cell_374
									} else {
										__let_scalar_308 := uint32(__arg0)
										_ = __let_scalar_308
										__let_scalar_309 := int64(__arg2)
										_ = __let_scalar_309
										__scalar_310 := uint32(__let_scalar_308)
										_ = __scalar_310
										__read_311 := __arg1
										_ = __read_311
										__scalar_312 := int64(__let_scalar_309)
										_ = __scalar_312
										__read_313 := __arg3
										_ = __read_313
										__donor_slot_314 := __donor
										_ = __donor_slot_314
										var __cell_315 *Constructor_Main_T
										if (__donor_slot_314) != (nil) {
											__cell_315 = __donor_slot_314
											__donor_slot_314 = nil
										} else {

										}
										if (__cell_315) == (nil) {
											__cell_315 = new(Constructor_Main_T)
										} else {

										}
										__cell_315.Rc = 1
										__cell_315.V0 = __scalar_310
										__cell_315.V1 = __read_311
										__cell_315.V2 = __scalar_312
										__cell_315.V3 = __read_313
										return __cell_315
									}
								}
							} else {
								__let_scalar_24 := uint32(__arg0)
								_ = __let_scalar_24
								__let_scalar_25 := int64(__arg2)
								_ = __let_scalar_25
								__scalar_26 := uint32(__let_scalar_24)
								_ = __scalar_26
								__read_27 := __arg1
								_ = __read_27
								__scalar_28 := int64(__let_scalar_25)
								_ = __scalar_28
								__read_29 := __arg3
								_ = __read_29
								__donor_slot_30 := __donor
								_ = __donor_slot_30
								var __cell_31 *Constructor_Main_T
								if (__donor_slot_30) != (nil) {
									__cell_31 = __donor_slot_30
									__donor_slot_30 = nil
								} else {

								}
								if (__cell_31) == (nil) {
									__cell_31 = new(Constructor_Main_T)
								} else {

								}
								__cell_31.Rc = 1
								__cell_31.V0 = __scalar_26
								__cell_31.V1 = __read_27
								__cell_31.V2 = __scalar_28
								__cell_31.V3 = __read_29
								return __cell_31
							}
						}
					}
				} else {
					if ((__arg3) != (nil)) && ((__arg3.V0) == (3558538316)) {
						if (__arg3.V1) != (nil) {
							if (__arg3.V1.V0) == (3558538316) {
								__let_scalar_391 := int64(__arg2)
								_ = __let_scalar_391
								__let_scalar_392 := int64(__arg3.V1.V2)
								_ = __let_scalar_392
								__let_scalar_393 := int64(__arg3.V2)
								_ = __let_scalar_393
								__scalar_394 := uint32(3558538316)
								_ = __scalar_394
								__scalar_395 := uint32(4250879068)
								_ = __scalar_395
								__read_396 := __arg1
								_ = __read_396
								__scalar_397 := int64(__let_scalar_391)
								_ = __scalar_397
								__read_398 := __arg3.V1.V1
								_ = __read_398
								__scalar_399 := int64(__let_scalar_392)
								_ = __scalar_399
								__scalar_400 := uint32(4250879068)
								_ = __scalar_400
								__read_401 := __arg3.V1.V3
								_ = __read_401
								__scalar_402 := int64(__let_scalar_393)
								_ = __scalar_402
								__read_403 := __arg3.V3
								_ = __read_403
								__donor_slot_406 := __donor
								_ = __donor_slot_406
								__dead_404 := __arg3
								_ = __dead_404
								__dead_405 := __arg3.V1
								_ = __dead_405
								__dead_404.Rc = 1
								__dead_404.V0 = __scalar_395
								__dead_404.V1 = __read_396
								__dead_404.V2 = __scalar_397
								__dead_404.V3 = __read_398
								__dead_405.Rc = 1
								__dead_405.V0 = __scalar_400
								__dead_405.V1 = __read_401
								__dead_405.V2 = __scalar_402
								__dead_405.V3 = __read_403
								var __cell_407 *Constructor_Main_T
								if (__donor_slot_406) != (nil) {
									__cell_407 = __donor_slot_406
									__donor_slot_406 = nil
								} else {

								}
								if (__cell_407) == (nil) {
									__cell_407 = new(Constructor_Main_T)
								} else {

								}
								__cell_407.Rc = 1
								__cell_407.V0 = __scalar_394
								__cell_407.V1 = __dead_404
								__cell_407.V2 = __scalar_399
								__cell_407.V3 = __dead_405
								return __cell_407
							} else {
								if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
									__let_scalar_408 := int64(__arg2)
									_ = __let_scalar_408
									__let_scalar_409 := int64(__arg3.V2)
									_ = __let_scalar_409
									__let_scalar_410 := int64(__arg3.V3.V2)
									_ = __let_scalar_410
									__scalar_411 := uint32(3558538316)
									_ = __scalar_411
									__scalar_412 := uint32(4250879068)
									_ = __scalar_412
									__read_413 := __arg1
									_ = __read_413
									__scalar_414 := int64(__let_scalar_408)
									_ = __scalar_414
									__read_415 := __arg3.V1
									_ = __read_415
									__scalar_416 := int64(__let_scalar_409)
									_ = __scalar_416
									__scalar_417 := uint32(4250879068)
									_ = __scalar_417
									__read_418 := __arg3.V3.V1
									_ = __read_418
									__scalar_419 := int64(__let_scalar_410)
									_ = __scalar_419
									__read_420 := __arg3.V3.V3
									_ = __read_420
									__donor_slot_423 := __donor
									_ = __donor_slot_423
									__dead_421 := __arg3
									_ = __dead_421
									__dead_422 := __arg3.V3
									_ = __dead_422
									__dead_421.Rc = 1
									__dead_421.V0 = __scalar_412
									__dead_421.V1 = __read_413
									__dead_421.V2 = __scalar_414
									__dead_421.V3 = __read_415
									__dead_422.Rc = 1
									__dead_422.V0 = __scalar_417
									__dead_422.V1 = __read_418
									__dead_422.V2 = __scalar_419
									__dead_422.V3 = __read_420
									var __cell_424 *Constructor_Main_T
									if (__donor_slot_423) != (nil) {
										__cell_424 = __donor_slot_423
										__donor_slot_423 = nil
									} else {

									}
									if (__cell_424) == (nil) {
										__cell_424 = new(Constructor_Main_T)
									} else {

									}
									__cell_424.Rc = 1
									__cell_424.V0 = __scalar_411
									__cell_424.V1 = __dead_421
									__cell_424.V2 = __scalar_416
									__cell_424.V3 = __dead_422
									return __cell_424
								} else {
									__let_scalar_383 := uint32(__arg0)
									_ = __let_scalar_383
									__let_scalar_384 := int64(__arg2)
									_ = __let_scalar_384
									__scalar_385 := uint32(__let_scalar_383)
									_ = __scalar_385
									__read_386 := __arg1
									_ = __read_386
									__scalar_387 := int64(__let_scalar_384)
									_ = __scalar_387
									__read_388 := __arg3
									_ = __read_388
									__donor_slot_389 := __donor
									_ = __donor_slot_389
									var __cell_390 *Constructor_Main_T
									if (__donor_slot_389) != (nil) {
										__cell_390 = __donor_slot_389
										__donor_slot_389 = nil
									} else {

									}
									if (__cell_390) == (nil) {
										__cell_390 = new(Constructor_Main_T)
									} else {

									}
									__cell_390.Rc = 1
									__cell_390.V0 = __scalar_385
									__cell_390.V1 = __read_386
									__cell_390.V2 = __scalar_387
									__cell_390.V3 = __read_388
									return __cell_390
								}
							}
						} else {
							if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
								__let_scalar_425 := int64(__arg2)
								_ = __let_scalar_425
								__let_scalar_426 := int64(__arg3.V2)
								_ = __let_scalar_426
								__let_scalar_427 := int64(__arg3.V3.V2)
								_ = __let_scalar_427
								__scalar_428 := uint32(3558538316)
								_ = __scalar_428
								__scalar_429 := uint32(4250879068)
								_ = __scalar_429
								__read_430 := __arg1
								_ = __read_430
								__scalar_431 := int64(__let_scalar_425)
								_ = __scalar_431
								__read_432 := __arg3.V1
								_ = __read_432
								__scalar_433 := int64(__let_scalar_426)
								_ = __scalar_433
								__scalar_434 := uint32(4250879068)
								_ = __scalar_434
								__read_435 := __arg3.V3.V1
								_ = __read_435
								__scalar_436 := int64(__let_scalar_427)
								_ = __scalar_436
								__read_437 := __arg3.V3.V3
								_ = __read_437
								__donor_slot_440 := __donor
								_ = __donor_slot_440
								__dead_438 := __arg3
								_ = __dead_438
								__dead_439 := __arg3.V3
								_ = __dead_439
								__dead_438.Rc = 1
								__dead_438.V0 = __scalar_429
								__dead_438.V1 = __read_430
								__dead_438.V2 = __scalar_431
								__dead_438.V3 = __read_432
								__dead_439.Rc = 1
								__dead_439.V0 = __scalar_434
								__dead_439.V1 = __read_435
								__dead_439.V2 = __scalar_436
								__dead_439.V3 = __read_437
								var __cell_441 *Constructor_Main_T
								if (__donor_slot_440) != (nil) {
									__cell_441 = __donor_slot_440
									__donor_slot_440 = nil
								} else {

								}
								if (__cell_441) == (nil) {
									__cell_441 = new(Constructor_Main_T)
								} else {

								}
								__cell_441.Rc = 1
								__cell_441.V0 = __scalar_428
								__cell_441.V1 = __dead_438
								__cell_441.V2 = __scalar_433
								__cell_441.V3 = __dead_439
								return __cell_441
							} else {
								__let_scalar_375 := uint32(__arg0)
								_ = __let_scalar_375
								__let_scalar_376 := int64(__arg2)
								_ = __let_scalar_376
								__scalar_377 := uint32(__let_scalar_375)
								_ = __scalar_377
								__read_378 := __arg1
								_ = __read_378
								__scalar_379 := int64(__let_scalar_376)
								_ = __scalar_379
								__read_380 := __arg3
								_ = __read_380
								__donor_slot_381 := __donor
								_ = __donor_slot_381
								var __cell_382 *Constructor_Main_T
								if (__donor_slot_381) != (nil) {
									__cell_382 = __donor_slot_381
									__donor_slot_381 = nil
								} else {

								}
								if (__cell_382) == (nil) {
									__cell_382 = new(Constructor_Main_T)
								} else {

								}
								__cell_382.Rc = 1
								__cell_382.V0 = __scalar_377
								__cell_382.V1 = __read_378
								__cell_382.V2 = __scalar_379
								__cell_382.V3 = __read_380
								return __cell_382
							}
						}
					} else {
						__let_scalar_16 := uint32(__arg0)
						_ = __let_scalar_16
						__let_scalar_17 := int64(__arg2)
						_ = __let_scalar_17
						__scalar_18 := uint32(__let_scalar_16)
						_ = __scalar_18
						__read_19 := __arg1
						_ = __read_19
						__scalar_20 := int64(__let_scalar_17)
						_ = __scalar_20
						__read_21 := __arg3
						_ = __read_21
						__donor_slot_22 := __donor
						_ = __donor_slot_22
						var __cell_23 *Constructor_Main_T
						if (__donor_slot_22) != (nil) {
							__cell_23 = __donor_slot_22
							__donor_slot_22 = nil
						} else {

						}
						if (__cell_23) == (nil) {
							__cell_23 = new(Constructor_Main_T)
						} else {

						}
						__cell_23.Rc = 1
						__cell_23.V0 = __scalar_18
						__cell_23.V1 = __read_19
						__cell_23.V2 = __scalar_20
						__cell_23.V3 = __read_21
						return __cell_23
					}
				}
			} else {
				if ((__arg3) != (nil)) && ((__arg3.V0) == (3558538316)) {
					if (__arg3.V1) != (nil) {
						if (__arg3.V1.V0) == (3558538316) {
							__let_scalar_458 := int64(__arg2)
							_ = __let_scalar_458
							__let_scalar_459 := int64(__arg3.V1.V2)
							_ = __let_scalar_459
							__let_scalar_460 := int64(__arg3.V2)
							_ = __let_scalar_460
							__scalar_461 := uint32(3558538316)
							_ = __scalar_461
							__scalar_462 := uint32(4250879068)
							_ = __scalar_462
							__read_463 := __arg1
							_ = __read_463
							__scalar_464 := int64(__let_scalar_458)
							_ = __scalar_464
							__read_465 := __arg3.V1.V1
							_ = __read_465
							__scalar_466 := int64(__let_scalar_459)
							_ = __scalar_466
							__scalar_467 := uint32(4250879068)
							_ = __scalar_467
							__read_468 := __arg3.V1.V3
							_ = __read_468
							__scalar_469 := int64(__let_scalar_460)
							_ = __scalar_469
							__read_470 := __arg3.V3
							_ = __read_470
							__donor_slot_473 := __donor
							_ = __donor_slot_473
							__dead_471 := __arg3
							_ = __dead_471
							__dead_472 := __arg3.V1
							_ = __dead_472
							__dead_471.Rc = 1
							__dead_471.V0 = __scalar_462
							__dead_471.V1 = __read_463
							__dead_471.V2 = __scalar_464
							__dead_471.V3 = __read_465
							__dead_472.Rc = 1
							__dead_472.V0 = __scalar_467
							__dead_472.V1 = __read_468
							__dead_472.V2 = __scalar_469
							__dead_472.V3 = __read_470
							var __cell_474 *Constructor_Main_T
							if (__donor_slot_473) != (nil) {
								__cell_474 = __donor_slot_473
								__donor_slot_473 = nil
							} else {

							}
							if (__cell_474) == (nil) {
								__cell_474 = new(Constructor_Main_T)
							} else {

							}
							__cell_474.Rc = 1
							__cell_474.V0 = __scalar_461
							__cell_474.V1 = __dead_471
							__cell_474.V2 = __scalar_466
							__cell_474.V3 = __dead_472
							return __cell_474
						} else {
							if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
								__let_scalar_475 := int64(__arg2)
								_ = __let_scalar_475
								__let_scalar_476 := int64(__arg3.V2)
								_ = __let_scalar_476
								__let_scalar_477 := int64(__arg3.V3.V2)
								_ = __let_scalar_477
								__scalar_478 := uint32(3558538316)
								_ = __scalar_478
								__scalar_479 := uint32(4250879068)
								_ = __scalar_479
								__read_480 := __arg1
								_ = __read_480
								__scalar_481 := int64(__let_scalar_475)
								_ = __scalar_481
								__read_482 := __arg3.V1
								_ = __read_482
								__scalar_483 := int64(__let_scalar_476)
								_ = __scalar_483
								__scalar_484 := uint32(4250879068)
								_ = __scalar_484
								__read_485 := __arg3.V3.V1
								_ = __read_485
								__scalar_486 := int64(__let_scalar_477)
								_ = __scalar_486
								__read_487 := __arg3.V3.V3
								_ = __read_487
								__donor_slot_490 := __donor
								_ = __donor_slot_490
								__dead_488 := __arg3
								_ = __dead_488
								__dead_489 := __arg3.V3
								_ = __dead_489
								__dead_488.Rc = 1
								__dead_488.V0 = __scalar_479
								__dead_488.V1 = __read_480
								__dead_488.V2 = __scalar_481
								__dead_488.V3 = __read_482
								__dead_489.Rc = 1
								__dead_489.V0 = __scalar_484
								__dead_489.V1 = __read_485
								__dead_489.V2 = __scalar_486
								__dead_489.V3 = __read_487
								var __cell_491 *Constructor_Main_T
								if (__donor_slot_490) != (nil) {
									__cell_491 = __donor_slot_490
									__donor_slot_490 = nil
								} else {

								}
								if (__cell_491) == (nil) {
									__cell_491 = new(Constructor_Main_T)
								} else {

								}
								__cell_491.Rc = 1
								__cell_491.V0 = __scalar_478
								__cell_491.V1 = __dead_488
								__cell_491.V2 = __scalar_483
								__cell_491.V3 = __dead_489
								return __cell_491
							} else {
								__let_scalar_450 := uint32(__arg0)
								_ = __let_scalar_450
								__let_scalar_451 := int64(__arg2)
								_ = __let_scalar_451
								__scalar_452 := uint32(__let_scalar_450)
								_ = __scalar_452
								__read_453 := __arg1
								_ = __read_453
								__scalar_454 := int64(__let_scalar_451)
								_ = __scalar_454
								__read_455 := __arg3
								_ = __read_455
								__donor_slot_456 := __donor
								_ = __donor_slot_456
								var __cell_457 *Constructor_Main_T
								if (__donor_slot_456) != (nil) {
									__cell_457 = __donor_slot_456
									__donor_slot_456 = nil
								} else {

								}
								if (__cell_457) == (nil) {
									__cell_457 = new(Constructor_Main_T)
								} else {

								}
								__cell_457.Rc = 1
								__cell_457.V0 = __scalar_452
								__cell_457.V1 = __read_453
								__cell_457.V2 = __scalar_454
								__cell_457.V3 = __read_455
								return __cell_457
							}
						}
					} else {
						if ((__arg3.V3) != (nil)) && ((__arg3.V3.V0) == (3558538316)) {
							__let_scalar_492 := int64(__arg2)
							_ = __let_scalar_492
							__let_scalar_493 := int64(__arg3.V2)
							_ = __let_scalar_493
							__let_scalar_494 := int64(__arg3.V3.V2)
							_ = __let_scalar_494
							__scalar_495 := uint32(3558538316)
							_ = __scalar_495
							__scalar_496 := uint32(4250879068)
							_ = __scalar_496
							__read_497 := __arg1
							_ = __read_497
							__scalar_498 := int64(__let_scalar_492)
							_ = __scalar_498
							__read_499 := __arg3.V1
							_ = __read_499
							__scalar_500 := int64(__let_scalar_493)
							_ = __scalar_500
							__scalar_501 := uint32(4250879068)
							_ = __scalar_501
							__read_502 := __arg3.V3.V1
							_ = __read_502
							__scalar_503 := int64(__let_scalar_494)
							_ = __scalar_503
							__read_504 := __arg3.V3.V3
							_ = __read_504
							__donor_slot_507 := __donor
							_ = __donor_slot_507
							__dead_505 := __arg3
							_ = __dead_505
							__dead_506 := __arg3.V3
							_ = __dead_506
							__dead_505.Rc = 1
							__dead_505.V0 = __scalar_496
							__dead_505.V1 = __read_497
							__dead_505.V2 = __scalar_498
							__dead_505.V3 = __read_499
							__dead_506.Rc = 1
							__dead_506.V0 = __scalar_501
							__dead_506.V1 = __read_502
							__dead_506.V2 = __scalar_503
							__dead_506.V3 = __read_504
							var __cell_508 *Constructor_Main_T
							if (__donor_slot_507) != (nil) {
								__cell_508 = __donor_slot_507
								__donor_slot_507 = nil
							} else {

							}
							if (__cell_508) == (nil) {
								__cell_508 = new(Constructor_Main_T)
							} else {

							}
							__cell_508.Rc = 1
							__cell_508.V0 = __scalar_495
							__cell_508.V1 = __dead_505
							__cell_508.V2 = __scalar_500
							__cell_508.V3 = __dead_506
							return __cell_508
						} else {
							__let_scalar_442 := uint32(__arg0)
							_ = __let_scalar_442
							__let_scalar_443 := int64(__arg2)
							_ = __let_scalar_443
							__scalar_444 := uint32(__let_scalar_442)
							_ = __scalar_444
							__read_445 := __arg1
							_ = __read_445
							__scalar_446 := int64(__let_scalar_443)
							_ = __scalar_446
							__read_447 := __arg3
							_ = __read_447
							__donor_slot_448 := __donor
							_ = __donor_slot_448
							var __cell_449 *Constructor_Main_T
							if (__donor_slot_448) != (nil) {
								__cell_449 = __donor_slot_448
								__donor_slot_448 = nil
							} else {

							}
							if (__cell_449) == (nil) {
								__cell_449 = new(Constructor_Main_T)
							} else {

							}
							__cell_449.Rc = 1
							__cell_449.V0 = __scalar_444
							__cell_449.V1 = __read_445
							__cell_449.V2 = __scalar_446
							__cell_449.V3 = __read_447
							return __cell_449
						}
					}
				} else {
					__let_scalar_8 := uint32(__arg0)
					_ = __let_scalar_8
					__let_scalar_9 := int64(__arg2)
					_ = __let_scalar_9
					__scalar_10 := uint32(__let_scalar_8)
					_ = __scalar_10
					__read_11 := __arg1
					_ = __read_11
					__scalar_12 := int64(__let_scalar_9)
					_ = __scalar_12
					__read_13 := __arg3
					_ = __read_13
					__donor_slot_14 := __donor
					_ = __donor_slot_14
					var __cell_15 *Constructor_Main_T
					if (__donor_slot_14) != (nil) {
						__cell_15 = __donor_slot_14
						__donor_slot_14 = nil
					} else {

					}
					if (__cell_15) == (nil) {
						__cell_15 = new(Constructor_Main_T)
					} else {

					}
					__cell_15.Rc = 1
					__cell_15.V0 = __scalar_10
					__cell_15.V1 = __read_11
					__cell_15.V2 = __scalar_12
					__cell_15.V3 = __read_13
					return __cell_15
				}
			}
		} else {
			__let_scalar_0 := uint32(__arg0)
			_ = __let_scalar_0
			__let_scalar_1 := int64(__arg2)
			_ = __let_scalar_1
			__scalar_2 := uint32(__let_scalar_0)
			_ = __scalar_2
			__read_3 := __arg1
			_ = __read_3
			__scalar_4 := int64(__let_scalar_1)
			_ = __scalar_4
			__read_5 := __arg3
			_ = __read_5
			__donor_slot_6 := __donor
			_ = __donor_slot_6
			var __cell_7 *Constructor_Main_T
			if (__donor_slot_6) != (nil) {
				__cell_7 = __donor_slot_6
				__donor_slot_6 = nil
			} else {

			}
			if (__cell_7) == (nil) {
				__cell_7 = new(Constructor_Main_T)
			} else {

			}
			__cell_7.Rc = 1
			__cell_7.V0 = __scalar_2
			__cell_7.V1 = __read_3
			__cell_7.V2 = __scalar_4
			__cell_7.V3 = __read_5
			return __cell_7
		}
	}
}

func Call_Main___gopurs_owned_balance_0(__arg0 uint32, __arg1 *Constructor_Main_T, __arg2 int64, __arg3 *Constructor_Main_T) *Constructor_Main_T {
	return Call_Main___gopurs_owned_balance_0_consume(__arg0, __arg1, __arg2, __arg3, nil)
}

type Constructor_Main_R struct {
	Rc uint32
}

type Constructor_Main_B struct {
	Rc uint32
}

type Constructor_Main_E struct {
	Rc uint32
}

type Constructor_Main_T struct {
	Rc uint32
	V0 uint32
	V1 *Constructor_Main_T
	V2 int64
	V3 *Constructor_Main_T
}

func Call_Main_max(x_0_loop int64, y_1_loop int64) int64 {
	var x_0 int64 = x_0_loop
	_ = x_0
	var y_1 int64 = y_1_loop
	_ = y_1
	var __t0 int64
	{
		if (x_0) > (y_1) {
			__t0 = x_0
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = y_1
	}
end_branch_0:
	return __t0
}

func Call_Main_depth(v_0_loop *Constructor_Main_T) int64 {
depth:
	for {
		if false {
			continue depth
		}
		var v_0 *Constructor_Main_T = v_0_loop
		_ = v_0
		var __t3 int64
		{
			if v_0 == nil {
				__t3 = int64(0)
				goto end_branch_3
			} else {

			}
		}
		{
			if v_0 != nil {
				// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
				__local_var_1_0 := Call_Main_depth((v_0).V1)
				_ = __local_var_1_0
				// TAST (Let): __local_var_2_1 shape=App(Var) bindingType=Int
				__local_var_2_1 := Call_Main_depth((v_0).V3)
				_ = __local_var_2_1
				var __t2 int64
				{
					if (__local_var_1_0) > (__local_var_2_1) {
						__t2 = __local_var_1_0
						goto end_branch_2
					} else {

					}
				}
				{
					__t2 = __local_var_2_1
				}
			end_branch_2:
				__t3 = (int64(1)) + (__t2)
				goto end_branch_3
			} else {

			}
		}
		{
			__t3 = func() int64 { panic("Failed pattern match") }()
		}
	end_branch_3:
		return __t3
	}
}

func Call_Main_balance(v_0_loop uint32, v1_1_loop *Constructor_Main_T, v2_2_loop int64, v3_3_loop *Constructor_Main_T) *Constructor_Main_T {
	var v_0 uint32 = v_0_loop
	_ = v_0
	var v1_1 *Constructor_Main_T = v1_1_loop
	_ = v1_1
	var v2_2 int64 = v2_2_loop
	_ = v2_2
	var v3_3 *Constructor_Main_T = v3_3_loop
	_ = v3_3
	var __t308 *Constructor_Main_T
	{
		if v_0 == 4250879068 {
			var __t307 *Constructor_Main_T
			{
				if v1_1 != nil {
					var __t265 *Constructor_Main_T
					{
						var __t_tag_12 uint32 = (v1_1).V0
						_ = __t_tag_12
						if uint32(__t_tag_12) == 3558538316 {
							var __t223 *Constructor_Main_T
							{
								var __t_tag_17 *Constructor_Main_T = (v1_1).V1
								_ = __t_tag_17
								if __t_tag_17 != nil {
									var __t126 *Constructor_Main_T
									{
										var __t_tag_22 uint32 = ((v1_1).V1).V0
										_ = __t_tag_22
										if uint32(__t_tag_22) == 3558538316 {
											// TAST (Let): __local_var_4_23 shape=Other bindingType=Any
											__local_var_4_23 := ((v1_1).V1).V1
											_ = __local_var_4_23
											// TAST (Let): __local_var_5_24 shape=Other bindingType=Any
											__local_var_5_24 := ((v1_1).V1).V3
											_ = __local_var_5_24
											// TAST (Let): __local_var_6_25 shape=Other bindingType=Any
											__local_var_6_25 := (v1_1).V3
											_ = __local_var_6_25
											// TAST (Let): __local_var_7_26 shape=Other bindingType=(ADT ["Main","Tree"] [])
											__local_var_7_26 := v3_3
											_ = __local_var_7_26
											// TAST (Let): __local_var_8_27 shape=Other bindingType=Any
											__local_var_8_27 := ((v1_1).V1).V2
											_ = __local_var_8_27
											// TAST (Let): __local_var_9_28 shape=Other bindingType=Any
											__local_var_9_28 := (v1_1).V2
											_ = __local_var_9_28
											// TAST (Let): __local_var_10_29 shape=Other bindingType=Int
											__local_var_10_29 := v2_2
											_ = __local_var_10_29
											__t126 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_23, __local_var_8_27, __local_var_5_24}), __local_var_9_28, (&Constructor_Main_T{1, 4250879068, __local_var_6_25, __local_var_10_29, __local_var_7_26})})
											goto end_branch_126
										} else {

										}
									}
									{
										var __t_tag_30 *Constructor_Main_T = (v1_1).V3
										_ = __t_tag_30
										if __t_tag_30 != nil {
											var __t84 *Constructor_Main_T
											{
												var __t_tag_35 uint32 = ((v1_1).V3).V0
												_ = __t_tag_35
												if uint32(__t_tag_35) == 3558538316 {
													// TAST (Let): __local_var_4_36 shape=Other bindingType=Any
													__local_var_4_36 := (v1_1).V1
													_ = __local_var_4_36
													// TAST (Let): __local_var_5_37 shape=Other bindingType=Any
													__local_var_5_37 := ((v1_1).V3).V1
													_ = __local_var_5_37
													// TAST (Let): __local_var_6_38 shape=Other bindingType=Any
													__local_var_6_38 := ((v1_1).V3).V3
													_ = __local_var_6_38
													// TAST (Let): __local_var_7_39 shape=Other bindingType=(ADT ["Main","Tree"] [])
													__local_var_7_39 := v3_3
													_ = __local_var_7_39
													// TAST (Let): __local_var_8_40 shape=Other bindingType=Any
													__local_var_8_40 := (v1_1).V2
													_ = __local_var_8_40
													// TAST (Let): __local_var_9_41 shape=Other bindingType=Any
													__local_var_9_41 := ((v1_1).V3).V2
													_ = __local_var_9_41
													// TAST (Let): __local_var_10_42 shape=Other bindingType=Int
													__local_var_10_42 := v2_2
													_ = __local_var_10_42
													__t84 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_36, __local_var_8_40, __local_var_5_37}), __local_var_9_41, (&Constructor_Main_T{1, 4250879068, __local_var_6_38, __local_var_10_42, __local_var_7_39})})
													goto end_branch_84
												} else {

												}
											}
											{
												var __t_and_44 bool = false
												if v3_3 != nil {

													var __t_tag_43 uint32 = (v3_3).V0
													_ = __t_tag_43
													__t_and_44 = (uint32(__t_tag_43) == 3558538316)
												}
												if __t_and_44 {
													var __t83 *Constructor_Main_T
													{
														var __t_tag_49 *Constructor_Main_T = (v3_3).V1
														_ = __t_tag_49
														if __t_tag_49 != nil {
															var __t72 *Constructor_Main_T
															{
																var __t_tag_54 uint32 = ((v3_3).V1).V0
																_ = __t_tag_54
																if uint32(__t_tag_54) == 3558538316 {
																	// TAST (Let): __local_var_4_55 shape=Other bindingType=(ADT ["Main","Tree"] [])
																	__local_var_4_55 := v1_1
																	_ = __local_var_4_55
																	// TAST (Let): __local_var_5_56 shape=Other bindingType=Any
																	__local_var_5_56 := ((v3_3).V1).V1
																	_ = __local_var_5_56
																	// TAST (Let): __local_var_6_57 shape=Other bindingType=Any
																	__local_var_6_57 := ((v3_3).V1).V3
																	_ = __local_var_6_57
																	// TAST (Let): __local_var_7_58 shape=Other bindingType=Any
																	__local_var_7_58 := (v3_3).V3
																	_ = __local_var_7_58
																	// TAST (Let): __local_var_8_59 shape=Other bindingType=Int
																	__local_var_8_59 := v2_2
																	_ = __local_var_8_59
																	// TAST (Let): __local_var_9_60 shape=Other bindingType=Any
																	__local_var_9_60 := ((v3_3).V1).V2
																	_ = __local_var_9_60
																	// TAST (Let): __local_var_10_61 shape=Other bindingType=Any
																	__local_var_10_61 := (v3_3).V2
																	_ = __local_var_10_61
																	__t72 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_55, __local_var_8_59, __local_var_5_56}), __local_var_9_60, (&Constructor_Main_T{1, 4250879068, __local_var_6_57, __local_var_10_61, __local_var_7_58})})
																	goto end_branch_72
																} else {

																}
															}
															{
																var __t_tag_62 *Constructor_Main_T = (v3_3).V3
																_ = __t_tag_62
																var __t_and_64 bool = false
																if __t_tag_62 != nil {

																	var __t_tag_63 uint32 = ((v3_3).V3).V0
																	_ = __t_tag_63
																	__t_and_64 = (uint32(__t_tag_63) == 3558538316)
																}
																if __t_and_64 {
																	// TAST (Let): __local_var_4_65 shape=Other bindingType=(ADT ["Main","Tree"] [])
																	__local_var_4_65 := v1_1
																	_ = __local_var_4_65
																	// TAST (Let): __local_var_5_66 shape=Other bindingType=Any
																	__local_var_5_66 := (v3_3).V1
																	_ = __local_var_5_66
																	// TAST (Let): __local_var_6_67 shape=Other bindingType=Any
																	__local_var_6_67 := ((v3_3).V3).V1
																	_ = __local_var_6_67
																	// TAST (Let): __local_var_7_68 shape=Other bindingType=Any
																	__local_var_7_68 := ((v3_3).V3).V3
																	_ = __local_var_7_68
																	// TAST (Let): __local_var_8_69 shape=Other bindingType=Int
																	__local_var_8_69 := v2_2
																	_ = __local_var_8_69
																	// TAST (Let): __local_var_9_70 shape=Other bindingType=Any
																	__local_var_9_70 := (v3_3).V2
																	_ = __local_var_9_70
																	// TAST (Let): __local_var_10_71 shape=Other bindingType=Any
																	__local_var_10_71 := ((v3_3).V3).V2
																	_ = __local_var_10_71
																	__t72 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_65, __local_var_8_69, __local_var_5_66}), __local_var_9_70, (&Constructor_Main_T{1, 4250879068, __local_var_6_67, __local_var_10_71, __local_var_7_68})})
																	goto end_branch_72
																} else {

																}
															}
															{
																// TAST (Let): __local_var_4_50 shape=Other bindingType=(ADT ["Main","Tree"] [])
																__local_var_4_50 := v1_1
																_ = __local_var_4_50
																// TAST (Let): __local_var_5_51 shape=Other bindingType=(ADT ["Main","Tree"] [])
																__local_var_5_51 := v3_3
																_ = __local_var_5_51
																// TAST (Let): __local_var_6_52 shape=Other bindingType=(ADT ["Main","Color"] [])
																__local_var_6_52 := v_0
																_ = __local_var_6_52
																// TAST (Let): __local_var_7_53 shape=Other bindingType=Int
																__local_var_7_53 := v2_2
																_ = __local_var_7_53
																__t72 = (&Constructor_Main_T{1, __local_var_6_52, __local_var_4_50, __local_var_7_53, __local_var_5_51})
															}
														end_branch_72:
															__t83 = __t72
															goto end_branch_83
														} else {

														}
													}
													{
														var __t_tag_73 *Constructor_Main_T = (v3_3).V3
														_ = __t_tag_73
														var __t_and_75 bool = false
														if __t_tag_73 != nil {

															var __t_tag_74 uint32 = ((v3_3).V3).V0
															_ = __t_tag_74
															__t_and_75 = (uint32(__t_tag_74) == 3558538316)
														}
														if __t_and_75 {
															// TAST (Let): __local_var_4_76 shape=Other bindingType=(ADT ["Main","Tree"] [])
															__local_var_4_76 := v1_1
															_ = __local_var_4_76
															// TAST (Let): __local_var_5_77 shape=Other bindingType=Any
															__local_var_5_77 := (v3_3).V1
															_ = __local_var_5_77
															// TAST (Let): __local_var_6_78 shape=Other bindingType=Any
															__local_var_6_78 := ((v3_3).V3).V1
															_ = __local_var_6_78
															// TAST (Let): __local_var_7_79 shape=Other bindingType=Any
															__local_var_7_79 := ((v3_3).V3).V3
															_ = __local_var_7_79
															// TAST (Let): __local_var_8_80 shape=Other bindingType=Int
															__local_var_8_80 := v2_2
															_ = __local_var_8_80
															// TAST (Let): __local_var_9_81 shape=Other bindingType=Any
															__local_var_9_81 := (v3_3).V2
															_ = __local_var_9_81
															// TAST (Let): __local_var_10_82 shape=Other bindingType=Any
															__local_var_10_82 := ((v3_3).V3).V2
															_ = __local_var_10_82
															__t83 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_76, __local_var_8_80, __local_var_5_77}), __local_var_9_81, (&Constructor_Main_T{1, 4250879068, __local_var_6_78, __local_var_10_82, __local_var_7_79})})
															goto end_branch_83
														} else {

														}
													}
													{
														// TAST (Let): __local_var_4_45 shape=Other bindingType=(ADT ["Main","Tree"] [])
														__local_var_4_45 := v1_1
														_ = __local_var_4_45
														// TAST (Let): __local_var_5_46 shape=Other bindingType=(ADT ["Main","Tree"] [])
														__local_var_5_46 := v3_3
														_ = __local_var_5_46
														// TAST (Let): __local_var_6_47 shape=Other bindingType=(ADT ["Main","Color"] [])
														__local_var_6_47 := v_0
														_ = __local_var_6_47
														// TAST (Let): __local_var_7_48 shape=Other bindingType=Int
														__local_var_7_48 := v2_2
														_ = __local_var_7_48
														__t83 = (&Constructor_Main_T{1, __local_var_6_47, __local_var_4_45, __local_var_7_48, __local_var_5_46})
													}
												end_branch_83:
													__t84 = __t83
													goto end_branch_84
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_31 shape=Other bindingType=(ADT ["Main","Tree"] [])
												__local_var_4_31 := v1_1
												_ = __local_var_4_31
												// TAST (Let): __local_var_5_32 shape=Other bindingType=(ADT ["Main","Tree"] [])
												__local_var_5_32 := v3_3
												_ = __local_var_5_32
												// TAST (Let): __local_var_6_33 shape=Other bindingType=(ADT ["Main","Color"] [])
												__local_var_6_33 := v_0
												_ = __local_var_6_33
												// TAST (Let): __local_var_7_34 shape=Other bindingType=Int
												__local_var_7_34 := v2_2
												_ = __local_var_7_34
												__t84 = (&Constructor_Main_T{1, __local_var_6_33, __local_var_4_31, __local_var_7_34, __local_var_5_32})
											}
										end_branch_84:
											__t126 = __t84
											goto end_branch_126
										} else {

										}
									}
									{
										var __t_and_86 bool = false
										if v3_3 != nil {

											var __t_tag_85 uint32 = (v3_3).V0
											_ = __t_tag_85
											__t_and_86 = (uint32(__t_tag_85) == 3558538316)
										}
										if __t_and_86 {
											var __t125 *Constructor_Main_T
											{
												var __t_tag_91 *Constructor_Main_T = (v3_3).V1
												_ = __t_tag_91
												if __t_tag_91 != nil {
													var __t114 *Constructor_Main_T
													{
														var __t_tag_96 uint32 = ((v3_3).V1).V0
														_ = __t_tag_96
														if uint32(__t_tag_96) == 3558538316 {
															// TAST (Let): __local_var_4_97 shape=Other bindingType=(ADT ["Main","Tree"] [])
															__local_var_4_97 := v1_1
															_ = __local_var_4_97
															// TAST (Let): __local_var_5_98 shape=Other bindingType=Any
															__local_var_5_98 := ((v3_3).V1).V1
															_ = __local_var_5_98
															// TAST (Let): __local_var_6_99 shape=Other bindingType=Any
															__local_var_6_99 := ((v3_3).V1).V3
															_ = __local_var_6_99
															// TAST (Let): __local_var_7_100 shape=Other bindingType=Any
															__local_var_7_100 := (v3_3).V3
															_ = __local_var_7_100
															// TAST (Let): __local_var_8_101 shape=Other bindingType=Int
															__local_var_8_101 := v2_2
															_ = __local_var_8_101
															// TAST (Let): __local_var_9_102 shape=Other bindingType=Any
															__local_var_9_102 := ((v3_3).V1).V2
															_ = __local_var_9_102
															// TAST (Let): __local_var_10_103 shape=Other bindingType=Any
															__local_var_10_103 := (v3_3).V2
															_ = __local_var_10_103
															__t114 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_97, __local_var_8_101, __local_var_5_98}), __local_var_9_102, (&Constructor_Main_T{1, 4250879068, __local_var_6_99, __local_var_10_103, __local_var_7_100})})
															goto end_branch_114
														} else {

														}
													}
													{
														var __t_tag_104 *Constructor_Main_T = (v3_3).V3
														_ = __t_tag_104
														var __t_and_106 bool = false
														if __t_tag_104 != nil {

															var __t_tag_105 uint32 = ((v3_3).V3).V0
															_ = __t_tag_105
															__t_and_106 = (uint32(__t_tag_105) == 3558538316)
														}
														if __t_and_106 {
															// TAST (Let): __local_var_4_107 shape=Other bindingType=(ADT ["Main","Tree"] [])
															__local_var_4_107 := v1_1
															_ = __local_var_4_107
															// TAST (Let): __local_var_5_108 shape=Other bindingType=Any
															__local_var_5_108 := (v3_3).V1
															_ = __local_var_5_108
															// TAST (Let): __local_var_6_109 shape=Other bindingType=Any
															__local_var_6_109 := ((v3_3).V3).V1
															_ = __local_var_6_109
															// TAST (Let): __local_var_7_110 shape=Other bindingType=Any
															__local_var_7_110 := ((v3_3).V3).V3
															_ = __local_var_7_110
															// TAST (Let): __local_var_8_111 shape=Other bindingType=Int
															__local_var_8_111 := v2_2
															_ = __local_var_8_111
															// TAST (Let): __local_var_9_112 shape=Other bindingType=Any
															__local_var_9_112 := (v3_3).V2
															_ = __local_var_9_112
															// TAST (Let): __local_var_10_113 shape=Other bindingType=Any
															__local_var_10_113 := ((v3_3).V3).V2
															_ = __local_var_10_113
															__t114 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_107, __local_var_8_111, __local_var_5_108}), __local_var_9_112, (&Constructor_Main_T{1, 4250879068, __local_var_6_109, __local_var_10_113, __local_var_7_110})})
															goto end_branch_114
														} else {

														}
													}
													{
														// TAST (Let): __local_var_4_92 shape=Other bindingType=(ADT ["Main","Tree"] [])
														__local_var_4_92 := v1_1
														_ = __local_var_4_92
														// TAST (Let): __local_var_5_93 shape=Other bindingType=(ADT ["Main","Tree"] [])
														__local_var_5_93 := v3_3
														_ = __local_var_5_93
														// TAST (Let): __local_var_6_94 shape=Other bindingType=(ADT ["Main","Color"] [])
														__local_var_6_94 := v_0
														_ = __local_var_6_94
														// TAST (Let): __local_var_7_95 shape=Other bindingType=Int
														__local_var_7_95 := v2_2
														_ = __local_var_7_95
														__t114 = (&Constructor_Main_T{1, __local_var_6_94, __local_var_4_92, __local_var_7_95, __local_var_5_93})
													}
												end_branch_114:
													__t125 = __t114
													goto end_branch_125
												} else {

												}
											}
											{
												var __t_tag_115 *Constructor_Main_T = (v3_3).V3
												_ = __t_tag_115
												var __t_and_117 bool = false
												if __t_tag_115 != nil {

													var __t_tag_116 uint32 = ((v3_3).V3).V0
													_ = __t_tag_116
													__t_and_117 = (uint32(__t_tag_116) == 3558538316)
												}
												if __t_and_117 {
													// TAST (Let): __local_var_4_118 shape=Other bindingType=(ADT ["Main","Tree"] [])
													__local_var_4_118 := v1_1
													_ = __local_var_4_118
													// TAST (Let): __local_var_5_119 shape=Other bindingType=Any
													__local_var_5_119 := (v3_3).V1
													_ = __local_var_5_119
													// TAST (Let): __local_var_6_120 shape=Other bindingType=Any
													__local_var_6_120 := ((v3_3).V3).V1
													_ = __local_var_6_120
													// TAST (Let): __local_var_7_121 shape=Other bindingType=Any
													__local_var_7_121 := ((v3_3).V3).V3
													_ = __local_var_7_121
													// TAST (Let): __local_var_8_122 shape=Other bindingType=Int
													__local_var_8_122 := v2_2
													_ = __local_var_8_122
													// TAST (Let): __local_var_9_123 shape=Other bindingType=Any
													__local_var_9_123 := (v3_3).V2
													_ = __local_var_9_123
													// TAST (Let): __local_var_10_124 shape=Other bindingType=Any
													__local_var_10_124 := ((v3_3).V3).V2
													_ = __local_var_10_124
													__t125 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_118, __local_var_8_122, __local_var_5_119}), __local_var_9_123, (&Constructor_Main_T{1, 4250879068, __local_var_6_120, __local_var_10_124, __local_var_7_121})})
													goto end_branch_125
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_87 shape=Other bindingType=(ADT ["Main","Tree"] [])
												__local_var_4_87 := v1_1
												_ = __local_var_4_87
												// TAST (Let): __local_var_5_88 shape=Other bindingType=(ADT ["Main","Tree"] [])
												__local_var_5_88 := v3_3
												_ = __local_var_5_88
												// TAST (Let): __local_var_6_89 shape=Other bindingType=(ADT ["Main","Color"] [])
												__local_var_6_89 := v_0
												_ = __local_var_6_89
												// TAST (Let): __local_var_7_90 shape=Other bindingType=Int
												__local_var_7_90 := v2_2
												_ = __local_var_7_90
												__t125 = (&Constructor_Main_T{1, __local_var_6_89, __local_var_4_87, __local_var_7_90, __local_var_5_88})
											}
										end_branch_125:
											__t126 = __t125
											goto end_branch_126
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_18 shape=Other bindingType=(ADT ["Main","Tree"] [])
										__local_var_4_18 := v1_1
										_ = __local_var_4_18
										// TAST (Let): __local_var_5_19 shape=Other bindingType=(ADT ["Main","Tree"] [])
										__local_var_5_19 := v3_3
										_ = __local_var_5_19
										// TAST (Let): __local_var_6_20 shape=Other bindingType=(ADT ["Main","Color"] [])
										__local_var_6_20 := v_0
										_ = __local_var_6_20
										// TAST (Let): __local_var_7_21 shape=Other bindingType=Int
										__local_var_7_21 := v2_2
										_ = __local_var_7_21
										__t126 = (&Constructor_Main_T{1, __local_var_6_20, __local_var_4_18, __local_var_7_21, __local_var_5_19})
									}
								end_branch_126:
									__t223 = __t126
									goto end_branch_223
								} else {

								}
							}
							{
								var __t_tag_127 *Constructor_Main_T = (v1_1).V3
								_ = __t_tag_127
								if __t_tag_127 != nil {
									var __t181 *Constructor_Main_T
									{
										var __t_tag_132 uint32 = ((v1_1).V3).V0
										_ = __t_tag_132
										if uint32(__t_tag_132) == 3558538316 {
											// TAST (Let): __local_var_4_133 shape=Other bindingType=Any
											__local_var_4_133 := (v1_1).V1
											_ = __local_var_4_133
											// TAST (Let): __local_var_5_134 shape=Other bindingType=Any
											__local_var_5_134 := ((v1_1).V3).V1
											_ = __local_var_5_134
											// TAST (Let): __local_var_6_135 shape=Other bindingType=Any
											__local_var_6_135 := ((v1_1).V3).V3
											_ = __local_var_6_135
											// TAST (Let): __local_var_7_136 shape=Other bindingType=(ADT ["Main","Tree"] [])
											__local_var_7_136 := v3_3
											_ = __local_var_7_136
											// TAST (Let): __local_var_8_137 shape=Other bindingType=Any
											__local_var_8_137 := (v1_1).V2
											_ = __local_var_8_137
											// TAST (Let): __local_var_9_138 shape=Other bindingType=Any
											__local_var_9_138 := ((v1_1).V3).V2
											_ = __local_var_9_138
											// TAST (Let): __local_var_10_139 shape=Other bindingType=Int
											__local_var_10_139 := v2_2
											_ = __local_var_10_139
											__t181 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_133, __local_var_8_137, __local_var_5_134}), __local_var_9_138, (&Constructor_Main_T{1, 4250879068, __local_var_6_135, __local_var_10_139, __local_var_7_136})})
											goto end_branch_181
										} else {

										}
									}
									{
										var __t_and_141 bool = false
										if v3_3 != nil {

											var __t_tag_140 uint32 = (v3_3).V0
											_ = __t_tag_140
											__t_and_141 = (uint32(__t_tag_140) == 3558538316)
										}
										if __t_and_141 {
											var __t180 *Constructor_Main_T
											{
												var __t_tag_146 *Constructor_Main_T = (v3_3).V1
												_ = __t_tag_146
												if __t_tag_146 != nil {
													var __t169 *Constructor_Main_T
													{
														var __t_tag_151 uint32 = ((v3_3).V1).V0
														_ = __t_tag_151
														if uint32(__t_tag_151) == 3558538316 {
															// TAST (Let): __local_var_4_152 shape=Other bindingType=(ADT ["Main","Tree"] [])
															__local_var_4_152 := v1_1
															_ = __local_var_4_152
															// TAST (Let): __local_var_5_153 shape=Other bindingType=Any
															__local_var_5_153 := ((v3_3).V1).V1
															_ = __local_var_5_153
															// TAST (Let): __local_var_6_154 shape=Other bindingType=Any
															__local_var_6_154 := ((v3_3).V1).V3
															_ = __local_var_6_154
															// TAST (Let): __local_var_7_155 shape=Other bindingType=Any
															__local_var_7_155 := (v3_3).V3
															_ = __local_var_7_155
															// TAST (Let): __local_var_8_156 shape=Other bindingType=Int
															__local_var_8_156 := v2_2
															_ = __local_var_8_156
															// TAST (Let): __local_var_9_157 shape=Other bindingType=Any
															__local_var_9_157 := ((v3_3).V1).V2
															_ = __local_var_9_157
															// TAST (Let): __local_var_10_158 shape=Other bindingType=Any
															__local_var_10_158 := (v3_3).V2
															_ = __local_var_10_158
															__t169 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_152, __local_var_8_156, __local_var_5_153}), __local_var_9_157, (&Constructor_Main_T{1, 4250879068, __local_var_6_154, __local_var_10_158, __local_var_7_155})})
															goto end_branch_169
														} else {

														}
													}
													{
														var __t_tag_159 *Constructor_Main_T = (v3_3).V3
														_ = __t_tag_159
														var __t_and_161 bool = false
														if __t_tag_159 != nil {

															var __t_tag_160 uint32 = ((v3_3).V3).V0
															_ = __t_tag_160
															__t_and_161 = (uint32(__t_tag_160) == 3558538316)
														}
														if __t_and_161 {
															// TAST (Let): __local_var_4_162 shape=Other bindingType=(ADT ["Main","Tree"] [])
															__local_var_4_162 := v1_1
															_ = __local_var_4_162
															// TAST (Let): __local_var_5_163 shape=Other bindingType=Any
															__local_var_5_163 := (v3_3).V1
															_ = __local_var_5_163
															// TAST (Let): __local_var_6_164 shape=Other bindingType=Any
															__local_var_6_164 := ((v3_3).V3).V1
															_ = __local_var_6_164
															// TAST (Let): __local_var_7_165 shape=Other bindingType=Any
															__local_var_7_165 := ((v3_3).V3).V3
															_ = __local_var_7_165
															// TAST (Let): __local_var_8_166 shape=Other bindingType=Int
															__local_var_8_166 := v2_2
															_ = __local_var_8_166
															// TAST (Let): __local_var_9_167 shape=Other bindingType=Any
															__local_var_9_167 := (v3_3).V2
															_ = __local_var_9_167
															// TAST (Let): __local_var_10_168 shape=Other bindingType=Any
															__local_var_10_168 := ((v3_3).V3).V2
															_ = __local_var_10_168
															__t169 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_162, __local_var_8_166, __local_var_5_163}), __local_var_9_167, (&Constructor_Main_T{1, 4250879068, __local_var_6_164, __local_var_10_168, __local_var_7_165})})
															goto end_branch_169
														} else {

														}
													}
													{
														// TAST (Let): __local_var_4_147 shape=Other bindingType=(ADT ["Main","Tree"] [])
														__local_var_4_147 := v1_1
														_ = __local_var_4_147
														// TAST (Let): __local_var_5_148 shape=Other bindingType=(ADT ["Main","Tree"] [])
														__local_var_5_148 := v3_3
														_ = __local_var_5_148
														// TAST (Let): __local_var_6_149 shape=Other bindingType=(ADT ["Main","Color"] [])
														__local_var_6_149 := v_0
														_ = __local_var_6_149
														// TAST (Let): __local_var_7_150 shape=Other bindingType=Int
														__local_var_7_150 := v2_2
														_ = __local_var_7_150
														__t169 = (&Constructor_Main_T{1, __local_var_6_149, __local_var_4_147, __local_var_7_150, __local_var_5_148})
													}
												end_branch_169:
													__t180 = __t169
													goto end_branch_180
												} else {

												}
											}
											{
												var __t_tag_170 *Constructor_Main_T = (v3_3).V3
												_ = __t_tag_170
												var __t_and_172 bool = false
												if __t_tag_170 != nil {

													var __t_tag_171 uint32 = ((v3_3).V3).V0
													_ = __t_tag_171
													__t_and_172 = (uint32(__t_tag_171) == 3558538316)
												}
												if __t_and_172 {
													// TAST (Let): __local_var_4_173 shape=Other bindingType=(ADT ["Main","Tree"] [])
													__local_var_4_173 := v1_1
													_ = __local_var_4_173
													// TAST (Let): __local_var_5_174 shape=Other bindingType=Any
													__local_var_5_174 := (v3_3).V1
													_ = __local_var_5_174
													// TAST (Let): __local_var_6_175 shape=Other bindingType=Any
													__local_var_6_175 := ((v3_3).V3).V1
													_ = __local_var_6_175
													// TAST (Let): __local_var_7_176 shape=Other bindingType=Any
													__local_var_7_176 := ((v3_3).V3).V3
													_ = __local_var_7_176
													// TAST (Let): __local_var_8_177 shape=Other bindingType=Int
													__local_var_8_177 := v2_2
													_ = __local_var_8_177
													// TAST (Let): __local_var_9_178 shape=Other bindingType=Any
													__local_var_9_178 := (v3_3).V2
													_ = __local_var_9_178
													// TAST (Let): __local_var_10_179 shape=Other bindingType=Any
													__local_var_10_179 := ((v3_3).V3).V2
													_ = __local_var_10_179
													__t180 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_173, __local_var_8_177, __local_var_5_174}), __local_var_9_178, (&Constructor_Main_T{1, 4250879068, __local_var_6_175, __local_var_10_179, __local_var_7_176})})
													goto end_branch_180
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_142 shape=Other bindingType=(ADT ["Main","Tree"] [])
												__local_var_4_142 := v1_1
												_ = __local_var_4_142
												// TAST (Let): __local_var_5_143 shape=Other bindingType=(ADT ["Main","Tree"] [])
												__local_var_5_143 := v3_3
												_ = __local_var_5_143
												// TAST (Let): __local_var_6_144 shape=Other bindingType=(ADT ["Main","Color"] [])
												__local_var_6_144 := v_0
												_ = __local_var_6_144
												// TAST (Let): __local_var_7_145 shape=Other bindingType=Int
												__local_var_7_145 := v2_2
												_ = __local_var_7_145
												__t180 = (&Constructor_Main_T{1, __local_var_6_144, __local_var_4_142, __local_var_7_145, __local_var_5_143})
											}
										end_branch_180:
											__t181 = __t180
											goto end_branch_181
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_128 shape=Other bindingType=(ADT ["Main","Tree"] [])
										__local_var_4_128 := v1_1
										_ = __local_var_4_128
										// TAST (Let): __local_var_5_129 shape=Other bindingType=(ADT ["Main","Tree"] [])
										__local_var_5_129 := v3_3
										_ = __local_var_5_129
										// TAST (Let): __local_var_6_130 shape=Other bindingType=(ADT ["Main","Color"] [])
										__local_var_6_130 := v_0
										_ = __local_var_6_130
										// TAST (Let): __local_var_7_131 shape=Other bindingType=Int
										__local_var_7_131 := v2_2
										_ = __local_var_7_131
										__t181 = (&Constructor_Main_T{1, __local_var_6_130, __local_var_4_128, __local_var_7_131, __local_var_5_129})
									}
								end_branch_181:
									__t223 = __t181
									goto end_branch_223
								} else {

								}
							}
							{
								var __t_and_183 bool = false
								if v3_3 != nil {

									var __t_tag_182 uint32 = (v3_3).V0
									_ = __t_tag_182
									__t_and_183 = (uint32(__t_tag_182) == 3558538316)
								}
								if __t_and_183 {
									var __t222 *Constructor_Main_T
									{
										var __t_tag_188 *Constructor_Main_T = (v3_3).V1
										_ = __t_tag_188
										if __t_tag_188 != nil {
											var __t211 *Constructor_Main_T
											{
												var __t_tag_193 uint32 = ((v3_3).V1).V0
												_ = __t_tag_193
												if uint32(__t_tag_193) == 3558538316 {
													// TAST (Let): __local_var_4_194 shape=Other bindingType=(ADT ["Main","Tree"] [])
													__local_var_4_194 := v1_1
													_ = __local_var_4_194
													// TAST (Let): __local_var_5_195 shape=Other bindingType=Any
													__local_var_5_195 := ((v3_3).V1).V1
													_ = __local_var_5_195
													// TAST (Let): __local_var_6_196 shape=Other bindingType=Any
													__local_var_6_196 := ((v3_3).V1).V3
													_ = __local_var_6_196
													// TAST (Let): __local_var_7_197 shape=Other bindingType=Any
													__local_var_7_197 := (v3_3).V3
													_ = __local_var_7_197
													// TAST (Let): __local_var_8_198 shape=Other bindingType=Int
													__local_var_8_198 := v2_2
													_ = __local_var_8_198
													// TAST (Let): __local_var_9_199 shape=Other bindingType=Any
													__local_var_9_199 := ((v3_3).V1).V2
													_ = __local_var_9_199
													// TAST (Let): __local_var_10_200 shape=Other bindingType=Any
													__local_var_10_200 := (v3_3).V2
													_ = __local_var_10_200
													__t211 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_194, __local_var_8_198, __local_var_5_195}), __local_var_9_199, (&Constructor_Main_T{1, 4250879068, __local_var_6_196, __local_var_10_200, __local_var_7_197})})
													goto end_branch_211
												} else {

												}
											}
											{
												var __t_tag_201 *Constructor_Main_T = (v3_3).V3
												_ = __t_tag_201
												var __t_and_203 bool = false
												if __t_tag_201 != nil {

													var __t_tag_202 uint32 = ((v3_3).V3).V0
													_ = __t_tag_202
													__t_and_203 = (uint32(__t_tag_202) == 3558538316)
												}
												if __t_and_203 {
													// TAST (Let): __local_var_4_204 shape=Other bindingType=(ADT ["Main","Tree"] [])
													__local_var_4_204 := v1_1
													_ = __local_var_4_204
													// TAST (Let): __local_var_5_205 shape=Other bindingType=Any
													__local_var_5_205 := (v3_3).V1
													_ = __local_var_5_205
													// TAST (Let): __local_var_6_206 shape=Other bindingType=Any
													__local_var_6_206 := ((v3_3).V3).V1
													_ = __local_var_6_206
													// TAST (Let): __local_var_7_207 shape=Other bindingType=Any
													__local_var_7_207 := ((v3_3).V3).V3
													_ = __local_var_7_207
													// TAST (Let): __local_var_8_208 shape=Other bindingType=Int
													__local_var_8_208 := v2_2
													_ = __local_var_8_208
													// TAST (Let): __local_var_9_209 shape=Other bindingType=Any
													__local_var_9_209 := (v3_3).V2
													_ = __local_var_9_209
													// TAST (Let): __local_var_10_210 shape=Other bindingType=Any
													__local_var_10_210 := ((v3_3).V3).V2
													_ = __local_var_10_210
													__t211 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_204, __local_var_8_208, __local_var_5_205}), __local_var_9_209, (&Constructor_Main_T{1, 4250879068, __local_var_6_206, __local_var_10_210, __local_var_7_207})})
													goto end_branch_211
												} else {

												}
											}
											{
												// TAST (Let): __local_var_4_189 shape=Other bindingType=(ADT ["Main","Tree"] [])
												__local_var_4_189 := v1_1
												_ = __local_var_4_189
												// TAST (Let): __local_var_5_190 shape=Other bindingType=(ADT ["Main","Tree"] [])
												__local_var_5_190 := v3_3
												_ = __local_var_5_190
												// TAST (Let): __local_var_6_191 shape=Other bindingType=(ADT ["Main","Color"] [])
												__local_var_6_191 := v_0
												_ = __local_var_6_191
												// TAST (Let): __local_var_7_192 shape=Other bindingType=Int
												__local_var_7_192 := v2_2
												_ = __local_var_7_192
												__t211 = (&Constructor_Main_T{1, __local_var_6_191, __local_var_4_189, __local_var_7_192, __local_var_5_190})
											}
										end_branch_211:
											__t222 = __t211
											goto end_branch_222
										} else {

										}
									}
									{
										var __t_tag_212 *Constructor_Main_T = (v3_3).V3
										_ = __t_tag_212
										var __t_and_214 bool = false
										if __t_tag_212 != nil {

											var __t_tag_213 uint32 = ((v3_3).V3).V0
											_ = __t_tag_213
											__t_and_214 = (uint32(__t_tag_213) == 3558538316)
										}
										if __t_and_214 {
											// TAST (Let): __local_var_4_215 shape=Other bindingType=(ADT ["Main","Tree"] [])
											__local_var_4_215 := v1_1
											_ = __local_var_4_215
											// TAST (Let): __local_var_5_216 shape=Other bindingType=Any
											__local_var_5_216 := (v3_3).V1
											_ = __local_var_5_216
											// TAST (Let): __local_var_6_217 shape=Other bindingType=Any
											__local_var_6_217 := ((v3_3).V3).V1
											_ = __local_var_6_217
											// TAST (Let): __local_var_7_218 shape=Other bindingType=Any
											__local_var_7_218 := ((v3_3).V3).V3
											_ = __local_var_7_218
											// TAST (Let): __local_var_8_219 shape=Other bindingType=Int
											__local_var_8_219 := v2_2
											_ = __local_var_8_219
											// TAST (Let): __local_var_9_220 shape=Other bindingType=Any
											__local_var_9_220 := (v3_3).V2
											_ = __local_var_9_220
											// TAST (Let): __local_var_10_221 shape=Other bindingType=Any
											__local_var_10_221 := ((v3_3).V3).V2
											_ = __local_var_10_221
											__t222 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_215, __local_var_8_219, __local_var_5_216}), __local_var_9_220, (&Constructor_Main_T{1, 4250879068, __local_var_6_217, __local_var_10_221, __local_var_7_218})})
											goto end_branch_222
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_184 shape=Other bindingType=(ADT ["Main","Tree"] [])
										__local_var_4_184 := v1_1
										_ = __local_var_4_184
										// TAST (Let): __local_var_5_185 shape=Other bindingType=(ADT ["Main","Tree"] [])
										__local_var_5_185 := v3_3
										_ = __local_var_5_185
										// TAST (Let): __local_var_6_186 shape=Other bindingType=(ADT ["Main","Color"] [])
										__local_var_6_186 := v_0
										_ = __local_var_6_186
										// TAST (Let): __local_var_7_187 shape=Other bindingType=Int
										__local_var_7_187 := v2_2
										_ = __local_var_7_187
										__t222 = (&Constructor_Main_T{1, __local_var_6_186, __local_var_4_184, __local_var_7_187, __local_var_5_185})
									}
								end_branch_222:
									__t223 = __t222
									goto end_branch_223
								} else {

								}
							}
							{
								// TAST (Let): __local_var_4_13 shape=Other bindingType=(ADT ["Main","Tree"] [])
								__local_var_4_13 := v1_1
								_ = __local_var_4_13
								// TAST (Let): __local_var_5_14 shape=Other bindingType=(ADT ["Main","Tree"] [])
								__local_var_5_14 := v3_3
								_ = __local_var_5_14
								// TAST (Let): __local_var_6_15 shape=Other bindingType=(ADT ["Main","Color"] [])
								__local_var_6_15 := v_0
								_ = __local_var_6_15
								// TAST (Let): __local_var_7_16 shape=Other bindingType=Int
								__local_var_7_16 := v2_2
								_ = __local_var_7_16
								__t223 = (&Constructor_Main_T{1, __local_var_6_15, __local_var_4_13, __local_var_7_16, __local_var_5_14})
							}
						end_branch_223:
							__t265 = __t223
							goto end_branch_265
						} else {

						}
					}
					{
						var __t_and_225 bool = false
						if v3_3 != nil {

							var __t_tag_224 uint32 = (v3_3).V0
							_ = __t_tag_224
							__t_and_225 = (uint32(__t_tag_224) == 3558538316)
						}
						if __t_and_225 {
							var __t264 *Constructor_Main_T
							{
								var __t_tag_230 *Constructor_Main_T = (v3_3).V1
								_ = __t_tag_230
								if __t_tag_230 != nil {
									var __t253 *Constructor_Main_T
									{
										var __t_tag_235 uint32 = ((v3_3).V1).V0
										_ = __t_tag_235
										if uint32(__t_tag_235) == 3558538316 {
											// TAST (Let): __local_var_4_236 shape=Other bindingType=(ADT ["Main","Tree"] [])
											__local_var_4_236 := v1_1
											_ = __local_var_4_236
											// TAST (Let): __local_var_5_237 shape=Other bindingType=Any
											__local_var_5_237 := ((v3_3).V1).V1
											_ = __local_var_5_237
											// TAST (Let): __local_var_6_238 shape=Other bindingType=Any
											__local_var_6_238 := ((v3_3).V1).V3
											_ = __local_var_6_238
											// TAST (Let): __local_var_7_239 shape=Other bindingType=Any
											__local_var_7_239 := (v3_3).V3
											_ = __local_var_7_239
											// TAST (Let): __local_var_8_240 shape=Other bindingType=Int
											__local_var_8_240 := v2_2
											_ = __local_var_8_240
											// TAST (Let): __local_var_9_241 shape=Other bindingType=Any
											__local_var_9_241 := ((v3_3).V1).V2
											_ = __local_var_9_241
											// TAST (Let): __local_var_10_242 shape=Other bindingType=Any
											__local_var_10_242 := (v3_3).V2
											_ = __local_var_10_242
											__t253 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_236, __local_var_8_240, __local_var_5_237}), __local_var_9_241, (&Constructor_Main_T{1, 4250879068, __local_var_6_238, __local_var_10_242, __local_var_7_239})})
											goto end_branch_253
										} else {

										}
									}
									{
										var __t_tag_243 *Constructor_Main_T = (v3_3).V3
										_ = __t_tag_243
										var __t_and_245 bool = false
										if __t_tag_243 != nil {

											var __t_tag_244 uint32 = ((v3_3).V3).V0
											_ = __t_tag_244
											__t_and_245 = (uint32(__t_tag_244) == 3558538316)
										}
										if __t_and_245 {
											// TAST (Let): __local_var_4_246 shape=Other bindingType=(ADT ["Main","Tree"] [])
											__local_var_4_246 := v1_1
											_ = __local_var_4_246
											// TAST (Let): __local_var_5_247 shape=Other bindingType=Any
											__local_var_5_247 := (v3_3).V1
											_ = __local_var_5_247
											// TAST (Let): __local_var_6_248 shape=Other bindingType=Any
											__local_var_6_248 := ((v3_3).V3).V1
											_ = __local_var_6_248
											// TAST (Let): __local_var_7_249 shape=Other bindingType=Any
											__local_var_7_249 := ((v3_3).V3).V3
											_ = __local_var_7_249
											// TAST (Let): __local_var_8_250 shape=Other bindingType=Int
											__local_var_8_250 := v2_2
											_ = __local_var_8_250
											// TAST (Let): __local_var_9_251 shape=Other bindingType=Any
											__local_var_9_251 := (v3_3).V2
											_ = __local_var_9_251
											// TAST (Let): __local_var_10_252 shape=Other bindingType=Any
											__local_var_10_252 := ((v3_3).V3).V2
											_ = __local_var_10_252
											__t253 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_246, __local_var_8_250, __local_var_5_247}), __local_var_9_251, (&Constructor_Main_T{1, 4250879068, __local_var_6_248, __local_var_10_252, __local_var_7_249})})
											goto end_branch_253
										} else {

										}
									}
									{
										// TAST (Let): __local_var_4_231 shape=Other bindingType=(ADT ["Main","Tree"] [])
										__local_var_4_231 := v1_1
										_ = __local_var_4_231
										// TAST (Let): __local_var_5_232 shape=Other bindingType=(ADT ["Main","Tree"] [])
										__local_var_5_232 := v3_3
										_ = __local_var_5_232
										// TAST (Let): __local_var_6_233 shape=Other bindingType=(ADT ["Main","Color"] [])
										__local_var_6_233 := v_0
										_ = __local_var_6_233
										// TAST (Let): __local_var_7_234 shape=Other bindingType=Int
										__local_var_7_234 := v2_2
										_ = __local_var_7_234
										__t253 = (&Constructor_Main_T{1, __local_var_6_233, __local_var_4_231, __local_var_7_234, __local_var_5_232})
									}
								end_branch_253:
									__t264 = __t253
									goto end_branch_264
								} else {

								}
							}
							{
								var __t_tag_254 *Constructor_Main_T = (v3_3).V3
								_ = __t_tag_254
								var __t_and_256 bool = false
								if __t_tag_254 != nil {

									var __t_tag_255 uint32 = ((v3_3).V3).V0
									_ = __t_tag_255
									__t_and_256 = (uint32(__t_tag_255) == 3558538316)
								}
								if __t_and_256 {
									// TAST (Let): __local_var_4_257 shape=Other bindingType=(ADT ["Main","Tree"] [])
									__local_var_4_257 := v1_1
									_ = __local_var_4_257
									// TAST (Let): __local_var_5_258 shape=Other bindingType=Any
									__local_var_5_258 := (v3_3).V1
									_ = __local_var_5_258
									// TAST (Let): __local_var_6_259 shape=Other bindingType=Any
									__local_var_6_259 := ((v3_3).V3).V1
									_ = __local_var_6_259
									// TAST (Let): __local_var_7_260 shape=Other bindingType=Any
									__local_var_7_260 := ((v3_3).V3).V3
									_ = __local_var_7_260
									// TAST (Let): __local_var_8_261 shape=Other bindingType=Int
									__local_var_8_261 := v2_2
									_ = __local_var_8_261
									// TAST (Let): __local_var_9_262 shape=Other bindingType=Any
									__local_var_9_262 := (v3_3).V2
									_ = __local_var_9_262
									// TAST (Let): __local_var_10_263 shape=Other bindingType=Any
									__local_var_10_263 := ((v3_3).V3).V2
									_ = __local_var_10_263
									__t264 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_257, __local_var_8_261, __local_var_5_258}), __local_var_9_262, (&Constructor_Main_T{1, 4250879068, __local_var_6_259, __local_var_10_263, __local_var_7_260})})
									goto end_branch_264
								} else {

								}
							}
							{
								// TAST (Let): __local_var_4_226 shape=Other bindingType=(ADT ["Main","Tree"] [])
								__local_var_4_226 := v1_1
								_ = __local_var_4_226
								// TAST (Let): __local_var_5_227 shape=Other bindingType=(ADT ["Main","Tree"] [])
								__local_var_5_227 := v3_3
								_ = __local_var_5_227
								// TAST (Let): __local_var_6_228 shape=Other bindingType=(ADT ["Main","Color"] [])
								__local_var_6_228 := v_0
								_ = __local_var_6_228
								// TAST (Let): __local_var_7_229 shape=Other bindingType=Int
								__local_var_7_229 := v2_2
								_ = __local_var_7_229
								__t264 = (&Constructor_Main_T{1, __local_var_6_228, __local_var_4_226, __local_var_7_229, __local_var_5_227})
							}
						end_branch_264:
							__t265 = __t264
							goto end_branch_265
						} else {

						}
					}
					{
						// TAST (Let): __local_var_4_8 shape=Other bindingType=(ADT ["Main","Tree"] [])
						__local_var_4_8 := v1_1
						_ = __local_var_4_8
						// TAST (Let): __local_var_5_9 shape=Other bindingType=(ADT ["Main","Tree"] [])
						__local_var_5_9 := v3_3
						_ = __local_var_5_9
						// TAST (Let): __local_var_6_10 shape=Other bindingType=(ADT ["Main","Color"] [])
						__local_var_6_10 := v_0
						_ = __local_var_6_10
						// TAST (Let): __local_var_7_11 shape=Other bindingType=Int
						__local_var_7_11 := v2_2
						_ = __local_var_7_11
						__t265 = (&Constructor_Main_T{1, __local_var_6_10, __local_var_4_8, __local_var_7_11, __local_var_5_9})
					}
				end_branch_265:
					__t307 = __t265
					goto end_branch_307
				} else {

				}
			}
			{
				var __t_and_267 bool = false
				if v3_3 != nil {

					var __t_tag_266 uint32 = (v3_3).V0
					_ = __t_tag_266
					__t_and_267 = (uint32(__t_tag_266) == 3558538316)
				}
				if __t_and_267 {
					var __t306 *Constructor_Main_T
					{
						var __t_tag_272 *Constructor_Main_T = (v3_3).V1
						_ = __t_tag_272
						if __t_tag_272 != nil {
							var __t295 *Constructor_Main_T
							{
								var __t_tag_277 uint32 = ((v3_3).V1).V0
								_ = __t_tag_277
								if uint32(__t_tag_277) == 3558538316 {
									// TAST (Let): __local_var_4_278 shape=Other bindingType=(ADT ["Main","Tree"] [])
									__local_var_4_278 := v1_1
									_ = __local_var_4_278
									// TAST (Let): __local_var_5_279 shape=Other bindingType=Any
									__local_var_5_279 := ((v3_3).V1).V1
									_ = __local_var_5_279
									// TAST (Let): __local_var_6_280 shape=Other bindingType=Any
									__local_var_6_280 := ((v3_3).V1).V3
									_ = __local_var_6_280
									// TAST (Let): __local_var_7_281 shape=Other bindingType=Any
									__local_var_7_281 := (v3_3).V3
									_ = __local_var_7_281
									// TAST (Let): __local_var_8_282 shape=Other bindingType=Int
									__local_var_8_282 := v2_2
									_ = __local_var_8_282
									// TAST (Let): __local_var_9_283 shape=Other bindingType=Any
									__local_var_9_283 := ((v3_3).V1).V2
									_ = __local_var_9_283
									// TAST (Let): __local_var_10_284 shape=Other bindingType=Any
									__local_var_10_284 := (v3_3).V2
									_ = __local_var_10_284
									__t295 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_278, __local_var_8_282, __local_var_5_279}), __local_var_9_283, (&Constructor_Main_T{1, 4250879068, __local_var_6_280, __local_var_10_284, __local_var_7_281})})
									goto end_branch_295
								} else {

								}
							}
							{
								var __t_tag_285 *Constructor_Main_T = (v3_3).V3
								_ = __t_tag_285
								var __t_and_287 bool = false
								if __t_tag_285 != nil {

									var __t_tag_286 uint32 = ((v3_3).V3).V0
									_ = __t_tag_286
									__t_and_287 = (uint32(__t_tag_286) == 3558538316)
								}
								if __t_and_287 {
									// TAST (Let): __local_var_4_288 shape=Other bindingType=(ADT ["Main","Tree"] [])
									__local_var_4_288 := v1_1
									_ = __local_var_4_288
									// TAST (Let): __local_var_5_289 shape=Other bindingType=Any
									__local_var_5_289 := (v3_3).V1
									_ = __local_var_5_289
									// TAST (Let): __local_var_6_290 shape=Other bindingType=Any
									__local_var_6_290 := ((v3_3).V3).V1
									_ = __local_var_6_290
									// TAST (Let): __local_var_7_291 shape=Other bindingType=Any
									__local_var_7_291 := ((v3_3).V3).V3
									_ = __local_var_7_291
									// TAST (Let): __local_var_8_292 shape=Other bindingType=Int
									__local_var_8_292 := v2_2
									_ = __local_var_8_292
									// TAST (Let): __local_var_9_293 shape=Other bindingType=Any
									__local_var_9_293 := (v3_3).V2
									_ = __local_var_9_293
									// TAST (Let): __local_var_10_294 shape=Other bindingType=Any
									__local_var_10_294 := ((v3_3).V3).V2
									_ = __local_var_10_294
									__t295 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_288, __local_var_8_292, __local_var_5_289}), __local_var_9_293, (&Constructor_Main_T{1, 4250879068, __local_var_6_290, __local_var_10_294, __local_var_7_291})})
									goto end_branch_295
								} else {

								}
							}
							{
								// TAST (Let): __local_var_4_273 shape=Other bindingType=(ADT ["Main","Tree"] [])
								__local_var_4_273 := v1_1
								_ = __local_var_4_273
								// TAST (Let): __local_var_5_274 shape=Other bindingType=(ADT ["Main","Tree"] [])
								__local_var_5_274 := v3_3
								_ = __local_var_5_274
								// TAST (Let): __local_var_6_275 shape=Other bindingType=(ADT ["Main","Color"] [])
								__local_var_6_275 := v_0
								_ = __local_var_6_275
								// TAST (Let): __local_var_7_276 shape=Other bindingType=Int
								__local_var_7_276 := v2_2
								_ = __local_var_7_276
								__t295 = (&Constructor_Main_T{1, __local_var_6_275, __local_var_4_273, __local_var_7_276, __local_var_5_274})
							}
						end_branch_295:
							__t306 = __t295
							goto end_branch_306
						} else {

						}
					}
					{
						var __t_tag_296 *Constructor_Main_T = (v3_3).V3
						_ = __t_tag_296
						var __t_and_298 bool = false
						if __t_tag_296 != nil {

							var __t_tag_297 uint32 = ((v3_3).V3).V0
							_ = __t_tag_297
							__t_and_298 = (uint32(__t_tag_297) == 3558538316)
						}
						if __t_and_298 {
							// TAST (Let): __local_var_4_299 shape=Other bindingType=(ADT ["Main","Tree"] [])
							__local_var_4_299 := v1_1
							_ = __local_var_4_299
							// TAST (Let): __local_var_5_300 shape=Other bindingType=Any
							__local_var_5_300 := (v3_3).V1
							_ = __local_var_5_300
							// TAST (Let): __local_var_6_301 shape=Other bindingType=Any
							__local_var_6_301 := ((v3_3).V3).V1
							_ = __local_var_6_301
							// TAST (Let): __local_var_7_302 shape=Other bindingType=Any
							__local_var_7_302 := ((v3_3).V3).V3
							_ = __local_var_7_302
							// TAST (Let): __local_var_8_303 shape=Other bindingType=Int
							__local_var_8_303 := v2_2
							_ = __local_var_8_303
							// TAST (Let): __local_var_9_304 shape=Other bindingType=Any
							__local_var_9_304 := (v3_3).V2
							_ = __local_var_9_304
							// TAST (Let): __local_var_10_305 shape=Other bindingType=Any
							__local_var_10_305 := ((v3_3).V3).V2
							_ = __local_var_10_305
							__t306 = (&Constructor_Main_T{1, 3558538316, (&Constructor_Main_T{1, 4250879068, __local_var_4_299, __local_var_8_303, __local_var_5_300}), __local_var_9_304, (&Constructor_Main_T{1, 4250879068, __local_var_6_301, __local_var_10_305, __local_var_7_302})})
							goto end_branch_306
						} else {

						}
					}
					{
						// TAST (Let): __local_var_4_268 shape=Other bindingType=(ADT ["Main","Tree"] [])
						__local_var_4_268 := v1_1
						_ = __local_var_4_268
						// TAST (Let): __local_var_5_269 shape=Other bindingType=(ADT ["Main","Tree"] [])
						__local_var_5_269 := v3_3
						_ = __local_var_5_269
						// TAST (Let): __local_var_6_270 shape=Other bindingType=(ADT ["Main","Color"] [])
						__local_var_6_270 := v_0
						_ = __local_var_6_270
						// TAST (Let): __local_var_7_271 shape=Other bindingType=Int
						__local_var_7_271 := v2_2
						_ = __local_var_7_271
						__t306 = (&Constructor_Main_T{1, __local_var_6_270, __local_var_4_268, __local_var_7_271, __local_var_5_269})
					}
				end_branch_306:
					__t307 = __t306
					goto end_branch_307
				} else {

				}
			}
			{
				// TAST (Let): __local_var_4_4 shape=Other bindingType=(ADT ["Main","Tree"] [])
				__local_var_4_4 := v1_1
				_ = __local_var_4_4
				// TAST (Let): __local_var_5_5 shape=Other bindingType=(ADT ["Main","Tree"] [])
				__local_var_5_5 := v3_3
				_ = __local_var_5_5
				// TAST (Let): __local_var_6_6 shape=Other bindingType=(ADT ["Main","Color"] [])
				__local_var_6_6 := v_0
				_ = __local_var_6_6
				// TAST (Let): __local_var_7_7 shape=Other bindingType=Int
				__local_var_7_7 := v2_2
				_ = __local_var_7_7
				__t307 = (&Constructor_Main_T{1, __local_var_6_6, __local_var_4_4, __local_var_7_7, __local_var_5_5})
			}
		end_branch_307:
			__t308 = __t307
			goto end_branch_308
		} else {

		}
	}
	{
		// TAST (Let): __local_var_4_0 shape=Other bindingType=(ADT ["Main","Tree"] [])
		__local_var_4_0 := v1_1
		_ = __local_var_4_0
		// TAST (Let): __local_var_5_1 shape=Other bindingType=(ADT ["Main","Tree"] [])
		__local_var_5_1 := v3_3
		_ = __local_var_5_1
		// TAST (Let): __local_var_6_2 shape=Other bindingType=(ADT ["Main","Color"] [])
		__local_var_6_2 := v_0
		_ = __local_var_6_2
		// TAST (Let): __local_var_7_3 shape=Other bindingType=Int
		__local_var_7_3 := v2_2
		_ = __local_var_7_3
		__t308 = (&Constructor_Main_T{1, __local_var_6_2, __local_var_4_0, __local_var_7_3, __local_var_5_1})
	}
end_branch_308:
	return __t308
}

func Call_Main_insert(x_0_loop int64, s_1_loop *Constructor_Main_T) *Constructor_Main_T {
	var x_0 int64 = x_0_loop
	_ = x_0
	var s_1 *Constructor_Main_T = s_1_loop
	_ = s_1
	var ins__3997675643_2_0_0 gopurs_runtime.Value
	_ = ins__3997675643_2_0_0
	var ins__3997675643_2_0_0_cell *gopurs_runtime.Value
	_ = ins__3997675643_2_0_0_cell
	// FALLBACK TCO: isLoop=false len=2
	var ins_2_1_1 gopurs_runtime.Value
	_ = ins_2_1_1
	var ins_2_1_1_cell *gopurs_runtime.Value
	_ = ins_2_1_1_cell
	// FALLBACK TCO: isLoop=false len=2
	ins__3997675643_2_0_0 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t4 *Constructor_Main_T
		{
			if v_3.Type == 9 && v_3.IntVal == 990467018 && v_3.UnsafePtr == nil {
				__t4 = (&Constructor_Main_T{1, 3558538316, (*Constructor_Main_T)(nil), x_0, (*Constructor_Main_T)(nil)})
				goto end_branch_4
			} else {

			}
		}
		{
			if v_3.Type == 9 && v_3.IntVal == 990467018 && v_3.UnsafePtr != nil {
				var __t3 *Constructor_Main_T
				{
					if (x_0) < ((*Constructor_Main_T)(v_3.UnsafePtr).V2) {
						__t3 = Call_Main_balance((*Constructor_Main_T)(v_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Main_T](gopurs_runtime.Apply((*ins__3997675643_2_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((*Constructor_Main_T)(v_3.UnsafePtr).V1)})), (*Constructor_Main_T)(v_3.UnsafePtr).V2, (*Constructor_Main_T)(v_3.UnsafePtr).V3)
						goto end_branch_3
					} else {

					}
				}
				{
					var __t2 *Constructor_Main_T
					{
						if (x_0) > ((*Constructor_Main_T)(v_3.UnsafePtr).V2) {
							__t2 = Call_Main_balance((*Constructor_Main_T)(v_3.UnsafePtr).V0, (*Constructor_Main_T)(v_3.UnsafePtr).V1, (*Constructor_Main_T)(v_3.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Main_T](gopurs_runtime.Apply((*ins__3997675643_2_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((*Constructor_Main_T)(v_3.UnsafePtr).V3)})))
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = (&Constructor_Main_T{1, (*Constructor_Main_T)(v_3.UnsafePtr).V0, (*Constructor_Main_T)(v_3.UnsafePtr).V1, (*Constructor_Main_T)(v_3.UnsafePtr).V2, (*Constructor_Main_T)(v_3.UnsafePtr).V3})
					}
				end_branch_2:
					__t3 = __t2
				}
			end_branch_3:
				__t4 = __t3
				goto end_branch_4
			} else {

			}
		}
		{
			__t4 = func() *Constructor_Main_T { panic("Failed pattern match") }()
		}
	end_branch_4:
		return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(__t4)}
	})
	ins__3997675643_2_0_0_cell = &ins__3997675643_2_0_0
	ins_2_1_1 = gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
		var __t7 *Constructor_Main_T
		{
			if v_3.Type == 9 && v_3.IntVal == 990467018 && v_3.UnsafePtr == nil {
				__t7 = (&Constructor_Main_T{1, 3558538316, (*Constructor_Main_T)(nil), x_0, (*Constructor_Main_T)(nil)})
				goto end_branch_7
			} else {

			}
		}
		{
			if v_3.Type == 9 && v_3.IntVal == 990467018 && v_3.UnsafePtr != nil {
				var __t6 *Constructor_Main_T
				{
					if (x_0) < ((*Constructor_Main_T)(v_3.UnsafePtr).V2) {
						__t6 = Call_Main_balance((*Constructor_Main_T)(v_3.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Main_T](gopurs_runtime.Apply((*ins__3997675643_2_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((*Constructor_Main_T)(v_3.UnsafePtr).V1)})), (*Constructor_Main_T)(v_3.UnsafePtr).V2, (*Constructor_Main_T)(v_3.UnsafePtr).V3)
						goto end_branch_6
					} else {

					}
				}
				{
					var __t5 *Constructor_Main_T
					{
						if (x_0) > ((*Constructor_Main_T)(v_3.UnsafePtr).V2) {
							__t5 = Call_Main_balance((*Constructor_Main_T)(v_3.UnsafePtr).V0, (*Constructor_Main_T)(v_3.UnsafePtr).V1, (*Constructor_Main_T)(v_3.UnsafePtr).V2, gopurs_runtime.CoerceToStruct[Constructor_Main_T](gopurs_runtime.Apply((*ins__3997675643_2_0_0_cell), gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer((*Constructor_Main_T)(v_3.UnsafePtr).V3)})))
							goto end_branch_5
						} else {

						}
					}
					{
						__t5 = (&Constructor_Main_T{1, (*Constructor_Main_T)(v_3.UnsafePtr).V0, (*Constructor_Main_T)(v_3.UnsafePtr).V1, (*Constructor_Main_T)(v_3.UnsafePtr).V2, (*Constructor_Main_T)(v_3.UnsafePtr).V3})
					}
				end_branch_5:
					__t6 = __t5
				}
			end_branch_6:
				__t7 = __t6
				goto end_branch_7
			} else {

			}
		}
		{
			__t7 = func() *Constructor_Main_T { panic("Failed pattern match") }()
		}
	end_branch_7:
		return gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(__t7)}
	})
	ins_2_1_1_cell = &ins_2_1_1
	// TAST (Let): __local_var_3_8 shape=App(Other) bindingType=(ADT ["Main","Tree"] [])
	__local_var_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Main_T](gopurs_runtime.Apply(ins__3997675643_2_0_0, gopurs_runtime.Value{Type: 9, IntVal: 990467018, UnsafePtr: unsafe.Pointer(s_1)}))
	_ = __local_var_3_8
	var __t10 *Constructor_Main_T
	{
		if __local_var_3_8 != nil {
			var __reuse_9 *Constructor_Main_T
			if ((__local_var_3_8) != (nil)) && (((__local_var_3_8).V0) == (4250879068)) {
				__reuse_9 = __local_var_3_8
			} else {
				__reuse_9 = (&Constructor_Main_T{1, 4250879068, (__local_var_3_8).V1, (__local_var_3_8).V2, (__local_var_3_8).V3})
			}
			__t10 = __reuse_9
			goto end_branch_10
		} else {

		}
	}
	{
		if __local_var_3_8 == nil {
			__t10 = (*Constructor_Main_T)(nil)
			goto end_branch_10
		} else {

		}
	}
	{
		__t10 = func() *Constructor_Main_T { panic("Failed pattern match") }()
	}
end_branch_10:
	return __t10
}

func Call_Main_buildTree(v_0_loop int64, v1_1_loop *Constructor_Main_T) *Constructor_Main_T {
buildTree:
	for {
		if false {
			continue buildTree
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 *Constructor_Main_T = v1_1_loop
		_ = v1_1
		var __t0 *Constructor_Main_T
		{
			if (v_0) == (int64(0)) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (int64(1))
			v1_1_loop = Call_Main_insert(v_0, v1_1)
			continue buildTree
			__t0 = func() *Constructor_Main_T { panic("unreachable") }()
		}
	end_branch_0:
		return __t0
	}
}

func Rebox_Main_1386611502_1636311157(in *Constructor_Data_Show_Show[gopurs_runtime.Value]) *Constructor_Data_Show_Show[int64] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[int64]{}
	out.V0 = in.V0
	return out
}

func Rebox_Main_1636311157_1386611502(in *Constructor_Data_Show_Show[int64]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
	out.V0 = in.V0
	return out
}
