package vo

type LoginReq struct {
	Username string `json:"username,omitempty"` //用户名
	Password string `json:"password,omitempty"` //密码
}

type RegReq struct {
	Username       string `json:"username,omitempty"`       //用户名
	Password       string `json:"password,omitempty"`       //密码
	RepeatPassword string `json:"repeatPassword,omitempty"` //重复密码
}

type ProfileReq struct {
	ID        int    `json:"id,omitempty"`        //用户ID
	Nickname  string `json:"nickname,omitempty"`  //昵称
	AvatarUrl string `json:"avatarUrl,omitempty"` //头像URL
	Slogan    string `json:"slogan,omitempty"`    //标语
	CoverUrl  string `json:"coverUrl,omitempty"`  //封面URL
	Password  string `json:"password,omitempty"`  //密码,不修改不要填写
}
