package vo

type S3VO struct {
	Domain          string `json:"domain,omitempty"`
	Bucket          string `json:"bucket,omitempty"`
	Region          string `json:"region,omitempty"`
	AccessKey       string `json:"accessKey,omitempty"`
	SecretKey       string `json:"secretKey,omitempty"`
	Endpoint        string `json:"endpoint,omitempty"`
	ThumbnailSuffix string `json:"thumbnailSuffix,omitempty"`
}

type SysConfigVO struct {
	AdminUserName string `json:"adminUserName,omitempty"`
	EnableS3      bool   `json:"enableS3,omitempty"`
	Favicon       string `json:"favicon,omitempty"`
	Title         string `json:"title,omitempty"`
	BeiAnNo       string `json:"beiAnNo,omitempty"`
	Css           string `json:"css,omitempty"`
	Js            string `json:"js,omitempty"`
	S3            S3VO   `json:"s3"`
}
