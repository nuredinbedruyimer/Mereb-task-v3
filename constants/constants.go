package constants

import "time"

var (
	TIME_OUT                   = time.Second * 25
	PERSON_DOES_NOT_EXIST      = "Person Does Not Exists !!!"
	PERSON_ALEADY_EXISTS       = "Person Already Exists !!!"
	INTERNAL_SERVER_ERROR      = "Internal Server Error !!!"
	DATABASE_CONNECTED_SUCCESS = "Database Connected Successfully !!!"
	PERSON_CREATED             = "Person Created Successfully !!!"
	PERSON_DELETED             = "Person Deleted Successfully !!!"
	PERSONS_FETCHED            = "Persons Fetched Successfully !!!"
	PERSON_FETCHED             = "Person Fetched Successfully !!!"
	PERSON_UPDATE              = "Person Update Successfully !!!"

	DATABASE_CONNECTED_FAILED = "error in database connection"
	ERROR_IN_DATA_BINDING     = "Error When We Bind Request Body To Go Struct"
	ERROR_STATUS              = "Failure"
	SUCCESS_STATUS            = "Success"
	VALIDATION_FAILED         = "Person Data Failed Validation"
)
