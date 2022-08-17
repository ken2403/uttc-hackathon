package common

// PreConditionError - 事前条件違反を表現する(400 BadRequestを判断するために使う)
// 仕様書レベルの違反は事前条件違反として扱う。
// それ以外のネットワークエラーなど、利用者やプログラマが気をつけていても発生するものは事後条件違反(任意のerror)として扱う。
type PreConditionError struct {
	Message string
}

func NewPreConditionError(message string) *PreConditionError {
	return &PreConditionError{Message: message}
}

// errorインタフェースを満たすために定義
func (receiver *PreConditionError) Error() string {
	return receiver.Message
}
