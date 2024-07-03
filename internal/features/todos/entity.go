package todos

type Todo struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
	Status      bool
}

type Hendler interface {
}

type Services interface {
}

type Query interface {
}