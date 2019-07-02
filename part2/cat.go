package main

type Cat struct {
	Name string `json:"name"`
	CatType string `json:"type"`
}

func CreateCat(name string,catType string) *Cat{
	return &Cat{Name:name,CatType:catType}
}
