package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"prometheusdemo/dao"
	"prometheusdemo/pkg/ginx"
	"strconv"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {

	return &UserHandler{
		db: db,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")

	ug.POST("/signup", u.SignUp)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	var req SignUpReq

	if err := ctx.Bind(&req); err != nil {
		return
	}

	userDao := dao.User{
		Email:    sql.NullString{Valid: len(req.Email) > 0, String: req.Email},
		Password: req.Password,
	}

	err := u.db.WithContext(ctx).Create(&userDao).Error
	if err != nil {
		ginx.Vector.WithLabelValues(strconv.Itoa(http.StatusBadRequest)).Inc()
		ctx.String(http.StatusBadRequest, "create user err %v", err)
		return
	}
	ginx.Vector.WithLabelValues(strconv.Itoa(http.StatusOK)).Inc()

	ctx.String(http.StatusOK, "注册成功")
}
