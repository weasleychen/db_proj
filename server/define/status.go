package define

const (
	OK = iota
	ErrorDuplicateUserName
	ErrorWrongPassword
	ErrorTokenInvalid
	ErrorTableIsOpened
	ErrorTableIsClosed
	ErrorTableIdExist
	ErrorTableIdNotExist
	TableIsNotInUse
	TableIsInUse
)
