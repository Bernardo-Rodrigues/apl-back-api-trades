package infra_test

import (
	"app/infra"
	"app/infra/bdd/steps"
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/sirupsen/logrus"
)

var scenariosMap = map[string]func(ctx *godog.ScenarioContext){
	"bdd/features": steps.InitializeGenerateTradesReportScenario,
}

func TestMain(m *testing.M) {
	go infra.Start()

	time.Sleep(5 * time.Second)

	runSuits()

	os.Exit(0)
}

func runSuits() {
	for path, scenario := range scenariosMap {
		opts := godog.Options{
			Output:    colors.Colored(os.Stdout),
			Format:    "pretty",
			Paths:     []string{path},
			Randomize: time.Now().UTC().UnixNano(),
		}

		suite := godog.TestSuite{
			Name:                "generating-trades-report",
			ScenarioInitializer: scenario,
			Options:             &opts,
		}.Run()

		if suite != 0 {
			os.Exit(suite)
		}
	}
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
}
