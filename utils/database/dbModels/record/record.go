package record

import "time"

type TansRecord struct {
	RecordID    int       `db:"id"`
	GroupID     int       `db:"g_id"`
	Date        time.Time `db:"day"`
	Payer       string    `db:"payer"`
	Spliters    []string  `db:"spliters"`
	Amount      float32   `db:"pay_amount"`
	Description string    `db:"description"`
}
