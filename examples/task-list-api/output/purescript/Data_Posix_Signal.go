package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Posix_Signal_SIGABRT gopurs_runtime.Value
var once_Data_Posix_Signal_SIGABRT sync.Once

func Get_Data_Posix_Signal_SIGABRT() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGABRT.Do(func() {
		cache_Data_Posix_Signal_SIGABRT = gopurs_runtime.Value{Type: 9, IntVal: int64(3640760862), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGABRT
}

var cache_Data_Posix_Signal_SIGALRM gopurs_runtime.Value
var once_Data_Posix_Signal_SIGALRM sync.Once

func Get_Data_Posix_Signal_SIGALRM() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGALRM.Do(func() {
		cache_Data_Posix_Signal_SIGALRM = gopurs_runtime.Value{Type: 9, IntVal: int64(1332741801), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGALRM
}

var cache_Data_Posix_Signal_SIGBUS gopurs_runtime.Value
var once_Data_Posix_Signal_SIGBUS sync.Once

func Get_Data_Posix_Signal_SIGBUS() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGBUS.Do(func() {
		cache_Data_Posix_Signal_SIGBUS = gopurs_runtime.Value{Type: 9, IntVal: int64(1978617951), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGBUS
}

var cache_Data_Posix_Signal_SIGCHLD gopurs_runtime.Value
var once_Data_Posix_Signal_SIGCHLD sync.Once

func Get_Data_Posix_Signal_SIGCHLD() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGCHLD.Do(func() {
		cache_Data_Posix_Signal_SIGCHLD = gopurs_runtime.Value{Type: 9, IntVal: int64(737752792), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGCHLD
}

var cache_Data_Posix_Signal_SIGCLD gopurs_runtime.Value
var once_Data_Posix_Signal_SIGCLD sync.Once

func Get_Data_Posix_Signal_SIGCLD() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGCLD.Do(func() {
		cache_Data_Posix_Signal_SIGCLD = gopurs_runtime.Value{Type: 9, IntVal: int64(305761328), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGCLD
}

var cache_Data_Posix_Signal_SIGCONT gopurs_runtime.Value
var once_Data_Posix_Signal_SIGCONT sync.Once

func Get_Data_Posix_Signal_SIGCONT() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGCONT.Do(func() {
		cache_Data_Posix_Signal_SIGCONT = gopurs_runtime.Value{Type: 9, IntVal: int64(2341479853), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGCONT
}

var cache_Data_Posix_Signal_SIGEMT gopurs_runtime.Value
var once_Data_Posix_Signal_SIGEMT sync.Once

func Get_Data_Posix_Signal_SIGEMT() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGEMT.Do(func() {
		cache_Data_Posix_Signal_SIGEMT = gopurs_runtime.Value{Type: 9, IntVal: int64(2502462151), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGEMT
}

var cache_Data_Posix_Signal_SIGFPE gopurs_runtime.Value
var once_Data_Posix_Signal_SIGFPE sync.Once

func Get_Data_Posix_Signal_SIGFPE() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGFPE.Do(func() {
		cache_Data_Posix_Signal_SIGFPE = gopurs_runtime.Value{Type: 9, IntVal: int64(3424215656), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGFPE
}

var cache_Data_Posix_Signal_SIGHUP gopurs_runtime.Value
var once_Data_Posix_Signal_SIGHUP sync.Once

func Get_Data_Posix_Signal_SIGHUP() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGHUP.Do(func() {
		cache_Data_Posix_Signal_SIGHUP = gopurs_runtime.Value{Type: 9, IntVal: int64(232722006), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGHUP
}

var cache_Data_Posix_Signal_SIGILL gopurs_runtime.Value
var once_Data_Posix_Signal_SIGILL sync.Once

func Get_Data_Posix_Signal_SIGILL() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGILL.Do(func() {
		cache_Data_Posix_Signal_SIGILL = gopurs_runtime.Value{Type: 9, IntVal: int64(3802329906), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGILL
}

var cache_Data_Posix_Signal_SIGINFO gopurs_runtime.Value
var once_Data_Posix_Signal_SIGINFO sync.Once

func Get_Data_Posix_Signal_SIGINFO() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGINFO.Do(func() {
		cache_Data_Posix_Signal_SIGINFO = gopurs_runtime.Value{Type: 9, IntVal: int64(3332358197), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGINFO
}

var cache_Data_Posix_Signal_SIGINT gopurs_runtime.Value
var once_Data_Posix_Signal_SIGINT sync.Once

func Get_Data_Posix_Signal_SIGINT() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGINT.Do(func() {
		cache_Data_Posix_Signal_SIGINT = gopurs_runtime.Value{Type: 9, IntVal: int64(681794280), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGINT
}

var cache_Data_Posix_Signal_SIGIO gopurs_runtime.Value
var once_Data_Posix_Signal_SIGIO sync.Once

func Get_Data_Posix_Signal_SIGIO() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGIO.Do(func() {
		cache_Data_Posix_Signal_SIGIO = gopurs_runtime.Value{Type: 9, IntVal: int64(2445008893), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGIO
}

var cache_Data_Posix_Signal_SIGIOT gopurs_runtime.Value
var once_Data_Posix_Signal_SIGIOT sync.Once

func Get_Data_Posix_Signal_SIGIOT() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGIOT.Do(func() {
		cache_Data_Posix_Signal_SIGIOT = gopurs_runtime.Value{Type: 9, IntVal: int64(186232457), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGIOT
}

var cache_Data_Posix_Signal_SIGKILL gopurs_runtime.Value
var once_Data_Posix_Signal_SIGKILL sync.Once

func Get_Data_Posix_Signal_SIGKILL() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGKILL.Do(func() {
		cache_Data_Posix_Signal_SIGKILL = gopurs_runtime.Value{Type: 9, IntVal: int64(3476794617), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGKILL
}

var cache_Data_Posix_Signal_SIGLOST gopurs_runtime.Value
var once_Data_Posix_Signal_SIGLOST sync.Once

func Get_Data_Posix_Signal_SIGLOST() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGLOST.Do(func() {
		cache_Data_Posix_Signal_SIGLOST = gopurs_runtime.Value{Type: 9, IntVal: int64(651594687), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGLOST
}

var cache_Data_Posix_Signal_SIGPIPE gopurs_runtime.Value
var once_Data_Posix_Signal_SIGPIPE sync.Once

func Get_Data_Posix_Signal_SIGPIPE() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGPIPE.Do(func() {
		cache_Data_Posix_Signal_SIGPIPE = gopurs_runtime.Value{Type: 9, IntVal: int64(2385806231), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGPIPE
}

var cache_Data_Posix_Signal_SIGPOLL gopurs_runtime.Value
var once_Data_Posix_Signal_SIGPOLL sync.Once

func Get_Data_Posix_Signal_SIGPOLL() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGPOLL.Do(func() {
		cache_Data_Posix_Signal_SIGPOLL = gopurs_runtime.Value{Type: 9, IntVal: int64(583450148), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGPOLL
}

var cache_Data_Posix_Signal_SIGPROF gopurs_runtime.Value
var once_Data_Posix_Signal_SIGPROF sync.Once

func Get_Data_Posix_Signal_SIGPROF() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGPROF.Do(func() {
		cache_Data_Posix_Signal_SIGPROF = gopurs_runtime.Value{Type: 9, IntVal: int64(3857536656), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGPROF
}

var cache_Data_Posix_Signal_SIGPWR gopurs_runtime.Value
var once_Data_Posix_Signal_SIGPWR sync.Once

func Get_Data_Posix_Signal_SIGPWR() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGPWR.Do(func() {
		cache_Data_Posix_Signal_SIGPWR = gopurs_runtime.Value{Type: 9, IntVal: int64(3766169998), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGPWR
}

var cache_Data_Posix_Signal_SIGQUIT gopurs_runtime.Value
var once_Data_Posix_Signal_SIGQUIT sync.Once

func Get_Data_Posix_Signal_SIGQUIT() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGQUIT.Do(func() {
		cache_Data_Posix_Signal_SIGQUIT = gopurs_runtime.Value{Type: 9, IntVal: int64(533014370), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGQUIT
}

var cache_Data_Posix_Signal_SIGSEGV gopurs_runtime.Value
var once_Data_Posix_Signal_SIGSEGV sync.Once

func Get_Data_Posix_Signal_SIGSEGV() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGSEGV.Do(func() {
		cache_Data_Posix_Signal_SIGSEGV = gopurs_runtime.Value{Type: 9, IntVal: int64(232395324), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGSEGV
}

var cache_Data_Posix_Signal_SIGSTKFLT gopurs_runtime.Value
var once_Data_Posix_Signal_SIGSTKFLT sync.Once

func Get_Data_Posix_Signal_SIGSTKFLT() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGSTKFLT.Do(func() {
		cache_Data_Posix_Signal_SIGSTKFLT = gopurs_runtime.Value{Type: 9, IntVal: int64(1465356297), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGSTKFLT
}

var cache_Data_Posix_Signal_SIGSTOP gopurs_runtime.Value
var once_Data_Posix_Signal_SIGSTOP sync.Once

func Get_Data_Posix_Signal_SIGSTOP() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGSTOP.Do(func() {
		cache_Data_Posix_Signal_SIGSTOP = gopurs_runtime.Value{Type: 9, IntVal: int64(2266362243), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGSTOP
}

var cache_Data_Posix_Signal_SIGSYS gopurs_runtime.Value
var once_Data_Posix_Signal_SIGSYS sync.Once

func Get_Data_Posix_Signal_SIGSYS() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGSYS.Do(func() {
		cache_Data_Posix_Signal_SIGSYS = gopurs_runtime.Value{Type: 9, IntVal: int64(1233964866), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGSYS
}

var cache_Data_Posix_Signal_SIGTERM gopurs_runtime.Value
var once_Data_Posix_Signal_SIGTERM sync.Once

func Get_Data_Posix_Signal_SIGTERM() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGTERM.Do(func() {
		cache_Data_Posix_Signal_SIGTERM = gopurs_runtime.Value{Type: 9, IntVal: int64(1894626581), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGTERM
}

var cache_Data_Posix_Signal_SIGTRAP gopurs_runtime.Value
var once_Data_Posix_Signal_SIGTRAP sync.Once

func Get_Data_Posix_Signal_SIGTRAP() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGTRAP.Do(func() {
		cache_Data_Posix_Signal_SIGTRAP = gopurs_runtime.Value{Type: 9, IntVal: int64(2192683852), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGTRAP
}

var cache_Data_Posix_Signal_SIGTSTP gopurs_runtime.Value
var once_Data_Posix_Signal_SIGTSTP sync.Once

func Get_Data_Posix_Signal_SIGTSTP() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGTSTP.Do(func() {
		cache_Data_Posix_Signal_SIGTSTP = gopurs_runtime.Value{Type: 9, IntVal: int64(2539377464), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGTSTP
}

var cache_Data_Posix_Signal_SIGTTIN gopurs_runtime.Value
var once_Data_Posix_Signal_SIGTTIN sync.Once

func Get_Data_Posix_Signal_SIGTTIN() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGTTIN.Do(func() {
		cache_Data_Posix_Signal_SIGTTIN = gopurs_runtime.Value{Type: 9, IntVal: int64(2244980060), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGTTIN
}

var cache_Data_Posix_Signal_SIGTTOU gopurs_runtime.Value
var once_Data_Posix_Signal_SIGTTOU sync.Once

func Get_Data_Posix_Signal_SIGTTOU() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGTTOU.Do(func() {
		cache_Data_Posix_Signal_SIGTTOU = gopurs_runtime.Value{Type: 9, IntVal: int64(1507781857), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGTTOU
}

var cache_Data_Posix_Signal_SIGUNUSED gopurs_runtime.Value
var once_Data_Posix_Signal_SIGUNUSED sync.Once

func Get_Data_Posix_Signal_SIGUNUSED() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGUNUSED.Do(func() {
		cache_Data_Posix_Signal_SIGUNUSED = gopurs_runtime.Value{Type: 9, IntVal: int64(1484085351), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGUNUSED
}

var cache_Data_Posix_Signal_SIGURG gopurs_runtime.Value
var once_Data_Posix_Signal_SIGURG sync.Once

func Get_Data_Posix_Signal_SIGURG() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGURG.Do(func() {
		cache_Data_Posix_Signal_SIGURG = gopurs_runtime.Value{Type: 9, IntVal: int64(1502848187), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGURG
}

var cache_Data_Posix_Signal_SIGUSR1 gopurs_runtime.Value
var once_Data_Posix_Signal_SIGUSR1 sync.Once

func Get_Data_Posix_Signal_SIGUSR1() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGUSR1.Do(func() {
		cache_Data_Posix_Signal_SIGUSR1 = gopurs_runtime.Value{Type: 9, IntVal: int64(517530686), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGUSR1
}

var cache_Data_Posix_Signal_SIGUSR2 gopurs_runtime.Value
var once_Data_Posix_Signal_SIGUSR2 sync.Once

func Get_Data_Posix_Signal_SIGUSR2() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGUSR2.Do(func() {
		cache_Data_Posix_Signal_SIGUSR2 = gopurs_runtime.Value{Type: 9, IntVal: int64(3761689693), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGUSR2
}

var cache_Data_Posix_Signal_SIGVTALRM gopurs_runtime.Value
var once_Data_Posix_Signal_SIGVTALRM sync.Once

func Get_Data_Posix_Signal_SIGVTALRM() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGVTALRM.Do(func() {
		cache_Data_Posix_Signal_SIGVTALRM = gopurs_runtime.Value{Type: 9, IntVal: int64(2190983275), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGVTALRM
}

var cache_Data_Posix_Signal_SIGWINCH gopurs_runtime.Value
var once_Data_Posix_Signal_SIGWINCH sync.Once

func Get_Data_Posix_Signal_SIGWINCH() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGWINCH.Do(func() {
		cache_Data_Posix_Signal_SIGWINCH = gopurs_runtime.Value{Type: 9, IntVal: int64(2253074784), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGWINCH
}

var cache_Data_Posix_Signal_SIGXCPU gopurs_runtime.Value
var once_Data_Posix_Signal_SIGXCPU sync.Once

func Get_Data_Posix_Signal_SIGXCPU() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGXCPU.Do(func() {
		cache_Data_Posix_Signal_SIGXCPU = gopurs_runtime.Value{Type: 9, IntVal: int64(1954817349), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGXCPU
}

var cache_Data_Posix_Signal_SIGXFSZ gopurs_runtime.Value
var once_Data_Posix_Signal_SIGXFSZ sync.Once

func Get_Data_Posix_Signal_SIGXFSZ() gopurs_runtime.Value {
	once_Data_Posix_Signal_SIGXFSZ.Do(func() {
		cache_Data_Posix_Signal_SIGXFSZ = gopurs_runtime.Value{Type: 9, IntVal: int64(1796895884), UnsafePtr: nil}
	})
	return cache_Data_Posix_Signal_SIGXFSZ
}

var cache_Data_Posix_Signal_toString gopurs_runtime.Value
var once_Data_Posix_Signal_toString sync.Once

func Get_Data_Posix_Signal_toString() gopurs_runtime.Value {
	once_Data_Posix_Signal_toString.Do(func() {
		cache_Data_Posix_Signal_toString = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Str(Call_Data_Posix_Signal_toString(uint32(s_0_box.IntVal)))
		})
	})
	return cache_Data_Posix_Signal_toString
}

var cache_Data_Posix_Signal_showSignal gopurs_runtime.Value
var once_Data_Posix_Signal_showSignal sync.Once

func Get_Data_Posix_Signal_showSignal() gopurs_runtime.Value {
	once_Data_Posix_Signal_showSignal.Do(func() {
		cache_Data_Posix_Signal_showSignal = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_Signal_1209612131_1386611502((&Constructor_Data_Show_Show[uint32]{1, Get_Data_Posix_Signal_toString()})))}
	})
	return cache_Data_Posix_Signal_showSignal
}

var cache_Data_Posix_Signal_fromString gopurs_runtime.Value
var once_Data_Posix_Signal_fromString sync.Once

func Get_Data_Posix_Signal_fromString() gopurs_runtime.Value {
	once_Data_Posix_Signal_fromString.Do(func() {
		cache_Data_Posix_Signal_fromString = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return func() gopurs_runtime.Value {
				_v := Call_Data_Posix_Signal_fromString(s_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
		})
	})
	return cache_Data_Posix_Signal_fromString
}

var cache_Data_Posix_Signal_eqSignal gopurs_runtime.Value
var once_Data_Posix_Signal_eqSignal sync.Once

func Get_Data_Posix_Signal_eqSignal() gopurs_runtime.Value {
	once_Data_Posix_Signal_eqSignal.Do(func() {
		cache_Data_Posix_Signal_eqSignal = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_Signal_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t77 bool
			{
				var __t_tag_3 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 3640760862 {
					var __t_tag_4 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_4
					__t77 = (uint32(__t_tag_4) == 3640760862)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_5 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_5
				if uint32(__t_tag_5) == 1332741801 {
					var __t_tag_6 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_6
					__t77 = (uint32(__t_tag_6) == 1332741801)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_7 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_7
				if uint32(__t_tag_7) == 1978617951 {
					var __t_tag_8 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_8
					__t77 = (uint32(__t_tag_8) == 1978617951)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_9 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_9
				if uint32(__t_tag_9) == 737752792 {
					var __t_tag_10 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_10
					__t77 = (uint32(__t_tag_10) == 737752792)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_11 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_11
				if uint32(__t_tag_11) == 305761328 {
					var __t_tag_12 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_12
					__t77 = (uint32(__t_tag_12) == 305761328)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_13 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_13
				if uint32(__t_tag_13) == 2341479853 {
					var __t_tag_14 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_14
					__t77 = (uint32(__t_tag_14) == 2341479853)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_15 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_15
				if uint32(__t_tag_15) == 2502462151 {
					var __t_tag_16 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_16
					__t77 = (uint32(__t_tag_16) == 2502462151)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_17 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_17
				if uint32(__t_tag_17) == 3424215656 {
					var __t_tag_18 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_18
					__t77 = (uint32(__t_tag_18) == 3424215656)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_19 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_19
				if uint32(__t_tag_19) == 232722006 {
					var __t_tag_20 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_20
					__t77 = (uint32(__t_tag_20) == 232722006)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_21 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_21
				if uint32(__t_tag_21) == 3802329906 {
					var __t_tag_22 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_22
					__t77 = (uint32(__t_tag_22) == 3802329906)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_23 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_23
				if uint32(__t_tag_23) == 3332358197 {
					var __t_tag_24 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_24
					__t77 = (uint32(__t_tag_24) == 3332358197)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_25 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_25
				if uint32(__t_tag_25) == 681794280 {
					var __t_tag_26 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_26
					__t77 = (uint32(__t_tag_26) == 681794280)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_27 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_27
				if uint32(__t_tag_27) == 2445008893 {
					var __t_tag_28 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_28
					__t77 = (uint32(__t_tag_28) == 2445008893)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_29 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_29
				if uint32(__t_tag_29) == 186232457 {
					var __t_tag_30 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_30
					__t77 = (uint32(__t_tag_30) == 186232457)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_31 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_31
				if uint32(__t_tag_31) == 3476794617 {
					var __t_tag_32 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_32
					__t77 = (uint32(__t_tag_32) == 3476794617)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_33 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_33
				if uint32(__t_tag_33) == 651594687 {
					var __t_tag_34 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_34
					__t77 = (uint32(__t_tag_34) == 651594687)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_35 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_35
				if uint32(__t_tag_35) == 2385806231 {
					var __t_tag_36 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_36
					__t77 = (uint32(__t_tag_36) == 2385806231)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_37 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_37
				if uint32(__t_tag_37) == 583450148 {
					var __t_tag_38 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_38
					__t77 = (uint32(__t_tag_38) == 583450148)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_39 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_39
				if uint32(__t_tag_39) == 3857536656 {
					var __t_tag_40 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_40
					__t77 = (uint32(__t_tag_40) == 3857536656)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_41 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_41
				if uint32(__t_tag_41) == 3766169998 {
					var __t_tag_42 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_42
					__t77 = (uint32(__t_tag_42) == 3766169998)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_43 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_43
				if uint32(__t_tag_43) == 533014370 {
					var __t_tag_44 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_44
					__t77 = (uint32(__t_tag_44) == 533014370)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_45 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_45
				if uint32(__t_tag_45) == 232395324 {
					var __t_tag_46 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_46
					__t77 = (uint32(__t_tag_46) == 232395324)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_47 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_47
				if uint32(__t_tag_47) == 1465356297 {
					var __t_tag_48 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_48
					__t77 = (uint32(__t_tag_48) == 1465356297)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_49 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_49
				if uint32(__t_tag_49) == 2266362243 {
					var __t_tag_50 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_50
					__t77 = (uint32(__t_tag_50) == 2266362243)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_51 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_51
				if uint32(__t_tag_51) == 1233964866 {
					var __t_tag_52 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_52
					__t77 = (uint32(__t_tag_52) == 1233964866)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_53 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_53
				if uint32(__t_tag_53) == 1894626581 {
					var __t_tag_54 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_54
					__t77 = (uint32(__t_tag_54) == 1894626581)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_55 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_55
				if uint32(__t_tag_55) == 2192683852 {
					var __t_tag_56 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_56
					__t77 = (uint32(__t_tag_56) == 2192683852)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_57 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_57
				if uint32(__t_tag_57) == 2539377464 {
					var __t_tag_58 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_58
					__t77 = (uint32(__t_tag_58) == 2539377464)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_59 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_59
				if uint32(__t_tag_59) == 2244980060 {
					var __t_tag_60 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_60
					__t77 = (uint32(__t_tag_60) == 2244980060)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_61 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_61
				if uint32(__t_tag_61) == 1507781857 {
					var __t_tag_62 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_62
					__t77 = (uint32(__t_tag_62) == 1507781857)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_63 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_63
				if uint32(__t_tag_63) == 1484085351 {
					var __t_tag_64 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_64
					__t77 = (uint32(__t_tag_64) == 1484085351)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_65 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_65
				if uint32(__t_tag_65) == 1502848187 {
					var __t_tag_66 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_66
					__t77 = (uint32(__t_tag_66) == 1502848187)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_67 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_67
				if uint32(__t_tag_67) == 517530686 {
					var __t_tag_68 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_68
					__t77 = (uint32(__t_tag_68) == 517530686)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_69 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_69
				if uint32(__t_tag_69) == 3761689693 {
					var __t_tag_70 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_70
					__t77 = (uint32(__t_tag_70) == 3761689693)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_71 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_71
				if uint32(__t_tag_71) == 2190983275 {
					var __t_tag_72 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_72
					__t77 = (uint32(__t_tag_72) == 2190983275)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_73 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_73
				if uint32(__t_tag_73) == 2253074784 {
					var __t_tag_74 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_74
					__t77 = (uint32(__t_tag_74) == 2253074784)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_75 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_75
				if uint32(__t_tag_75) == 1954817349 {
					var __t_tag_76 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_76
					__t77 = (uint32(__t_tag_76) == 1954817349)
					goto end_branch_77
				} else {

				}
			}
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				var __t_and_2 bool = false
				if uint32(__t_tag_0) == 1796895884 {

					var __t_tag_1 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_1
					__t_and_2 = (uint32(__t_tag_1) == 1796895884)
				}
				__t77 = __t_and_2
			}
		end_branch_77:
			return gopurs_runtime.Bool(__t77)
		})})))}
	})
	return cache_Data_Posix_Signal_eqSignal
}

var cache_Data_Posix_Signal_ordSignal gopurs_runtime.Value
var once_Data_Posix_Signal_ordSignal sync.Once

func Get_Data_Posix_Signal_ordSignal() gopurs_runtime.Value {
	once_Data_Posix_Signal_ordSignal.Do(func() {
		cache_Data_Posix_Signal_ordSignal = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_Signal_3730953251_4177771502((&Constructor_Data_Ord_Ord[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_Signal_3768443459_3790796878(Rebox_Data_Posix_Signal_3790796878_3768443459(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Posix_Signal_eqSignal()))))}
		}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
			var __t151 gopurs_runtime.Value
			{
				var __t_tag_0 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_0
				if uint32(__t_tag_0) == 3640760862 {
					var __t2 gopurs_runtime.Value
					{
						var __t_tag_1 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_1
						if uint32(__t_tag_1) == 3640760862 {
							__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_2
						} else {

						}
					}
					{
						__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_2:
					__t151 = __t2
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_3 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_3
				if uint32(__t_tag_3) == 3640760862 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_4 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_4
				if uint32(__t_tag_4) == 1332741801 {
					var __t6 gopurs_runtime.Value
					{
						var __t_tag_5 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_5
						if uint32(__t_tag_5) == 1332741801 {
							__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_6
						} else {

						}
					}
					{
						__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_6:
					__t151 = __t6
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_7 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_7
				if uint32(__t_tag_7) == 1332741801 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_8 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_8
				if uint32(__t_tag_8) == 1978617951 {
					var __t10 gopurs_runtime.Value
					{
						var __t_tag_9 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_9
						if uint32(__t_tag_9) == 1978617951 {
							__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_10
						} else {

						}
					}
					{
						__t10 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_10:
					__t151 = __t10
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_11 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_11
				if uint32(__t_tag_11) == 1978617951 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_12 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_12
				if uint32(__t_tag_12) == 737752792 {
					var __t14 gopurs_runtime.Value
					{
						var __t_tag_13 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_13
						if uint32(__t_tag_13) == 737752792 {
							__t14 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_14
						} else {

						}
					}
					{
						__t14 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_14:
					__t151 = __t14
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_15 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_15
				if uint32(__t_tag_15) == 737752792 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_16 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_16
				if uint32(__t_tag_16) == 305761328 {
					var __t18 gopurs_runtime.Value
					{
						var __t_tag_17 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_17
						if uint32(__t_tag_17) == 305761328 {
							__t18 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_18
						} else {

						}
					}
					{
						__t18 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_18:
					__t151 = __t18
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_19 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_19
				if uint32(__t_tag_19) == 305761328 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_20 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_20
				if uint32(__t_tag_20) == 2341479853 {
					var __t22 gopurs_runtime.Value
					{
						var __t_tag_21 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_21
						if uint32(__t_tag_21) == 2341479853 {
							__t22 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_22
						} else {

						}
					}
					{
						__t22 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_22:
					__t151 = __t22
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_23 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_23
				if uint32(__t_tag_23) == 2341479853 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_24 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_24
				if uint32(__t_tag_24) == 2502462151 {
					var __t26 gopurs_runtime.Value
					{
						var __t_tag_25 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_25
						if uint32(__t_tag_25) == 2502462151 {
							__t26 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_26
						} else {

						}
					}
					{
						__t26 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_26:
					__t151 = __t26
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_27 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_27
				if uint32(__t_tag_27) == 2502462151 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_28 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_28
				if uint32(__t_tag_28) == 3424215656 {
					var __t30 gopurs_runtime.Value
					{
						var __t_tag_29 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_29
						if uint32(__t_tag_29) == 3424215656 {
							__t30 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_30
						} else {

						}
					}
					{
						__t30 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_30:
					__t151 = __t30
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_31 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_31
				if uint32(__t_tag_31) == 3424215656 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_32 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_32
				if uint32(__t_tag_32) == 232722006 {
					var __t34 gopurs_runtime.Value
					{
						var __t_tag_33 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_33
						if uint32(__t_tag_33) == 232722006 {
							__t34 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_34
						} else {

						}
					}
					{
						__t34 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_34:
					__t151 = __t34
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_35 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_35
				if uint32(__t_tag_35) == 232722006 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_36 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_36
				if uint32(__t_tag_36) == 3802329906 {
					var __t38 gopurs_runtime.Value
					{
						var __t_tag_37 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_37
						if uint32(__t_tag_37) == 3802329906 {
							__t38 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_38
						} else {

						}
					}
					{
						__t38 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_38:
					__t151 = __t38
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_39 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_39
				if uint32(__t_tag_39) == 3802329906 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_40 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_40
				if uint32(__t_tag_40) == 3332358197 {
					var __t42 gopurs_runtime.Value
					{
						var __t_tag_41 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_41
						if uint32(__t_tag_41) == 3332358197 {
							__t42 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_42
						} else {

						}
					}
					{
						__t42 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_42:
					__t151 = __t42
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_43 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_43
				if uint32(__t_tag_43) == 3332358197 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_44 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_44
				if uint32(__t_tag_44) == 681794280 {
					var __t46 gopurs_runtime.Value
					{
						var __t_tag_45 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_45
						if uint32(__t_tag_45) == 681794280 {
							__t46 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_46
						} else {

						}
					}
					{
						__t46 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_46:
					__t151 = __t46
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_47 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_47
				if uint32(__t_tag_47) == 681794280 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_48 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_48
				if uint32(__t_tag_48) == 2445008893 {
					var __t50 gopurs_runtime.Value
					{
						var __t_tag_49 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_49
						if uint32(__t_tag_49) == 2445008893 {
							__t50 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_50
						} else {

						}
					}
					{
						__t50 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_50:
					__t151 = __t50
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_51 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_51
				if uint32(__t_tag_51) == 2445008893 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_52 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_52
				if uint32(__t_tag_52) == 186232457 {
					var __t54 gopurs_runtime.Value
					{
						var __t_tag_53 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_53
						if uint32(__t_tag_53) == 186232457 {
							__t54 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_54
						} else {

						}
					}
					{
						__t54 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_54:
					__t151 = __t54
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_55 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_55
				if uint32(__t_tag_55) == 186232457 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_56 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_56
				if uint32(__t_tag_56) == 3476794617 {
					var __t58 gopurs_runtime.Value
					{
						var __t_tag_57 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_57
						if uint32(__t_tag_57) == 3476794617 {
							__t58 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_58
						} else {

						}
					}
					{
						__t58 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_58:
					__t151 = __t58
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_59 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_59
				if uint32(__t_tag_59) == 3476794617 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_60 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_60
				if uint32(__t_tag_60) == 651594687 {
					var __t62 gopurs_runtime.Value
					{
						var __t_tag_61 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_61
						if uint32(__t_tag_61) == 651594687 {
							__t62 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_62
						} else {

						}
					}
					{
						__t62 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_62:
					__t151 = __t62
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_63 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_63
				if uint32(__t_tag_63) == 651594687 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_64 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_64
				if uint32(__t_tag_64) == 2385806231 {
					var __t66 gopurs_runtime.Value
					{
						var __t_tag_65 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_65
						if uint32(__t_tag_65) == 2385806231 {
							__t66 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_66
						} else {

						}
					}
					{
						__t66 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_66:
					__t151 = __t66
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_67 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_67
				if uint32(__t_tag_67) == 2385806231 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_68 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_68
				if uint32(__t_tag_68) == 583450148 {
					var __t70 gopurs_runtime.Value
					{
						var __t_tag_69 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_69
						if uint32(__t_tag_69) == 583450148 {
							__t70 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_70
						} else {

						}
					}
					{
						__t70 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_70:
					__t151 = __t70
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_71 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_71
				if uint32(__t_tag_71) == 583450148 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_72 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_72
				if uint32(__t_tag_72) == 3857536656 {
					var __t74 gopurs_runtime.Value
					{
						var __t_tag_73 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_73
						if uint32(__t_tag_73) == 3857536656 {
							__t74 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_74
						} else {

						}
					}
					{
						__t74 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_74:
					__t151 = __t74
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_75 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_75
				if uint32(__t_tag_75) == 3857536656 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_76 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_76
				if uint32(__t_tag_76) == 3766169998 {
					var __t78 gopurs_runtime.Value
					{
						var __t_tag_77 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_77
						if uint32(__t_tag_77) == 3766169998 {
							__t78 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_78
						} else {

						}
					}
					{
						__t78 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_78:
					__t151 = __t78
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_79 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_79
				if uint32(__t_tag_79) == 3766169998 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_80 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_80
				if uint32(__t_tag_80) == 533014370 {
					var __t82 gopurs_runtime.Value
					{
						var __t_tag_81 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_81
						if uint32(__t_tag_81) == 533014370 {
							__t82 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_82
						} else {

						}
					}
					{
						__t82 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_82:
					__t151 = __t82
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_83 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_83
				if uint32(__t_tag_83) == 533014370 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_84 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_84
				if uint32(__t_tag_84) == 232395324 {
					var __t86 gopurs_runtime.Value
					{
						var __t_tag_85 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_85
						if uint32(__t_tag_85) == 232395324 {
							__t86 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_86
						} else {

						}
					}
					{
						__t86 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_86:
					__t151 = __t86
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_87 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_87
				if uint32(__t_tag_87) == 232395324 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_88 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_88
				if uint32(__t_tag_88) == 1465356297 {
					var __t90 gopurs_runtime.Value
					{
						var __t_tag_89 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_89
						if uint32(__t_tag_89) == 1465356297 {
							__t90 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_90
						} else {

						}
					}
					{
						__t90 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_90:
					__t151 = __t90
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_91 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_91
				if uint32(__t_tag_91) == 1465356297 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_92 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_92
				if uint32(__t_tag_92) == 2266362243 {
					var __t94 gopurs_runtime.Value
					{
						var __t_tag_93 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_93
						if uint32(__t_tag_93) == 2266362243 {
							__t94 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_94
						} else {

						}
					}
					{
						__t94 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_94:
					__t151 = __t94
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_95 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_95
				if uint32(__t_tag_95) == 2266362243 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_96 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_96
				if uint32(__t_tag_96) == 1233964866 {
					var __t98 gopurs_runtime.Value
					{
						var __t_tag_97 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_97
						if uint32(__t_tag_97) == 1233964866 {
							__t98 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_98
						} else {

						}
					}
					{
						__t98 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_98:
					__t151 = __t98
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_99 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_99
				if uint32(__t_tag_99) == 1233964866 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_100 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_100
				if uint32(__t_tag_100) == 1894626581 {
					var __t102 gopurs_runtime.Value
					{
						var __t_tag_101 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_101
						if uint32(__t_tag_101) == 1894626581 {
							__t102 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_102
						} else {

						}
					}
					{
						__t102 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_102:
					__t151 = __t102
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_103 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_103
				if uint32(__t_tag_103) == 1894626581 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_104 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_104
				if uint32(__t_tag_104) == 2192683852 {
					var __t106 gopurs_runtime.Value
					{
						var __t_tag_105 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_105
						if uint32(__t_tag_105) == 2192683852 {
							__t106 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_106
						} else {

						}
					}
					{
						__t106 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_106:
					__t151 = __t106
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_107 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_107
				if uint32(__t_tag_107) == 2192683852 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_108 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_108
				if uint32(__t_tag_108) == 2539377464 {
					var __t110 gopurs_runtime.Value
					{
						var __t_tag_109 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_109
						if uint32(__t_tag_109) == 2539377464 {
							__t110 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_110
						} else {

						}
					}
					{
						__t110 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_110:
					__t151 = __t110
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_111 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_111
				if uint32(__t_tag_111) == 2539377464 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_112 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_112
				if uint32(__t_tag_112) == 2244980060 {
					var __t114 gopurs_runtime.Value
					{
						var __t_tag_113 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_113
						if uint32(__t_tag_113) == 2244980060 {
							__t114 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_114
						} else {

						}
					}
					{
						__t114 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_114:
					__t151 = __t114
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_115 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_115
				if uint32(__t_tag_115) == 2244980060 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_116 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_116
				if uint32(__t_tag_116) == 1507781857 {
					var __t118 gopurs_runtime.Value
					{
						var __t_tag_117 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_117
						if uint32(__t_tag_117) == 1507781857 {
							__t118 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_118
						} else {

						}
					}
					{
						__t118 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_118:
					__t151 = __t118
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_119 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_119
				if uint32(__t_tag_119) == 1507781857 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_120 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_120
				if uint32(__t_tag_120) == 1484085351 {
					var __t122 gopurs_runtime.Value
					{
						var __t_tag_121 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_121
						if uint32(__t_tag_121) == 1484085351 {
							__t122 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_122
						} else {

						}
					}
					{
						__t122 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_122:
					__t151 = __t122
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_123 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_123
				if uint32(__t_tag_123) == 1484085351 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_124 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_124
				if uint32(__t_tag_124) == 1502848187 {
					var __t126 gopurs_runtime.Value
					{
						var __t_tag_125 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_125
						if uint32(__t_tag_125) == 1502848187 {
							__t126 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_126
						} else {

						}
					}
					{
						__t126 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_126:
					__t151 = __t126
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_127 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_127
				if uint32(__t_tag_127) == 1502848187 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_128 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_128
				if uint32(__t_tag_128) == 517530686 {
					var __t130 gopurs_runtime.Value
					{
						var __t_tag_129 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_129
						if uint32(__t_tag_129) == 517530686 {
							__t130 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_130
						} else {

						}
					}
					{
						__t130 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_130:
					__t151 = __t130
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_131 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_131
				if uint32(__t_tag_131) == 517530686 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_132 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_132
				if uint32(__t_tag_132) == 3761689693 {
					var __t134 gopurs_runtime.Value
					{
						var __t_tag_133 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_133
						if uint32(__t_tag_133) == 3761689693 {
							__t134 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_134
						} else {

						}
					}
					{
						__t134 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_134:
					__t151 = __t134
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_135 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_135
				if uint32(__t_tag_135) == 3761689693 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_136 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_136
				if uint32(__t_tag_136) == 2190983275 {
					var __t138 gopurs_runtime.Value
					{
						var __t_tag_137 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_137
						if uint32(__t_tag_137) == 2190983275 {
							__t138 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_138
						} else {

						}
					}
					{
						__t138 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_138:
					__t151 = __t138
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_139 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_139
				if uint32(__t_tag_139) == 2190983275 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_140 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_140
				if uint32(__t_tag_140) == 2253074784 {
					var __t142 gopurs_runtime.Value
					{
						var __t_tag_141 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_141
						if uint32(__t_tag_141) == 2253074784 {
							__t142 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_142
						} else {

						}
					}
					{
						__t142 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_142:
					__t151 = __t142
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_143 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_143
				if uint32(__t_tag_143) == 2253074784 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_144 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_144
				if uint32(__t_tag_144) == 1954817349 {
					var __t146 gopurs_runtime.Value
					{
						var __t_tag_145 uint32 = uint32(y_1.IntVal)
						_ = __t_tag_145
						if uint32(__t_tag_145) == 1954817349 {
							__t146 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
							goto end_branch_146
						} else {

						}
					}
					{
						__t146 = gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}
					}
				end_branch_146:
					__t151 = __t146
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_147 uint32 = uint32(y_1.IntVal)
				_ = __t_tag_147
				if uint32(__t_tag_147) == 1954817349 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				var __t_tag_148 uint32 = uint32(x_0.IntVal)
				_ = __t_tag_148
				var __t_and_150 bool = false
				if uint32(__t_tag_148) == 1796895884 {

					var __t_tag_149 uint32 = uint32(y_1.IntVal)
					_ = __t_tag_149
					__t_and_150 = (uint32(__t_tag_149) == 1796895884)
				}
				if __t_and_150 {
					__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
					goto end_branch_151
				} else {

				}
			}
			{
				__t151 = gopurs_runtime.Value{Type: 9, IntVal: int64(func() uint32 { panic("Failed pattern match") }()), UnsafePtr: nil}
			}
		end_branch_151:
			return __t151
		})})))}
	})
	return cache_Data_Posix_Signal_ordSignal
}

type Constructor_Data_Posix_Signal_SIGABRT struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGALRM struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGBUS struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGCHLD struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGCLD struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGCONT struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGEMT struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGFPE struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGHUP struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGILL struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGINFO struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGINT struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGIO struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGIOT struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGKILL struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGLOST struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGPIPE struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGPOLL struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGPROF struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGPWR struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGQUIT struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGSEGV struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGSTKFLT struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGSTOP struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGSYS struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGTERM struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGTRAP struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGTSTP struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGTTIN struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGTTOU struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGUNUSED struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGURG struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGUSR1 struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGUSR2 struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGVTALRM struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGWINCH struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGXCPU struct {
	Rc uint32
}

type Constructor_Data_Posix_Signal_SIGXFSZ struct {
	Rc uint32
}

func Call_Data_Posix_Signal_toString(s_0_loop uint32) string {
	var s_0 uint32 = s_0_loop
	_ = s_0
	var __t0 string
	{
		if s_0 == 3640760862 {
			__t0 = "SIGABRT"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 1332741801 {
			__t0 = "SIGALRM"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 1978617951 {
			__t0 = "SIGBUS"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 737752792 {
			__t0 = "SIGCHLD"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 305761328 {
			__t0 = "SIGCLD"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 2341479853 {
			__t0 = "SIGCONT"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 2502462151 {
			__t0 = "SIGEMT"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 3424215656 {
			__t0 = "SIGFPE"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 232722006 {
			__t0 = "SIGHUP"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 3802329906 {
			__t0 = "SIGILL"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 3332358197 {
			__t0 = "SIGINFO"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 681794280 {
			__t0 = "SIGINT"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 2445008893 {
			__t0 = "SIGIO"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 186232457 {
			__t0 = "SIGIOT"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 3476794617 {
			__t0 = "SIGKILL"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 651594687 {
			__t0 = "SIGLOST"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 2385806231 {
			__t0 = "SIGPIPE"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 583450148 {
			__t0 = "SIGPOLL"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 3857536656 {
			__t0 = "SIGPROF"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 3766169998 {
			__t0 = "SIGPWR"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 533014370 {
			__t0 = "SIGQUIT"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 232395324 {
			__t0 = "SIGSEGV"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 1465356297 {
			__t0 = "SIGSTKFLT"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 2266362243 {
			__t0 = "SIGSTOP"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 1233964866 {
			__t0 = "SIGSYS"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 1894626581 {
			__t0 = "SIGTERM"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 2192683852 {
			__t0 = "SIGTRAP"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 2539377464 {
			__t0 = "SIGTSTP"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 2244980060 {
			__t0 = "SIGTTIN"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 1507781857 {
			__t0 = "SIGTTOU"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 1484085351 {
			__t0 = "SIGUNUSED"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 1502848187 {
			__t0 = "SIGURG"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 517530686 {
			__t0 = "SIGUSR1"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 3761689693 {
			__t0 = "SIGUSR2"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 2190983275 {
			__t0 = "SIGVTALRM"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 2253074784 {
			__t0 = "SIGWINCH"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 1954817349 {
			__t0 = "SIGXCPU"
			goto end_branch_0
		} else {

		}
	}
	{
		if s_0 == 1796895884 {
			__t0 = "SIGXFSZ"
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = func() string { panic("Failed pattern match") }()
	}
end_branch_0:
	return __t0
}

func Call_Data_Posix_Signal_fromString(s_0_loop string) struct {
	V0 gopurs_runtime.Value
	V1 bool
} {
	var s_0 string = s_0_loop
	_ = s_0
	var __t0 *Constructor_Data_Maybe_Just[uint32]
	{
		if (s_0) == ("SIGABRT") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(3640760862), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGALRM") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(1332741801), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGBUS") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(1978617951), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGCHLD") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(737752792), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGCLD") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(305761328), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGCONT") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2341479853), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGEMT") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2502462151), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGFPE") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(3424215656), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGHUP") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(232722006), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGILL") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(3802329906), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGINFO") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(3332358197), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGINT") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(681794280), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGIO") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2445008893), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGIOT") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(186232457), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGKILL") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(3476794617), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGLOST") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(651594687), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGPIPE") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2385806231), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGPOLL") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(583450148), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGPROF") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(3857536656), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGPWR") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(3766169998), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGQUIT") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(533014370), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGSEGV") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(232395324), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGSTKFLT") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(1465356297), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGSTOP") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2266362243), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGSYS") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(1233964866), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGTERM") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(1894626581), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGTRAP") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2192683852), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGTSTP") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2539377464), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGTTIN") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2244980060), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGTTOU") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(1507781857), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGUNUSED") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(1484085351), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGURG") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(1502848187), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGUSR1") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(517530686), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGUSR2") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(3761689693), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGVTALRM") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2190983275), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGWINCH") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(2253074784), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGXCPU") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(1954817349), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		if (s_0) == ("SIGXFSZ") {
			__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct {
					V0 gopurs_runtime.Value
					V1 bool
				}{gopurs_runtime.Value{Type: 9, IntVal: int64(1796895884), UnsafePtr: nil}, true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
			goto end_branch_0
		} else {

		}
	}
	{
		__t0 = Rebox_Data_Posix_Signal_3094389156_622082505(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
			_v := struct {
				V0 gopurs_runtime.Value
				V1 bool
			}{gopurs_runtime.Value{}, false}
			if _v.V1 {
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
			}
			return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
		}()))
	}
end_branch_0:
	return func() struct {
		V0 gopurs_runtime.Value
		V1 bool
	} { _v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Posix_Signal_622082505_3094389156(__t0))}; if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
		return struct {
			V0 gopurs_runtime.Value
			V1 bool
		}{V0: (*(*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr)).V0, V1: true}
	}; return struct {
		V0 gopurs_runtime.Value
		V1 bool
	}{V0: gopurs_runtime.Value{}, V1: false} }()
}

func Rebox_Data_Posix_Signal_1209612131_1386611502(in *Constructor_Data_Show_Show[uint32]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Show_Show[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Posix_Signal_3094389156_622082505(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[uint32] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[uint32]{}
	out.V0 = uint32(in.V0.IntVal)
	return out
}

func Rebox_Data_Posix_Signal_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Ord_Ord[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Posix_Signal_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[gopurs_runtime.Value])(unsafe.Pointer(in))
}

func Rebox_Data_Posix_Signal_3790796878_3768443459(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[uint32] {
	if in == nil {
		return nil
	}
	return (*Constructor_Data_Eq_Eq[uint32])(unsafe.Pointer(in))
}

func Rebox_Data_Posix_Signal_622082505_3094389156(in *Constructor_Data_Maybe_Just[uint32]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil {
		return nil
	}
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
	out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
	return out
}
