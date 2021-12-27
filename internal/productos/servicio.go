package products

type Service interface {
	GetAll() ([]Product, error)
	Store(Nombre string, Color string, Precio float64, Stock int, Codigo int, Publicado bool) (Product, error)
	Update(Id int, Nombre string, Color string, Precio float64, Stock int, Codigo int, Publicado bool) (Product, error)
	UpdateName(Id int, Nombre string) (Product, error)
	Delete(Id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(Nombre string, Color string, Precio float64, Stock int, Codigo int, Publicado bool) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, Nombre, Color, Precio, Stock, Codigo, Publicado)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}

func (s *service) Update(Id int, Nombre string, Color string, Precio float64, Stock int, Codigo int, Publicado bool) (Product, error) {
	return s.repository.Update(Id, Nombre, Color, Precio, Stock, Codigo, Publicado)
}

func (s *service) UpdateName(Id int, Nombre string) (Product, error) {
	return s.repository.UpdateName(Id, Nombre)
}

func (s *service) Delete(Id int) error {
	return s.repository.Delete(Id)
}
