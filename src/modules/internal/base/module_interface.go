package base

import "ADPwn/core/model"

type ADPwnModule interface {
	GetName() string
	GetDescription() string
	GetVersion() string
	GetAuthor() string
	Execute(project model.Project, options []string) error

	DependsOn() int
}
