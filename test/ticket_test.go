package test

import (
	"testing"

	"github.com/doduyphatgmo/tokoin-test/lib/meta"
	"github.com/doduyphatgmo/tokoin-test/lib/searchmod"
	"github.com/doduyphatgmo/tokoin-test/models"
	"github.com/stretchr/testify/assert"
)

type testTicket struct {
	input          meta.SearchEntry
	expectedResult testTicketResult
	actualResult   testTicketResult
}

type testTicketResult struct {
	ticketTotal        int
	assigneeUserTotal  int
	submittedUserTotal int
	orgTotal           int
}

var testTicketCases = []testTicket{
	{
		input:          meta.SearchEntry{Field: models.TicketFieldID, Value: "fa3a37e3-942e-4048-81bc-d0d7e79cb686"},
		expectedResult: testTicketResult{1, 1, 1, 1},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldURL, Value: "http://initech.tokoin.io.com/api/v2/tickets/140e0cd4-c31b-4e90-833d-c42a12d4b713.json"},
		expectedResult: testTicketResult{1, 1, 1, 1},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldExternalID, Value: "dd0229db-7361-4d4e-bbe3-235024436c4b"},
		expectedResult: testTicketResult{1, 1, 1, 1},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldCreatedAt, Value: "2016-03-12T11:33:46 -11:00"},
		expectedResult: testTicketResult{1, 1, 1, 1},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldType, Value: "incidentttt"},
		expectedResult: testTicketResult{0, 0, 0, 0},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldSubject, Value: "A Nuisance in Macedonia"},
		expectedResult: testTicketResult{1, 1, 1, 1},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldDescription, Value: "Incididunt et laborum tempor nulla. Exercitation laborum est nulla aliquip fugiat nisi excepteur do labore Lorem in."},
		expectedResult: testTicketResult{1, 1, 1, 1},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldPriority, Value: "normalll"},
		expectedResult: testTicketResult{0, 0, 0, 0},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldStatus, Value: "pending"},
		expectedResult: testTicketResult{45, 44, 45, 44},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldSubmitterID, Value: "45"},
		expectedResult: testTicketResult{2, 2, 2, 2},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldAssigneeID, Value: "30"},
		expectedResult: testTicketResult{2, 2, 2, 2},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldOrgID, Value: "121"},
		expectedResult: testTicketResult{5, 5, 5, 5},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldTags, Value: "California"},
		expectedResult: testTicketResult{14, 14, 14, 14},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldHasIncidents, Value: "false"},
		expectedResult: testTicketResult{101, 97, 101, 98},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldDueAt, Value: "2016-08-06T04:15:59 -10:00"},
		expectedResult: testTicketResult{1, 1, 1, 1},
	},
	{
		input:          meta.SearchEntry{Field: models.TicketFieldVia, Value: "webbbbb"},
		expectedResult: testTicketResult{0, 0, 0, 0},
	},
}

func TestSearchTickets(t *testing.T) {
	for _, testCase := range testTicketCases {
		ticketResultList, err := searchmod.SearchTickets(testCase.input)
		assert.Nil(t, err)
		assert.Equal(
			t,
			testCase.expectedResult,
			getActualTicketResult(ticketResultList),
		)
	}
}

func getActualTicketResult(ticketResultList []searchmod.TicketResult) (actualResult testTicketResult) {
	for _, ticketResult := range ticketResultList {
		if ticketResult.Ticket.ID != "" {
			actualResult.ticketTotal += 1
		}
		if ticketResult.AssigneeUser.ID > 0 {
			actualResult.assigneeUserTotal += 1
		}
		if ticketResult.SubmittedUser.ID > 0 {
			actualResult.submittedUserTotal += 1
		}
		if ticketResult.Org.ID > 0 {
			actualResult.orgTotal += 1
		}
	}
	return actualResult
}
