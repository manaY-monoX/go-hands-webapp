package tests

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/categories"
	"service-exercise/domain/model/products"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("エンティティ Product\n", func() {
	Describe("NewProduct()関数を検証する\n", func() {
		// 値オブジェクトの変数を準備する
		var (
			validId       *products.ProductId
			validName     *products.ProductName
			validPrice    *products.ProductPrice
			validCategory *categories.Category
		)
		BeforeEach(func() {
			// 有効なインスタンスをセットアップする
			validId = products.NewProductIdWithUUID()
			validName, _ = products.NewProductName("ボールペン")
			validPrice, _ = products.NewProductPrice(100)
			categoryId := categories.NewCategoryIdWithUUID()
			categoryName, _ := categories.NewCategoryName("文房具")
			validCategory, _ = categories.NewCategory(categoryId, categoryName)
		})
		When("すべてのパラメータが有効な場合\n", func() {
			It("エラーなくProductのインスタンスが返される\n", func() {
				product, err := products.NewProduct(validId, validName, validPrice, validCategory)
				Expect(err).NotTo(HaveOccurred())
				Expect(product).NotTo(BeNil())
				Expect(product.Id()).To(Equal(validId))
				Expect(product.Name()).To(Equal(validName))
				Expect(product.Price()).To(Equal(validPrice))
				Expect(product.Category()).To(Equal(validCategory))
			})
		})
		When("idがnilの場合\n", func() {
			It("エラーが返される\n", func() {
				product, err := products.NewProduct(nil, validName, validPrice, validCategory)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errortype.NewDomainError("商品Idは、必須です。")))
				Expect(product).To(BeNil())
			})
		})
	})
	Describe("Equals()メソッドの検証\n", func() {
		var (
			product1 *products.Product
			product2 *products.Product
		)
		BeforeEach(func() {
			// テスト用のProductインスタンスをセットアップする
			id1, _ := products.NewProductId("123e4567-e89b-12d3-a456-426614174000")
			name, _ := products.NewProductName("ボールペン")
			price, _ := products.NewProductPrice(100)
			categoryId := categories.NewCategoryIdWithUUID()
			categoryName, _ := categories.NewCategoryName("文房具")
			category, _ := categories.NewCategory(categoryId, categoryName)
			product1, _ = products.NewProduct(id1, name, price, category)

			id2, _ := products.NewProductId("123e4567-e89b-12d3-a456-426614174055")
			product2, _ = products.NewProduct(id2, name, price, category)
		})
		When("2つの商品が同じIDを持つ場合\n", func() {
			It("trueが返される\n", func() {
				duplicateProduct, _ := products.NewProduct(product1.Id(), product1.Name(), product1.Price(), product1.Category())
				Expect(product1.Equals(duplicateProduct)).To(BeTrue())
			})
		})
		When("2つの商品のIDが異なる場合\n", func() {
			It("falseが返される\n", func() {
				Expect(product1.Equals(product2)).To(BeFalse())
			})
		})
		When("nilを渡した場合\n", func() {
			It("falseが返される", func() {
				Expect(product1.Equals(nil)).To(BeFalse())
			})
		})
		When("別の型のインスタンスと比較した場合\n", func() {
			It("falseが返される\n", func() {
				Expect(product1.Equals(product2.Category())).To(BeFalse())
			})
		})
	})
	Describe("ChangeName()メソッドの検証", func() {
		var product *products.Product

		BeforeEach(func() {
			// テスト用のProductインスタンスを初期化する
			productId := products.NewProductIdWithUUID()
			productName, _ := products.NewProductName("初期商品名")
			productPrice, _ := products.NewProductPrice(100)
			categoryId := categories.NewCategoryIdWithUUID()
			categoryName, _ := categories.NewCategoryName("文房具")
			category, _ := categories.NewCategory(categoryId, categoryName)
			product, _ = products.NewProduct(productId, productName, productPrice, category)
		})
		When("有効なProductNameが指定された場合\n", func() {
			It("商品名は正常に変更される\n", func() {
				newName, _ := products.NewProductName("ボールペン(青)")
				err := product.ChangeName(newName)
				Expect(err).To(BeNil())
			})
		})
		When("nilが渡された場合\n", func() {
			It("エラーを返す\n", func() {
				err := product.ChangeName(nil)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errortype.NewDomainError("商品名は、必須です。")))
			})
		})
	})
	Describe("ChangePrice()メソッドの検証", func() {
		var product *products.Product

		BeforeEach(func() {
			// テスト用のProductインスタンスを初期化する
			productId := products.NewProductIdWithUUID()
			productName, _ := products.NewProductName("初期商品名")
			productPrice, _ := products.NewProductPrice(100)
			categoryId := categories.NewCategoryIdWithUUID()
			categoryName, _ := categories.NewCategoryName("文房具")
			category, _ := categories.NewCategory(categoryId, categoryName)
			product, _ = products.NewProduct(productId, productName, productPrice, category)
		})
		When("有効なProductPriceが指定された場合\n", func() {
			It("商品単価は正常に変更される\n", func() {
				newPrice, _ := products.NewProductPrice(120)
				err := product.ChangePrice(newPrice)
				Expect(err).To(BeNil())
			})
		})
		When("nilが渡された場合\n", func() {
			It("エラーを返す\n", func() {
				err := product.ChangePrice(nil)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errortype.NewDomainError("商品単価は、必須です。")))
			})
		})
	})
	Describe("ChangeCategory()メソッドの検証", func() {
		var product *products.Product

		BeforeEach(func() {
			// テスト用のProductインスタンスを初期化する
			productId := products.NewProductIdWithUUID()
			productName, _ := products.NewProductName("初期商品名")
			productPrice, _ := products.NewProductPrice(100)
			categoryId := categories.NewCategoryIdWithUUID()
			categoryName, _ := categories.NewCategoryName("文房具")
			category, _ := categories.NewCategory(categoryId, categoryName)
			product, _ = products.NewProduct(productId, productName, productPrice, category)
		})
		When("有効なCategoryが指定された場合\n", func() {
			It("商品カテゴリは正常に変更される\n", func() {
				categoryId := categories.NewCategoryIdWithUUID()
				categoryName, _ := categories.NewCategoryName("文房具")
				category, _ := categories.NewCategory(categoryId, categoryName)
				err := product.ChangeCategory(category)
				Expect(err).To(BeNil())
			})
		})
		When("nilが渡された場合\n", func() {
			It("エラーを返す\n", func() {
				err := product.ChangeCategory(nil)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(errortype.NewDomainError("商品カテゴリは、必須です。")))
			})
		})
	})
})
