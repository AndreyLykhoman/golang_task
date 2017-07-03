package article

type ArticlesList struct{
	ArArticle [] Article
}

func (list *ArticlesList) AddArticle(p Article) *ArticlesList {
	list.ArArticle = append(list.ArArticle, p)
	return list
}

func (list *ArticlesList) GetPost() []Article {
	return list.ArArticle
}


