package entities

type Bag struct {
	Colour   string
	Children []*Bag
}

func NewBag(colour string) *Bag {
	return &Bag{Colour: colour}
}

func (b *Bag) AddChild(child *Bag) {
	if len(b.Children) == 0 || b.Children == nil {
		b.Children = make([]*Bag, 1)
		b.Children[0] = child
	} else {
		b.Children = append(b.Children, child)
	}
}
