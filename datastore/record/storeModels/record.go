package storeModels

type RecordCreateParams {
  GroupID     int
  PayerID     string
  Spliters    []string
  Amount      int
  Description string
}

type RecordUpdateParams {
  PayerID     string
  Spliters    []string
  Amount      int
  description string
}

type Record
