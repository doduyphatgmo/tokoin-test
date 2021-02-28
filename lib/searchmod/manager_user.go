package searchmod

import (
	"errors"
	"fmt"
	"strconv"

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
	userByIdMap          = make(map[uint64]models.User)
	userByUrlMap         = make(map[string][]models.User)
	userByExternalIdMap  = make(map[string]models.User)
	userByNameMap        = make(map[string][]models.User)
	userByAliasMap       = make(map[string][]models.User)
	userByCreatedAtMap   = make(map[string][]models.User)
	userByActiveMap      = make(map[bool][]models.User)
	userByVerifiedMap    = make(map[bool][]models.User)
	userBySharedMap      = make(map[bool][]models.User)
	userByLocaleMap      = make(map[string][]models.User)
	userByTimezoneMap    = make(map[string][]models.User)
	userByLastLoginAtMap = make(map[string][]models.User)
	userByEmailMap       = make(map[string][]models.User)
	userByPhoneMap       = make(map[string][]models.User)
	userBySignatureMap   = make(map[string][]models.User)
	userByOrgIdMap       = make(map[uint64][]models.User)
	userByTagMap         = make(map[string][]models.User)
	userBySuspendedMap   = make(map[bool][]models.User)
	userByRoleMap        = make(map[string][]models.User)
)

var searchableUserFieldMap = make(map[string]bool)

func initUserList() {
	err := utils.ReadJsonFile(pathDataUsers, &userList)
	if err != nil {
		fmt.Println(err)
	}
	mapUserData()
	utils.ConvertStrListToMap(models.SearchableUserFieldList, searchableUserFieldMap)
}

func mapUserData() {
	for _, user := range userList {
		userByIdMap[user.ID] = user
		userByUrlMap[user.URL] = append(userByUrlMap[user.URL], user)
		userByExternalIdMap[user.ExternalID] = user
		userByNameMap[user.Name] = append(userByNameMap[user.Name], user)
		userByAliasMap[user.Alias] = append(userByAliasMap[user.Alias], user)
		userByCreatedAtMap[user.CreatedAt] = append(userByCreatedAtMap[user.CreatedAt], user)
		userByActiveMap[user.Active] = append(userByActiveMap[user.Active], user)
		userByVerifiedMap[user.Verified] = append(userByVerifiedMap[user.Verified], user)
		userBySharedMap[user.Shared] = append(userBySharedMap[user.Shared], user)
		userByLocaleMap[user.Locale] = append(userByLocaleMap[user.Locale], user)
		userByTimezoneMap[user.Timezone] = append(userByTimezoneMap[user.Timezone], user)
		userByLastLoginAtMap[user.LastLogicAt] = append(userByLastLoginAtMap[user.LastLogicAt], user)
		userByEmailMap[user.Email] = append(userByEmailMap[user.Email], user)
		userByPhoneMap[user.Phone] = append(userByPhoneMap[user.Phone], user)
		userBySignatureMap[user.Signature] = append(userBySignatureMap[user.Signature], user)
		userByOrgIdMap[user.OrgID] = append(userByOrgIdMap[user.OrgID], user)
		mapUserByTag(user)
		userBySuspendedMap[user.Suspended] = append(userBySuspendedMap[user.Suspended], user)
		userByRoleMap[user.Role] = append(userByRoleMap[user.Role], user)
	}
}

func mapUserByTag(user models.User) {
	for _, tag := range user.Tags {
		userByTagMap[tag] = append(userByTagMap[tag], user)
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
	case models.UserFieldURL:
		userList = userByUrlMap[searchEntry.Value]
		break
	case models.UserFieldExternalID:
		if user, ok := userByExternalIdMap[searchEntry.Value]; ok {
			userList = append(userList, user)
		}
		break
	case models.UserFieldName:
		userList = userByNameMap[searchEntry.Value]
		break
	case models.UserFieldAlias:
		userList = userByAliasMap[searchEntry.Value]
		break
	case models.UserFieldCreatedAt:
		userList = userByCreatedAtMap[searchEntry.Value]
		break
	case models.UserFieldActive:
		isActive, err := strconv.ParseBool(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		userList = userByActiveMap[isActive]
		break
	case models.UserFieldVerified:
		isVerified, err := strconv.ParseBool(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		userList = userByVerifiedMap[isVerified]
		break
	case models.UserFieldShared:
		isShared, err := strconv.ParseBool(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		userList = userBySharedMap[isShared]
		break
	case models.UserFieldLocale:
		userList = userByLocaleMap[searchEntry.Value]
		break
	case models.UserFieldTimezone:
		userList = userByTimezoneMap[searchEntry.Value]
		break
	case models.UserFieldLastLoginAt:
		userList = userByLastLoginAtMap[searchEntry.Value]
		break
	case models.UserFieldEmail:
		userList = userByEmailMap[searchEntry.Value]
		break
	case models.UserFieldPhone:
		userList = userByPhoneMap[searchEntry.Value]
		break
	case models.UserFieldSignature:
		userList = userBySignatureMap[searchEntry.Value]
		break
	case models.UserFieldOrgID:
		orgID, err := utils.ParseUint64(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		userList = userByOrgIdMap[orgID]
		break
	case models.UserFieldTags:
		userList = userByTagMap[searchEntry.Value]
		break
	case models.UserFieldSuspended:
		isSuspended, err := strconv.ParseBool(searchEntry.Value)
		if err != nil {
			return nil, err
		}
		userList = userBySuspendedMap[isSuspended]
		break
	case models.UserFieldRole:
		userList = userByRoleMap[searchEntry.Value]
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
