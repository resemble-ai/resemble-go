package response

import (
	"os"
)

func loadTestData() *os.File {
	f, err := os.Open("test.wav")
	if err != nil {
		panic(err)
	}
	return f
}
