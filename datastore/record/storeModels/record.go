package storeModels

import "time"

type RecordCreateParams struct {
  GroupID     int
  Date        time.Time
  PayerID     string
  Spliters    []string
  Amount      float32
  Description string
  UpdatedAt   time.Time
  DeletedAt   time.Time
}

type RecordUpdateParams struct {
  Date        time.Time
  PayerID     string
  Spliters    []string
  Amount      float32
  Description string
  UpdatedAt   time.Time
  DeletedAt   time.Time
}

type RecordRetrieveParams struct {
  HostID  string
  GuestID string
  GroupID int
}

type RecordRetrieveResult struct {
  RecordID    int
  GroupID     int
  Date        time.Time
  Payer       string
  Spliters    []string
  Amount      float32
  Description string
}
