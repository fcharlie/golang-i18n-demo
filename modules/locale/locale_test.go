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
