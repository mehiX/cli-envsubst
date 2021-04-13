package transformers

import (
	"fmt"
	"os"
	"strings"
)

func LookupEnvCaseIns(s string) string {
	u := strings.ToUpper(s)
	v, ok := os.LookupEnv(u)
	if !ok {
		return fmt.Sprintf("??%s??", u)
	}

	return v
}
