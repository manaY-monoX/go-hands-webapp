package adapter

import "service-exercise/domain/model/products"

// 商品エンティティと他のモデルを相互変換するインターフェイス
type ProductAdapter[T any] interface {
	// エンティティから他のモデルへ変換する
	Convert(source *products.Product) (T, error)
	// 他のモデルからエンティティを復元する
	Restore(source T) (*products.Product, error)
}

// ジェネリクス型パラメータ（型引数）
// 構造体を作成するときに、具体的な型を指定する。
// 型引数は、型パラメータと呼ばれる。
