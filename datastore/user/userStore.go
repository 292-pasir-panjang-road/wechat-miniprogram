package user

import (
  mDB             "mediocris/utils/database"
  mStore          "mediocris/datastore"
  converter       "mediocris/datastore/helper"
  userStoreModels "mediocris/datastore/user/storeModels"
  userDBModels    "mediocris/utils/database/dbModels"
  storeErr        "mediocris/datastore/error"
)

// Defines related SQL queries
// Currently the table is quite small and simple, therefore no need for update
// Allow user to permenently remove him/herself
const (
  sqlInsertUser      = `INSERT INTO MUser VALUES ($1)`
  sqlSelectUser      = `SELECT * FROM User`
  sqlSelectUserWhere = ` WHERE w_id = $1`
  sqlDeleteUserWhere = `DELETE FROM MUser WHERE w_id = $1`
)

// Defines user store struct
type UserStore struct {
  DB mDB.Database
}

// Constructor of UserStore
func NewUserStore(db mDB.Database) mStore.Store {
  return UserStore{db}
}

// Creates a new user record, return errors if any
func (s UserStore) Create(args interface{}) (interface{}, error) {

  // try to user converter to convert service model to store model
  converted, err := converter.serviceToStore(args)
  if err != nil {
    return nil, err
  }

  // Then cast it to correct data type
  casted, ok := converted.(userStoreModels.CreateParams)
  if !ok {
    return nil, storeErr.ErrConverterError
  }

  // execute insertion
  _, err = s.DB.Exec(sqlInsertUser, casted.wechatID)
  if err != nil {
    return nil, err
  }

  return casted.wechatID, nil
}

// Retrieves users. Since currently it does not make sense getting one user,
// we just assume retrieving all users first
func (s UserStore) Retrieve(_ interface{}) (interface{}, error) {
  var dbUsers []*userDBModels.UserAccount
  var storeUsers []*userStoreModels.UserRecord
  err := s.DB.SelectMany(&dbUsers, sqlSelectUser)
  if err != nil {
    return nil, err
  }

  for _, dbUser := range dbUsers {
    storeUsers = append(storeUsers, converter.dbToStore(dbUser))
  }
  return storeUsers, nil
}

// Updates a user records. Currently no need for update
func (s UserStore) Update(_ string, _ interface{}) (interface{}, error) {
  return nil, nil
}

// Deletes a user record.
func (s UserStore) Delete(id string, _ interface{}) (interface{}, error) {
  _, err := s.DB.Exec(sqlDeleteUserWhere, id)
  if err != nil {
    return nil, err
  }
  return nil, nil
}
