package steps

import "github.com/cucumber/godog"

func thatTheInitialConditionsAreMet() {}

func theProcessExecutes() {}

func thisShouldHappen() {}

func InitializeGenerateTradesReportScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^que as condicoes iniciais sejam satisfeitas$`, thatTheInitialConditionsAreMet)
	ctx.Step(`^o processo executar$`, theProcessExecutes)
	ctx.Step(`^isso deve acontecer$`, thisShouldHappen)
}
