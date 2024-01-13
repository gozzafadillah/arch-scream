package user_management

type User struct {
	MongoID   string  `bson:"_id,omitempty" json:"mongo_id,omitempty"`
	ID        int     `bson:"id,omitempty" json:"id,omitempty"`
	FirstName string  `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string  `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Email     string  `bson:"email,omitempty" json:"email,omitempty"`
	Password  string  `bson:"password,omitempty" json:"password,omitempty"`
	Fee       float64 `bson:"fee,omitempty" json:"fee,omitempty"`
}

type Login struct {
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}

type Departerment struct {
	ID   int     `bson:"_id,omitempty" json:"id,omitempty"`
	Name string  `bson:"name,omitempty" json:"name,omitempty"`
	User []*User `bson:"user,omitempty" json:"user,omitempty"`
}
