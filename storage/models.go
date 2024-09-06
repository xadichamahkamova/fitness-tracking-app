// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package storage

import (
	"database/sql"

	"github.com/sqlc-dev/pqtype"
)

type Exercise struct {
	ID        int32
	WorkoutID sql.NullInt32
	Name      sql.NullString
}

type Image struct {
	ID     int32
	UserID sql.NullInt32
	Url    sql.NullString
}

type PasswordReset struct {
	ID        int32
	UserEmail sql.NullString
	UserToken sql.NullString
}

type Set struct {
	ID          int32
	ExerciseID  sql.NullInt32
	Repetitions sql.NullInt32
	Weight      sql.NullFloat64
}

type User struct {
	ID           int32
	Username     sql.NullString
	Email        sql.NullString
	PasswordHash sql.NullString
	Profile      pqtype.NullRawMessage
}

type Workout struct {
	ID     int32
	UserID sql.NullInt32
	Date   sql.NullString
}
