package parser

//TODO: check html5 tags and add to parser if that need
import (
	"parser/entity/articles"
	"parser/entity/db"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"fmt"
)

type HTML struct {
	Html string
}


func (html HTML) clearHtml() (HTML) {
	html.Html = strings.Replace(html.Html, "  ", "", -1)
	html.Html = strings.Replace(html.Html, "\n", "", -1)
	html.Html = strings.Replace(html.Html, "</div>", "\n</div>", -1)
	html.Html = strings.Replace(html.Html, "</article>", "\n</article>", -1)
	return html
}


func (html *HTML) SplitBy(sep string)(ls LinesList){
	temp := strings.Split(html.Html, sep)
	var lines []Line
	for _, strLine := range temp{
		lines = append(lines, Line{
			len(strLine),
			strLine,
		})
	}
	ls.Lines = lines
	return
}



func  SelectorRemove(d *goquery.Document, selectors [] string) *goquery.Document{
	for _, selector := range selectors  {
		d.Find(selector).Remove()
	}
	return d
}

func WorkWithArticle( article *article.Article,  db db.DB)  {
	doc, err := goquery.NewDocument(article.GetURL())
	if err != nil {
		log.Fatal(err)
	}
	doc = SelectorRemove(doc, []string{
		"script",
		"noscript",
		"style",
		"form",
		"br",
		"aside",
		"nav",
	})
	body := doc.Find("body")

	value, err := body.Html()
	if err != nil {
		panic(err)
	}
	var html HTML = HTML{
		value,
	}
	var clearedHtml HTML = html.clearHtml()
	var linesList LinesList = clearedHtml.SplitBy("\n")
	linesList = linesList.SortLinesBiggestFromLess()
	linesList = linesList.GetNBiggest(5)
	var linesValue string = linesList.getLinesValue()
	topicList := db.GetTopics()
	var mapTopicsCount = topicList.CountTopicsInString(linesValue)

	article.SetText(body.Text())
	fmt.Println(article.GetURL(), mapTopicsCount)
}


