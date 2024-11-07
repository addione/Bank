package src

type srcDiContainer struct {
	bank *bank
}

func NewSrcDI() *srcDiContainer {
	return &srcDiContainer{
		bank: newBank(),
	}
}

func (sdi *srcDiContainer) GetBank() *bank {
	return newBank()
}
