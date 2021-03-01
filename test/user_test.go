package test

import (
	"testing"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/searchmod"
	"github.com/doduyphatgmo/tokoin-test/models"
	"github.com/stretchr/testify/assert"
)

type testUser struct {
	input          meta.SearchEntry
	expectedResult testUserResult
	actualResult   testUserResult
}

type testUserResult struct {
	userListLength            int
	orgTotal                  int
	submittedTicketListLength int
	assigneeTicketListLength  int
}

var testUserCases = []testUser{
	{
		input:          meta.SearchEntry{Field: models.UserFieldID, Value: "1"},
		expectedResult: testUserResult{1, 1, 2, 2},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldURL, Value: "http://initech.tokoin.io.com/api/v2/users/2.json"},
		expectedResult: testUserResult{1, 1, 0, 2},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldExternalID, Value: "85c599c1-ebab-474d-a4e6-32f1c06e8730"},
		expectedResult: testUserResult{1, 1, 5, 2},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldName, Value: "Rose Newton"},
		expectedResult: testUserResult{1, 1, 4, 0},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldAlias, Value: "Mr Ola"},
		expectedResult: testUserResult{1, 1, 2, 4},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldCreatedAt, Value: "2016-04-04T01:30:49 -10:00"},
		expectedResult: testUserResult{1, 1, 0, 4},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldActive, Value: "false"},
		expectedResult: testUserResult{36, 35, 94, 108},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldVerified, Value: "true"},
		expectedResult: testUserResult{26, 26, 61, 61},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldShared, Value: "true"},
		expectedResult: testUserResult{28, 28, 78, 60},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldLocale, Value: "en-AU"},
		expectedResult: testUserResult{32, 32, 93, 86},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldTimezone, Value: "Falkland Islands (Malvinas)"},
		expectedResult: testUserResult{1, 1, 3, 3},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldLastLoginAt, Value: "2013-05-11T07:33:30 -10:00"},
		expectedResult: testUserResult{1, 1, 3, 3},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldEmail, Value: "martinajoseph@flotonic.com"},
		expectedResult: testUserResult{1, 1, 3, 3},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldPhone, Value: "8294-733-398"},
		expectedResult: testUserResult{1, 1, 0, 5},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldSignature, Value: "Don't Worry Be Happy!"},
		expectedResult: testUserResult{75, 72, 199, 195},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldOrgID, Value: "122"},
		expectedResult: testUserResult{4, 4, 13, 10},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldTags, Value: "Fairhaven"},
		expectedResult: testUserResult{1, 1, 5, 2},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldSuspended, Value: "False"},
		expectedResult: testUserResult{39, 37, 96, 102},
	},
	{
		input:          meta.SearchEntry{Field: models.UserFieldRole, Value: "agent"},
		expectedResult: testUserResult{25, 23, 62, 60},
	},
}

func TestSearchUsers(t *testing.T) {
	for _, testCase := range testUserCases {
		userResultList, err := searchmod.SearchUsers(testCase.input)
		assert.Nil(t, err)
		assert.Equal(
			t,
			testCase.expectedResult,
			getActualUserResult(userResultList),
		)
	}
}

func getActualUserResult(userResultList []searchmod.UserResult) (actualResult testUserResult) {
	actualResult.userListLength = len(userResultList)
	for _, userResult := range userResultList {
		if userResult.Org.ID > 0 {
			actualResult.orgTotal += 1
		}
		actualResult.submittedTicketListLength += len(userResult.SubmittedTicketList)
		actualResult.assigneeTicketListLength += len(userResult.AssigneeTicketList)
	}
	return actualResult
}
