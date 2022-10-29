package vagrant

type pool struct {
	vagrants []*aliveVagrant
}

func (p *pool) add(e ...*aliveVagrant) {
	p.vagrants = append(p.vagrants, e...)
}

func (p *pool) vagrant() *aliveVagrant {
	return p.vagrants[0]
}

func (p *pool) size() int {
	return len(p.vagrants)
}

func (p *pool) elements() []*aliveVagrant {
	return p.vagrants
}
