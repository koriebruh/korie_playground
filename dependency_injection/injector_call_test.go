package dependency_injection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInjector(t *testing.T) {
	wowHandler := InitializeWowController(true)
	assert.NotNil(t, wowHandler)
}
