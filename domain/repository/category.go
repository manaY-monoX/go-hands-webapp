package repository

import "service-exercise/domain/model/categories"

// 商品カテゴリをCRUD操作するリポジトリインターフェイス
// DBアクセス機能を持つ構造体を型引数に指定する。
type CategtoryRepository[T any] interface {
	// すべての商品カテゴリを取得する
	FindAll(db T) ([]*categories.Category, error)
}

// T: ジェネリクス型パラメータ（型引数）
// 構造体を作成するときに、具体的な型を指定する。
