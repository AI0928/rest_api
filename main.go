package main

import (
	"net/http"
	"rest_api/database" //自作だから自分のパッケージ
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type User struct {
	Id        int            `json:"id" param:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Post struct {
	Id        int            `json:"id" param:"id"`
	Content   string         `json:"content"`
	User_id   int            `json:"user_id" param:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Like struct {
	Id        int            `json:"id"`
	Post_id   int            `json:"post_id" param:"post_id"`
	User_id   int            `json:"user_id" param:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Comment struct {
	Id        int            `json:"id"`
	Comment   string         `json:"comment"`
	Post_id   int            `json:"post_id" param:"post_id"`
	User_id   int            `json:"user_id" param:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Relationship struct {
	Id             int            `json:"id"`
	Follow_user_id int            `json:"follow_user_id" param:"follow_user_id"`
	User_id        int            `json:"user_id" param:"user_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

func getUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Take(&user)
	return c.JSON(http.StatusOK, user)
}

func getUsers(c echo.Context) error {
	users := []User{}
	database.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func updateUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	database.DB.Save(&user)
	return c.JSON(http.StatusOK, user)
}

func createUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	database.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	database.DB.Delete(&User{}, id)
	return c.NoContent(http.StatusNoContent)
}

func getPosts(c echo.Context) error {
	posts := []Post{}
	database.DB.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}

func getPost(c echo.Context) error {
	post := Post{}
	if err := c.Bind(&post); err != nil {
		return err
	}
	database.DB.Take(&post)
	return c.JSON(http.StatusOK, post)
}

func updatePost(c echo.Context) error {
	post := Post{}
	if err := c.Bind(&post); err != nil {
		return err
	}

	database.DB.Save(&post)
	return c.JSON(http.StatusOK, post)
}

func createPost(c echo.Context) error {
	post := Post{}
	if err := c.Bind(&post); err != nil {
		return err
	}

	database.DB.Create(&post)
	return c.JSON(http.StatusCreated, post)
}

func deletePost(c echo.Context) error {
	id := c.Param("id")
	database.DB.Delete(&Post{}, id)
	return c.NoContent(http.StatusNoContent)
}

func getLike(c echo.Context) error {
	likes := []Like{}
	if err := c.Bind(&likes); err != nil {
		return err
	}
	database.DB.Where("user_id = ?", c.Param("user_id")).Find(&likes)
	return c.JSON(http.StatusOK, likes)
}

func getLikes(c echo.Context) error {
	likes := []Like{}
	database.DB.Find(&likes)
	return c.JSON(http.StatusOK, likes)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/users", getUsers)          //ユーザー一覧取得
	e.GET("/users/:id", getUser)       //任意(一人）のユーザー情報の取得
	e.PUT("/users/:id", updateUser)    //任意（一人）のユーザー情報の更新
	e.POST("/users", createUser)       //ユーザーの作成
	e.DELETE("/users/:id", deleteUser) //任意（一人）のユーザーの削除

	e.GET("/posts", getPosts)          //投稿の一覧取得
	e.GET("/posts/:id", getPost)       //任意の投稿の取得
	e.PUT("/posts/:id", updatePost)    //任意（一人）のユーザー情報の更新
	e.POST("/posts", createPost)       //ユーザーの作成
	e.DELETE("/posts/:id", deletePost) //任意（一人）のユーザーの削除

	e.GET("/likes", getLikes) //任意のユーザーのいいねを取得する

	e.Logger.Fatal(e.Start(":3000")) // コンテナ側の開放ポートと一緒にすること
}
