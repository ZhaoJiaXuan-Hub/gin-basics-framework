package entity

func (SystemAccount) TableName() string {
	return "system_account"
}

type SystemAccount struct {
	Id         int    `json:"id"`
	UserName   string `json:"user_name"`
	Number     int    `json:"number"`
	NickName   string `json:"nick_name"`
	Avatar     string `json:"avatar"`
	Password   string `json:"password"`
	Salt       string `json:"salt"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	DelStatus  int    `json:"delStatus"`
	CreateTime int64  `json:"createTime"`
}
