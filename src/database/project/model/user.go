package model

type User struct {
	UID      string   `json:"uid,omitempty"`
	Username string   `json:"username"`
	Password string   `json:"password,omitempty"`
	NTLNHash string   `json:"ntlmHash,omitempty"`
	DType    []string `json:"dgraph.type,omitempty"`
}

func NewPasswordUser(username string, password string) *User {
	return &User{
		Username: username,
		Password: password,
		DType:    []string{"User"},
	}
}

func NewHashUser(username string, ntlmHash string) *User {
	return &User{
		Username: username,
		NTLNHash: ntlmHash,
		DType:    []string{"User"},
	}
}
