package types

type Student struct {
	ID	int `json:"id"`
	Name	string `validate:"required" json:"name"`
	Email	string `validate:"required" json:"email"`
	Age	int `validate:"required" json:"age"`
}
