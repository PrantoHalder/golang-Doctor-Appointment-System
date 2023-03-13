package storage

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	ID        int          `form:"ID" db:"id"`
	FirstName string       `form:"FirstName" db:"first_name"`
	LastName  string       `form:"LastName" db:"last_name"`
	Email     string       `form:"Email" db:"email"`
	Username  string       `form:"Username" db:"username"`
	Password  string       `form:"Password" db:"password"`
	Role      string       `form:"Role" db:"role"`
	Status    bool         `form:"Status" db:"status"`
	CreatedAt time.Time    `form:"Created_at" db:"created_at"`
	UpdatedAt time.Time    `form:"Updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `form:"Deleted_at" db:"deleted_at"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.FirstName,
		validation.Required.Error("fast name can not be blank"),
		validation.Length(3, 45).Error("fast name must be between 3 to 45 characters"),
	),
		validation.Field(&u.LastName,
			validation.Required.Error("last name can not be blank"),
			validation.Length(3, 45).Error("last name must be between 3 to 45 characters"),
		),
		validation.Field(&u.Username,
			validation.Required.Error("username cannot be blank"),
			validation.Length(4, 10).Error("fast name must be between 4 to 10 characters"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("Email cannot be blank"),
			is.Email.Error("email should be in valid format"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password cannot be blank"),
			validation.Length(6, 8).Error("fast name must be between 6 to 8 characters"),
			validation.Required.When(u.ID == 0).Error("unable to set password"),
		),
	)
}

type Status struct {
	Username string `db:"username"`
	User User
}
func (u Status) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.Username,
		validation.Required.Error("fast name can not be blank"),
	),
	)
}