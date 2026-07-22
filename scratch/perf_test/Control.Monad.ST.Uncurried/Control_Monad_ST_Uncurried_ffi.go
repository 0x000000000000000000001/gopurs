package Control_Monad_ST_Uncurried

import "gopurs/output/gopurs_runtime"

var MkSTFn1 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), gopurs_runtime.Value{})
	})
})

var MkSTFn2 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), gopurs_runtime.Value{})
		})
	})
})

var MkSTFn3 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(c gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b), c), gopurs_runtime.Value{})
			})
		})
	})
})

var MkSTFn4 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var MkSTFn5 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var MkSTFn6 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var MkSTFn7 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var MkSTFn8 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var MkSTFn9 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var MkSTFn10 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var RunSTFn1 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(fn, a)
		})
	})
})

var RunSTFn2 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(b gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Apply(gopurs_runtime.Apply(fn, a), b)
			})
		})
	})
})

var RunSTFn3 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var RunSTFn4 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var RunSTFn5 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var RunSTFn6 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var RunSTFn7 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var RunSTFn8 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var RunSTFn9 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

var RunSTFn10 = gopurs_runtime.Func(func(fn gopurs_runtime.Value) gopurs_runtime.Value {
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

