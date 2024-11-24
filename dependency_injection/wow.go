package dependency_injection

// REPOSITORY ADD CONSTRUCTOR(PROVIDER)

type WowRepository struct {
}

func NewWowRepository() *WowRepository {
	return &WowRepository{}
}

// SERVICE

type WowService struct {
	DB bool
	*WowRepository
}

func NewWowService(DB bool, wowRepository *WowRepository) *WowService {
	return &WowService{DB: DB, WowRepository: wowRepository}
}

// CONTROLLER

type WowController struct {
	*WowService
}

func NewWowController(wowService *WowService) *WowController {
	return &WowController{WowService: wowService}
}
