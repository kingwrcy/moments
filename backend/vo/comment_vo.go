package vo

type AddCommentReq struct {
	Content    string `json:"content,omitempty"`    //正文
	ReplyTo    string `json:"replyTo,omitempty"`    //回复给某人
	ReplyEmail string `json:"replyEmail,omitempty"` //回复邮箱
	Username   string `json:"username,omitempty"`   //作者用户名,未登录时才有效
	Email      string `json:"email,omitempty"`      //作者邮箱
	Website    string `json:"website,omitempty"`    //作者网站

	MemoID  int32  `json:"memoId,omitempty"`   //回复的memoID
	GuestID string `json:"guest_id,omitempty"` // 访客ID,未登录时才有效
	Token   string `json:"token,omitempty"`    //google recaptcha的token,开启的话必填
}
