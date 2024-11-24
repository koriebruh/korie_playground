//go:build wireinject
// +build wireinject

package dependency_injection

import "github.com/google/wire"

func InitializeWowController(DB bool) *WowController {
	wire.Build(
		NewWowRepository,
		NewWowService,
		NewWowController,
	)
	return nil
}
