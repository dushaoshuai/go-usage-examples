package iterator

type Iterator interface {
	HasNext() bool
	GetNext() *Elem
}

type Elem struct {
	Name string
}

type Collection struct {
	data []*Elem
}

func NewCollection(elems ...*Elem) *Collection {
	c := &Collection{}
	for _, e := range elems {
		c.data = append(c.data, e)
	}
	return c
}

// forward traversal
type f struct {
	*Collection
	index int
}

func (es *Collection) ForwardIterator() Iterator {
	return &f{Collection: es, index: 0}
}

func (f *f) HasNext() bool {
	return f.index < len(f.data)
}

func (f *f) GetNext() *Elem {
	if !f.HasNext() {
		return nil
	}
	e := f.data[f.index]
	f.index++
	return e
}

// reverse traversal
type r struct {
	*Collection
	index int
}

func (es *Collection) ReverseIterator() Iterator {
	return &r{
		Collection: es,
		index:      len(es.data) - 1,
	}
}

func (r *r) HasNext() bool {
	return r.index >= 0
}

func (r *r) GetNext() *Elem {
	if !r.HasNext() {
		return nil
	}
	e := r.data[r.index]
	r.index--
	return e
}
