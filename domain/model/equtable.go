package model

// 等価性を検証インターフェイス
type Equatable interface {
	// 引数: other Equatable:Equatableインターフェースを実装した構造体
	// 戻り値: true:等価, false:非等価
	Equals(other Equatable) bool
}
