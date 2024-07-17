package repository

type example struct {
}

type IExampleRepository interface {
	GetAll() map[string]interface{}
}

func NewExampleRepository() IExampleRepository {
	return &example{}
}

func (r *example) GetAll() map[string]interface{} {
	return map[string]interface{}{
		"id":   1,
		"name": "Example Name",
		"age":  12,
	}
}
