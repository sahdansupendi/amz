package supplier

type Service interface {
	RegisterSupplier(input RegisterSupllierInput) (Supplier, error)
	GetSuppliers(supplierID int) ([]Supplier, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterSupplier(input RegisterSupllierInput) (Supplier, error) {
	supplier := Supplier{}
	supplier.NamaSupplier = input.NamaSupplier
	supplier.Alamat = input.Alamat
	supplier.NoHp = input.NoHp
	supplier.Email = input.Email

	newsupplier, err := s.repository.Save(supplier)

	if err != nil {
		return supplier, nil
	}

	return newsupplier, nil

}

func (s *service) GetSuppliers(supplierID int) ([]Supplier, error) {
	supplier, err := s.repository.FindAll()
	if err != nil {
		return supplier, err
	}
	return supplier, nil
}
