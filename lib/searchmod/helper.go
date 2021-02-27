package searchmod

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
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
		value := reflect.ValueOf(org)
		type2 := value.Type()
		for i := 0; i < type2.NumField(); i++ {
			key := type2.Field(i).Tag.Get("json")
			v := value.Field(i).Interface()
			length := strings.Repeat(" ", 50 - len(key))
			s := fmt.Sprintf("%v%v%v", key, length, v)
			fmt.Printf(s + "\n")
		}
	}
}