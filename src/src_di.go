package src

import "github.com/addione/New/manager"

type srcDiContainer struct {
	bank               *bank
	managerDIContainer *manager.ManagerDIContainer
	controllers        Controllers
}

type Controllers struct {
	userController *userController
}

func NewContollersDI() Controllers {
	cdi := Controllers{
		userController: NewUserController(),
	}
	return cdi
}

func NewSrcDI() *srcDiContainer {

	sdi := &srcDiContainer{
		bank:        newBank(),
		controllers: NewContollersDI(),
	}
	sdi.managerDIContainer = manager.NewManagerDIContainer()
	sdi.bank = newBank()
	return sdi
}

func (sdi *srcDiContainer) GetBank() *bank {
	return newBank()
}

func (sdi *srcDiContainer) GetUserController() *userController {
	return sdi.controllers.userController
}
