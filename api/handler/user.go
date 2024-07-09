package handler

import (
	"root/api/token"

	pb "root/genprotos"

	"github.com/gin-gonic/gin"
)

// LoginUser 		handles the Login  User
// @Summary 		/LoginUser
// @Description 	LoginUser page
// @Tags 			User
// @Accept  		json
// @Security  		BearerAuth
// @Produce  		json
// @Param   		Create  body  pb.LoginUserRequest   true     "Create"
// @Success 		200 {object}  string  "LoginUser Successful"
// @Failure 		401 {string}  string   "Error while LoginUserd"
// @Router 			/user/login [post]
func (h *Handler) LoginUser(ctx *gin.Context) {
	user := &pb.LoginUserRequest{}
	err := ctx.ShouldBindJSON(user)
	if err != nil {
		panic(err)
	}
	res, err := h.User.LoginUser(ctx, user)
	if err != nil {
		panic(err)
	}
	t := token.GenereteJWTToken(res)
	ctx.JSON(200, t)
}
