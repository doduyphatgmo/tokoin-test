package main

import (
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/searchmod"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
)

func main() {
	for {
		fmt.Print(meta.MsgSearchOptions)
		selectedOptionStr, err := utils.GetConsoleInput()
		if err != nil {
			fmt.Printf("Error has occurred: %v\n", err)
		}
		if selectedOptionStr == meta.OptionQuit {
			fmt.Println("stoppppppp")
			break
		}
		err = searchmod.ExecuteSelectedOption(selectedOptionStr)
		if err != nil {
			fmt.Printf("Error just occurred: %v. Please try again\n", err)
		}
	}
}

