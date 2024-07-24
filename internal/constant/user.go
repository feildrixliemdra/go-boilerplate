package constant

import "fmt"

var (
	ErrEmailAlreadyRegistered = fmt.Errorf("email already registered")
	ErrUserNotFound           = fmt.Errorf("user not found")
)
