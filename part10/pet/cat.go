package pet

type Cat struct {
	Pet
}

func CreateCat(name string,catType string) *Cat{
	return &Cat{Pet{Name:name,Type:catType}}
}
