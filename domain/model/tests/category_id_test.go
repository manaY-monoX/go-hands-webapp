package tests

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/categories"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("値オブジェクト: CategoryId", func() {
	Describe("NewCategoryIdWithUUID()関数を検証する \n", func() {
		It("有効なUUIDのCategoryIdが作成される \n", func() {
			categoryId := categories.NewCategoryIdWithUUID()
			// CategoryIdがnilでないことを検証する
			Expect(categoryId).NotTo(BeNil())
			// CategoryIdのvalueがUUID形式であることを検証する
			_, err := uuid.Parse(categoryId.Value())
			Expect(err).NotTo(HaveOccurred())
			// DepartmentIdのvalueが36⽂字であることを検証する
			Expect(len(categoryId.Value())).To(Equal(36))
		})
	})

	Describe("NewCategoryId()関数を検証する \n", func() {
		When("有効なUUIDを使⽤した場合 \n", func() {
			It("エラーなくCategoryIdが返される \n", func() {
				uuid := "123e4567-e89b-12d3-a456-426614174000" // 仮の有効なUUID
				categoryId, err := categories.NewCategoryId(uuid)
				// エラーが発⽣していないことを検証する
				Expect(err).NotTo(HaveOccurred())
				// nilでないことを検証する
				Expect(categoryId).NotTo(BeNil())
				// valueフィールドのUUIDと仮のUUIDが等価であることを検証する
				Expect(categoryId.Value()).To(Equal(uuid))
			})
		})
	})
	When("空文字の場合 \n", func() {
		It("エラーが返される \n", func() {
			categoryId, err := categories.NewCategoryId("")
			// エラーであることを検証する
			Expect(err).To(HaveOccurred())
			// エラーがerrortype.DomainError型であることを検証する
			Expect(err).To(BeAssignableToTypeOf(&errortype.DomainError{}))
			// エラーメッセージを検証する
			Expect(err.Error()).To(Equal("商品カテゴリIdは、空文字列であってはなりません。"))
			// categoryId変数がnilであることを検証する
			Expect(categoryId).To(BeNil())
		})
	})
	When("36⽂字でない場合\n", func() {
		It("エラーが返される \n", func() {
			uuid := "12345" // 不正な⻑さのUUID
			categoryId, err := categories.NewCategoryId(uuid)
			// エラーであることを検証する
			Expect(err).To(HaveOccurred())
			// エラーがerrortype.DomainError型であることを検証する
			Expect(err).To(BeAssignableToTypeOf(&errortype.DomainError{}))
			// エラーメッセージを検証する
			Expect(err.Error()).To(Equal("商品カテゴリIdは、36文字でなければなりません。"))
			// categoryId変数がnilであることを検証する
			Expect(categoryId).To(BeNil())
		})
	})
	When("UUID形式でない場合\n", func() {
		It("エラーが返される \n", func() {
			uuid := "yy3e45z--e89b-12d3-a456-426614174000" // UUIDフォーマットではない
			categoryId, err := categories.NewCategoryId(uuid)
			// エラーであることを検証する
			Expect(err).To(HaveOccurred())
			// エラーがerrortype.DomainError型であることを検証する
			Expect(err).To(BeAssignableToTypeOf(&errortype.DomainError{}))
			// エラーメッセージを検証する
			Expect(err.Error()).To(Equal("商品カテゴリIdは、UUID形式でなければなりません。"))
			// categoryId変数がnilであることを検証する
			Expect(categoryId).To(BeNil())
		})
	})
	Describe("Equals()メソッドの検証 \n", func() {
		When("同じ値を持つ2つのCategoryIdインスタンスを⽐較する場合 \n", func() {
			It("trueが返される\n", func() {
				uuid := "123e4567-e89b-12d3-a456-426614174000" // 有効なUUID
				categoryId1, err1 := categories.NewCategoryId(uuid)
				categoryId2, err2 := categories.NewCategoryId(uuid)
				// エラーになっていないことを検証する
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
				// categoryId1とcategoryId2の比較結果がtrueであることを検証する
				Expect(categoryId1.Equals(categoryId2)).To(BeTrue())
			})
		})
		When("値が異なる2つのCategoryIdインスタンスを比較する場合 \n", func() {
			It("falseが返される \n", func() {
				uuid1 := "123e4567-e89b-12d3-a456-426614174000" // 有効なUUID
				uuid2 := "123e4567-e89b-12d3-a456-426614174001" // 別の有効なUUID
				categoryId1, err1 := categories.NewCategoryId(uuid1)
				categoryId2, err2 := categories.NewCategoryId(uuid2)
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
				Expect(categoryId1.Equals(categoryId2)).To(BeFalse())
			})
		})
		When("CategoryIdインスタンスと異なる型のインスタンスを⽐較する場合 \n", func() {
			It("falseが返される \n", func() {
				uuid := "123e4567-e89b-12d3-a456-426614174000" // 有効なUUID
				categoryId, err := categories.NewCategoryId(uuid)
				Expect(err).NotTo(HaveOccurred())
				// 仮に異なる型のオブジェクトを作成
				nonCategoryId, err := categories.NewCategoryName("文房具")
				Expect(err).NotTo(HaveOccurred())
				// CategoryIdと明らかに異なる型との⽐較
				Expect(categoryId.Equals(nonCategoryId)).To(BeFalse())
			})
		})
	})
})
