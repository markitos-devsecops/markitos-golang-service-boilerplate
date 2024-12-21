package domain

type BoilerRepository interface {
	Create(boiler *Boiler) error
	Delete(id *string) error
	Update(boiler *Boiler) error
	One(id *string) (*Boiler, error)
	List() ([]*Boiler, error)
	SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*Boiler, error)
}
