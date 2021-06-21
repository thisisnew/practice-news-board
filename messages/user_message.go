package messages

type User struct {
	UserId    uint   `json:"userId"`
	UserName  string `json:"userName"`
	UserPw    string `json:"userPw"`
	UserEmail string `json:"userEmail"`
	UserRank  uint   `json:"userRank"`
	UserState bool   `json:"userState"`
	IsLogin   bool   `json:"isLogin"`
}
