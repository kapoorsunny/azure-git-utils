package file

import (
	"os"
)

func IsExist(file string) bool {
	//fmt.Println(file)
	if _, err := os.Stat(file); err == nil {
		return true

	} else if os.IsNotExist(err) {
		return false

	} else {
		return false
	}
}
