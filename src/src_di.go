package src

import "github.com/addione/New/manager"

type srcDiContainer struct {
	bank               *bank
	userController     *userController
	managerDIContainer *manager.ManagerDIContainer
}

func NewSrcDI() *srcDiContainer {

	sdi := &srcDiContainer{
		bank:           newBank(),
		userController: NewUserController(),
	}
	sdi.managerDIContainer = manager.NewManagerDIContainer()
	sdi.bank = newBank()
	return sdi
}

func (sdi *srcDiContainer) GetBank() *bank {
	return newBank()
}

func (sdi *srcDiContainer) GetUserController() *userController {
	return NewUserController()
}
