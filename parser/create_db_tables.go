package main

import (
	"parser/entity/db"
)

func CreateDatabasesTables()  {
	mysq := db.Open()
	defer mysq.Close()

	siteTable := GetSiteTableFields()
	mysq.CreateTable("site", siteTable, "", "id")

	topicTable := GetTopicTableFields()
	mysq.CreateTable("topic", topicTable, "","id")

	tagTable := GetTagTableFields()
	mysq.CreateTable("tag", tagTable, "", "id")

	articleTable := GetArticleTableFields()
	mysq.CreateTable("article", articleTable, "","id")

	topicToTagTable := GetTopicToTagTableFields()
	mysq.CreateTable("topic_to_tag", topicToTagTable, "KEY topic_id (topic_id),KEY tag_id (tag_id)")


	topicToArticleTable := GetTopicToArticleTableFields()
	mysq.CreateTable("topic_to_article", topicToArticleTable, "KEY topic_id (topic_id),KEY article_id (article_id)")

	/*mysq.Query("ALTER TABLE topic ADD CONSTRAINT fk_topic_to_tag_topic FOREIGN KEY (id) REFERENCES  topic_to_tag (topic_id) ON DELETE CASCADE ON UPDATE CASCADE;")
	mysq.Query("ALTER TABLE tag ADD CONSTRAINT fk_topic_to_tag_tag FOREIGN KEY (id) REFERENCES topic_to_tag (tag_id) ON DELETE CASCADE ON UPDATE CASCADE;")
	mysq.Query("ALTER TABLE topic ADD CONSTRAINT fk_topic_to_article_topic FOREIGN KEY (id) REFERENCES topic_to_article (topic_id) ON DELETE CASCADE ON UPDATE CASCADE;")
	mysq.Query("ALTER TABLE article ADD CONSTRAINT fk_topic_to_article_article FOREIGN KEY (id) REFERENCES topic_to_article (article_id) ON DELETE CASCADE ON UPDATE CASCADE;")

*/

}

func GetSiteTableFields() []db.TabelsFiels  {
	return []db.TabelsFiels{
		db.TabelsFiels{
			"id",
			"INT",
			false,
			true,
			"",
		},
		db.TabelsFiels{
			"name",
			"VARCHAR(100)",
			false,
			false,
			"",
		},
		db.TabelsFiels{
			"url",
			"VARCHAR(200)",
			false,
			false,
			"",
		},
		db.TabelsFiels{
			"link_selector",
			"VARCHAR(200)",
			false,
			false,
			"",
		},
	}
}

func GetTopicTableFields() []db.TabelsFiels  {
	return []db.TabelsFiels{
		db.TabelsFiels{
			"id",
			"INT",
			false,
			true,
			"",
		},
		db.TabelsFiels{
			"name",
			"VARCHAR(100)",
			false,
			false,
			"",
		},

	}
}

func GetTagTableFields() []db.TabelsFiels  {
	return []db.TabelsFiels{
		db.TabelsFiels{
			"id",
			"INT",
			false,
			true,
			"",
		},
		db.TabelsFiels{
			"name",
			"VARCHAR(100)",
			false,
			false,
			"",
		},

	}
}

func GetArticleTableFields() []db.TabelsFiels  {
	return []db.TabelsFiels{
		db.TabelsFiels{
			"id",
			"INT",
			false,
			true,
			"",
		},
		db.TabelsFiels{
			"site_id",
			"INT",
			false,
			false,
			"",
		},
		db.TabelsFiels{
			"title",
			"VARCHAR(300)",
			false,
			false,
			"",
		},
		db.TabelsFiels{
			"text",
			"TEXT",
			false,
			false,
			"",
		},
		db.TabelsFiels{
			"date",
			"DATE",
			true,
			false,
			"DEFAULT CURRENT_TIMESTAMP",
		},
	}
}

func GetTopicToTagTableFields() []db.TabelsFiels  {
	return []db.TabelsFiels{
		db.TabelsFiels{
			"topic_id",
			"INT",
			false,
			false,
			"",
		},
		db.TabelsFiels{
			"tag_id",
			"INT",
			false,
			false,
			"",
		},

	}
}

func GetTopicToArticleTableFields() []db.TabelsFiels  {
	return []db.TabelsFiels{
		db.TabelsFiels{
			"topic_id",
			"INT",
			false,
			false,
			"",
		},
		db.TabelsFiels{
			"article_id",
			"INT",
			false,
			false,
			"",
		},

	}
}


func main() {
	CreateDatabasesTables()
	mysq := db.Open()
	defer mysq.Close()
	mysq.SetTopicWithTags("Go" , []string{
		"go",
		"golang",
		"javascript",
		"программирование",
		"benchmark",
		"ооп",
		"docker",
		"html",
		"rust",
	} )
}