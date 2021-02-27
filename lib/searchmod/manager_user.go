package searchmod

import (
	"fmt"

	"github.com/doduyphatgmo/tokoin-test/lib/utils"
	"github.com/doduyphatgmo/tokoin-test/models"
)

var userList []models.User
var userByOrgIdMap = make(map[uint64][]models.User)

func initUserList() {
	err := utils.ReadJsonFile(pathDataUsers, &userList)
	if err != nil {
		fmt.Println(err)
	}
	mapUserData()
}

func mapUserData()  {
	for _, user := range userList {
		userByOrgIdMap[user.OrgID] = append(userByOrgIdMap[user.OrgID], user)
	}
}
