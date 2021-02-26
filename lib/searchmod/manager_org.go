package searchmod

import (
	"fmt"
	"strconv"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

var OrgList []models.Organization
var orgMapById = make(map[uint64]models.Organization)

func initOrganizationList() {
	err := utils.ReadJsonFile(pathDataOrganizations, &OrgList)
	if err != nil {
		fmt.Println(err)
	}
	mapOrgData()
}

func mapOrgData() {
	for _, org := range OrgList {
		orgMapById[org.ID] = org
	}
}

func searchOrg(searchEntry meta.SearchEntry) (orgList []models.Organization, err error) {
	switch searchEntry.Field {
	case models.OrgFieldID:
		orgID, err := strconv.ParseUint(searchEntry.Value, 10, 64)
		if err != nil {
			return nil, err
		}
		if org, ok := orgMapById[orgID]; ok {
			orgList = append(orgList, org)
		}
		break
	}
	return
}
