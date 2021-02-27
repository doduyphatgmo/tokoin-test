package searchmod

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
)

func ExecuteSearchOption(optionStr string) error {
	if optionStr == meta.OptionQuit {
		return errors.New("stoppppppp")
	}
	option, err := strconv.Atoi(optionStr)
	if err != nil {
		return err
	}
	switch option {
	case meta.OptionSearch:
		err = executeOptionSearch()
		break
	case meta.OptionViewFields:
		//viewListFields()
		break
	}
	return nil
}

func executeOptionSearch() error {
	fmt.Println(meta.MsgSelectItem)
	inputItemStr, err := utils.GetConsoleInput()
	if err != nil {
		return err
	}
	inputItem, err := strconv.Atoi(inputItemStr)
	if err != nil {
		return err
	}

	searchEntry, err := GetSearchEntry()
	if err != nil {
		return err
	}

	switch inputItem {
	case meta.ItemOrganizations:
		orgList, err := searchOrg(searchEntry)
		if err != nil {
			return err
		}
		printOrgResult(orgList)
		break
	//case meta.ItemUsers:
	//	searchEntry, err := GetSearchEntry()
	//	if err != nil {
	//		return err
	//	}
	//	break
	//case meta.ItemTickets:
	//	searchEntry, err := GetSearchEntry()
	//	if err != nil {
	//		return err
	//	}
	//	break
	}
	return nil
}

