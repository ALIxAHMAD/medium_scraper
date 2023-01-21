package botdatabase

type BotDataBase interface {
	AddUser(chatId int64, user User) error
	IsNewUser(chatId int64) (bool, error)
	GetUser(chatId int64) (*User, error)
	UpdateUser(chatId int64, command string) error
}

type User struct {
	FirstName   string
	LastName    string
	LastCommand string
}
