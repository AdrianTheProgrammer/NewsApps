package routes

import (
	"newsapps/configs"
	articles "newsapps/internal/features/articles"
	comments "newsapps/internal/features/comments"
	users "newsapps/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uh users.UHandlers, ah articles.AHandlers, ch comments.CHandlers) {
	e.POST("/login", uh.Login)
	e.POST("/register", uh.CreateUser)
	e.GET("/articles", ah.ShowAllArticles())
	e.GET("/articles/:id", ah.ReadArticle())
	UsersRoute(e, uh)
	ArticlesRoute(e, ah)
	CommentsRoute(e, ch)
}

func UsersRoute(e *echo.Echo, uh users.UHandlers) {
	u := e.Group("/users")
	u.Use(JWTConfig())
	u.GET("/settings", uh.ReadUser)
	u.PUT("/edit", uh.UpdateUser)
	u.DELETE("/deactivate", uh.DeleteUser)
}

func ArticlesRoute(e *echo.Echo, ah articles.AHandlers) {
	a := e.Group("/articles")
	a.Use(JWTConfig())
	a.POST("/post", ah.CreateArticle())
	a.PUT("/:id/edit", ah.UpdateArticle())
	a.DELETE("/:id/delete", ah.DeleteArticle())

}

func CommentsRoute(e *echo.Echo, ch comments.CHandlers) {
	c := e.Group("/comments")
	c.Use(JWTConfig())
	c.POST("/post", ch.CreateComment())
	c.PUT("/:cid/edit", ch.UpdateComment())
	c.DELETE("/:cid/delete", ch.DeleteComment())

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

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(configs.ImportPasskey()),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
