package user_management

type UserService interface {
	GetUser(id int) (User, error)
	GetUsers() ([]User, error)
	UpdateUser(user User, id int) (User, error)
	DeleteUser(id int) error
	Register(user User) (User, error)
	Login(email string, password string) (string, error)
}

type UserRepository interface {
	GetUser(id int) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUsers() ([]User, error)
	UpdateUser(user User, id int) (User, error)
	DeleteUser(id int) error
	Register(user User) (User, error)
	Login(email string, password string) (string, error)
}

type DepartermentService interface {
	GetDeparterment(id int) (Departerment, error)
	GetDeparterments() ([]Departerment, error)
	UpdateDeparterment(departerment Departerment, id int) (Departerment, error)
	DeleteDeparterment(id int) error
}

type DepartermentRepository interface {
	GetDeparterment(id int) (Departerment, error)
	GetDeparterments() ([]Departerment, error)
	UpdateDeparterment(departerment Departerment, id int) (Departerment, error)
	DeleteDeparterment(id int) error
}
