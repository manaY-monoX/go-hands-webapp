package repimpl

import (
	"fmt"
	"service-exercise/domain/model/categories"
	"service-exercise/domain/repository"

	"gorm.io/gorm"
)

type categoryRepositoryImpl2 struct{}

// すべての商品カテゴリを取得するレシーバメソッド
// FindAll()は、gorm.DB型のポインタを引数として、商品カテゴリのスライス(リスト)を返す
func (c *categoryRepositoryImpl2) FindAll(db *gorm.DB) ([]*categories.Category, error) {
	fmt.Println("categoryRepositoryImpl2のFindAll()が実行されました。")
	return nil, nil
}

// コンストラクタ
// コンストラクタは、型パラメータを持つ関数である
func NewcategoryRepositoryImpl2() repository.CategtoryRepository[*gorm.DB] {
	fmt.Println("categoryRepositoryImpl2のコンストラクタが実行されました。")
	fmt.Println("categoryRepositoryImpl2の型パラメータは、*gorm.DB型です。")
	fmt.Println("型パラメータとは、型パラメータは、関数の引数として型を指定することができる。")
	return &categoryRepositoryImpl2{}
}
