package storage

import (
	"database/sql"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	NotFound = "sql: no rows in result set"
)

type DoctorDetailsList struct {
	ID         int    `db:"id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	DoctorType string `db:"doctortype"`
	Degree     string `db:"degree"`
	Gender     string `db:"gender"`
}
type User struct {
	ID        int          `db:"id"`
	FirstName string       `db:"first_name"`
	LastName  string       `db:"last_name"`
	Email     string       `db:"email"`
	Username  string       `db:"username"`
	Password  string       `db:"password"`
	Role      string       `db:"role"`
	Is_active bool         `db:"is_active"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
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
			validation.Match(regexp.MustCompile(`^user$|^admin$|^doctor$`)).
				Error("Role should be user or admin or doctor"),
		),
		validation.Field(&u.Password,
			validation.Required.Error("password cannot be blank"),
			validation.Length(6, 8).Error("fast name must be between 6 to 8 characters"),
			validation.Required.When(u.ID == 0).Error("unable to set password"),
		),
	)
}

type Register struct {
	ID        int          `db:"id"`
	FirstName string       `db:"first_name"`
	LastName  string       `db:"last_name"`
	Email     string       `db:"email"`
	Username  string       `db:"username"`
	Password  string       `db:"password"`
	Role      string       `db:"role"`
	Is_active bool         `db:"is_active"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (u Register) Validate() error {
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

type DoctorDetails struct {
	ID           int          `db:"id"`
	UserID       int          `db:"userid"`
	DoctorTypeID int          `db:"doctortypeid"`
	Degree       string       `db:"degree"`
	Gender       string       `db:"gender"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

func (u DoctorDetails) Validate() error {
	return validation.ValidateStruct(&u, validation.Field(&u.Degree,
		validation.Required.Error("DoctorTypeID can not be blank"),
	),
		validation.Field(&u.Gender,
			validation.Required.Error("Gender cannot be blank"),
			validation.Match(regexp.MustCompile(`^Male$|^Female$|^Others$`)).
				Error("Gender should be Male or Female or Others"),
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
	return validation.ValidateStruct(&u, validation.Field(&u.StartAt,
		validation.Required.Error("DoctorDetailsID can not be blank"),
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

type Appointment struct {
	ID              int    `db:"id"`
	UserID          int    `db:"userid"`
	DoctorDetailsID int    `db:"doctordetailsid"`
	ScheduleID      int    `db:"schduleid"`
	Is_Appointed    bool   `db:"is_appointed"`
	TimeSlot        string `db:"timeslot"`
}

type UpdateStatus struct {
	ID        int  `db:"id"`
	Is_active bool `db:"is_active"`
}
type ShowDoctorToPatient struct {
	ID         int    `db:"id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Degree     string `db:"degree"`
	DoctorType string `db:"doctortype"`
	Gender     string `db:"gender"`
}
type AppontmentStatus struct {
	ID           int    `db:"id"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	Is_Appointed bool   `db:"is_appointed"`
	TimeSlot     string `db:"timeslot"`
}
