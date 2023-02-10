package botdatabase

type BotDataBase interface {
	AddUser(chatId int64, user User) error
	IsNewUser(chatId int64) (bool, error)
	GetUser(chatId int64) (*User, error)
	UpdateUserCommand(chatId int64, command string) error
	UpdateUserArticles(chatId int64, articles map[string]string) error
}

type User struct {
	FirstName    string
	LastName     string
	LastCommand  string
	LastArticles map[string]string
}
