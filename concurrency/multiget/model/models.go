package model

// UserAssignedPermission defines the permission assigned to an user
type UserAssignedPermission struct {
	ID             int           `json:"id"`
	Key            string        `json:"key"`
	ApplicationKey string        `json:"application_key"`
	ApproverRoles  []interface{} `json:"approver_roles"`
	ApproverUser   interface{}   `json:"approver_user"`
}

// Attribute describe an groot-attribute entity
type Attribute struct {
	ID    int           `json:"id"`
	Value []interface{} `json:"value"`
	Key   string        `json:"key"`
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DateCreated string `json:"date_created"`
	Active      bool   `json:"active"`
	Key         string `json:"key"`
}

// User represents an groot-user entity
type User struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	LastName    string      `json:"last_name"`
	LdapUser    string      `json:"ldap_user"`
	MeliUser    string      `json:"meli_user"`
	ClientID    string      `json:"client_id"`
	DateCreated string      `json:"date_created"`
	Active      bool        `json:"active"`
	Roles       []Role      `json:"roles"`
	Attributes  []Attribute `json:"attributes"`
}
