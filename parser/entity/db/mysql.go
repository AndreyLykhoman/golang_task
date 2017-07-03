package db
//TODO: Write struct and interfase to work with mysql
import(
	"parser/entity/articles"
	"parser/entity/site"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const(
	DRIVER_NAME    = "mysql"
	LOGIN          = "golang"
	PASSWORD       = "andrey31"
	DATABASE       = "articles"
	NULL           = "NULL"
	NOT_NULL       = "NOT NULL"
	AUTO_INCREMENT = "AUTO_INCREMENT"
)

type TabelsFiels struct {
	Name string
	Type string
	IsNull bool
	AUTO_INCREMENT bool
	Aditional__params string;
}

type Mysql struct {
	dataBase *sql.DB
}


func Open() (db Mysql) {
	connect, err :=  sql.Open(DRIVER_NAME, fmt.Sprintf("%v:%v@/%v", LOGIN, PASSWORD, DATABASE) )
	if err != nil {
	panic(err.Error())
	}
	err = connect.Ping()
	if err != nil {
		panic(err.Error())
	}
	db = Mysql{connect}
	return
}

func (db *Mysql) Close(){
	db.dataBase.Close();
}

func (db *Mysql) Query (query string){
	_, err := db.dataBase.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func (db *Mysql) CreateTable (name string, fields []TabelsFiels, key string ,primary_key... string ){
	var biteFields []byte;
	for _, field := range fields {
		var strNull string
		var strAuto string

		if field.IsNull {
			strNull = NULL
		} else {
			strNull = NOT_NULL
		}
		if field.AUTO_INCREMENT {
			strAuto = AUTO_INCREMENT
		}
		srtField := fmt.Sprintf("%v %v %v %v, ", field.Name, field.Type, strNull, strAuto)
		biteFields = append(biteFields, []byte(srtField)...)
	}
	var bitePrimary []byte;
	if primary_key != nil {
		for  index, key := range primary_key {
			if index > 0{
				bitePrimary = append(bitePrimary, []byte(", ")...)
			}
			bitePrimary = append(bitePrimary, []byte(key)...)
		}
		strPrimary := fmt.Sprintf("PRIMARY KEY ( %v )", string(bitePrimary))
		biteFields = append(biteFields, []byte(strPrimary)...)
	}

	if key != ""{
		biteFields = append(biteFields, []byte(key)...)
	}
	strQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (%v)", name, string(biteFields))
	_, err := db.dataBase.Query(strQuery)
	if err != nil {
		panic(err.Error())
	}
}

func (db Mysql) GetTopics() ( tl article.TopicsList) {
	//TODO: rewrite that code to use db

	_, err := db.dataBase.Query("SELECT *  FROM site")
	if err != nil {
		panic(err)
	}
	strTag := []string{
		"php",
		"pyhton",
		"java",
		"администрирование",
	}
	for id, name := range strTag {
		tl.AddNewTopics(article.Topic{id,name, []string{"php", "laravel"}})
	}
	return
}

func (db Mysql) GetRemoteSiteSlice() (sites [] site.RemoteSite) {

	row, err := db.dataBase.Query("SELECT url, link_selector FROM site")
	if err != nil {
		panic(err)
	}
	for  row.Next() {
		var url string
		var link_selector string
		err = row.Scan(&url, &link_selector)
		sites = append(sites, site.RemoteSite{
			url,
			link_selector,
		})

	}
	return
}

func (db *Mysql) SetTopicWithTags(topic string, tags []string)  {

	stmt, err := db.dataBase.Prepare("INSERT INTO `articles`.`topic` (`name`) VALUES (?);")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(topic)
	if err != nil {
		panic(err)
	}

	topic_id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	var tagIds [] int;
	for _, tag := range tags {
		stmt, err = db.dataBase.Prepare("INSERT INTO `articles`.`tag` (`name`) VALUES (?);")
		if err != nil {
			panic(err)
		}
		res, err := stmt.Exec(tag)
		if err != nil {
			panic(err)
		}

		tag_id, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		tagIds = append(tagIds, int(tag_id))

	}

	//var biteInsert[] byte;
	for _, tagId := range tagIds {
		//biteInsert = append(biteInsert, []byte(fmt.Sprintf(
		//	"INSERT INTO `articles`.`topic_to_tag` (`topic_id`, `tag_id`) VALUES ( %v, %v );",
		//	topic_id,
		//	tagId) )...)
		_, err = db.dataBase.Query(fmt.Sprintf(
			"INSERT INTO `articles`.`topic_to_tag` (`topic_id`, `tag_id`) VALUES ( %v, %v );",
			topic_id,
			tagId) )
		if err != nil {
			panic(err)
		}
	}
	//fmt.Println(string(biteInsert))
	//_, err = db.dataBase.Query(string(biteInsert))
	//if err != nil {
	//	panic(err)
	//}

}