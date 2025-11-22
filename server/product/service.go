package product

import "ecommerce-server/domain"

type service struct {
	prdctRepo ProductRepo
}

func NerService(prdctRepo ProductRepo) Service {
	return  &service{
		prdctRepo: prdctRepo,
	}
}

	func (svc *service) Create(prd domain.Product) (*domain.Product, error){
		return  svc.prdctRepo.Create(prd)
	}
	func (svc *service) Get(id int) (*domain.Product, error){
		return  svc.prdctRepo.Get(id)
	}
	func (svc *service) List(page, limit int64) ([]*domain.Product, error){
		return  svc.prdctRepo.List(page, limit)
	}
	func (svc *service) Count() (int64, error) {
		return svc.prdctRepo.Count()
	}
	func (svc *service) Update(prd domain.Product) (*domain.Product, error){
		return  svc.prdctRepo.Update(prd)
	}
	func (svc *service) Delete(id int) error {
		return  svc.prdctRepo.Delete(id)
	}