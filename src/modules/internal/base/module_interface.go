package base

import "ADPwn/database/model"

type ADPwnModule interface {
	GetName() string
	GetDescription() string
	GetVersion() string
	GetAuthor() string
	Execute(project model.Project, options []string) error
}
