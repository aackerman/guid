package guid

import (
	"fmt"
	"os"
)

func Generate() (string, error) {
	file, err := os.Open("/dev/urandom")
	defer file.Close()

	if err != nil {
		return "", err
	}

	b := make([]byte, 16)
	file.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

