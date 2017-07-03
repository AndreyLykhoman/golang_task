package article

import "strings"

type TopicsList struct {
	Topics [] Topic
}

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