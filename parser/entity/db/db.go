package db

import "parser/entity/articles"

type DB interface {
	GetTopics() article.TopicsList
}
