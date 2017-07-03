package main
//TODO: rewrite that file to build site's article list parser controller
//ToDo: rebase to main folder list_post

//TODO: rewite application code to use chanel and go routine
import (

	"parser/entity/parser"
	"parser/entity/site"
	"parser/entity/db"
	"sync"
	"parser/entity/articles"
)


func postScrape(removeSite site.RemoteSite) {
	var wg sync.WaitGroup
	articlesList := removeSite.GetListArticles()
	articles := articlesList.GetPost()

	wg.Add(len(articles))
	for index , articleItem := range articles{
		go func(article article.Article , ind int) {
			defer wg.Done()
			mysql := db.Open()
			defer mysql.Close()
			parser.WorkWithArticle(&article, mysql)
		}(articleItem, index)
	}

	wg.Wait()

}
func main() {


	mysql := db.Open()
	defer mysql.Close()
	var siteWG sync.WaitGroup
	var remoteSites []site.RemoteSite = mysql.GetRemoteSiteSlice()
	siteWG.Add(len(remoteSites))
	for _,remoteSite := range remoteSites {
		go func(site site.RemoteSite){
			defer siteWG.Done()
			postScrape(site)
		}(remoteSite)
	}
	siteWG.Wait()
}