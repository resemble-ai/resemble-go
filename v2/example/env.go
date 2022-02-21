package example

import (
	"fmt"
	"os"
)

func LoadConfigByKey(k string) string {
	v, f := os.LookupEnv(k)
	if !f {
		panic(fmt.Sprintf("key: %s not found on your environment, Please set it", k))
	}

	return v
}
