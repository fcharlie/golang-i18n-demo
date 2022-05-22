//go:build !windows

package locale

import (
	"os"
	"strings"
)

var localeEnvs = []string{"LC_ALL", "LC_MESSAGES", "LANG"}

func parseEnvLc(s string) string {
	x := strings.Split(s, ".")
	// "C" means "ANSI-C" and "POSIX", if locale set to C, we can simple
	// set returned language to "en_US"
	if x[0] == "C" {
		return "en-US"
	}
	return strings.Replace(x[0], "_", "-", 1)
}

func defaultLocaleName() string {
	for _, v := range localeEnvs {
		if e, ok := os.LookupEnv(v); ok {
			return parseEnvLc(e)
		}
	}
	return "en-US"
}
