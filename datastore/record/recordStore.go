package record

import (
  mDB               "mediocris/utils/database"
  mStore            "mediocris/datastore"
  converter         "mediocris/datastore/helper"
  recordStoreModels "mediocris/datastore/record/storeModels"
  recordDBModels    "mediocris/utils/database/dbModels"
  storeErr          "mediocris/datastore/error"
)

const (
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
  SQL_DELETE_RECORD       = `UPDATE `
)

type RecordStore struct {
  DB mDB.Database
}

func NewRecordStore(db mDB.Database) mStore.Store {
  return RecordStore{db}
}

// Creates a record. all arguments are included in args
func (s RecordStore) Create(args interface{}) (interface{}, error) {

}

// Retrieves record(s).
func (s RecordStore) Retrieve(args interface{}) (interface{}, error) {
  converted, err := helper.serviceToStore(args)
  if err != nil {
    return nil, err
  }
}

// Updates a record. id should be specified
func (s RecordStore) Update(id string, args interface{}) (interface{}, error) {

}

// Deletes a record. id should be specified
func (s RecordStore) Delete(id string, args interface{}) (interface{}, error) {

}
