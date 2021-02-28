package main

import (
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/searchmod"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	for {
		fmt.Print(meta.MsgSearchOptions)
		selectedOptionStr, err := utils.GetConsoleInput()
		if err != nil {
			fmt.Printf("Error has occurred: %v\n", err)
		}
		err = searchmod.ExecuteSelectedOption(selectedOptionStr)
		if err != nil {
			fmt.Printf("Error just occurred: %v. Please try again\n", err)
		}
	}
}

