package logger

var activeMock bool
type Mock struct {
	Err []error
	Msg []string
}
var mock Mock

func StartMock() {
	activeMock = true
}

func StopMock() {
	activeMock = false
}

func GetMock() *Mock{
	return &mock
}

func ClearMock() {
	mock.Err = nil
	mock.Msg = nil
}

