package searchmod

import (
	"errors"
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

type orgResult struct {
	Org        models.Organization `json:"org"`
	UserList   []models.User       `json:"user_list"`
	TicketList []models.Ticket     `json:"ticket_list"`
}

var orgList []models.Organization

var (
	orgByIdMap         = make(map[uint64]models.Organization)
	orgByUrlMap        = make(map[string][]models.Organization)
	orgByExternalIdMap = make(map[string][]models.Organization)
)

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
		orgByUrlMap[org.URL] = append(orgByUrlMap[org.URL], org)
		orgByExternalIdMap[org.ExternalID] = append(orgByExternalIdMap[org.ExternalID], org)
	}
}

func searchOrg(searchEntry meta.SearchEntry) (orgResultList []orgResult, err error) {
	if !models.SearchableOrgFieldsMap[searchEntry.Field] {
		return nil, errors.New("wrong field, try again")
	}
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
	case models.OrgFieldURL:
		orgList = orgByUrlMap[searchEntry.Value]
		break
	case models.OrgFieldExternalID:
		orgList = orgByExternalIdMap[searchEntry.Value]
		break
	}

	if len(orgList) > 0 {
		orgResultList = transformOrgList(orgList)
	}
	return orgResultList, nil
}

func transformOrgList(orgList []models.Organization) []orgResult {
	var orgResultList []orgResult
	for _, org := range orgList {
		var orgResult orgResult
		orgResult.Org = org
		orgResult.UserList = userByOrgIdMap[org.ID]
		orgResult.TicketList = ticketByOrgIdMap[org.ID]
		orgResultList = append(orgResultList, orgResult)
	}
	return orgResultList
}
