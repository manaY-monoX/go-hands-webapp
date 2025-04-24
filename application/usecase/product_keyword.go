package usecase

import "service-exercise/domain/model/products"

// ユースケース:[商品を検索する] インターフェイス
type ProductKeyword interface {
	Execute(keyword string) ([]*products.Product, error)
}
