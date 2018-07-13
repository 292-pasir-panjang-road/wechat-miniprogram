package servicemodels

import "time"

// DetailRetrieveParams wraps the params used to retrieve detailed info
type DetailRetrieveParams struct {
	HostID  string
	GuestID string
	GroupID int
}

// DetailBetweenUsers wraps the detail info result between users
type DetailBetweenUsers struct {
	HostID  string
	GuestID string
	Records []*userRecord
}

// userRecord wraps the record used for detail info
// between users
type userRecord struct {
	RecordID    int
	GroupID     int
	Date        time.Time
	Amount      float32
	Description string
}

// GroupDetails wraps the detail info result for a group
type GroupDetails struct {
	HostID    string
	GroupID   int
	GroupName string
	Members   map[string]string
	Records   []*groupRecord
}

// groupRecord wraps the record used for detail info
// between user and group
type groupRecord struct {
	RecordID    int
	Date        time.Time
	Amount      float32
	Payer       string
	Spliters    map[string]string
	Description string
}
