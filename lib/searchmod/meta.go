package searchmod

const (
	pathDataOrganizations = "data/organizations.json"
	pathDataUsers         = "data/users.json"
	pathDataTickets       = "data/tickets.json"
)

func init() {
	initOrgList()
	initUserList()
}
