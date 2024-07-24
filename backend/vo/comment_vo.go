package vo

type AddCommentReq struct {
	Content  string `json:"content,omitempty"`
	ReplyTo  string `json:"replyTo,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Website  string `json:"website,omitempty"`
	MemoID   int32  `json:"memoId,omitempty"`
	Token    string `json:"token,omitempty"`
}
