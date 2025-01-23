package repository

type Repo interface {
	Migrate(any) error
	Create(any) error
	ReadAll(any) error
	Update(any, uint) error
	Delete(any, uint) error
}
