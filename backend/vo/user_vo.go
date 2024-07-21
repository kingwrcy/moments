package vo

type LoginReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type RegReq struct {
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	RepeatPassword string `json:"repeatPassword,omitempty"`
}

type ProfileReq struct {
	ID        int    `json:"id,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Slogan    string `json:"slogan,omitempty"`
	CoverUrl  string `json:"coverUrl,omitempty"`
	Password  string `json:"password,omitempty"`
}
