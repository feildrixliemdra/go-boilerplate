package service

type example struct {
	Option
}

type IExampleService interface {
	GetAll() map[string]interface{}
}

func NewExampleService(opt Option) IExampleService {
	return &example{
		opt,
	}
}

func (e *example) GetAll() map[string]interface{} {
	return e.Repository.ExampleRepository.GetAll()
}
