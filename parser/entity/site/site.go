package site

import (
	"parser/entity/articles"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)
type RemoteSite struct {
	URL string
	LinksSelector string
}

func (r *RemoteSite) GetListArticles() article.ArticlesList {
	arArticle := []article.Article{}
	articlesList := article.ArticlesList{arArticle,}
	doc, err := goquery.NewDocument(r.URL)
	if err != nil {
		log.Fatal(err)
	}
	body := doc.Find("body")
	selScript := body.Find("script")
	newBody := body.NotSelection(selScript)


	newBody.Find(r.LinksSelector).Each(func(index int, item *goquery.Selection) {
		link, _ := item.Attr("href")
		newArticle  := article.Article {}
		newArticle.SetTitle(strings.TrimSpace(item.Text()))
		newArticle.SetURL(link)
		articlesList.AddArticle(newArticle)
	})
	return articlesList
}


func (r *RemoteSite) GetUrl() string{
	return  r.URL
}

func (r *RemoteSite) SetUrl(s string)  {
	r.URL = s
}

func (r *RemoteSite) GetLinksSelector() string{
	return  r.LinksSelector
}

func (r *RemoteSite) SetLinksSelector(s string)  {
	r.LinksSelector = s
}

