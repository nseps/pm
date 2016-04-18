package pm

import "time"

type Comment struct {
	Author string
	Text string
	time.Time
}
