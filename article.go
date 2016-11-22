package main

import (
  "github.com/gin-gonic/gin"
  "time"
  "strconv"
)

type Article struct {
  Id int64 `db:"article_id"`
  Created int64
  Title string
  Content string
}

func createArticle(title, body string) Article {
  article := Article{
    Created:    time.Now().UnixNano(),
    Title:      title,
    Content:    body,
  }

  err := dbmap.Insert(&article)
  checkErr(err, "Insert failed")
  return article
}

func getArticle(article_id int) Article {
  article := Article{}
  err := dbmap.SelectOne(&article, "select * from articles where article_id=?", article_id)
  checkErr(err, "SelectOne failed")
  return article
}

func ArticlesList(c *gin.Context) {
  var articles []Article
  _, err := dbmap.Select(&articles, "select * from articles order by article_id")
  checkErr(err, "Select failed")
  content := gin.H{}
  for k, v := range articles {
    content[strconv.Itoa(k)] = v
  }
  c.JSON(200, content)
}

func ArticlesDetail(c *gin.Context) {
  article_id := c.Params.ByName("id")
  a_id, _ := strconv.Atoi(article_id)
  article := getArticle(a_id)
  content := gin.H{"title": article.Title, "content": article.Content}
  c.JSON(200, content)
}

func ArticlePost(c *gin.Context) {
  var json Article

  c.Bind(&json)
  article := createArticle(json.Title, json.Content)
  if article.Title == json.Title {
    content := gin.H{
      "result": "Success",
      "title": article.Title,
      "content": article.Content,
    }
    c.JSON(201, content)
  } else {
    c.JSON(500, gin.H{"result": "An error occured"})
  }
}
