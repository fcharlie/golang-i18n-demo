package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fcharlie/golang-i18n-demo/modules/locale"
)

func main() {
	os.Setenv("LC_ALL", "zh_CN.UTF8")
	_ = locale.DelayInitializeLocale()
	fmt.Fprintf(os.Stderr, "load ok={%v}\n", locale.LoadString("ok"))
	locale.Fprintf(os.Stderr, "current os '%s'\n", runtime.GOOS)
}
