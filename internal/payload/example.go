package payload

type CreateExampleRequest struct {
	Name string `json:"name" binding:"required" validate:"required"`
	Age  int    `json:"age" binding:"required,gte=17" binding:"required,gte=17"`
}
