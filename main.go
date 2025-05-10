package puppy

type Puppy struct {
	Name   string
	Breed  string
	Age    int
	Weight int
}

func (p *Puppy) Bark() string {
	return "Woof! Woof!"
}
