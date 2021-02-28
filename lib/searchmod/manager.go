package searchmod

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

func ExecuteSelectedOption(optionStr string) error {
	option, err := strconv.Atoi(optionStr)
	if err != nil {
		return err
	}
	switch option {
	case meta.OptionSearch:
		err = executeOptionSearch()
		if err != nil {
			return err
		}
		break
	case meta.OptionViewFields:
		viewSearchableFields()
		break
	default:
		return errors.New("wrong option. try again")
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

	isRetry := true
	for isRetry {
		searchEntry, err := getSearchEntry()
			if err != nil {
				return err
			}

		switch inputItem {
		case meta.ItemOrganizations:
			orgResultList, err := searchOrg(searchEntry)
			if err != nil {
				fmt.Println(err)
				break
			}
			printOrgResult(orgResultList)
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
		isRetry, err = isRetrySearching()
		if err != nil {
			return err
		}
	}
	return nil
}

func viewSearchableFields() {
	fmt.Println(strings.Repeat("-", meta.DisplaySpaceLength))
	fmt.Println("Search Users with")
	printSearchableFields(models.SearchableUserFieldsMap)

	fmt.Println(strings.Repeat("-", meta.DisplaySpaceLength))
	fmt.Println("Search Tickets with")
	printSearchableFields(models.SearchableTicketFieldsMap)

	fmt.Println(strings.Repeat("-", meta.DisplaySpaceLength))
	fmt.Println("Search Orgs with")
	printSearchableFields(models.SearchableOrgFieldsMap)
}

func printSearchableFields(searchableFieldsMap map[string]bool) {
	for searchableField, _ := range searchableFieldsMap{
		fmt.Println(searchableField)
	}
}