package searchmod

import (
	"errors"
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

type userResult struct {
	User                models.User
	AssigneeTicketList  []models.Ticket
	SubmittedTicketList []models.Ticket
	Org                 models.Organization
}

var userList []models.User

var (
	userByIdMap    = make(map[uint64]models.User)
	userByOrgIdMap = make(map[uint64][]models.User)
)

var searchableUserFieldMap = make(map[string]bool)

func initUserList() {
	err := utils.ReadJsonFile(pathDataUsers, &userList)
	if err != nil {
		fmt.Println(err)
	}
	mapUserData()
	convertSearchableListToMap(models.SearchableUserFieldList, searchableUserFieldMap)
}

func mapUserData() {
	for _, user := range userList {
		userByIdMap[user.ID] = user
		userByOrgIdMap[user.OrgID] = append(userByOrgIdMap[user.OrgID], user)
	}
}

func searchUser(searchEntry meta.SearchEntry) (userResultList []userResult, err error) {
	if !searchableUserFieldMap[searchEntry.Field] {
		return nil, errors.New("invalid term, please try again")
	}
	var userList []models.User
	switch searchEntry.Field {
	case models.UserFieldID:
		userID, err := utils.ParseUint64(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		if user, ok := userByIdMap[userID]; ok {
			userList = append(userList, user)
		}
		break
	}
	if len(userList) > 0 {
		userResultList = transformUserList(userList)
	}
	return userResultList, nil
}

func transformUserList(userList []models.User) (userResultList []userResult) {
	for _, user := range userList {
		var userResult userResult
		userResult.User = user
		userResult.AssigneeTicketList = ticketByAssigneeIdMap[user.ID]
		userResult.SubmittedTicketList = ticketBySubmitterIdMap[user.ID]
		userResult.Org = orgByIdMap[user.OrgID]
		userResultList = append(userResultList, userResult)
	}
	return userResultList
}

func printUserResult(userResultList []userResult) {
	if len(userResultList) == 0 {
		fmt.Println(meta.SearchNoResult)
	}
	fmt.Println("")
	for i, userResult := range userResultList {
		fmt.Println(fmt.Sprintf("User %v:", i+1))
		printDisplayModel(userResult.User)
		printOrgName(userResult.Org)
		printTicketSubject(userResult.AssigneeTicketList, DisplayKeyAssigneeTicket)
		printTicketSubject(userResult.SubmittedTicketList, DisplayKeySubmittedTicket)
	}
}
