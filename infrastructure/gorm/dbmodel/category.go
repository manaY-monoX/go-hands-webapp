package dbmodel

// 商品カテゴリ用GORMモデル
type CategoryModel struct {
	ID uint `gorm:"primaryKey"`
	// カテゴリId
	ObjID string `gorm:"uniqueIndex;type:varchar(36)"`
	// カテゴリ名
	Name string `gorm:"type:varchar(20)"`
}

// GORMにこのモデルが使用するテーブル名を教える
func (CategoryModel) TableName() string {
	return "category"
}
