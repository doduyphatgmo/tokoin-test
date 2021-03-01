package test

import (
	"testing"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/searchmod"
	"github.com/doduyphatgmo/tokoin-test/models"
	"github.com/stretchr/testify/assert"
)

type testOrg struct {
	input          meta.SearchEntry
	expectedResult testOrgResult
}

type testOrgResult struct {
	orgListLength    int
	userListLength   int
	ticketListLength int
}

var testOrgCases = []testOrg{
	{
		input:          meta.SearchEntry{Field: models.OrgFieldID, Value: "567"},
		expectedResult: testOrgResult{0, 0, 0},
	},
	{
		input:          meta.SearchEntry{Field: models.OrgFieldID, Value: "105"},
		expectedResult: testOrgResult{1, 2, 11},
	},
	{
		input:          meta.SearchEntry{Field: models.OrgFieldURL, Value: "http://initech.tokoin.io.com/api/v2/organizations/105.json"},
		expectedResult: testOrgResult{1, 2, 11},
	},
	{
		input:          meta.SearchEntry{Field: models.OrgFieldExternalID, Value: "52f12203-6112-4fb9-aadc-70a6c816d605"},
		expectedResult: testOrgResult{1, 2, 11},
	},
	{
		input:          meta.SearchEntry{Field: models.OrgFieldName, Value: "Qualitern"},
		expectedResult: testOrgResult{1, 4, 7},
	},
	{
		input:          meta.SearchEntry{Field: models.OrgFieldDomainNames, Value: "gology.com"},
		expectedResult: testOrgResult{1, 4, 7},
	},
	{
		input:          meta.SearchEntry{Field: models.OrgFieldCreatedAt, Value: "2016-07-23T09:48:02 -10:00"},
		expectedResult: testOrgResult{1, 4, 7},
	},
	{
		input:          meta.SearchEntry{Field: models.OrgFieldDetails, Value: "Artisân"},
		expectedResult: testOrgResult{2, 9, 12},
	},
	{
		input:          meta.SearchEntry{Field: models.OrgFieldSharedTickets, Value: "true"},
		expectedResult: testOrgResult{10, 33, 79},
	},
	{
		input:          meta.SearchEntry{Field: models.OrgFieldTags, Value: "Burton"},
		expectedResult: testOrgResult{1, 5, 5},
	},
}

func TestSearchOrgs(t *testing.T) {
	for _, testCase := range testOrgCases {
		orgResultList, err := searchmod.SearchOrgs(testCase.input)
		assert.Nil(t, err)
		assert.Equal(
			t,
			testCase.expectedResult,
			getActualOrgResult(orgResultList),
		)
	}
}

func getActualOrgResult(orgResultList []searchmod.OrgResult) (actualResult testOrgResult) {
	actualResult.orgListLength = len(orgResultList)
	for _, orgResult := range orgResultList {
		actualResult.userListLength += len(orgResult.UserList)
		actualResult.ticketListLength += len(orgResult.TicketList)
	}
	return actualResult
}
