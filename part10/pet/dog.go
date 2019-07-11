package pet

type Dog struct {
	Pet
}

func CreateDog(name string,catType string) *Dog{
	return &Dog{Pet{Name:name,Type:catType}}
}

