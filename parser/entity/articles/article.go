package article

type Article struct{
	Title string
	Url   string
	Text  string
	Topic Topic
}

func (a *Article) GetTitle() string {
	return a.Title
}

func (a *Article) SetTitle(s string)  {
	a.Title = s
}

func (a *Article) GetURL() string {
	return a.Url
}

func (a *Article) SetURL(s string)  {
	a.Url = s
}

func (a *Article) GetText() string {
	return a.Text
}

func (a *Article) SetText(s string)  {
	a.Text = s
}

func (a *Article) GetTopic() Topic {
	return a.Topic
}

func (a *Article) SetTopic(t Topic )  {
	a.Topic = t
}

