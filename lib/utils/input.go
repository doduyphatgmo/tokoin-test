package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
)

func GetConsoleInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	if input == meta.QuitInput {
		panic(fmt.Sprintf("\n%v", "Exit application"))
	}
	return input, nil
}
