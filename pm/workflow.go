package pm

type State struct {
	Id int
	Name string

}

type Workflow struct {
	States []string
	TransitionMatrix map[string][]string
}