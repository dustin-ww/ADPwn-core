package adpwn

type Parameter struct {
	RunID    string // RunID hinzufügen
	Inputs   map[string]interface{}
	Metadata map[string]string
}
