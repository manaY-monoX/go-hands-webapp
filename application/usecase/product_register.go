package usecase

import "service-exercise/domain/model/products"

// ユースケース:[商品を登録する] インターフェイス
type ProductRegister interface {
	// 画面で入力された商品情報を受け取り、永続化する
	Execute(newProduct *products.Product) (err error)
}
