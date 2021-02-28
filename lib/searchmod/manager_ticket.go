package searchmod

import (
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

var ticketList []models.Ticket

var (
	ticketByOrgIdMap       = make(map[uint64][]models.Ticket)
	ticketByAssigneeIdMap  = make(map[uint64][]models.Ticket)
	ticketBySubmitterIdMap = make(map[uint64][]models.Ticket)
)

var searchableTicketFieldMap = make(map[string]bool)

func initTicketList() {
	err := utils.ReadJsonFile(pathDataTickets, &ticketList)
	if err != nil {
		fmt.Println(err)
	}
	mapTicketData()
	convertSearchableListToMap(models.SearchableTicketFieldList, searchableTicketFieldMap)
}

func mapTicketData() {
	for _, ticket := range ticketList {
		ticketByOrgIdMap[ticket.OrgID] = append(ticketByOrgIdMap[ticket.OrgID], ticket)
		ticketByAssigneeIdMap[ticket.AssigneeId] = append(ticketByAssigneeIdMap[ticket.AssigneeId], ticket)
		ticketBySubmitterIdMap[ticket.SubmitterId] = append(ticketBySubmitterIdMap[ticket.SubmitterId], ticket)
	}
}
