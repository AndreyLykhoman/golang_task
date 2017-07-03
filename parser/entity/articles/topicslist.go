package article

import "strings"

type TopicsList struct {
	Topics [] Topic
}
// TODO: переписать на реализацию с тегами используя map(string) int как набор тегов и написать структуру в которую поместить текст подсчитанный на вхождение слов. после этого найти максимально совпадение тегов категорий к тексту
func (tl TopicsList) CountTopicsInString(source string)(resTagMap map[string] int)  {
	resTagMap = make(map[string] int)
	loverString := strings.ToLower(source)
	for _, topic := range tl.Topics {
		count := strings.Count(loverString, topic.Name)
		resTagMap[topic.Name] = count
	}
	return
}

func (tl *TopicsList) AddNewTopics( t... Topic) {
	tl.Topics = append(tl.Topics, t...)
}