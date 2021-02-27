package searchmod

import (
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

type resultOrg struct {
	Org     models.Organization `json:"org"`
	Users   []models.User       `json:"users"`
	Tickets []models.Ticket     `json:"tickets"`
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

func searchOrg(searchEntry meta.SearchEntry) (resultOrgList []resultOrg, err error) {
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
	for _,org := range orgList {
		var resultOrg resultOrg
		resultOrg.Org = org
		resultOrg.Users = userByOrgIdMap[org.ID]
		resultOrgList = append(resultOrgList, resultOrg)
	}
	return resultOrgList, nil
}
