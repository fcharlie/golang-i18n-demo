package locale

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func TestFS(t *testing.T) {
	_ = DelayInitializeLocale()
	fmt.Fprintf(os.Stderr, "load ok={%v}\n", LoadString("ok"))
	Fprintf(os.Stderr, "current os '%s'\n", runtime.GOOS)
}

func TestLANG(t *testing.T) {
	os.Setenv("LC_ALL", "zh_CN.UTF8")
	_ = DelayInitializeLocale()
	fmt.Fprintf(os.Stderr, "load ok={%v}\n", LoadString("ok"))
	Fprintf(os.Stderr, "current os '%s'\n", runtime.GOOS)
}
