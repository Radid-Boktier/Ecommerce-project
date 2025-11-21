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
	func (svc *service) List() ([]*domain.Product, error){
		return  svc.prdctRepo.List()
	}
	func (svc *service) Update(prd domain.Product) (*domain.Product, error){
		return  svc.prdctRepo.Update(prd)
	}
	func (svc *service) Delete(id int) error {
		return  svc.prdctRepo.Delete(id)
	}