package apperror

import "fmt"

// error representing a conflict between entities.
type ConflictError struct {
	// the type of entity. e.g. User, File
	Entity string
	// name of the field colliding. e.g. ID, alias
	CollisionField string
	// the value colliding. e.g. 123456, gh7Qz5
	FieldValue string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("%s with %s \"%s\" already exists", e.Entity, e.CollisionField, e.FieldValue)
}

type DatabaseError struct {
	Err error
}

func (e *DatabaseError) Error() string {
	return e.Err.Error()
}
