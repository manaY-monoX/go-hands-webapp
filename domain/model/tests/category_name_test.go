package tests

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/categories"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("値オブジェクト CategoryName \n", func() {
	Describe("NewCategoryName()関数を検証する \n", func() {
		When("空文字ではなく、20文字以内のカテゴリ名の場合\n", func() {
			It("エラーなしで有効なCategoryNameインスタンスが返される", func() {
				value := "電化製品"
				categoryName, err := categories.NewCategoryName(value)
				// エラーが発⽣していないことを検証する
				Expect(err).NotTo(HaveOccurred())
				// nilでないことを検証する
				Expect(categoryName).NotTo(BeNil())
				// 変数valueと同じ値であることを検証する
				Expect(categoryName.Value()).To(Equal(value))
			})
		})
		When("空文字の場合\n", func() {
			It("エラーが返される\n", func() {
				_, err := categories.NewCategoryName("")
				// エラーであることを検証する
				Expect(err).To(HaveOccurred())
				// エラーメッセージを検証する
				Expect(err).To(MatchError(errortype.NewDomainError("商品カテゴリ名は、空文字列であってはなりません。")))
			})
		})
		When("値が20文字を超えていた場合\n", func() {
			It("エラーが返される\n", func() {
				value := "非常に長い商品カテゴリ名をここに入力します。"
				_, err := categories.NewCategoryName(value)
				// エラーであることを検証する
				Expect(err).To(HaveOccurred())
				// エラーメッセージを検証する
				Expect(err).To(MatchError(errortype.NewDomainError("商品カテゴリ名は、20文字以内である必要があります。")))
			})
		})
	})

	Describe("Equals()メソッドの検証 \n", func() {
		When("同じ値を持つ2つのCategoryNameインスタンスを⽐較する場合 \n", func() {
			It("trueが返される \n", func() {
				value := "電化製品"
				categoryName1, err1 := categories.NewCategoryName(value)
				categoryName2, err2 := categories.NewCategoryName(value)
				// エラーになっていないことを検証する
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
				// categoryName1とcategoryName2の比較結果がtrueであることを検証する
				Expect(categoryName1.Equals(categoryName2)).To(BeTrue())
			})
		})
		When("値が異なる2つのCategoryNameインスタンスを比較する場合 \n", func() {
			It("falseが返される \n", func() {
				value1 := "家電製品"
				value2 := "飲料水"
				categoryName1, err1 := categories.NewCategoryName(value1)
				categoryName2, err2 := categories.NewCategoryName(value2)
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
				Expect(categoryName1.Equals(categoryName2)).To(BeFalse())
			})
		})
		When("CategoryNameインスタンスと異なる型のインスタンスを⽐較する場合 \n", func() {
			It("falseが返される\n", func() {
				value := "家電製品"
				categoryName, err := categories.NewCategoryName(value)
				Expect(err).NotTo(HaveOccurred())
				// 仮に異なる型のオブジェクトを作成
				nonCategoryName, err := categories.NewCategoryId("123e4567-e89b-12d3-a456-426614174000")
				Expect(err).NotTo(HaveOccurred())
				// CategoryNameと明らかに異なる型との⽐較
				Expect(categoryName.Equals(nonCategoryName)).To(BeFalse())
			})
		})
	})
})
