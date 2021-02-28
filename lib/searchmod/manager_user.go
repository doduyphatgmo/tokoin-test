package searchmod

import (
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

type userResult struct {
	User models.User
}

var userList []models.User
var userByOrgIdMap = make(map[uint64][]models.User)

var searchableUserFieldMap = make(map[string]bool)

func initUserList() {
	err := utils.ReadJsonFile(pathDataUsers, &userList)
	if err != nil {
		fmt.Println(err)
	}
	mapUserData()
	convertSearchableListToMap(models.SearchableUserFieldList, searchableUserFieldMap)
}

func mapUserData()  {
	for _, user := range userList {
		userByOrgIdMap[user.OrgID] = append(userByOrgIdMap[user.OrgID], user)
	}
}

//func searchUser() {
//
//}
