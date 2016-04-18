package pm

type Project struct {
	Id int
	Title string
	Description string
	Trackers []Tracker
	Issues []Issue
}

func NewProject(title string) *Project {
	return &Project{Id: 1, Title: title}
}
