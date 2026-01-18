package config

import (
	"os"
	"strconv"
)

// Runtime specific variables
type runtimeConfigs struct {
	TestVar1 int
	TestVar2 string
}

// Generic struct that holds specific structs
type Config struct {
	RuntimeConfigs *runtimeConfigs
	// Add more structs as needed
}

// Load configs from .env
func Load() *Config {
	return &Config{
		RuntimeConfigs: loadRuntimeConfigs(),
	}
}

// Load runtime configs from .env
func loadRuntimeConfigs() *runtimeConfigs {
	return &runtimeConfigs{
		TestVar1: loadIntfromEnv("TEST_VAR_1", 999),
		TestVar2: loadStrfromEnv("TEST_VAR_2", "some-value"),
	}
}

// Returns the value of varName. If it's empty or not present, then
// returns defVal
func loadIntfromEnv(varName string, defVal int) int {
	i, err := strconv.ParseInt(os.Getenv(varName), 10, 32)
	if err != nil {
		return defVal
	} else {
		return int(i)
	}
}

// Returns the value of varName even if it's empty (if present).
// Otherwise returns defVal
func loadStrfromEnv(varName, defVal string) string {
	val, present := os.LookupEnv(varName)
	if present {
		return val
	} else {
		return defVal
	}
}