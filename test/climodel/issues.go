package climodel

import "time"

type Issues struct {
	Name       string    `db:"name" ch:"name"`
	Owner      string    `db:"owner" ch:"owner"`
	Category   string    `db:"category" ch:"category"`
	Module     string    `db:"module" ch:"module"`
	Project    string    `db:"project" ch:"project"`
	DueDate    time.Time `db:"due_date" ch:"due_date"`
	StartDate  time.Time `db:"start_date" ch:"start_date"`
	Status     string    `db:"status" ch:"status"`
	EndDate    time.Time `db:"end_date" ch:"end_date"`
	Progress   int16     `db:"progress" ch:"progress"`
	Key        string    `db:"key" ch:"key"`
	CreateDate time.Time `db:"create_date" ch:"create_date"`
}
