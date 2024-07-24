package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"go-boilerplate/internal/constant"
	"go-boilerplate/internal/payload"
	"go-boilerplate/internal/util"
)

type IUserHandler interface {
	GetAll(c *gin.Context)
	GetDetail(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
type user struct {
	Option
}

func NewUserHandler(opt Option) IUserHandler {
	return &user{
		Option: opt,
	}
}

func (h *user) GetAll(c *gin.Context) {
	result, err := h.Service.UserService.GetList(c)
	if err != nil {
		log.Errorf("error get user list %v", err) // log error
		util.ErrInternalResponse(c, err)

		return
	}

	util.GeneralSuccessResponse(c, "success get user data", result)

	return
}

func (h *user) GetDetail(c *gin.Context) {
	id := c.Param("id")

	uintID, err := cast.ToUint64E(id)
	if err != nil {
		log.Warnf("error parsing user id from param %v", err)

		util.ErrBindResponse(c, err)

		return
	}

	result, err := h.Service.UserService.GetByID(c, uintID)
	if err != nil {
		if errors.Is(err, constant.ErrUserNotFound) {
			log.Warnf("user id not found %v", uintID)
			util.ErrBadRequestResponse(c, err.Error())

			return
		}

		util.ErrInternalResponse(c, err)

		return
	}

	util.GeneralSuccessResponse(c, "success get user detail", result)

	return
}

func (h *user) Create(c *gin.Context) {
	req := payload.CreateUserRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	err = h.Service.UserService.Create(c, req)
	if err != nil {
		if errors.Is(constant.ErrEmailAlreadyRegistered, err) {
			log.Warnf("%s already registered", req.Email)

			util.ErrBadRequestResponse(c, err.Error())
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success create user", nil)

	return
}

func (h *user) Update(c *gin.Context) {
	id := c.Param("id")

	uintID, err := cast.ToUint64E(id)
	if err != nil {
		log.Warnf("error parsing user id from param %v", err)

		util.ErrBindResponse(c, err)

		return
	}

	req := payload.UpdateUserRequest{
		ID: uintID,
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		util.ErrBindResponse(c, err)
		return
	}

	err = h.Service.UserService.Update(c, req)
	if err != nil {
		if errors.Is(constant.ErrUserNotFound, err) {
			log.Warnf("user id not found %v", req.ID)

			util.ErrBadRequestResponse(c, err.Error())
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success update user", nil)

	return
}

func (h *user) Delete(c *gin.Context) {
	id := c.Param("id")

	uintID, err := cast.ToUint64E(id)
	if err != nil {
		log.Warnf("error parsing user id from param %v", err)

		util.ErrBindResponse(c, err)

		return
	}

	err = h.Service.UserService.Delete(c, uintID)
	if err != nil {
		if errors.Is(constant.ErrUserNotFound, err) {
			log.Warnf("user id not found %v", uintID)

			util.ErrBadRequestResponse(c, err.Error())
			return
		}

		util.ErrInternalResponse(c, err)
		return
	}

	util.GeneralSuccessResponse(c, "success delete user", nil)

	return
}
