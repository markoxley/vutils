package vutils

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/digital-dream-labs/vector-go-sdk/pkg/vector"
)

// GetVector returns a vector object based on the sdk_config.ini file.
func GetVector() (*vector.Vector, error) {
	botToken := ""
	botTarget := ""

	home := os.Getenv("HOME")
	f, err := ioutil.ReadFile(home + "/.anki_vector/sdk_config.ini")
	if err != nil {
		return nil, err
	}
	var re = regexp.MustCompile(`(?m)^\s*([^\s=]+)\s*=\s*(\S+)$`)

	for _, m := range re.FindAllStringSubmatch(string(f), -1) {
		switch m[1] {
		case "ip":
			botTarget = m[2]
			if !strings.HasPrefix(botTarget, ":443") {
				botTarget += ":443"
			}
		case "guid":
			botToken = m[2]
		}
	}
	return vector.New(
		vector.WithTarget(botTarget),
		vector.WithToken(botToken),
	)
}
