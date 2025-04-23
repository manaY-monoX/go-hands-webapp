package repository

import "service-exercise/domain/model/products"

// 商品のCRUD操作リポジトリインターフェイス
type ProductRepository[T any] interface {
	// 指定された商品の存在確認
	Exists(db T, name *products.ProductName) (bool, error)
	// 新しい商品を永続化(登録)する
	Create(db T, product *products.Product) error
	// 商品をキーワード検索する（LIKE: 部分一致）
	FindByNameLike(db T, keyword string) ([]*products.Product, error)
}

// ジェネリクス型パラメータ（型引数）
// 構造体を作成するときに、具体的な型を指定する。
// 型引数は、型パラメータと呼ばれる。
