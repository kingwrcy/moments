package vo

type AddCommentReq struct {
	Content  string `json:"content,omitempty"`  //正文
	ReplyTo  string `json:"replyTo,omitempty"`  //回复给某人
	Username string `json:"username,omitempty"` //作者用户名,未登录时才有效
	Email    string `json:"email,omitempty"`    //作者邮箱
	Website  string `json:"website,omitempty"`  //作者网站
	MemoID   int32  `json:"memoId,omitempty"`   //回复的memoID
	Token    string `json:"token,omitempty"`    //google recaptcha的token,开启的话必填
}
