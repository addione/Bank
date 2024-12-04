package helpers

type HelpersDIContainer struct {
	hashing   *Hashing
	jwtHelper *JwtHelper
	envHelper *EnvHelper
}

func NewHelpersDIContainer() *HelpersDIContainer {
	return &HelpersDIContainer{
		hashing:   newHashing(),
		jwtHelper: newJwtHelper(),
		envHelper: newEnvHelper(),
	}
}

func (hdi *HelpersDIContainer) GetHashing() *Hashing {
	return hdi.hashing
}

func (hdi *HelpersDIContainer) GetJwtTokenHelper() *JwtHelper {
	return hdi.jwtHelper
}
