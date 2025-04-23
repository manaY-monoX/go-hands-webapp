package categories

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model"

	"github.com/google/uuid"
)

// 商品カテゴリIdを表す値オブジェクト
type CategoryId struct {
	value string // 小文字だからprivateなフィールド
}

// ゲッター（privateなフィールドの値を取得するメソッド）
func (c *CategoryId) Value() string {
	return c.value
}

// Equatableインターフェイスの実装
func (c *CategoryId) Equals(other model.Equatable) bool {
	// 型アサーション
	// otherがCategoryId型かどうかを確認
	otherCategoryId, ok := other.(*CategoryId)
	if !ok {
		return false
	}
	// 値の比較
	return c.value == otherCategoryId.value
}

// 新しいUUIDを生成し、その値を持つCategoryIdを返す
// 新しいCategoryId（商品カテゴリId）を生成するメソッド
func NewCategoryIdWithUUID() *CategoryId {
	categoryId := CategoryId{value: uuid.NewString()}
	return &categoryId
}

// コンストラクタ
// 既に存在しているUUIDを指定して、CategoryIdを生成するメソッド
func NewCategoryId(value string) (*CategoryId, error) {
	if value == "" {
		return nil, errortype.NewDomainError("商品カテゴリIdは、空文字列であってはなりません。")
	}
	if len(value) != 36 {
		return nil, errortype.NewDomainError("商品カテゴリIdは、36文字でなければなりません。")
	}
	if _, err := uuid.Parse(value); err != nil {
		return nil, errortype.NewDomainError("商品カテゴリIdは、UUID形式でなければなりません。")
	}
	categoryId := CategoryId{value: value}
	return &categoryId, nil
}
