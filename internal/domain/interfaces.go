package domain

type Entity interface {
	ID() string
}

type Codex interface {
	Init() error
}
type StorageManager interface {
	LoadCodex(id string, codex interface{}) error
	SaveCodex(id string, data interface{}) error
}
