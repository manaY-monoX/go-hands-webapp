package errortype

// ドメインエラー型
type DomainError struct {
	message string
}

// errorインターフェイス（Goの標準のインターフェースの一つ）の実装
func (e *DomainError) Error() string {
	return e.message
}

// コンストラクタ
func NewDomainError(message string) *DomainError {
	return &DomainError{message: message}
}
