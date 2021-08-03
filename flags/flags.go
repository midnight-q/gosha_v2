package flags

import (
	"flag"
	"os"
	"regexp"
)

var IsDev = flag.Bool("dev", false, "Enable development mode")
var Auth = flag.Bool("auth", false, "Disable authorisation")

var _ = ParseFlags()

func ParseFlags() error {

	isTest, _ := regexp.MatchString(`\.test$`, os.Args[0])

	if !isTest {
		flag.Parse()
	}

	return nil
}
