package interfaces

type Repository interface {
	LoadFromStorage(id string) error
	SaveToStorage(id string) error
}
