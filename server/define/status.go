package define

const (
	OK = iota
	ErrorDuplicateUserName
	ErrorWrongPassword
	ErrorTokenInvalid
	ErrorTableIsOpened
	ErrorTableIdInvalid
	ErrorTableIsClosed
	TableIsClosed
	TableIsOpened
)
