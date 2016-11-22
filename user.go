package main

import (
  "github.com/gin-gonic/gin"
  "time"
)

type User struct {
  Id int64 `db:"user_id"`
  Created int64
  Name string
  Username string
}

func createUser(name, username string) User {
  user := User{
    Created:    time.Now().UnixNano(),
    Name: name,
    Username: username,
  }

  err := dbmap.Insert(&user)
  checkErr(err, "Insert failed")
  return user
}

func getUser(user_id int) User {
  user := User{}
  err := dbmap.SelectOne(&user, "select* from user where user_id=?", user_id)
  checkErr(err, "SelectOne failed")
  return user
}

func UserPost(c *gin.Context) {
  var json User

  c.Bind(&json)
  user := createUser(json.Name, json.Username)
  if user.Username == json.Username {
    content := gin.H{
      "result": "Success",
      "name": user.Name,
      "username": user.Username,
    }
    c.JSON(201, content)
  } else {
    c.JSON(500, gin.H{"result": "An error occured"})
  }
}
