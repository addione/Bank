package src

import (
	"github.com/addione/New/manager"
)

type srcDiContainer struct {
	bank               *bank
	managerDIContainer *manager.ManagerDIContainer
	controllers        Controllers
}

type Controllers struct {
	userController *userController
}

func NewContollersDI(sdi *srcDiContainer) Controllers {
	cdi := Controllers{
		userController: NewUserController(sdi),
	}
	return cdi
}

func NewSrcDI() *srcDiContainer {

	sdi := &srcDiContainer{
		bank: newBank(),
	}
	sdi.managerDIContainer = manager.NewManagerDIContainer()
	sdi.controllers = NewContollersDI(sdi)
	sdi.bank = newBank()
	return sdi
}

func (sdi *srcDiContainer) GetBank() *bank {
	return newBank()
}

func (sdi *srcDiContainer) GetUserController() *userController {
	return sdi.controllers.userController
}
