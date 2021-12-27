package productos

type Product struct {
	Id int 
	Nombre string
	Color string
	Precio float64
	Stock int
	Codigo string
	Publicado bool
}

var ps []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(Id int, Nombre string, Color string, Precio float64, Stock int, Codigo int, Publicado bool) (Product, error)
	LastID() (int, error)
	Update(Id int, Nombre string, Color string, Precio float64, Stock int, Codigo int, Publicado bool) (Product, error)
	UpdateName(Id int, Nombre string) (Product, error)
	Delete(Id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(Id int, Nombre string, Color string, Precio float64, Stock int, Codigo int, Publicado bool) (Product, error) {
	p := Product{Id, Nombre, Color, Precio, Stock, Codigo, Publicado}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

func (r *repository) Update(Id int, Nombre string, Color string, Precio float64, Stock int, Codigo int, Publicado bool) (Product, error) {
	p := Product{Nombre: Nombre, Color: Color, Color: Color, Precio: Precio, Stock: Stock, Codigo: Codigo, Publicado: Publicado}
	updated := false
	for i := range ps {
		if ps[i].Id == Id {
			p.Id = Id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", Id)
	}
	return p, nil
}

func (r *repository) UpdateName(Id int, Nombre string) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].Id == Id {
			ps[i].Nombre = Nombre
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", Id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == Id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("producto %d no encontrado", Id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil
}
