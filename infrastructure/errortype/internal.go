// internal.go：　内部エラーを定義

package errortype

type InternalError struct {
	message string
}

// インターフェース
func (e *InternalError) Error() string {
	return e.message
}

// コンストラクタ
func NewInternalError(message string) *InternalError {
	return &InternalError{message: message}
}
