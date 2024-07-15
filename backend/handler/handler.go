package handler

import (
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

type BizError = int
type h = map[string]any

const (
	SUCCESS      BizError = 0
	Fail         BizError = 1
	ParamError   BizError = 2
	TokenInvalid BizError = 3
	TokenMissing BizError = 4
)

var errorMessageMap = map[BizError]string{
	SUCCESS:      "成功",
	Fail:         "失败",
	ParamError:   "参数错误",
	TokenInvalid: "Token无效",
	TokenMissing: "Token缺失",
}

type BaseHandler struct {
	injector do.Injector
	cfg      vo.SysConfig
	db       *gorm.DB
	log      zerolog.Logger
}

func NewBaseHandler(injector do.Injector) (BaseHandler, error) {
	return BaseHandler{
		injector: injector,
		cfg:      do.MustInvoke[vo.SysConfig](injector),
		db:       do.MustInvoke[*gorm.DB](injector),
		log:      do.MustInvoke[zerolog.Logger](injector),
	}, nil
}

type Resp[T any] struct {
	Code    BizError `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Data    T        `json:"data,omitempty"`
}

func SuccessResp[T any](c echo.Context, data T) error {
	resp := Resp[T]{
		Code:    SUCCESS,
		Message: "",
		Data:    data,
	}
	return c.JSON(200, resp)
}

func FailResp(c echo.Context, code BizError) error {
	resp := Resp[any]{
		Code:    code,
		Message: errorMessageMap[code],
		Data:    nil,
	}
	return c.JSON(200, resp)
}

func FailRespWithMsg(c echo.Context, code BizError, message string) error {
	resp := Resp[any]{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	return c.JSON(200, resp)
}
