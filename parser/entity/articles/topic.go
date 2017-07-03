package article

type Topic struct{
	Id   int
	Name string
	Tags [] string
}

func (t *Topic) GetID() int{
	return t.Id
}

func (t *Topic) SetId(id int) {
	t.Id = id
}

func (t *Topic) GetName() string  {
	return t.Name
}

func (t *Topic) SetName( s string)  {
	t.Name = s
}

func (t *Topic) GetTags() [] string {
	return t.Tags
}