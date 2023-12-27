package datastructure

type Complain struct {
	Id               int    `json:"id"`
	Phone            int    `json:"phone"`
	TelegramUserName string `json:"telegram_user_name"`
	TelegramUserId   string `json:"telegram_user_id"`
	MainReason       string `json:"main_reason"`
	Description      string `json:"description"`
}

type UserComplain struct {
	Id         int
	ComplainId int
	UserId     int
}
