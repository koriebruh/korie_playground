// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependency_injection

// Injectors from injector.go:

func InitializeWowController(DB bool) *WowController {
	wowRepository := NewWowRepository()
	wowService := NewWowService(DB, wowRepository)
	wowController := NewWowController(wowService)
	return wowController
}
