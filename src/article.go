package article

import(
	"time"
	_ "text/template" //remove blank identifier to remove unused compiler error
	"io/ioutil"
	_ "strings" //remove blank identifier to remove unused compiler error
	"net/url"
	"strconv"
	"encoding/json"
)


type Article struct{
	Title string
	Author string
	Date time.Time
	Tags []string
	Content []byte //content should be template style: see documentation details on golang site
	Views int //this information might have to handled by a database
	UniqueViews int //this information might need to be handled by a database
}


func SaveJSONArticle(a Article) {
		b, err := json.Marshal(a)
		if err != nil {
			log.Fatal(err)
			return
		}
		filePath = '../static/articles' + a.Title
		// 0644 means overwrite
		ioutil.writeFile(filePath, b, 0644)
}
/*
Creates a pointer to an article. Article format must have format above.
Parsing isn't set up, but it could be done using this format:
Title:
Author:
Date:
	date
tags:
	tag1, tag2, tag3
Content:
	Content
Detail to change as no articles are written yet.
*/
func LoadJSONArticle(articleId int) *Article{
	b, err := 

}

/*
Creates a URL version of an article's title. Signatures use date and title for uniqueness in URL.
The URL is not totally valid, but the handler will enable its use.
*/
func (a *Article) parseTitle() string{
	date := "/" + strconv.Itoa(a.Date.Year()) + "/" + strconv.Itoa(int(a.Date.Month())) + "/" + strconv.Itoa(a.Date.Day())
	return url.QueryEscape(date + "/" + a.Title + ".html")
}