package serviceModels

import "time"

type DetailRetrieveParams struct {
	HostID  string
	GuestID string
	GroupID int
}

type DetailBetweenUsers struct {
	HostID  string
	GuestID string
	Records []*UserRecord
}

type UserRecord struct {
	RecordID    int
	GroupID     int
	Date        time.Time
	Amount      float32
	Description string
}

type GroupDetails struct {
	HostID    string
	GroupID   int
	GroupName string
	Members   map[string]string
	Records   []*GroupRecord
}

type GroupRecord struct {
	RecordID    int
	Date        time.Time
	Amount      float32
	Payer       string
	Spliters    map[string]string
	Description string
}
