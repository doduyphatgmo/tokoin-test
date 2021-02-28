package models

type User struct {
	ID          uint64   `json:"_id"`
	URL         string   `json:"url"`
	ExternalID  string   `json:"external_id"`
	Name        string   `json:"name"`
	Alias       string   `json:"alias"`
	CreatedAt   string   `json:"created_at"`
	Active      bool     `json:"active"`
	Verified    bool     `json:"verified"`
	Shared      bool     `json:"shared"`
	Locale      string   `json:"locale"`
	Timezone    string   `json:"timezone"`
	LastLogicAt string   `json:"last_login_at"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	Signature   string   `json:"signature"`
	OrgID       uint64   `json:"organization_id"`
	Tags        []string `json:"tags"`
	Suspended   bool     `json:"suspended"`
	Role        string   `json:"role"`
}

const (
	UserFieldID          = "_id"
	UserFieldURL         = "url"
	UserFieldExternalID  = "external_id"
	UserFieldName        = "name"
	UserFieldAlias       = "alias"
	UserFieldCreatedAt   = "created_at"
	UserFieldActive      = "active"
	UserFieldVerified    = "verified"
	UserFieldShared      = "shared"
	UserFieldLocale      = "locale"
	UserFieldTimezone    = "timezone"
	UserFieldLastLogicAt = "last_login_at"
	UserFieldEmail       = "email"
	UserFieldPhone       = "phone"
	UserFieldSignature   = "signature"
	UserFieldOrgID       = "organization_id"
	UserFieldTags        = "tags"
	UserFieldSuspended   = "suspended"
	UserFieldRole        = "role"
)

var SearchableUserFieldsMap = map[string]bool{
	UserFieldID:          true,
	UserFieldURL:         true,
	UserFieldExternalID:  true,
	UserFieldName:        true,
	UserFieldAlias:       true,
	UserFieldCreatedAt:   true,
	UserFieldActive:      true,
	UserFieldVerified:    true,
	UserFieldShared:      true,
	UserFieldLocale:      true,
	UserFieldTimezone:    true,
	UserFieldLastLogicAt: true,
	UserFieldEmail:       true,
	UserFieldPhone:       true,
	UserFieldSignature:   true,
	UserFieldOrgID:       true,
	UserFieldTags:        true,
	UserFieldSuspended:   true,
	UserFieldRole:        true,
}
