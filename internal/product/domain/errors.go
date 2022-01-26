package domain

var (
	ErrProductItemNotFound = NewRecordNotFound("El ID del producto no existe")
	ErrCategoryItemNotFound = NewRecordNotFound("El ID de la categoria no existe")
)

type ErrRecordNotFound struct {
	Message string
}

func (e *ErrRecordNotFound) Error() string {
	return e.Message
}

func NewRecordNotFound(message string) *ErrRecordNotFound {
	return &ErrRecordNotFound{message}
}
