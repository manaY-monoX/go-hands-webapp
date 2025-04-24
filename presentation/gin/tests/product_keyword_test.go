package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("商品キーワード検索のルーティング品質検証\n", Ordered, func() {
	var env *PresentationTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})
	When("存在する商品キーワード入力する", func() {
		var req *http.Request
		var resp *httptest.ResponseRecorder
		BeforeEach(func() {
			// 疑似的なリクエストを生成する
			req, _ = http.NewRequest("GET", fmt.Sprintf("/product/keyword/%s", "蛍光"), nil)
			// ResponseRecorder(レスポンスの記録)を生成する
			resp = httptest.NewRecorder()
			// リクエストをルーターに送信
			env.router.Engine.ServeHTTP(resp, req)
		})
		It("ステータスが200で期待されるJSONが返される\n", func() {
			// レスポンスを検証
			Expect(resp.Code).To(Equal(http.StatusOK))
			expectedJSON := `[
				{
					"productId": "dc7243af-c2ce-4136-bd5d-c6b28ee0a20a",
					"productName": "蛍光ペン(黄)",
					"productPrice": "130",
					"category": {
						"categoryId": "b1524011-b6af-417e-8bf2-f449dd58b5c0",
						"categoryName": "文房具"
					}
				},
				{
					"productId": "83fbc81d-2498-4da6-b8c2-54878d3b67ff",
					"productName": "蛍光ペン(赤)",
					"productPrice": "130",
					"category": {
						"categoryId": "b1524011-b6af-417e-8bf2-f449dd58b5c0",
						"categoryName": "文房具"
					}
				},
				{
					"productId": "ee4b3752-3fbd-45fc-afb5-8f37c3f701c9",
					"productName": "蛍光ペン(青)",
					"productPrice": "130",
					"category": {
						"categoryId": "b1524011-b6af-417e-8bf2-f449dd58b5c0",
						"categoryName": "文房具"
					}
				},
				{
					"productId": "35cb51a7-df79-4771-9939-7f32c19bca45",
					"productName": "蛍光ペン(緑)",
					"productPrice": "130",
					"category": {
						"categoryId": "b1524011-b6af-417e-8bf2-f449dd58b5c0",
						"categoryName": "文房具"
					}
				}
			]`
			// レスポンスボディが期待されるJSONと一致することを検証
			Expect(resp.Body.String()).To(MatchJSON(expectedJSON))
		})
	})
	When("存在しない商品キーワード入力する", func() {
		var req *http.Request
		var resp *httptest.ResponseRecorder
		BeforeEach(func() {
			// 疑似的なリクエストを生成する
			req, _ = http.NewRequest("GET", fmt.Sprintf("/product/keyword/%s", "川"), nil)
			// ResponseRecorder(レスポンスの記録)を生成する
			resp = httptest.NewRecorder()
			// リクエストをルーターに送信
			env.router.Engine.ServeHTTP(resp, req)
		})
		It("ステータスが404で期待されるJSONが返される\n", func() {
			// レスポンスを検証
			Expect(resp.Code).To(Equal(http.StatusNotFound))
			Expect(resp.Body.String()).To(Equal("\"キーワード:'川'に該当する商品は見つかりませんでした。\""))
		})
	})
})
