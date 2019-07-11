package pet

type Hamster struct {
	Pet
}

func CreateHamster(name string,catType string) *Hamster{
	return &Hamster{Pet{Name:name,Type:catType}}
}
