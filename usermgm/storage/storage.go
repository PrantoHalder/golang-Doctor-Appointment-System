package storage

import (
	"database/sql"
	"regexp"
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
	Is_active bool         `form:"Is_active" db:"is_active"`
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
		validation.Field(&u.Role,
			validation.Required.Error("Role cannot be blank"),
			validation.Match(regexp.MustCompile(`^user$`)).
				Error("Role should be user"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password cannot be blank"),
			validation.Length(6, 8).Error("fast name must be between 6 to 8 characters"),
			validation.Required.When(u.ID == 0).Error("unable to set password"),
		),
	)
}

type Patient struct {
	ID        int          `form:"ID" db:"id"`
	FirstName string       `form:"FirstName" db:"first_name"`
	LastName  string       `form:"LastName" db:"last_name"`
	Email     string       `form:"Email" db:"email"`
	Username  string       `form:"Username" db:"username"`
	Password  string       `form:"Password" db:"password"`
	Role      string       `form:"Role" db:"role"`
	Is_active bool         `form:"Is_active" db:"is_active"`
	CreatedAt time.Time    `form:"Created_at" db:"created_at"`
	UpdatedAt time.Time    `form:"Updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `form:"Deleted_at" db:"deleted_at"`
}

func (u Patient) Validate() error {
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

type Admin struct {
	ID        int          `form:"ID" db:"id"`
	FirstName string       `form:"FirstName" db:"first_name"`
	LastName  string       `form:"LastName" db:"last_name"`
	Email     string       `form:"Email" db:"email"`
	Username  string       `form:"Username" db:"username"`
	Password  string       `form:"Password" db:"password"`
	Role      string       `form:"Role" db:"role"`
	Is_active bool         `form:"Is_active" db:"is_active"`
	CreatedAt time.Time    `form:"Created_at" db:"created_at"`
	UpdatedAt time.Time    `form:"Updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `form:"Deleted_at" db:"deleted_at"`
}

func (u Admin) Validate() error {
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
		validation.Field(&u.Role,
			validation.Required.Error("Role cannot be blank"),
			validation.Match(regexp.MustCompile(`^admin$`)).
				Error("Role should be admin"),
		),
		validation.Field(&u.Role,
			validation.Required.Error("Role cannot be blank"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password cannot be blank"),
			validation.Length(6, 8).Error("fast name must be between 6 to 8 characters"),
			validation.Required.When(u.ID == 0).Error("unable to set password"),
		),
	)
}

type DoctorU struct {
	ID        int          `form:"ID" db:"id"`
	FirstName string       `form:"FirstName" db:"first_name"`
	LastName  string       `form:"LastName" db:"last_name"`
	Email     string       `form:"Email" db:"email"`
	Username  string       `form:"Username" db:"username"`
	Password  string       `form:"Password" db:"password"`
	Role      string       `form:"Role" db:"role"`
	Is_active bool         `form:"Is_active" db:"is_active"`
	CreatedAt time.Time    `form:"Created_at" db:"created_at"`
	UpdatedAt time.Time    `form:"Updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `form:"Deleted_at" db:"deleted_at"`
}

func (u DoctorU) Validate() error {
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
		validation.Field(&u.Role,
			validation.Required.Error("Role cannot be blank"),
			validation.Match(regexp.MustCompile(`^doctor$`)).
				Error("Role should be doctor"),
		),
		validation.Field(&u.Role,
			validation.Required.Error("Role cannot be blank"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password cannot be blank"),
			validation.Length(6, 8).Error("fast name must be between 6 to 8 characters"),
			validation.Required.When(u.ID == 0).Error("unable to set password"),
		),
	)
}

type Doctor struct {
	ID           int          `db:"id"`
	UserID       int          `db:"userid"`
	DoctorTypeID int          `db:"doctortypeid"`
	Degree       string       `db:"degree"`
	Gender       string       `db:"gender"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

func (u Doctor) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.DoctorTypeID,
		validation.Required.Error("DoctorTypeID can not be blank"),
	),
		validation.Field(&u.Degree,
			validation.Required.Error("Degree can not be blank"),
		),
		validation.Field(&u.Gender,
			validation.Required.Error("Gender cannot be blank"),
			validation.Match(regexp.MustCompile(`^Male$|^Female$|^Others$`)).
				Error("Gender should be Male or Female or Others"),
		),

		validation.Field(&u.UserID,
			validation.Required.Error("UserID cannot be blank"),
		),
	)
}

type Status struct {
	Username string `db:"username"`
	User     User
}

func (u Status) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.Username,
		validation.Required.Error("fast name can not be blank"),
	),
	)
}

type Login struct {
	Username string
	Password string
}

type Edit struct {
	ID int `form:"ID" db:"id"`
}

func (u Edit) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.ID,
		validation.Required.Error("id can not be blank"),
	),
	)
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Username,
			validation.Required.Error("The username field is required."),
		),
		validation.Field(&l.Password,
			validation.Required.Error("The password field is required."),
		),
	)
}

type DoctorType struct {
	ID         int          `db:"id"`
	DoctorType string       `form:"DoctorType" db:"doctortype"`
	CreatedAt  time.Time    `form:"Created_at" db:"created_at"`
	UpdatedAt  time.Time    `form:"Updated_at" db:"updated_at"`
	DeletedAt  sql.NullTime `form:"Deleted_at" db:"deleted_at"`
}

func (u DoctorType) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.DoctorType,
		validation.Required.Error("id can not be blank"),
	),
	)
}

type UpdateUser struct {
	ID        int    `form:"ID" db:"id"`
	FirstName string `form:"FirstName" db:"first_name"`
	LastName  string `form:"LastName" db:"last_name"`
	Email     string `form:"Email" db:"email"`
	Is_active bool   `form:"Is_active" db:"is_active"`
}

func (u UpdateUser) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.FirstName,
		validation.Required.Error("fast name can not be blank"),
		validation.Length(3, 45).Error("fast name must be between 3 to 45 characters"),
	),
		validation.Field(&u.LastName,
			validation.Required.Error("last name can not be blank"),
			validation.Length(3, 45).Error("last name must be between 3 to 45 characters"),
		),
		validation.Field(&u.Email,
			validation.Required.Error("Email cannot be blank"),
			is.Email.Error("email should be in valid format"),
		),
	)
}

type UserFilter struct {
	SearchTerm string `pb:"1"`
	Offset     int    `pb:"1"`
	Limit      int    `pb:"1"`
}
type Schedule struct {
	ID              int       `db:"id"`
	DoctorDetailsID int       `db:"doctorid"`
	StartAt         time.Time `db:"startat"`
	EndAt           time.Time `db:"endat"`
	WorkDays        string    `db:"workdays"`
	Address         string    `db:"address"`
	Phone           string    `db:"phone"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
func (u Schedule) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.DoctorDetailsID,
		validation.Required.Error("DoctorDetailsID can not be blank"),
	),
		validation.Field(&u.StartAt,
			validation.Required.Error("StartAt can not be blank"),
		),
		validation.Field(&u.EndAt,
			validation.Required.Error("EndAt cannot be blank"),
		),
		validation.Field(&u.WorkDays,
			validation.Required.Error("WorkDays cannot be blank"),
		),
		validation.Field(&u.Address,
			validation.Required.Error("Address cannot be blank"),
		),
		validation.Field(&u.Phone,
			validation.Required.Error("Phone cannot be blank"),
		),
	)
}