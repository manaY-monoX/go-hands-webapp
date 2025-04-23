package tests

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/products"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("値オブジェクト ProductId\n", func() {
	Describe("NewProductIdWithUUID()関数を検証する\n", func() {
		It("有効なUUIDのProductIdが作成される \n", func() {
			productId := products.NewProductIdWithUUID()
			// ProductIdがnilでないことを検証する
			Expect(productId).NotTo(BeNil())
			// ProductIdのvalueがUUID形式であることを検証する
			_, err := uuid.Parse(productId.Value())
			Expect(err).NotTo(HaveOccurred())
			// ProductIdのvalueが36⽂字であることを検証する
			Expect(len(productId.Value())).To(Equal(36))
		})
	})
	Describe("NewProductId()関数を検証する\n", func() {
		When("有効なUUIDを使⽤した場合 \n", func() {
			It("エラーなくProductIdが返される \n", func() {
				uuid := "123e4567-e89b-12d3-a456-426614174000" // 仮の有効なUUID
				productId, err := products.NewProductId(uuid)
				// エラーが発⽣していないことを検証する
				Expect(err).NotTo(HaveOccurred())
				// nilでないことを検証する
				Expect(productId).NotTo(BeNil())
				// valueフィールドのUUIDと仮のUUIDが等価であることを検証する
				Expect(productId.Value()).To(Equal(uuid))
			})
		})
		When("空文字を使用の場合\n", func() {
			It("エラーが返される\n", func() {
				productId, err := products.NewProductId("")
				// エラーであることを検証する
				Expect(err).To(HaveOccurred())
				// エラーがerrortype.DomainError型であることを検証する
				Expect(err).To(BeAssignableToTypeOf(&errortype.DomainError{}))
				// エラーメッセージを検証する
				Expect(err.Error()).To(Equal("商品Idは、空文字列であってはなりません。"))
				// categoryId変数がnilであることを検証する
				Expect(productId).To(BeNil())
			})
		})
		When("36⽂字でない場合\n", func() {
			It("エラーが返される \n", func() {
				uuid := "12345" // 不正な⻑さのUUID
				productId, err := products.NewProductId(uuid)
				// エラーであることを検証する
				Expect(err).To(HaveOccurred())
				// エラーがerrortype.DomainError型であることを検証する
				Expect(err).To(BeAssignableToTypeOf(&errortype.DomainError{}))
				// エラーメッセージを検証する
				Expect(err.Error()).To(Equal("商品Idは、36文字でなければなりません。"))
				// categoryId変数がnilであることを検証する
				Expect(productId).To(BeNil())
			})
		})
		When("UUID形式でない場合\n", func() {
			It("エラーが返される \n", func() {
				uuid := "yy3e45z--e89b-12d3-a456-426614174000" // UUIDフォーマットではない
				productId, err := products.NewProductId(uuid)
				// エラーであることを検証する
				Expect(err).To(HaveOccurred())
				// エラーがerrortype.DomainError型であることを検証する
				Expect(err).To(BeAssignableToTypeOf(&errortype.DomainError{}))
				// エラーメッセージを検証する
				Expect(err.Error()).To(Equal("商品Idは、UUID形式でなければなりません。"))
				// categoryId変数がnilであることを検証する
				Expect(productId).To(BeNil())
			})
		})
	})

	Describe("Equals()メソッドの検証 \n", func() {
		When("同じ値を持つ2つのProductIdインスタンスを⽐較する場合 \n", func() {
			It("trueが返される\n", func() {
				uuid := "123e4567-e89b-12d3-a456-426614174000" // 有効なUUID
				productId1, err1 := products.NewProductId(uuid)
				productId2, err2 := products.NewProductId(uuid)
				// エラーになっていないことを検証する
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
				// productId1とproductId2の比較結果がtrueであることを検証する
				Expect(productId1.Equals(productId2)).To(BeTrue())
			})
		})
		When("値が異なる2つのProductIdインスタンスを比較する場合 \n", func() {
			It("falseが返される \n", func() {
				uuid1 := "123e4567-e89b-12d3-a456-426614174000" // 有効なUUID
				uuid2 := "123e4567-e89b-12d3-a456-426614174001" // 別の有効なUUID
				productId1, err1 := products.NewProductId(uuid1)
				productId2, err2 := products.NewProductId(uuid2)
				Expect(err1).NotTo(HaveOccurred())
				Expect(err2).NotTo(HaveOccurred())
				Expect(productId1.Equals(productId2)).To(BeFalse())
			})
		})
		When("ProductIdインスタンスと異なる型のインスタンスを⽐較する場合 \n", func() {
			It("falseが返される \n", func() {
				uuid := "123e4567-e89b-12d3-a456-426614174000" // 有効なUUID
				productId, err := products.NewProductId(uuid)
				Expect(err).NotTo(HaveOccurred())
				// 仮に異なる型のオブジェクトを作成
				nonProductId, err := products.NewProductName("ボールペン")
				Expect(err).NotTo(HaveOccurred())
				// ProductIdと明らかに異なる型との⽐較
				Expect(productId.Equals(nonProductId)).To(BeFalse())
			})
		})
	})
})
