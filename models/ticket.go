package models

type Ticket struct {
	ID           string   `json:"_id"`
	URL          string   `json:"url"`
	ExternalID   string   `json:"external_id"`
	CreatedAt    string   `json:"created_at"`
	Type         string   `json:"type"`
	Subject      string   `json:"subject"`
	Description  string   `json:"description"`
	Priority     string   `json:"priority"`
	Status       string   `json:"status"`
	SubmitterId  uint64   `json:"submitter_id"`
	AssigneeId   uint64   `json:"assignee_id"`
	OrgID        uint64   `json:"organization_id"`
	Tags         []string `json:"tags"`
	HasIncidents bool     `json:"has_incidents"`
	DueAt        string   `json:"due_at"`
	Via          string   `json:"via"`
}

const (
	TicketFieldID           = "_id"
	TicketFieldURL          = "url"
	TicketFieldExternalID   = "external_id"
	TicketFieldCreatedAt    = "created_at"
	TicketFieldType         = "type"
	TicketFieldSubject      = "subject"
	TicketFieldDescription  = "description"
	TicketFieldPriority     = "priority"
	TicketFieldStatus       = "status"
	TicketFieldRecipient    = "recipient"
	TicketFieldSubmitterId  = "submitter_id"
	TicketFieldAssigneeId   = "assignee_id"
	TicketFieldOrgID        = "organization_id"
	TicketFieldTags         = "tags"
	TicketFieldHasIncidents = "has_incidents"
	TicketFieldDueAt        = "due_at"
	TicketFieldVia          = "via"
	TicketFieldRequesterID    = "requester_id"
)

var SearchableTicketFieldList = []string{
	TicketFieldID,
	TicketFieldURL,
	TicketFieldExternalID,
	TicketFieldCreatedAt,
	TicketFieldType,
	TicketFieldSubject,
	TicketFieldDescription,
	TicketFieldPriority,
	TicketFieldStatus,
	TicketFieldRecipient,
	TicketFieldSubmitterId,
	TicketFieldAssigneeId,
	TicketFieldOrgID,
	TicketFieldTags,
	TicketFieldHasIncidents,
	TicketFieldDueAt,
	TicketFieldVia,
	TicketFieldRequesterID,
}
