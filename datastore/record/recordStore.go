package record

import (
  mDB               "mediocris/utils/database"
  mStore            "mediocris/datastore"
  helper            "mediocris/datastore/helper"
  recordStoreModels "mediocris/datastore/record/storeModels"
  recordDBModels    "mediocris/utils/database/dbModels"
  storeErr          "mediocris/datastore/error"
)

const (
  NON_GROUP_ID            = -1
  NON_USER_ID             = ""

  SQL_SELECT_RECORD       = `SELECT * FROM record`
  SQL_WHERE_USER_RELATED  = ` WHERE payer = $1 OR $1=ANY(spliters)`
  SQL_WHERE_GROUP_RELATED = ` WHERE payer = $1 OR $1=ANY(spliters) AND g_id = $2`
  SQL_WHERE_BETWEEN_USERS = ` WHERE (payer = $1 AND $2=ANY(spliters)) OR (payer = $2 AND $1=ANY(spliters))`

  SQL_CREATE_GROUP_RECORD = `INSERT INTO record
                            (g_id, day, payer, spliters, pay_amount, description, updated_at, deleted_at)
                            VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
  SQL_CREATE_NO_GROUP     = `INSERT INTO record
                            (day, payer, spliters, pay_amount, description, updated_at, deleted_at)
                            VALUES($1, $2, $3, $4, $5, $6, $7)`
)

type RecordStore struct {
  DB mDB.Database
}

func NewRecordStore(db mDB.Database) mStore.Store {
  return RecordStore{db}
}

// Creates a record. all arguments are included in args
func (s RecordStore) Create(args interface{}) (interface{}, error) {
  return nil, nil
}

// Retrieves record(s).
func (s RecordStore) Retrieve(args interface{}) (interface{}, error) {
  converted, err := helper.ServiceToStore(args)
  if err != nil {
    return nil, err
  }
  retrieveParams := converted.(storeModels.RecordRetrieveParams)

  if retrieveParams.HostID == NON_USER_ID {
    return nil, storeErr.ErrInvalidQuery
  }

  // When no group and guest id is provided (records for a user)
  if retrieveParams.GroupID == NON_GROUP_ID && retrieveParams.GuestID == NON_USER_ID {
    return retrieveIndividualRecords(retrieveParams.HostID)
  }

  // When group specified but no guest id (records for a user in a specific group)
  if retrieveParams.GroupID != NON_GROUP_ID && retrieveParams.GuestID == NON_USER_ID {
    return retrieveInGroupRecords(retrieveParams.HostID, retrieveParams.GroupID)
  }

  // When two user ids are specified (records between two specific users)
  if retrieveParams.GuestID != NON_USER_ID {
    return retrieveBetweenUserRecords(retrieveParams.HostID, retrieveParams.GuestID)
  }

  return nil, storeErr.ErrNotSupportedQuery
}

// Updates a record. id should be specified
func (s RecordStore) Update(id string, args interface{}) (interface{}, error) {
  return nil, nil
}

// Deletes a record. id should be specified
func (s RecordStore) Delete(id string, args interface{}) (interface{}, error) {
  return nil, nil
}

// Retrieves records that are related to a specific individual
// Returns an array of references to the records
func retrieveIndividualRecords(db mDB.Database, userID string) ([]*storeModels.RecordRetrieveResult, error) {
  var dbResult []recordDBModels.TansRecord
  err := db.SelectMany(&dbResult, SQL_SELECT_RECORD+SQL_WHERE_USER_RELATED, userID)
  if err != nil {
    return nil, storeErr.ErrDBRetrievalError
  }
  return helper.DBRecordsToStore(dbResult), nil
}

// Retrieves records that are related to a specific individual in a specific group
// Returns an array of references to the records
func retrieveInGroupRecords(db mDB.Database, userID string, groupID int) ([]*storeModels.RecordRetrieveResult, error) {
  var dbResult []recordDBModels.TansRecord
  err := db.SelectMany(&dbResult, SQL_SELECT_RECORD+SQL_WHERE_GROUP_RELATED, userID, groupID)
  if err != nil {
    return nil, storeErr.ErrDBRetrievalError
  }
  return helper.DBRecordsToStore(dbResult), nil
}

// Retrieves records that are between two specific users
// Returns an array of references to the records
func retrieveBetweenUserRecords(db mDB.Database, hostID string, guestID int) ([]*storeModels.RecordRetrieveResult, error) {
  var dbResult []recordDBModels.TansRecord
  err := db.SelectMany(&dbResult, SQL_SELECT_RECORD+SQL_WHERE_BETWEEN_USERS, hostID, guestID)
  if err != nil {
    return nil, storeErr.ErrDBRetrievalError
  }
  return helper.DBRecordsToStore(dbResult), nil
}
