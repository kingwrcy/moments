package vo

type S3VO struct {
	Domain          string `json:"domain,omitempty"`          //S3 域名
	Bucket          string `json:"bucket,omitempty"`          //桶名称
	Region          string `json:"region,omitempty"`          //地区
	AccessKey       string `json:"accessKey,omitempty"`       //AK
	SecretKey       string `json:"secretKey,omitempty"`       //SK
	Endpoint        string `json:"endpoint,omitempty"`        //S3接口地址
	ThumbnailSuffix string `json:"thumbnailSuffix,omitempty"` //图片后缀
}

type SysConfigVO struct {
	AdminUserName         string `json:"adminUserName,omitempty"`    //管理员名称
	EnableS3              bool   `json:"enableS3"`                   //是否启用S3
	Favicon               string `json:"favicon,omitempty"`          //favicon
	Title                 string `json:"title,omitempty"`            //标题
	BeiAnNo               string `json:"beiAnNo,omitempty"`          //备案号码
	Css                   string `json:"css,omitempty"`              //自定义css
	Js                    string `json:"js,omitempty"`               //自定义js
	EnableGoogleRecaptcha bool   `json:"enableGoogleRecaptcha"`      //是否启用google recaptcha
	GoogleSiteKey         string `json:"googleSiteKey,omitempty"`    //google recaptcha siteKey
	EnableComment         bool   `json:"enableComment"`              //是否启用评论
	MaxCommentLength      int    `json:"maxCommentLength,omitempty"` //发言最大长度
	MemoMaxHeight         int    `json:"memoMaxHeight,omitempty"`    //单个memo的最大高度,单位px
	CommentOrder          string `json:"commentOrder,omitempty"`     //评论展示的顺序,asc:顺序,desc:逆序
	TimeFormat            string `json:"timeFormat,omitempty"`       //时间格式
	EnableRegister        bool   `json:"enableRegister,omitempty"`   //是否开启注册用户
}

type FullSysConfigVO struct {
	AdminUserName         string `json:"adminUserName,omitempty"`    //管理员名称
	EnableS3              bool   `json:"enableS3"`                   //是否启用S3
	Favicon               string `json:"favicon,omitempty"`          //favicon
	Title                 string `json:"title,omitempty"`            //标题
	BeiAnNo               string `json:"beiAnNo,omitempty"`          //备案号码
	Css                   string `json:"css,omitempty"`              //自定义css
	Js                    string `json:"js,omitempty"`               //自定义js
	S3                    S3VO   `json:"s3"`                         //S3相关信息
	EnableGoogleRecaptcha bool   `json:"enableGoogleRecaptcha"`      //是否启用google recaptcha
	GoogleSiteKey         string `json:"googleSiteKey,omitempty"`    //google recaptcha siteKey
	GoogleSecretKey       string `json:"googleSecretKey,omitempty"`  //google recaptcha secretKey
	EnableComment         bool   `json:"enableComment"`              //是否启用评论
	MaxCommentLength      int    `json:"maxCommentLength,omitempty"` //发言最大长度
	MemoMaxHeight         int    `json:"memoMaxHeight,omitempty"`    //单个memo的最大高度,单位px
	CommentOrder          string `json:"commentOrder,omitempty"`     //评论展示的顺序,asc:顺序,desc:逆序
	TimeFormat            string `json:"timeFormat,omitempty"`       //时间格式
	EnableRegister        bool   `json:"enableRegister,omitempty"`   //是否开启注册用户
}
