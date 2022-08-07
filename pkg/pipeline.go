package pkg

import (
	"github.com/analogj/sunlight/pkg/config"
)

type Pipeline struct {
	Config config.Interface
}

func (p *Pipeline) Start(configData config.Interface) error {
	// Initialize Pipeline.
	p.Config = configData
	return nil
}
