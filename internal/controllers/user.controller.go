package controllers

// import (
// 	services "gin-skeleton/internal/services"

// 	"github.com/gin-gonic/gin"
// )

// type UserController interface {
// 	GetAllUserData(c *gin.Context)
// 	AddUserData(c *gin.Context)
// 	GetUserById(c *gin.Context)
// 	UpdateUserData(c *gin.Context)
// 	DeleteUser(c *gin.Context)
// }

// type UserControllerImpl struct {
// 	svc services.UserService
// }

// func UserControllerInit(userService service.UserService) *UserControllerImpl {
// 	return &UserControllerImpl{
// 		svc: userService,
// 	}
// }

// func (u UserControllerImpl) GetAllUserData(c *gin.Context) {
// 	u.svc.GetAllUser(c)
// }

// func (u UserControllerImpl) AddUserData(c *gin.Context) {
// 	u.svc.AddUserData(c)
// }

// func (u UserControllerImpl) GetUserById(c *gin.Context) {
// 	u.svc.GetUserById(c)
// }

// func (u UserControllerImpl) UpdateUserData(c *gin.Context) {
// 	u.svc.UpdateUserData(c)
// }

// func (u UserControllerImpl) DeleteUser(c *gin.Context) {
// 	u.svc.DeleteUser(c)
// }
