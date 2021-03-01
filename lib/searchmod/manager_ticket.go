package searchmod

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

type TicketResult struct {
	Ticket        models.Ticket
	AssigneeUser  models.User
	SubmittedUser models.User
	Org           models.Organization
}

var ticketList []models.Ticket

var (
	ticketByIdMap           = make(map[string]models.Ticket)
	ticketByUrlMap          = make(map[string][]models.Ticket)
	ticketByExternalIdMap   = make(map[string]models.Ticket)
	ticketByCreatedAtMap    = make(map[string][]models.Ticket)
	ticketByTypeMap         = make(map[string][]models.Ticket)
	ticketBySubjectMap      = make(map[string][]models.Ticket)
	ticketByDescriptionMap  = make(map[string][]models.Ticket)
	ticketByPriorityMap     = make(map[string][]models.Ticket)
	ticketByStatusMap       = make(map[string][]models.Ticket)
	ticketBySubmitterIdMap  = make(map[uint64][]models.Ticket)
	ticketByAssigneeIdMap   = make(map[uint64][]models.Ticket)
	ticketByOrgIdMap        = make(map[uint64][]models.Ticket)
	ticketByTagMap          = make(map[string][]models.Ticket)
	ticketByHasIncidentsMap = make(map[bool][]models.Ticket)
	ticketByDueAtMap        = make(map[string][]models.Ticket)
	ticketByViaMap          = make(map[string][]models.Ticket)
)

var searchableTicketFieldMap = make(map[string]bool)

func InitTicketList(path string) {
	err := utils.ReadJsonFile(path, &ticketList)
	if err != nil {
		fmt.Println(err)
	}
	mapTicketData()
	utils.ConvertStrListToMap(models.SearchableTicketFieldList, searchableTicketFieldMap)
}

func mapTicketData() {
	for _, ticket := range ticketList {
		ticketByIdMap[ticket.ID] = ticket
		ticketByUrlMap[ticket.URL] = append(ticketByUrlMap[ticket.URL], ticket)
		ticketByExternalIdMap[ticket.ExternalID] = ticket
		ticketByCreatedAtMap[ticket.CreatedAt] = append(ticketByCreatedAtMap[ticket.CreatedAt], ticket)
		ticketByTypeMap[ticket.Type] = append(ticketByTypeMap[ticket.Type], ticket)
		ticketBySubjectMap[ticket.Subject] = append(ticketBySubjectMap[ticket.Subject], ticket)
		ticketByDescriptionMap[ticket.Description] = append(ticketByDescriptionMap[ticket.Description], ticket)
		ticketByPriorityMap[ticket.Priority] = append(ticketByPriorityMap[ticket.Priority], ticket)
		ticketByStatusMap[ticket.Status] = append(ticketByStatusMap[ticket.Status], ticket)
		ticketBySubmitterIdMap[ticket.SubmitterId] = append(ticketBySubmitterIdMap[ticket.SubmitterId], ticket)
		ticketByAssigneeIdMap[ticket.AssigneeId] = append(ticketByAssigneeIdMap[ticket.AssigneeId], ticket)
		ticketByOrgIdMap[ticket.OrgID] = append(ticketByOrgIdMap[ticket.OrgID], ticket)
		mapTicketByTag(ticket)
		ticketByHasIncidentsMap[ticket.HasIncidents] = append(ticketByHasIncidentsMap[ticket.HasIncidents], ticket)
		ticketByDueAtMap[ticket.DueAt] = append(ticketByDueAtMap[ticket.DueAt], ticket)
		ticketByViaMap[ticket.Via] = append(ticketByViaMap[ticket.Via], ticket)
	}
}

func mapTicketByTag(ticket models.Ticket) {
	for _, tag := range ticket.Tags {
		ticketByTagMap[tag] = append(ticketByTagMap[tag], ticket)
	}
}

func SearchTickets(searchEntry meta.SearchEntry) (ticketResultList []TicketResult, err error) {
	if !searchableTicketFieldMap[searchEntry.Field] {
		return nil, errors.New("invalid term, please try again")
	}
	var ticketList []models.Ticket
	switch searchEntry.Field {
	case models.TicketFieldID:
		if ticket, ok := ticketByIdMap[searchEntry.Value]; ok {
			ticketList = append(ticketList, ticket)
		}
		break
	case models.TicketFieldURL:
		ticketList = ticketByUrlMap[searchEntry.Value]
		break
	case models.TicketFieldExternalID:
		if ticket, ok := ticketByExternalIdMap[searchEntry.Value]; ok {
			ticketList = append(ticketList, ticket)
		}
		break
	case models.TicketFieldCreatedAt:
		ticketList = ticketByCreatedAtMap[searchEntry.Value]
		break
	case models.TicketFieldType:
		ticketList = ticketByTypeMap[searchEntry.Value]
		break
	case models.TicketFieldSubject:
		ticketList = ticketBySubjectMap[searchEntry.Value]
		break
	case models.TicketFieldDescription:
		ticketList = ticketByDescriptionMap[searchEntry.Value]
		break
	case models.TicketFieldPriority:
		ticketList = ticketByPriorityMap[searchEntry.Value]
		break
	case models.TicketFieldStatus:
		ticketList = ticketByStatusMap[searchEntry.Value]
		break
	case models.TicketFieldSubmitterID:
		submitterId, err := utils.ParseUint64(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		ticketList = ticketBySubmitterIdMap[submitterId]
		break
	case models.TicketFieldAssigneeID:
		assigneeId, err := utils.ParseUint64(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		ticketList = ticketByAssigneeIdMap[assigneeId]
		break
	case models.TicketFieldOrgID:
		orgId, err := utils.ParseUint64(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		ticketList = ticketByOrgIdMap[orgId]
		break
	case models.TicketFieldTags:
		ticketList = ticketByTagMap[searchEntry.Value]
		break
	case models.TicketFieldHasIncidents:
		hasIncidents, err := strconv.ParseBool(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		ticketList = ticketByHasIncidentsMap[hasIncidents]
		break
	case models.TicketFieldDueAt:
		ticketList = ticketByDueAtMap[searchEntry.Value]
		break
	case models.TicketFieldVia:
		ticketList = ticketByViaMap[searchEntry.Value]
		break
	}

	if len(ticketList) > 0 {
		ticketResultList = transformTicketList(ticketList)
	}
	return ticketResultList, nil
}

func transformTicketList(ticketList []models.Ticket) (ticketResultList []TicketResult) {
	for _, ticket := range ticketList {
		var ticketResult TicketResult
		ticketResult.Ticket = ticket
		ticketResult.AssigneeUser = userByIdMap[ticket.AssigneeId]
		ticketResult.SubmittedUser = userByIdMap[ticket.SubmitterId]
		ticketResult.Org = orgByIdMap[ticket.OrgID]
		ticketResultList = append(ticketResultList, ticketResult)
	}
	return ticketResultList
}

func printTicketResult(ticketResultList []TicketResult) {
	if len(ticketResultList) == 0 {
		fmt.Println(meta.SearchNoResult)
	}
	fmt.Println("")
	for i, ticketResult := range ticketResultList {
		fmt.Println(fmt.Sprintf("Ticket %v:", i+1))
		printDisplayModel(ticketResult.Ticket)
		printOrgName(ticketResult.Org)
		printUserName(ticketResult.SubmittedUser, DisplayKeySubmitterName)
		printUserName(ticketResult.AssigneeUser, DisplayKeyAssigneeName)
	}
}
