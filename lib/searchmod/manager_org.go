package searchmod

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

type OrgResult struct {
	Org        models.Organization `json:"org"`
	UserList   []models.User       `json:"user_list"`
	TicketList []models.Ticket     `json:"ticket_list"`
}

var orgList []models.Organization

var (
	orgByIdMap            = make(map[uint64]models.Organization)
	orgByUrlMap           = make(map[string][]models.Organization)
	orgByExternalIdMap    = make(map[string]models.Organization)
	orgByNameMap          = make(map[string][]models.Organization)
	orgByDomainNameMap    = make(map[string][]models.Organization)
	orgByCreatedAtMap     = make(map[string][]models.Organization)
	orgByDetailsMap       = make(map[string][]models.Organization)
	orgBySharedTicketsMap = make(map[bool][]models.Organization)
	orgByTagMap           = make(map[string][]models.Organization)
)

var searchableOrgFieldMap = make(map[string]bool)

func InitOrgList(path string) {
	err := utils.ReadJsonFile(path, &orgList)
	if err != nil {
		fmt.Println(err)
	}
	mapOrgData()
	utils.ConvertStrListToMap(models.SearchableOrgFieldList, searchableOrgFieldMap)
}

func mapOrgData() {
	for _, org := range orgList {
		orgByIdMap[org.ID] = org
		orgByUrlMap[org.URL] = append(orgByUrlMap[org.URL], org)
		orgByExternalIdMap[org.ExternalID] = org
		orgByNameMap[org.Name] = append(orgByNameMap[org.Name], org)
		mapOrgByDomainName(org)
		orgByCreatedAtMap[org.CreatedAt] = append(orgByNameMap[org.CreatedAt], org)
		orgByDetailsMap[org.Details] = append(orgByDetailsMap[org.Details], org)
		orgBySharedTicketsMap[org.SharedTickets] = append(orgBySharedTicketsMap[org.SharedTickets], org)
		mapOrgByTag(org)
	}
}

func mapOrgByDomainName(org models.Organization) {
	for _, domainName := range org.DomainNames {
		orgByDomainNameMap[domainName] = append(orgByDomainNameMap[domainName], org)
	}
}

func mapOrgByTag(org models.Organization) {
	for _, tag := range org.Tags {
		orgByTagMap[tag] = append(orgByTagMap[tag], org)
	}
}

func SearchOrgs(searchEntry meta.SearchEntry) (orgResultList []OrgResult, err error) {
	if !searchableOrgFieldMap[searchEntry.Field] {
		return nil, errors.New("invalid term, please try again")
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
		if org, ok := orgByExternalIdMap[searchEntry.Value]; ok {
			orgList = append(orgList, org)
		}
		break
	case models.OrgFieldName:
		orgList = orgByNameMap[searchEntry.Value]
		break
	case models.OrgFieldDomainNames:
		orgList = orgByDomainNameMap[searchEntry.Value]
		break
	case models.OrgFieldCreatedAt:
		orgList = orgByCreatedAtMap[searchEntry.Value]
		break
	case models.OrgFieldDetails:
		orgList = orgByDetailsMap[searchEntry.Value]
		break
	case models.OrgFieldSharedTickets:
		orgSharedTicket, err := strconv.ParseBool(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		orgList = orgBySharedTicketsMap[orgSharedTicket]
		break
	case models.OrgFieldTags:
		orgList = orgByTagMap[searchEntry.Value]
		break
	}

	if len(orgList) > 0 {
		orgResultList = transformOrgList(orgList)
	}
	return orgResultList, nil
}

func transformOrgList(orgList []models.Organization) (orgResultList []OrgResult) {
	for _, org := range orgList {
		var orgResult OrgResult
		orgResult.Org = org
		orgResult.UserList = userByOrgIdMap[org.ID]
		orgResult.TicketList = ticketByOrgIdMap[org.ID]
		orgResultList = append(orgResultList, orgResult)
	}
	return orgResultList
}

func printOrgResult(orgResultList []OrgResult) {
	if len(orgResultList) == 0 {
		fmt.Println(meta.SearchNoResult)
	}
	fmt.Println("")
	for i, orgResult := range orgResultList {
		fmt.Println(fmt.Sprintf("Organization %v:", i+1))
		printDisplayModel(orgResult.Org)
		printUserNameInList(orgResult.UserList)
		printTicketSubject(orgResult.TicketList, DisplayKeyTicket)
	}
}
