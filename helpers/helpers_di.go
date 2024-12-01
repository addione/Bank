package helpers

type HelpersDIContainer struct {
	Hashing *Hashing
}

func NewHelpersDIContainer() *HelpersDIContainer {
	return &HelpersDIContainer{
		Hashing: newHashing(),
	}
}

func (hdi *HelpersDIContainer) GetHashing() *Hashing {
	return hdi.Hashing
}
