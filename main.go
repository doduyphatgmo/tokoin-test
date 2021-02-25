package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/doduyphatgmo/tokoin-test/meta"
	"github.com/doduyphatgmo/tokoin-test/utils"
)

func main() {
	err := initMenuSearch()
	if err != nil {
		fmt.Printf("Error has occurred: %v\n", err)
	}
}

func initMenuSearch() error {
	fmt.Print(meta.MsgSearchOptions)
	inputOptionStr, err := utils.GetConsoleInput()
	if err != nil {
		return err
	}
	err = searchOption(inputOptionStr)
	if err != nil {
		return err
	}
	return nil
}

func searchOption(optionStr string) error {
	optionStr = strings.TrimSpace(optionStr)
	if optionStr == meta.OptionQuit {
		return errors.New("stoppppppp")
	}
	option, err := strconv.Atoi(optionStr)
	if err != nil {
		return err
	}
	switch option {
	case meta.OptionSearch:
		err = selectObject()
		break
	case meta.OptionViewFields:
		//viewListFields()
		break
	}
	return nil
}

func selectObject() error {
	fmt.Println(meta.MsgSelectObject)
	selectedInputItemStr, err := utils.GetConsoleInput()
	if err != nil {
		return err
	}
	selectedInputItem, err := strconv.Atoi(strings.TrimSpace(selectedInputItemStr))
	if err != nil {
		return err
	}
	switch selectedInputItem {
	case meta.ItemUsers:
		fmt.Println("Item users")
		break
	case meta.ItemTickets:
		fmt.Println("Item ticket")
		break
	case meta.ItemOrganizations:
		fmt.Println("Item Organizations")
		break
	}
	return nil
}
