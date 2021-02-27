package searchmod

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

func GetSearchEntry() (searchEntry meta.SearchEntry, err error) {
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

func printOrgResult(orgResultList []orgResult) {
	if len(orgResultList) == 0 {
		fmt.Println(meta.SearchNoResult)
	}
	for _, orgResult := range orgResultList {
		printDisplayModel(orgResult.Org)
		printUsername(orgResult.UserList)
		printTicketSubject(orgResult.TicketList)
	}
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
		key2 := fmt.Sprintf("user_%v", index)
		length := strings.Repeat(" ", 50-len(key2))
		s2 := fmt.Sprintf("%v%v%v", key2, length, user.Name)
		fmt.Printf(s2 + "\n")
	}
}

func printTicketSubject(ticketList []models.Ticket) {
	for index, ticket := range ticketList {
		key2 := fmt.Sprintf("ticket_%v", index)
		length := strings.Repeat(" ", 50-len(key2))
		s2 := fmt.Sprintf("%v%v%v", key2, length, ticket.Subject)
		fmt.Printf(s2 + "\n")
	}
}
