package firstlogin

import "io/ioutil"

// #include <locale.h>
import "C"

func (f *FirstLogin) generateLocale(locale string) error {
	C.setlocale(0, C.CString(locale))
	return ioutil.WriteFile("/etc/locale.conf", []byte("LANG="+locale), 0644)
}
