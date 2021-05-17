package config

import (
	"encoding/json"
	"log"

	sdkConfig "github.com/probr/probr-sdk/config"
	"github.com/probr/probr-sdk/config/setter"
	"github.com/probr/probr-sdk/utils"
)

// Vars is a stateful object containing the variables required to execute this pack
var Vars varOptions

// Init will set values with the content retrieved from a filepath, env vars, or defaults
func (ctx *varOptions) Init() (err error) {
	if ctx.VarsFile != "" {
		ctx.decode()
		if err != nil {
			log.Printf("[ERROR] %v", err)
			return
		}
	} else {
		log.Printf("[DEBUG] No vars file provided, unexpected behavior may occur")
	}
	sdkConfig.GlobalConfig.VarsFile = ctx.VarsFile
	sdkConfig.GlobalConfig.Init()

	ctx.ServicePacks.Wireframe.setEnvAndDefaults()

	log.Printf("[DEBUG] Config initialized by %s", utils.CallerName(1))
	return
}

// decode uses an SDK helper to create a YAML file decoder,
// parse the file to an object, then extracts the values from
// ServicePacks.Wireframe into this context
func (ctx *varOptions) decode() (err error) {
	configDecoder, file, err := sdkConfig.NewConfigDecoder(ctx.VarsFile)
	if err != nil {
		return
	}
	err = configDecoder.Decode(&ctx)
	file.Close()
	return err
}

// LogConfigState will write the config file to the write directory
func (ctx *varOptions) LogConfigState() {
	json, _ := json.MarshalIndent(ctx, "", "  ")
	log.Printf("[INFO] Config State: %s", json)
}

func (ctx *varOptions) Tags() string {
	return sdkConfig.ParseTags(ctx.ServicePacks.Wireframe.TagInclusions, ctx.ServicePacks.Wireframe.TagExclusions)
}

// setEnvOrDefaults will set value from os.Getenv and default to the specified value
func (ctx *wireframe) setEnvAndDefaults() {
	// Notes on SetVar's values:
	// 1. Pointer to local object; will be overwritten by env or default if empty
	// 2. Name of env var to check
	// 3. Default value to set if flags, vars file, and env have not provided a value

	setter.SetVar(&ctx.Pass, "PROBR_PASS_WIREFRAME_TESTS", "false")
}
