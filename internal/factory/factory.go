package factory

import (
	"fmt"
	"newsapps/configs"
	a_hnd "newsapps/internal/features/articles/handlers"
	a_rep "newsapps/internal/features/articles/repositories"
	a_srv "newsapps/internal/features/articles/services"
	c_hnd "newsapps/internal/features/comments/handlers"
	c_rep "newsapps/internal/features/comments/repositories"
	c_srv "newsapps/internal/features/comments/services"
	u_hnd "newsapps/internal/features/users/handlers"
	u_rep "newsapps/internal/features/users/repositories"
	u_srv "newsapps/internal/features/users/services"
	"newsapps/internal/routes"
	"newsapps/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo) {
	db := configs.ConnectDB()

	MigrateDB(db)

	pu := utils.NewPassUtil()
	tu := utils.NewTokenUtil()

	uq := u_rep.NewUsersQry(db)
	us := u_srv.NewUsersSrv(uq, pu, tu)
	uh := u_hnd.NewUsersHand(us, tu)

	aq := a_rep.NewArticlesQry(db)
	as := a_srv.NewArticlesSrv(aq)
	ah := a_hnd.NewArticlesHand(as)

	cq := c_rep.NewCommentsQry(db)
	cs := c_srv.NewCommentsSrv(cq)
	ch := c_hnd.NewCommentsHand(cs)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routes.InitRoute(e, uh, ah, ch)
}

func MigrateDB(db *gorm.DB) {
	var menu int
	fmt.Print("Masukkan '1' untuk Migrasi Database: ")
	fmt.Scanln(&menu)

	if menu == 1 {
		err := db.AutoMigrate(&u_rep.Users{}, &c_rep.Comments{}, &a_rep.Articles{})
		if err != nil {
			fmt.Println(err)
		}
	}
}
