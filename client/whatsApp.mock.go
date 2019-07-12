package client

type WppConnectionMock struct {}

func InitMock() *WppConnectionMock {
	return &WppConnectionMock{}
}

func (wpp *WppConnectionMock) Send(phoneNumber, text string) (string, error) {
	return "mock", nil
}
