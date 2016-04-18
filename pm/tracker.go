package pm

type Tracker struct {
	Id int
	Name string
	Workflow Workflow
	Fields map[string]string
}
