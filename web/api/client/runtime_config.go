package client

import (
	"github.com/raymao96/komari/database/clients"
	v2 "github.com/raymao96/komari/protocol/v2"
)

func getClientRuntimeConfig(uuid string) (*v2.ConfigParams, error) {
	return clients.RuntimeConfigForAgent(uuid)
}
