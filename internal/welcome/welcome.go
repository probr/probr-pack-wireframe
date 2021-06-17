// Package welcome provides the implementation required to execute the BDD tests described in container_registry_access.feature file
package welcome

import (
	"github.com/cucumber/godog"

	"github.com/probr/probr-pack-wireframe/internal/config"
	"github.com/probr/probr-pack-wireframe/internal/summary"
	audit "github.com/probr/probr-sdk/audit"
	"github.com/probr/probr-sdk/probeengine"
	"github.com/probr/probr-sdk/utils"
)

type probeStruct struct{}

// scenarioState holds the steps and state for any scenario in this probe
type scenarioState struct {
	name        string
	currentStep string
	audit       *audit.Scenario
	probe       *audit.Probe
}

// Probe meets the service pack interface for adding the logic from this file
var Probe probeStruct
var scenario scenarioState

func (scenario *scenarioState) testShouldRun() error {
	// Standard auditing logic to ensures panics are also audited
	stepTrace, payload, err := utils.AuditPlaceholders()
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = utils.ReformatError("[ERROR] Unexpected behavior occured: %s", panicErr)
		}
		scenario.audit.AuditScenarioStep(scenario.currentStep, stepTrace.String(), payload, err)
	}()

	stepTrace.WriteString("Validate that test is intended to pass; ")

	payload = struct {
		TestShouldPass string
	}{
		config.Vars.ServicePacks.Wireframe.Pass,
	}

	if config.Vars.ServicePacks.Wireframe.Pass == "false" {
		err = utils.ReformatError("Config state ServicePacks.Wireframe.Pass was false")
	}

	return err
}

func (scenario *scenarioState) visitorIsRespondedToAtTheDoor(arrival, response string) error {
	// Supported values for 'response':
	//	'welcomed'
	//	'rejected'

	// Standard auditing logic to ensures panics are also audited
	stepTrace, payload, err := utils.AuditPlaceholders()
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = utils.ReformatError("[ERROR] Unexpected behavior occured: %s", panicErr)
		}
		scenario.audit.AuditScenarioStep(scenario.currentStep, stepTrace.String(), payload, err)
	}()
	family := []string{"mom", "dad", "sister", "brother"}
	var properResponse bool

	// Validate input values
	switch response {
	case "welcomed":
		for _, familyMember := range family {
			if arrival == familyMember {
				properResponse = true // welcome the family inside!
			}
		}
	case "rejected":
		properResponse = true // reject anyone who isn't family!
		for _, familyMember := range family {
			if arrival == familyMember {
				properResponse = false
			}
		}
	default:
		err = utils.ReformatError("Unexpected value provided for expectedResponse: '%s' Expected values: ['welcomed', 'rejected']", response)
		return err
	}

	payload = struct {
		Arrival           string
		Response          string
		ResponseWasProper bool
	}{
		Arrival:           arrival,
		Response:          response,
		ResponseWasProper: properResponse,
	}

	if !properResponse {
		err = utils.ReformatError("Shouldn't have %s %s!")
	}

	return err
}

// Name presents the name of this probe for external reference
func (probe probeStruct) Name() string {
	// The return value for `Name` should match the probe directory
	// and it's feature file, so each may be properly addressed for
	// packing and opening the files.
	return "wireframe"
}

// Path presents the path of these feature files for external reference
func (probe probeStruct) Path() string {
	// this should reference the probe parent directory (usually `internal/<probe-name>`)
	return probeengine.GetFeaturePath("internal", probe.Name())
}

// ProbeInitialize handles any overall Test Suite initialisation steps.  This is registered with the
// test handler as part of the init() function.
func (probe probeStruct) ProbeInitialize(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
	})

	ctx.AfterSuite(func() {
	})
}

// ScenarioInitialize provides initialization logic before each scenario is executed
func (probe probeStruct) ScenarioInitialize(ctx *godog.ScenarioContext) {

	ctx.BeforeScenario(func(s *godog.Scenario) {
		scenario.name = s.Name
		scenario.probe = summary.State.GetProbeLog(probe.Name())
		scenario.audit = summary.State.GetProbeLog(probe.Name()).InitializeAuditor(s.Name, s.Tags)
		probeengine.LogScenarioStart(s)
	})

	ctx.BeforeStep(func(st *godog.Step) {
		// placeholder
	})

	// Background
	ctx.Step(`^the config states that this test should run$`, scenario.testShouldRun)

	// Steps
	ctx.Step(`^"([^"]*)" is "([^"]*)" when they come to the door$`, scenario.visitorIsRespondedToAtTheDoor)

	ctx.AfterStep(func(st *godog.Step, err error) {
		// placeholder
	})

	ctx.AfterScenario(func(s *godog.Scenario, err error) {
		probeengine.LogScenarioEnd(s)
	})

}
