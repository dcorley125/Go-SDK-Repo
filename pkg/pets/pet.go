package pets

type Pet struct {
	Id   *int64  `json:"id,omitempty" required:"true"`
	Name *string `json:"name,omitempty" required:"true"`
	Tag  *string `json:"tag,omitempty"`
}

func (p *Pet) SetId(id int64) {
	p.Id = &id
}

func (p *Pet) GetId() *int64 {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *Pet) SetName(name string) {
	p.Name = &name
}

func (p *Pet) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *Pet) SetTag(tag string) {
	p.Tag = &tag
}

func (p *Pet) GetTag() *string {
	if p == nil {
		return nil
	}
	return p.Tag
}
