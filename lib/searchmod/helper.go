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

func printOrgResult(orgList []resultOrg) {
	for _, org := range orgList {
		value := reflect.ValueOf(org.Org)
		type2 := value.Type()
		for i := 0; i < type2.NumField(); i++ {
			key := type2.Field(i).Tag.Get("json")
			v := value.Field(i).Interface()

			rt := reflect.TypeOf(v)
			if rt.Kind() == reflect.Slice {
				a := fmt.Sprintf("%q", v)
				v = strings.Join(strings.Split(a, " "), ", ")
			}

			length := strings.Repeat(" ", 50-len(key))
			s := fmt.Sprintf("%v%v%s", key, length, v)
			fmt.Printf(s + "\n")
		}
		printOrgUser(org.UserList)
		printOrgTicket(org.TicketList)
	}
}

func printOrgUser(userList []models.User) {
	for index, user := range userList {
		key2 := fmt.Sprintf("user_%v", index)
		length := strings.Repeat(" ", 50-len(key2))
		s2 := fmt.Sprintf("%v%v%v", key2, length, user.Name)
		fmt.Printf(s2 + "\n")
	}
}

func printOrgTicket(ticketList []models.Ticket) {
	for index, ticket := range ticketList {
		key2 := fmt.Sprintf("ticket_%v", index)
		length := strings.Repeat(" ", 50-len(key2))
		s2 := fmt.Sprintf("%v%v%v", key2, length, ticket.Subject)
		fmt.Printf(s2 + "\n")
	}
}
