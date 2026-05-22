package entity

type Category struct {
	id   string
	name string
	icon string
}

func NewCategory(id, name, icon string) *Category {
	return &Category{id: id, name: name, icon: icon}
}

func (c *Category) ID() string   { return c.id }
func (c *Category) Name() string { return c.name }
func (c *Category) Icon() string { return c.icon }
