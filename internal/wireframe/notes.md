# Probe Notes

This section of the Service Pack Wireframe contains a single probe.

Multiple directories such as this will result in mutliple probes.

Each probe should meet the SDK's probeengine interface `Probe`

```go
type Probe interface {
	ProbeInitialize(*godog.TestSuiteContext)
	ScenarioInitialize(*godog.ScenarioContext)
	Name() string
	Path() string
}
```
