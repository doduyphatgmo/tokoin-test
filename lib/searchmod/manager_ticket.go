package searchmod

import (
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

var ticketList []models.Ticket
var ticketByOrgIdMap = make(map[uint64][]models.Ticket)

func initTicketList() {
	err := utils.ReadJsonFile(pathDataTickets, &ticketList)
	if err != nil {
		fmt.Println(err)
	}
	mapTicketData()
}

func mapTicketData()  {
	for _, ticket := range ticketList {
		ticketByOrgIdMap[ticket.OrgID] = append(ticketByOrgIdMap[ticket.OrgID], ticket)
	}
}
