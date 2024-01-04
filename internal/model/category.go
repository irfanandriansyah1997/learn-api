package model

type Category struct {
	ID   int
	Name string
}

type CategoryCreateArgs struct {
	Name string
}

type CategoryUpdateArgs struct {
	ID   int
	Name string
}

type CategoryResponse struct {
	ID   int
	Name string
}
