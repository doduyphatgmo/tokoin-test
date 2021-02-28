package searchmod

const (
	pathDataOrganizations = "data/organizations.json"
	pathDataUsers         = "data/users.json"
	pathDataTickets       = "data/tickets.json"
)

const (
	TitleItemUsers   = "Users"
	TitleItemTickets = "Tickets"
	TitleItemOrgs    = "Organizations"
)

const (
	DisplayKeyTicket          = "ticket"
	DisplayKeyAssigneeTicket  = "assignee_ticket"
	DisplayKeySubmittedTicket = "submitted_ticket"
	DisplayKeyOrgName         = "organization_name"
)

func init() {
	initOrgList()
	initUserList()
	initTicketList()
}
