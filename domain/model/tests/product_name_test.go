package tests

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/products"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("値オブジェクト ProductName\n", func() {
	Describe("NewProductName()関数を検証する\n", func() {
		When("空文字ではなく、30文字以内のカテゴリ名の場合\n", func() {
			It("エラーなしで有効なProductNameインスタンスが返される", func() {
				value := "ボールペン"
				productName, err := products.NewProductName(value)
				// エラーが発⽣していないことを検証する
				Expect(err).NotTo(HaveOccurred())
				// nilでないことを検証する
				Expect(productName).NotTo(BeNil())
				// 変数valueと同じ値であることを検証する
				Expect(productName.Value()).To(Equal(value))
			})
		})
		When("空文字の場合\n", func() {
			It("エラーが返される\n", func() {
				_, err := products.NewProductName("")
				// エラーであることを検証する
				Expect(err).To(HaveOccurred())
				// エラーメッセージを検証する
				Expect(err).To(MatchError(errortype.NewDomainError("商品名は、空文字列であってはなりません。")))
			})
		})
		When("値が30文字を超えていた場合\n", func() {
			It("エラーが返される\n", func() {
				value := "非常に長い商品名をここに入力します!!!!!!!!!!!!!!!!!!!。"
				_, err := products.NewProductName(value)
				// エラーであることを検証する
				Expect(err).To(HaveOccurred())
				// エラーメッセージを検証する
				Expect(err).To(MatchError(errortype.NewDomainError("商品名は、30文字以内である必要があります。")))
			})
		})
	})
	Describe("Equals()メソッドの検証 \n", func() {
		When("同じ値を持つ2つのProductNameインスタンスを⽐較する場合 \n", func() {
			It("trueが返される \n", func() {
				value := "ボールペン"
				productName1, err1 := products.NewProductName(value)
				productName2, err2 := products.NewProductName(value)
				// エラーになっていないことを検証する
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
				// productName1とproductName2の比較結果がtrueであることを検証する
				Expect(productName1.Equals(productName2)).To(BeTrue())
			})
		})
		When("値が異なる2つのProductNameインスタンスを比較する場合 \n", func() {
			It("falseが返される \n", func() {
				value1 := "ボールペン"
				value2 := "蛍光ペン"
				productName1, err1 := products.NewProductName(value1)
				productName2, err2 := products.NewProductName(value2)
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
				Expect(productName1.Equals(productName2)).To(BeFalse())
			})
		})
		When("ProductNameインスタンスと異なる型のインスタンスを⽐較する場合 \n", func() {
			It("falseが返される\n", func() {
				value := "ボールペン"
				productName, err := products.NewProductName(value)
				Expect(err).NotTo(HaveOccurred())
				// 仮に異なる型のオブジェクトを作成
				nonProductName, err := products.NewProductId("123e4567-e89b-12d3-a456-426614174000")
				Expect(err).NotTo(HaveOccurred())
				// ProductNameと明らかに異なる型との⽐較
				Expect(productName.Equals(nonProductName)).To(BeFalse())
			})
		})
	})
})
