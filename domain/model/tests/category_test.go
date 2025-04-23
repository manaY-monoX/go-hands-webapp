package tests

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/categories"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("エンティティ Category\n", func() {
	Describe("NewCategory()関数を検証する\n", func() {
		// 値オブジェクトの変数を準備する
		var (
			validCategoryId   *categories.CategoryId
			validCategoryName *categories.CategoryName
		)
		BeforeEach(func() {
			// 正しいCategoryIdとCategoryNameインスタンスをセットアップする
			validCategoryId = categories.NewCategoryIdWithUUID()
			validCategoryName, _ = categories.NewCategoryName("有効なカテゴリ名")
		})
		When("有効なIDと名前を指定した場合\n", func() {
			It("エラーなくCategoryのインスタンスが生成される\n", func() {
				category, err := categories.NewCategory(validCategoryId, validCategoryName)
				Expect(err).NotTo(HaveOccurred())
				Expect(category).NotTo(BeNil())
				Expect(category.Id()).To(Equal(validCategoryId))
				Expect(category.Name()).To(Equal(validCategoryName))
			})
		})
		When("idにnilが渡された場合\n", func() {
			It("エラーが返される\n", func() {
				category, err := categories.NewCategory(nil, validCategoryName)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errortype.NewDomainError("カテゴリIdは、必須です。")))
				Expect(category).To(BeNil())
			})
		})
		When("nameにnilを指定した場合\n", func() {
			It("エラーなくCategoryのインスタンスが生成される\n", func() {
				category, err := categories.NewCategory(validCategoryId, nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(category).NotTo(BeNil())
				Expect(category.Id()).To(Equal(validCategoryId))
				Expect(category.Name()).To(BeNil())
			})
		})
	})

	Describe("Equals()メソッドを検証する\n", func() {
		var category1, category2 *categories.Category

		BeforeEach(func() {
			// Categoryインスタンスを初期化する
			categoryId1, _ := categories.NewCategoryId("123e4567-e89b-12d3-a456-426614174000")
			category1, _ = categories.NewCategory(categoryId1, nil)

			categoryId2, _ := categories.NewCategoryId("123e4567-e89b-12d3-a456-426614174010")
			category2, _ = categories.NewCategory(categoryId2, nil)
		})
		When("2つのカテゴリが同じIDを持つ場合\n", func() {
			It("trueが返される\n", func() {
				categoryId, _ := categories.NewCategory(category1.Id(), category1.Name())
				Expect(category1.Equals(categoryId)).To(BeTrue())
			})
		})
		When("2つのカテゴリのIDが異なる場合\n", func() {
			It("falseが返される", func() {
				Expect(category1.Equals(category2)).To(BeFalse())
			})
		})
		When("nilを渡した場合\n", func() {
			It("falseが返される", func() {
				Expect(category1.Equals(nil)).To(BeFalse())
			})
		})
		When("異なる型のインスタンスと比較した場合\n", func() {
			It("falseが返される", func() {
				name, _ := categories.NewCategoryName("家電製品")
				Expect(category1.Equals(name)).To(BeFalse())
			})
		})
	})

	Describe("ChangeName()メソッドの検証\n", func() {
		var category *categories.Category

		BeforeEach(func() {
			//有効なCategoryIdとCategoryNameでCategoryインスタンスを初期化する
			categoryId, _ := categories.NewCategoryId("123e4567-e89b-12d3-a456-426614174000")
			categoryName, _ := categories.NewCategoryName("食料品")
			category, _ = categories.NewCategory(categoryId, categoryName)
		})
		When("有効なCategoryNameが指定された場合\n", func() {
			It("カテゴリ名前が正常に変更される\n", func() {
				newCategoryName, _ := categories.NewCategoryName("Food")
				err := category.ChangeName(newCategoryName)
				Expect(err).NotTo(HaveOccurred())
				Expect(category.Name()).To(Equal(newCategoryName))
			})
		})
		When("nilが指定された場合\n", func() {
			It("エラーが返される", func() {
				err := category.ChangeName(nil)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errortype.NewDomainError("カテゴリ名は、必須です。")))
			})
		})
	})
})
