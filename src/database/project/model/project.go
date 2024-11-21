package model

type Project struct {
	UUID string `db:"UUID"`
	Name string `db:"name"`
}
