package routes

import (
	"newsapps/configs"
	articles "newsapps/internal/features/articles"
	comments "newsapps/internal/features/comments"
	users "newsapps/internal/features/users"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.UHandlers, ah articles.Handlers, ch comments.Handlers) {
	e.POST("/login", uh.Login)
	e.POST("/register", uh.CreateUser)
	e.GET("/articles", ah.ShowAllArticles())
	e.GET("/articles/:id", ah.ReadArticle())
	e.GET("/articles/:id/comment", ah.ReadArticle())

	UsersRoute(e, uh)
	ArticlesRoute(e, ah, ch)
	// CommentsRoute(e, ch)
}

func UsersRoute(e *echo.Echo, uh users.UHandlers) {
	u := e.Group("/users")
	u.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(configs.ImportPasskey()),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	u.GET("/settings/:id", uh.ReadUser)
	u.PUT("/edit/:id", uh.UpdateUser)
	u.POST("/deactivate/:id", uh.DeleteUser)
}

func ArticlesRoute(e *echo.Echo, ah articles.Handlers, ch comments.Handlers) {
	a := e.Group("/articles")
	a.Use(echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(configs.ImportPasskey()),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	))
	a.POST("/post", ah.CreateArticle())
	a.PUT("/:id/edit", ah.UpdateArticle())
	a.DELETE("/:id/delete", ah.DeleteArticle())
	a.GET("/:id/comment", ch.CreateComment())
	a.POST("/:id/comment/post", ch.CreateComment())
	a.PUT("/:id/comment/:cid/update", ch.UpdateComment())
	a.DELETE("/:id/comment/:cid/delete", ch.DeleteComment())
}

// func CommentsRoute(e *echo.Echo, ch comments.CHandlers) {
// 	c := e.Group("/comments")
// 	c.Use(echojwt.WithConfig(
// 		echojwt.Config{
// 			SigningKey:    []byte(configs.ImportPasskey()),
// 			SigningMethod: jwt.SigningMethodHS256.Name,
// 		},
// 	))
// 	c.POST("/post", ch.CreateComment)
// 	// Read Comment tidak ada karena ide saya readnya sekalian sama read article (Read 1 Article, ke-read semua commentnya)
// 	c.PUT("/edit/:id", ch.UpdateComment)
// 	c.DELETE("/delete/:id", ch.DeleteComment)
// }
