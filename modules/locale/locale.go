package locale

import (
	"embed"
	"fmt"

	"github.com/BurntSushi/toml"
)

//go:embed languages
var langFS embed.FS

type LangTable map[string]string

var (
	localeName = defaultLocaleName()
	langTable  = make(map[string]any)
)

func DelayInitializeLocale() error {
	langfile := fmt.Sprint("languages/", localeName, ".toml")
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

func LocaleName() string {
	return localeName
}

func LoadString(k string) string {
	v := langTable[k]
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
