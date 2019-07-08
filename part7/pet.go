package main


type Pet struct {
	Name string `json:"name"`
	Type string `json:"type"`
}



type Cat struct {
	Pet
}

func CreateCat(name string,catType string) *Cat{
	return &Cat{Pet{Name:name,Type:catType}}
}

type Dog struct {
	Pet
}

func CreateDog(name string,catType string) *Dog{
	return &Dog{Pet{Name:name,Type:catType}}
}

type Hamster struct {
	Pet
}

func CreateHamster(name string,catType string) *Hamster{
	return &Hamster{Pet{Name:name,Type:catType}}
}


