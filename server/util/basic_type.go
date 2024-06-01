package util

func NewInt32(value int32) *int32 {
	p := new(int32)
	*p = value
	return p
}

func NewBool(value bool) *bool {
	p := new(bool)
	*p = value
	return p
}

func NewString(value string) *string {
	p := new(string)
	*p = value
	return p
}
