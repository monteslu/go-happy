package main

import (
  "github.com/gin-gonic/gin"
  "database/sql"
  "github.com/coopernurse/gorp"
  _ "github.com/mattn/go-sqlite3"
  "log"
)

var dbmap = initDb()

func main(){

  defer dbmap.Db.Close()

  router := gin.Default()
  router.GET("/articles", ArticlesList)
  router.POST("/articles", ArticlePost)
  router.GET("/articles/:id", ArticlesDetail)

  router.POST("/user", UserPost)

  router.Run(":8000")
}



func initDb() *gorp.DbMap {
  db, err := sql.Open("sqlite3", "db.sqlite3")
  checkErr(err, "sql.Open failed")

  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

  dbmap.AddTableWithName(Article{}, "articles").SetKeys(true, "Id")
  dbmap.AddTableWithName(User{}, "user").SetKeys(true, "Id")

  err = dbmap.CreateTablesIfNotExists()
  checkErr(err, "Create tables failed")

  return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
    log.Fatalln(msg, err)
  }
}
