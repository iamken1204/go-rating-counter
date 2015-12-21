package rating_counter

import (
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
