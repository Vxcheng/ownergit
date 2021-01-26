package usecase

import "ownergit/external_libs/gomockery/domain"

type cImpl struct {
	C domain.C
}

func (c *cImpl) Collect(a domain.A) domain.B {
	return c.C.CreateB(a)
}
