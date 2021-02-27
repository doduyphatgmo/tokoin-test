package main

import (
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/searchmod"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
)

func main() {
	fmt.Print(meta.MsgSearchOptions)
	inputOptionStr, err := utils.GetConsoleInput()
	if err != nil {
		fmt.Printf("Error has occurred: %v\n", err)
	}
	err = searchmod.ExecuteSearchOption(inputOptionStr)
	if err != nil {
		fmt.Printf("Error has occurred: %v\n", err)
	}
}

