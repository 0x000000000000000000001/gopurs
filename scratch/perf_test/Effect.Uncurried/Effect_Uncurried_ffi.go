package Effect_Uncurried

import "gopurs/output/gopurs_runtime"

var MkEffectFn1 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), gopurs_runtime.Value{})
	})
})

var MkEffectFn2 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), gopurs_runtime.Value{})
		})
	})
})

var MkEffectFn3 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), gopurs_runtime.Value{})
			})
		})
	})
})

var MkEffectFn4 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), gopurs_runtime.Value{})
				})
			})
		})
	})
})

var MkEffectFn5 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), gopurs_runtime.Value{})
					})
				})
			})
		})
	})
})

var MkEffectFn6 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), f), gopurs_runtime.Value{})
						})
					})
				})
			})
		})
	})
})

var MkEffectFn7 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(g gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), f), g), gopurs_runtime.Value{})
							})
						})
					})
				})
			})
		})
	})
})

var MkEffectFn8 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(g gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(h gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), f), g), h), gopurs_runtime.Value{})
								})
							})
						})
					})
				})
			})
		})
	})
})

var MkEffectFn9 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(g gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(h gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(i gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), f), g), h), i), gopurs_runtime.Value{})
									})
								})
							})
						})
					})
				})
			})
		})
	})
})

var MkEffectFn10 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(g gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(h gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(i gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(j gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), f), g), h), i), j), gopurs_runtime.Value{})
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})
})

var RunEffectFn1 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(fn, a)
		})
	})
})

var RunEffectFn2 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b)
			})
		})
	})
})

var RunEffectFn3 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c)
				})
			})
		})
	})
})

var RunEffectFn4 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d)
					})
				})
			})
		})
	})
})

var RunEffectFn5 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e)
						})
					})
				})
			})
		})
	})
})

var RunEffectFn6 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), f)
							})
						})
					})
				})
			})
		})
	})
})

var RunEffectFn7 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(g gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), f), g)
								})
							})
						})
					})
				})
			})
		})
	})
})

var RunEffectFn8 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(g gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(h gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), f), g), h)
									})
								})
							})
						})
					})
				})
			})
		})
	})
})

var RunEffectFn9 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(g gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(h gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(i gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), f), g), h), i)
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})
})

var RunEffectFn10 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(e gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(f gopurs_runtime.Value) gopurs_runtime.Value {
							return gopurs_runtime.Func(func(g gopurs_runtime.Value) gopurs_runtime.Value {
								return gopurs_runtime.Func(func(h gopurs_runtime.Value) gopurs_runtime.Value {
									return gopurs_runtime.Func(func(i gopurs_runtime.Value) gopurs_runtime.Value {
										return gopurs_runtime.Func(func(j gopurs_runtime.Value) gopurs_runtime.Value {
											return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
												return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), d), e), f), g), h), i), j)
											})
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})
})

