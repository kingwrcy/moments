package vo

type ListMemoReq struct {
	Page   int    `json:"page,omitempty"`
	Size   int    `json:"size,omitempty"`
	Tag    string `json:"tag,omitempty"`
	Source string `json:"source,omitempty"`
}
