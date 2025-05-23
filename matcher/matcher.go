package matcher

import (
	"github.com/VincentBrodin/suddig/configs"
)

type Matcher struct {
	config configs.Config
}

func New(config configs.Config) *Matcher {
	return &Matcher{config: config}
}
