package model

type User struct {
	UID      string   `json:"uid,omitempty"`
	Name     string   `json:"name"`
	Password string   `json:"password,omitempty"`
	NTLNHash string   `json:"ntlmHash,omitempty"`
	IsAdmin  bool     `json:"isAdmin"`
	DType    []string `json:"dgraph.type,omitempty"`
}

func NewUser(username string, password string, ntlmHash string, isAdmin bool) *User {
	return &User{
		Name:     username,
		NTLNHash: ntlmHash,
		Password: password,
		IsAdmin:  isAdmin,
		DType:    []string{"User"},
	}
}
