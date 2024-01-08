package model

type Category struct {
	ID   int
	Name string
}

type CategoryCreateArgs struct {
	Name string `validate:"required,max=200,min=5" json:"name"`
}

type CategoryUpdateArgs struct {
	ID   int    `validate:"required"`
	Name string `validate:"required,max=200,min=5" json:"name"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
