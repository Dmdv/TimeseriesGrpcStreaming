package initializers

import (
	"github.com/gobuffalo/envy"
)

func InitializeEnvs() {
	envy.Load()
}
