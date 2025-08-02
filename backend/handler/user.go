package handler

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kingwrcy/moments/db"
	"github.com/kingwrcy/moments/vo"
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	base BaseHandler
}

type loginSuccessDTO struct {
	Token    string `json:"token,omitempty"`    // token
	Username string `json:"username,omitempty"` //用户名
	Id       int32  `json:"id,omitempty"`       //用户ID
}

type userListResp struct {
	List   []db.User `json:"list"`
	HasNext bool      `json:"hasNext"`
}

type updateUserReq struct {
	Id        int32  `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Slogan    string `json:"slogan"`
	AvatarUrl string `json:"avatarUrl"`
	CoverUrl  string `json:"coverUrl"`
	Password  string `json:"password,omitempty"`
}



func NewUserHandler(injector do.Injector) *UserHandler {
	return &UserHandler{do.MustInvoke[BaseHandler](injector)}
}

// Login godoc
//
//	@Tags		User
//	@Summary	用户登录
//	@Accept		json
//	@Produce	json
//	@Param		object	body		vo.LoginReq	true	"用户登录"
//	@Success	200		{object}	loginSuccessDTO
//	@Router		/api/user/login [post]
func (u UserHandler) Login(c echo.Context) error {
	var req vo.LoginReq
	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}

	var user db.User
	err = u.base.db.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return FailRespWithMsg(c, Fail, "用户不存在或密码不正确")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return FailRespWithMsg(c, Fail, "用户不存在或密码不正确")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"userId":   user.Id,
	})

	tokenString, err := token.SignedString([]byte(u.base.cfg.JwtKey))
	if err != nil {
		u.base.log.Error().Msgf("生成jwt token异常:%s", err)
		return FailRespWithMsg(c, Fail, "登录异常")
	}
	return SuccessResp(c, loginSuccessDTO{
		Token:    tokenString,
		Username: user.Username,
		Id:       user.Id,
	})
}

// Reg godoc
//
//	@Tags		User
//	@Summary	用户注册
//	@Accept		json
//	@Produce	json
//	@Param		object	body	vo.RegReq	true	"用户注册"
//	@Success	200
//	@Router		/api/user/reg [post]
func (u UserHandler) Reg(c echo.Context) error {
	var (
		req         vo.RegReq
		count       int64
		user        db.User
		now         = time.Now()
		sysConfig   db.SysConfig
		sysConfigVO vo.FullSysConfigVO
	)

	u.base.db.First(&sysConfig)
	_ = json.Unmarshal([]byte(sysConfig.Content), &sysConfigVO)

	if !sysConfigVO.EnableRegister {
		return FailRespWithMsg(c, Fail, "当前未开启注册用户")
	}

	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}

	if len(req.Username) < 3 {
		return FailRespWithMsg(c, Fail, "用户名最少3个字符")
	}
	if req.Password != req.RepeatPassword {
		return FailRespWithMsg(c, Fail, "两次密码不一致")
	}
	u.base.db.Table("User").Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return FailRespWithMsg(c, Fail, "用户名已存在")
	}
	user.Username = req.Username
	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		u.base.log.Error().Msgf("密码加密异常:%s", err)
		return FailRespWithMsg(c, Fail, "密码加密异常")
	}
	user.Password = string(pwd)
	user.CreatedAt = &now
	user.UpdatedAt = &now
	user.Nickname = req.Username
	user.AvatarUrl = "/avatar.webp"
	user.Slogan = "修道者，逆天而行，注定要一生孤独。"
	user.CoverUrl = "/cover.webp"
	if err := u.base.db.Save(&user).Error; err != nil {
		u.base.log.Error().Msgf("注册用户异常:%s", err)
		return FailRespWithMsg(c, Fail, "注册用户异常")
	}
	return SuccessResp(c, h{})
}

// ProfileForUser godoc
//
//	@Tags		User
//	@Summary	获取指定用户信息
//	@Accept		json
//	@Produce	json
//	@param		string	path		string	true	"用户名"
//	@Success	200		{object}	db.User
//	@Router		/api/user/profile/{username} [post]
func (u UserHandler) ProfileForUser(c echo.Context) error {
	username := c.Param("username")
	var user db.User
	u.base.db.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl", "email").Find(&user, "username = ?", username)
	return SuccessResp(c, user)
}

// Profile godoc
//
//	@Tags			User
//	@Summary		获取用户信息
//	@Description	当前如果已经登录了,获取当前用户信息,否则获取管理员的用户信息
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	db.User
//	@Router			/user/profile [post]
func (u UserHandler) Profile(c echo.Context) error {

	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser == nil {
		u.base.db.Select("username", "nickname", "slogan", "id", "avatarUrl", "coverUrl", "email").First(&currentUser)
	}

	return SuccessResp(c, currentUser)
}

// SaveProfile godoc
//
//	@Tags		User
//	@Summary	保存用户信息
//	@Accept		json
//	@Produce	json
//	@Param		object		body	vo.ProfileReq	true	"保存用户信息"
//	@Param		x-api-token	header	string			true	"登录TOKEN"
//	@Success	200
//	@Router		/api/user/saveProfile [post]
func (u UserHandler) SaveProfile(c echo.Context) error {
	var (
		req  vo.ProfileReq
		user db.User
	)
	err := c.Bind(&req)
	if err != nil {
		return FailResp(c, ParamError)
	}
	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser == nil {
		return FailResp(c, TokenMissing)
	}
	u.base.db.Find(&user, currentUser.Id)
	if req.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			return FailResp(c, Fail)
		}
		user.Password = string(password)
	}
	user.Nickname = req.Nickname
	user.AvatarUrl = req.AvatarUrl
	user.Slogan = req.Slogan
	user.CoverUrl = req.CoverUrl
	user.Email = req.Email

	if err := u.base.db.Save(&user).Error; err != nil {
		return FailResp(c, Fail)
	}
	return SuccessResp(c, h{})
}

// isAdmin 检查当前用户是否为管理员
func (u UserHandler) isAdmin(c echo.Context) bool {
	context := c.(CustomContext)
	currentUser := context.CurrentUser()
	if currentUser == nil {
		return false
	}

	return currentUser.Id == 1
}

// UserList godoc
//
//	@Tags		User
//	@Summary	管理员获取用户列表
//	@Description	管理员获取所有注册用户列表，支持分页和排序
//	@Accept		json
//	@Produce	json
//	@Param		object		body		object	false	"分页、排序和搜索参数"
//	@Param		object.page		body		int		false	"页码，默认为1"
//	@Param		object.size		body		int		false	"每页数量，默认为20"
//	@Param		object.sort		body		string	false	"排序方式，asc: 升序，desc: 降序，默认为desc"
//	@Param		object.keyword	body		string	false	"搜索关键词"
//	@Param		x-api-token	header		string	true	"登录TOKEN"
//	@Success	200		{object}	userListResp
//	@Router		/api/user/list [post]
func (u UserHandler) UserList(c echo.Context) error {
	if !u.isAdmin(c) {
		return FailRespWithMsg(c, Fail, "无权限访问")
	}

	type listReq struct {
		Page    int    `json:"page"`
		Size    int    `json:"size"`
		Sort    string `json:"sort"`
		Keyword string `json:"keyword"`
	}

	var req listReq
	if err := c.Bind(&req); err != nil {
		return FailResp(c, ParamError)
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	size := req.Size
	if size <= 0 {
		size = 20
	}
	sort := req.Sort
	if sort != "asc" && sort != "desc" {
		sort = "desc"
	}

	var users []db.User
	var total int64

	offset := (page - 1) * size

	// 构建查询条件
	query := u.base.db.Model(&db.User{})
	if req.Keyword != "" {
		query = query.Where("username LIKE ? OR nickname LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	query.Count(&total)
	query.Order("id " + sort).Limit(size).Offset(offset).Find(&users)

	hasNext := offset+len(users) < int(total)

	return SuccessResp(c, userListResp{
		List:   users,
		HasNext: hasNext,
	})
}

// GetUser godoc
//
//	@Tags		User
//	@Summary	管理员获取用户详情
//	@Description	管理员获取指定用户的详细信息
//	@Accept		json
//	@Produce	json
//	@Param		id			path		int		true	"用户ID"
//	@Param		x-api-token	header		string	true	"登录TOKEN"
//	@Success	200			{object}	db.User
//	@Router		/api/user/{id} [post]
func (u UserHandler) GetUser(c echo.Context) error {
	if !u.isAdmin(c) {
		return FailRespWithMsg(c, Fail, "无权限访问")
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return FailResp(c, ParamError)
	}

	var user db.User
	if err := u.base.db.First(&user, int32(id)).Error; err != nil {
		return FailRespWithMsg(c, Fail, "用户不存在")
	}

	return SuccessResp(c, user)
}

// UpdateUser godoc
//
//	@Tags		User
//	@Summary	管理员更新用户信息
//	@Description	管理员更新指定用户的基本信息
//	@Accept		json
//	@Produce	json
//	@Param		object		body		updateUserReq	true	"用户信息"
//	@Param		x-api-token	header		string				true	"登录TOKEN"
//	@Success	200
//	@Router		/api/user/update [post]
func (u UserHandler) UpdateUser(c echo.Context) error {
	if !u.isAdmin(c) {
		return FailRespWithMsg(c, Fail, "无权限访问")
	}

	var req updateUserReq
	if err := c.Bind(&req); err != nil {
		return FailResp(c, ParamError)
	}

	var user db.User
	if err := u.base.db.First(&user, req.Id).Error; err != nil {
		return FailRespWithMsg(c, Fail, "用户不存在")
	}

	// 检查用户名是否已存在
	if req.Username != user.Username {
		var count int64
		u.base.db.Model(&db.User{}).Where("username = ? AND id != ?", req.Username, req.Id).Count(&count)
		if count > 0 {
			return FailRespWithMsg(c, Fail, "用户名已存在")
		}
	}

	user.Username = req.Username
	user.Nickname = req.Nickname
	user.Email = req.Email
	user.Slogan = req.Slogan
	user.AvatarUrl = req.AvatarUrl
	user.CoverUrl = req.CoverUrl
	user.UpdatedAt = new(time.Time)
	*user.UpdatedAt = time.Now()

	// 如果有密码更新，处理密码加密
	if req.Password != "" {
		if len(req.Password) < 6 {
			return FailRespWithMsg(c, Fail, "密码长度不能少于6位")
		}
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			return FailRespWithMsg(c, Fail, "密码加密失败")
		}
		user.Password = string(password)
	}

	if err := u.base.db.Save(&user).Error; err != nil {
		return FailRespWithMsg(c, Fail, "更新用户失败")
	}

	return SuccessResp(c, h{})
}

// DeleteUser godoc
//
//	@Tags		User
//	@Summary	管理员删除用户
//	@Description	管理员删除指定用户（软删除）
//	@Accept		json
//	@Produce	json
//	@Param		id			path		int		true	"用户ID"
//	@Param		x-api-token	header		string	true	"登录TOKEN"
//	@Success		200
//	@Router		/api/user/{id} [post]
func (u UserHandler) DeleteUser(c echo.Context) error {
	if !u.isAdmin(c) {
		return FailRespWithMsg(c, Fail, "无权限访问")
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return FailResp(c, ParamError)
	}

	// 获取当前登录用户ID
	context, ok := c.(CustomContext)
	if !ok {
		return FailRespWithMsg(c, Fail, "无法获取当前用户信息")
	}
	currentUser := context.CurrentUser()
	if currentUser == nil {
		return FailRespWithMsg(c, Fail, "无法获取当前用户信息")
	}
	currentUserId := currentUser.Id

	// 检查是否尝试删除自己
	if int32(id) == currentUserId {
		return FailRespWithMsg(c, Fail, "不能删除自己的账户")
	}

	var user db.User
	if err := u.base.db.First(&user, int32(id)).Error; err != nil {
		return FailRespWithMsg(c, Fail, "用户不存在")
	}

	// 检查是否为管理员账户
	if user.Id == 1 {
		return FailRespWithMsg(c, Fail, "不能删除管理员账户")
	}

	// 删除用户相关的数据
	u.base.db.Where("userId = ?", user.Id).Delete(&db.Memo{})
	u.base.db.Where("author = ?", user.Id).Delete(&db.Comment{})

	// 删除用户
	if err := u.base.db.Delete(&user).Error; err != nil {
		return FailRespWithMsg(c, Fail, "删除用户失败")
	}

	return SuccessResp(c, h{})
}
