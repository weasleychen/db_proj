package define

const (
	OK = iota
	ErrorWrongPassword
	ErrorTokenInvalid
	ErrorTableIsOpened
	ErrorTableIsClosed
	ErrorTableIdExist
	ErrorTableIdNotExist
	TableIsNotInUse
	TableIsInUse
	ErrorCreateUser // 占位 暂时不清楚db会返回什么错误
	ErrorNoSuchUin
	ErrorNoSuchPhoneNumber
	ErrorNoSuchEmail
	ErrorDishIdNotExist
	ErrorNoSuchVipLevel
)
