package test

import "github.com/doduyphatgmo/tokoin-test/lib/searchmod"

const (
	PathDataOrgs    = "../data/organizations.json"
	PathDataUsers   = "../data/users.json"
	PathDataTickets = "../data/tickets.json"
)

func init() {
	searchmod.InitOrgList(PathDataOrgs)
	searchmod.InitUserList(PathDataUsers)
	searchmod.InitTicketList(PathDataTickets)
}
