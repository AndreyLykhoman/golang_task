package main

import (
	"parser/entity/db"
	"parser/entity/site"
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

	mysq.Query("ALTER TABLE topic_to_tag ADD CONSTRAINT fk_topic_to_tag_topic FOREIGN KEY (topic_id) REFERENCES  topic (id) ON DELETE CASCADE ON UPDATE CASCADE;")
	mysq.Query("ALTER TABLE topic_to_tag ADD CONSTRAINT fk_topic_to_tag_tag FOREIGN KEY (tag_id) REFERENCES tag (id) ON DELETE CASCADE ON UPDATE CASCADE;")
	mysq.Query("ALTER TABLE topic_to_article ADD CONSTRAINT fk_topic_to_article_topic FOREIGN KEY (topic_id) REFERENCES topic (id) ON DELETE CASCADE ON UPDATE CASCADE;")
	mysq.Query("ALTER TABLE topic_to_article ADD CONSTRAINT fk_topic_to_article_article FOREIGN KEY (article_id) REFERENCES article (id) ON DELETE CASCADE ON UPDATE CASCADE;")


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
	//CreateDatabasesTables()
	mysq := db.Open()
	defer mysq.Close()
	//mysq.SetTopicWithTags("Go" , []string{
	//	"go",
	//	"golang",
	//	"javascript",
	//	"программирование",
	//	"benchmark",
	//	"ооп",
	//	"docker",
	//	"html",
	//	"rust",
	//} )
	remoteSite := []site.RemoteSite{
		site.RemoteSite{"https://habrahabr.ru/", ".post__title_link"},
		site.RemoteSite{"https://geektimes.ru/", ".post__title_link"},
	}
	mysq.SetRemoteSiteSlice(remoteSite)

}