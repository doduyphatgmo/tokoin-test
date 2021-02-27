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