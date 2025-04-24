package tests

import (
	"service-exercise/presentation"
	"service-exercise/presentation/gin/preparation"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestPresentationSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "プレゼンテーション層のテストスイート")
}

// プレゼンテーション層のテストに必要な環境
type PresentationTestEnvironment struct {
	router *preparation.Router
	app    *fxtest.App
}

// テストに必要な環境を準備する
func SetupTestEnvironment() *PresentationTestEnvironment {
	var env PresentationTestEnvironment
	app := fxtest.New(
		GinkgoT(),
		presentation.PresentationModule,
		fx.Populate(&env.router),
	)
	env.app = app
	env.app.RequireStart()
	return &env
}

// fxコンテナを停止する
func TeardownTestEnvironment(env *PresentationTestEnvironment) {
	env.app.RequireStop()
}
