package locale

import (
	"fmt"
	"os"
	"testing"
)

func TestFS(t *testing.T) {
	_ = DelayInitializeLocale()
	fmt.Fprintf(os.Stderr, "load ok={%v}\n", LoadString("ok"))
}
