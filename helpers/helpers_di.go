package helpers

type HelpersDIContainer struct {
	hashing   *Hashing
	jwtHelper *JwtHelper
}

func NewHelpersDIContainer() *HelpersDIContainer {
	return &HelpersDIContainer{
		hashing:   newHashing(),
		jwtHelper: newJwtHelper(),
	}
}

func (hdi *HelpersDIContainer) GetHashing() *Hashing {
	return hdi.hashing
}

func (hdi *HelpersDIContainer) GetJwtTokenHelper() *JwtHelper {
	return hdi.jwtHelper
}
