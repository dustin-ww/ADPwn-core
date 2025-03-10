package interfaces

type ADPwnModule interface {
	GetName() string
	GetDescription() string
	GetVersion() string
	GetAuthor() string
	GetExecutionMetric() string
	// ... other methods...
}
