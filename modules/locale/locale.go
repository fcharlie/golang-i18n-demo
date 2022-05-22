package locale

import (
	"embed"
	"fmt"
	"io"

	"github.com/BurntSushi/toml"
)

//go:embed languages
var langFS embed.FS

var (
	langTable = make(map[string]any)
)

func DelayInitializeLocale() error {
	langfile := fmt.Sprint("languages/", defaultLocaleName(), ".toml")
	fd, err := langFS.Open(langfile)
	if err != nil {
		return err
	}
	defer fd.Close()
	if _, err := toml.NewDecoder(fd).Decode(&langTable); err != nil {
		return err
	}
	return nil
}

func DefaultLocaleName() string {
	return defaultLocaleName()
}

func LoadString(k string) string {
	if v, ok := langTable[k]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	if localeFmt := LoadString(format); len(localeFmt) != 0 {
		format = localeFmt
	}
	return fmt.Fprintf(w, format, a...)
}
