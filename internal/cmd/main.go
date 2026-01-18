package main

import (
	"fmt"
	"throughput-tester-client/internal/config"
)

func main() {
	config := config.Load()
	fmt.Println(config.RuntimeConfigs.TestVar1)
}