package storeModels

import "time"

type RecordCreateParams struct {
  GroupID     int
  Date        time.Time
  PayerID     string
  Spliters    []string
  Amount      int
  Description string
  UpdatedAt   time.Time
  DeletedAt   time.Time
}

type RecordUpdateParams struct {
  Date        time.Time
  PayerID     string
  Spliters    []string
  Amount      int
  Description string
  UpdatedAt   time.Time
  DeletedAt   time.Time
}

type RecordRetrieveParams struct {
  Type    string
  HostID  string
  GuestID string
  GroupID int
}
