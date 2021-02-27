package searchmod

import (
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

type resultOrg struct {
	Org        models.Organization `json:"org"`
	UserList   []models.User       `json:"user_list"`
	TicketList []models.Ticket     `json:"ticket_list"`
}

var orgList []models.Organization
var orgByIdMap = make(map[uint64]models.Organization)

func initOrgList() {
	err := utils.ReadJsonFile(pathDataOrganizations, &orgList)
	if err != nil {
		fmt.Println(err)
	}
	mapOrgData()
}

func mapOrgData() {
	for _, org := range orgList {
		orgByIdMap[org.ID] = org
	}
}

func searchOrg(searchEntry meta.SearchEntry) ([]resultOrg, error) {
	var orgList []models.Organization
	switch searchEntry.Field {
	case models.OrgFieldID:
		orgID, err := utils.ParseUint64(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		if org, ok := orgByIdMap[orgID]; ok {
			orgList = append(orgList, org)
		}
		break
	}
	resultOrgList := transformOrgList(orgList)
	return resultOrgList, nil
}

func transformOrgList(orgList []models.Organization) []resultOrg {
	var resultOrgList []resultOrg
	var resultOrg resultOrg
	for _, org := range orgList {
		resultOrg.Org = org
		resultOrg.UserList = userByOrgIdMap[org.ID]
		resultOrg.TicketList = ticketByOrgIdMap[org.ID]
		resultOrgList = append(resultOrgList, resultOrg)
	}
	return resultOrgList
}
