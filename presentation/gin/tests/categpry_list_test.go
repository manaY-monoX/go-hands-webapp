package tests

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("商品カテゴリ一覧のルーティング品質検証\n", Ordered, func() {
	var env *PresentationTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})
	When("商品カテゴリ一覧を取得する\n", func() {
		var req *http.Request
		var resp *httptest.ResponseRecorder
		BeforeEach(func() {
			// 疑似的なリクエストを生成する
			req, _ = http.NewRequest("GET", "/category/list", nil)
			// ResponseRecorder(レスポンスの記録)を生成する
			resp = httptest.NewRecorder()
			// リクエストをルーターに送信
			env.router.Engine.ServeHTTP(resp, req)
		})
		It("ステータスが200で期待されるJSONが返される\n", func() {
			// レスポンスを検証
			Expect(resp.Code).To(Equal(http.StatusOK))
			// 期待されるJSONレスポンス
			expectedJSON := `[
				{
					"categoryId": "b1524011-b6af-417e-8bf2-f449dd58b5c0",
					"categoryName": "文房具"
				},
				{
					"categoryId": "762bd1ea-9700-4bab-a28d-6cbebf20ddc2",
					"categoryName": "雑貨"
				},
				{
					"categoryId": "c05b1952-3bdf-4449-9b83-d0d123a667ce",
					"categoryName": "パソコン周辺機器"
				}
		  ]`
			// レスポンスボディが期待されるJSONと一致することを検証
			Expect(resp.Body.String()).To(MatchJSON(expectedJSON))
		})
	})
})
