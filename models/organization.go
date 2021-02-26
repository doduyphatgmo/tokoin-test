package models

type Organization struct {
	ID            uint64   `json:"_id"`
	URL           string   `json:"url"`
	ExternalID    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}

const (
	OrgFieldID            = "_id"
	OrgFieldURL           = "url"
	OrgFieldExternalID    = "external_id"
	OrgFieldName          = "name"
	OrgFieldDomainNames   = "domain_names"
	OrgFieldCreatedAt     = "created_at"
	OrgFieldDetails       = "details"
	OrgFieldSharedTickets = "shared_tickets"
	OrgFieldTags          = "tags"
)
