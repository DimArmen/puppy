package puppy

immport (
	"github.com/DimArmen/dog"
)
type Puppy struct {
	Name   string
	Breed  string
	Age    int
	Weight int
}

func (p *Puppy) Bark() string {
	return dog.Bark(p.Name)
}
