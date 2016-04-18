package pm

import "time"

type Issue struct {
	Id int
	Tracker *Tracker
	Title string
	Status string
	CreatedOn time.Time
	Fields map[string]string
}
