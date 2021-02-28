package searchmod

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

func getSearchEntry() (searchEntry meta.SearchEntry, err error) {
	field, err := getSearchField()
	if err != nil {
		return
	}
	value, err := getSearchValue()
	if err != nil {
		return
	}
	searchEntry.Field = field
	searchEntry.Value = value
	return
}

func getSearchField() (string, error) {
	fmt.Println(meta.MsgSearchField)
	inputField, err := utils.GetConsoleInput()
	if err != nil {
		return "", err
	}
	return inputField, nil
}

func getSearchValue() (string, error) {
	fmt.Println(meta.MsgSearchValue)
	inputValue, err := utils.GetConsoleInput()
	if err != nil {
		return "", err
	}
	return inputValue, nil
}

func isRetrySearching() (bool, error) {
	fmt.Println(meta.MsgRetrySearching)
	inputStr, err := utils.GetConsoleInput()
	if err != nil {
		return false, err
	}
	input, err := strconv.Atoi(inputStr)
	if err != nil {
		return false, err
	}
	return input == meta.RetrySearchingInput, nil
}

func printDisplayModel(model interface{}) {
	modelValue := reflect.ValueOf(model)
	if modelValue.Kind() != reflect.Struct {
		println("can't print")
	} else {
		modelType := modelValue.Type()
		for i := 0; i < modelType.NumField(); i++ {
			displayKey := modelType.Field(i).Tag.Get("json")
			displayValue := modelValue.Field(i).Interface()

			if reflect.TypeOf(displayValue).Kind() == reflect.Slice {
				quoteValue := fmt.Sprintf("%q", displayValue)
				displayValue = strings.Join(strings.Split(quoteValue, " "), ", ")
			}

			lengthSpace := strings.Repeat(" ", meta.DisplaySpaceLength-len(displayKey))
			result := fmt.Sprintf("%v%v%v", displayKey, lengthSpace, displayValue)
			fmt.Printf(result + "\n")
		}
	}
}

func printUsername(userList []models.User) {
	for index, user := range userList {
		userNumber := fmt.Sprintf("user_%v", index)
		length := strings.Repeat(" ", meta.DisplaySpaceLength-len(userNumber))
		user := fmt.Sprintf("%v%v%v", userNumber, length, user.Name)
		fmt.Printf(user + "\n")
	}
}

func printTicketSubject(ticketList []models.Ticket, displayKey string) {
	for index, ticket := range ticketList {
		ticketNumber := fmt.Sprintf("%v_%v", displayKey, index)
		length := strings.Repeat(" ", meta.DisplaySpaceLength-len(ticketNumber))
		ticket := fmt.Sprintf("%v%v%v", ticketNumber, length, ticket.Subject)
		fmt.Printf(ticket + "\n")
	}
}

func printOrgName(org models.Organization) {
	length := strings.Repeat(" ", meta.DisplaySpaceLength-len(DisplayKeyOrgName))
	orgName := fmt.Sprintf("%v%v%v", DisplayKeyOrgName, length, org.Name)
	fmt.Printf(orgName + "\n")
}

func printSearchableFieldList(searchableFieldList []string) {
	for _, searchableField := range searchableFieldList {
		fmt.Println(searchableField)
	}
	fmt.Println("")
}

func printSearchableTitle(item string) {
	fmt.Println(strings.Repeat("-", meta.DisplayDashLength))
	fmt.Println(fmt.Sprintf("Search %v with", item))
}
