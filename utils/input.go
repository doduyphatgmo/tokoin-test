package utils

import (
	"bufio"
	"os"
)

func GetConsoleInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input, nil
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	fmt.Println( )
	//}
	//if scanner.Err() != nil {
	//	return nil, scanner.Err()
	//}
	//return scanner, nil
}
