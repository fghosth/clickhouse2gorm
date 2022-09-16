package climodel

import "time"

type Issues struct {
	Name       string    `db:"name"`
	Owner      string    `db:"owner"`
	Category   string    `db:"category"`
	Epics      string    `db:"epics"`
	Project    string    `db:"project"`
	DueDate    time.Time `db:"due_date"`
	StartDate  time.Time `db:"start_date"`
	Status     string    `db:"status"`
	EndDate    time.Time `db:"end_date"`
	Progress   int16     `db:"progress"`
	Key        string    `db:"key"`
	CreateDate time.Time `db:"create_date"`
}
