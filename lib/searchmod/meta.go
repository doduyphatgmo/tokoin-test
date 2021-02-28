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

func init() {
	initOrgList()
	initUserList()
	initTicketList()
}
