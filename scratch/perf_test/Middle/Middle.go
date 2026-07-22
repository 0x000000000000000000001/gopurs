package Middle

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Test "gopurs/output/Test"
)

var middle gopurs_runtime.Value
var once_middle sync.Once
func Get_middle() gopurs_runtime.Value {
	once_middle.Do(func() {
		middle = pkg_Test.Get_test()
	})
	return middle
}


