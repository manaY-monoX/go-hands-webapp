package tests

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/products"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("値オブジェクト ProductPrice\n", func() {
	Describe("NewProductPrice()関数の検証\n", func() {
		When("値が有効範囲内の場合\n", func() {
			It("値50に対してエラーなしでProductPriceが返される\n", func() {
				productPrice, err := products.NewProductPrice(50)
				Expect(err).NotTo(HaveOccurred())
				Expect(productPrice).NotTo(BeNil())
				Expect(productPrice.Value()).To(Equal(50))
			})
			It("値9999に対してエラーなしでProductPriceが返される\n", func() {
				productPrice, err := products.NewProductPrice(9999)
				Expect(err).NotTo(HaveOccurred())
				Expect(productPrice).NotTo(BeNil())
				Expect(productPrice.Value()).To(Equal(9999))
			})
		})
		When("値が50未満の場合\n", func() {
			It("エラーが返される\n", func() {
				_, err := products.NewProductPrice(49)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errortype.NewDomainError("商品単価は、50以上10000未満でなければなりません。")))
			})
		})
		When("値が10000より大きいの場合\n", func() {
			It("エラーが返される\n", func() {
				_, err := products.NewProductPrice(10000)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errortype.NewDomainError("商品単価は、50以上10000未満でなければなりません。")))
			})
		})
	})
	Describe("Equals()メソッドの検証 \n", func() {
		When("同じ値を持つ2つのProductPriceインスタンスを⽐較する場合 \n", func() {
			It("trueが返される \n", func() {
				value := 200
				productPrice1, err1 := products.NewProductPrice(value)
				productPrice2, err2 := products.NewProductPrice(value)
				// エラーになっていないことを検証する
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
				// productPrice1とproductPrice2の比較結果がtrueであることを検証する
				Expect(productPrice1.Equals(productPrice2)).To(BeTrue())
			})
		})
		When("値が異なる2つのProductPriceインスタンスを比較する場合\n", func() {
			It("falseが返される\n", func() {
				value1 := 150
				value2 := 200
				productPrice1, err1 := products.NewProductPrice(value1)
				productPrice2, err2 := products.NewProductPrice(value2)
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
				Expect(productPrice1.Equals(productPrice2)).To(BeFalse())
			})
		})
		When("ProductPriceインスタンスと異なる型のインスタンスを⽐較する場合 \n", func() {
			It("falseが返される\n", func() {
				value := 100
				productPrice, err := products.NewProductPrice(value)
				Expect(err).NotTo(HaveOccurred())
				// 仮に異なる型のオブジェクトを作成
				nonProductPrice, err := products.NewProductName("ボールペン")
				Expect(err).NotTo(HaveOccurred())
				// ProductNameと明らかに異なる型との⽐較
				Expect(productPrice.Equals(nonProductPrice)).To(BeFalse())
			})
		})
	})
})
