package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"service-exercise/presentation/dto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("商品登録のルーティング品質検証\n", Ordered, func() {
	var env *PresentationTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})
	When("存在しない商品を登録する\n", func() {
		var req *http.Request
		var resp *httptest.ResponseRecorder
		BeforeEach(func() {
			category := dto.NewCatgeoryDTO("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
			product := dto.NewProductDTO("", "消しゴム", "150", category)
			jsonBody, _ := json.Marshal(product)
			req, _ = http.NewRequest("POST", "/product/register", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			resp = httptest.NewRecorder()
			env.router.Engine.ServeHTTP(resp, req)
		})
		It("ステータスが200で期待される商品データがJSONで返される\n", func() {
			Expect(resp.Code).To(Equal(http.StatusOK))
			expectedJSON := `
			{
				"productId": "",
				"productName": "消しゴム",
				"productPrice": "150",
				"category": {
				  "categoryId": "b1524011-b6af-417e-8bf2-f449dd58b5c0",
				  "categoryName": "文房具"
				}
			}
			`
			Expect(resp.Body.String()).To(MatchJSON(expectedJSON))
		})
	})
	When("登録済み商品を登録する\n", func() {
		var req *http.Request
		var resp *httptest.ResponseRecorder
		BeforeEach(func() {
			category := dto.NewCatgeoryDTO("b1524011-b6af-417e-8bf2-f449dd58b5c0", "文房具")
			product := dto.NewProductDTO("", "消しゴム", "150", category)
			jsonBody, _ := json.Marshal(product)
			req, _ = http.NewRequest("POST", "/product/register", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			resp = httptest.NewRecorder()
			env.router.Engine.ServeHTTP(resp, req)
		})
		It("ステータスが400で期待されるエラーメッセージが返される\n", func() {
			Expect(resp.Code).To(Equal(http.StatusBadRequest))
			Expect(resp.Body.String()).To(Equal("\"商品:消しゴムは、既に登録済です。\""))
		})
	})
})
