package main

import (
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/searchmod"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
)

const (
	PathDataOrgs    = "data/organizations.json"
	PathDataUsers   = "data/users.json"
	PathDataTickets = "data/tickets.json"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	searchmod.InitOrgList(PathDataOrgs)
	searchmod.InitUserList(PathDataUsers)
	searchmod.InitTicketList(PathDataTickets)
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

