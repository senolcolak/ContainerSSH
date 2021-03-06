package steps

import (
	"fmt"
	"github.com/containerssh/containerssh/test/config"
)

func (scenario *Scenario) StartConfigServer() error {
	if scenario.ConfigServer != nil {
		return fmt.Errorf("config server is already running")
	}
	scenario.ConfigServer = config.NewMemoryConfigServer()
	return scenario.ConfigServer.Start()
}

func (scenario *Scenario) StopConfigServer() error {
	if scenario.ConfigServer == nil {
		return fmt.Errorf("config server is not running")
	}
	err := scenario.ConfigServer.Stop()
	if err != nil {
		return err
	}
	scenario.ConfigServer = nil
	return nil
}
