package adapter

import "service-exercise/domain/model/categories"

// 商品カテゴリエンティティと他のモデルの相互変換インターフェイス
type CategoryAdapter[T any] interface {
	// エンティティを他のモデルに変換する
	Convert(source *categories.Category) (T, error)
	// 他のモデルからエンティティを復元する
	Restore(source T) (*categories.Category, error)
}

// ジェネリクス型パラメータ（型引数）
// 構造体を作成するときに、具体的な型を指定する。
// 型引数は、型パラメータと呼ばれる。
