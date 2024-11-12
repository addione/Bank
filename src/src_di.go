package src

import "github.com/addione/New/manager"

type srcDiContainer struct {
	bank               *bank
	managerDIContainer *manager.ManagerDIContainer
}

func NewSrcDI() *srcDiContainer {

	sdi := &srcDiContainer{
		bank: newBank(),
	}
	sdi.managerDIContainer = manager.NewManagerDIContainer()
	sdi.bank = newBank()
	return sdi
}

func (sdi *srcDiContainer) GetBank() *bank {
	return newBank()
}
