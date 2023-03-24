package postgres

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"main.go/usermgm/storage"
	"sort"
	"testing"
)

func TestRegister(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "REGISTER_PATIENT_SUCCESS",
			in: storage.User{
				FirstName: "Rahim",
				LastName:  "Khan",
				Email:     "rahim@gamil.com",
				Username:  "rahim",
				Password:  "12345678",
			},
			want: &storage.User{
				FirstName: "Rahim",
				LastName:  "Khan",
				Email:     "rahim@gamil.com",
				Username:  "rahim",
				Password:  "12345678",
				Role:      "user",
				Is_active: true,
			},
			wantErr: false,
		},
		{
			name: "REGISTER_PATIENT_EMAIL_UNIQUE_FAILURE",
			in: storage.User{
				FirstName: "Karim",
				LastName:  "Khanna",
				Email:     "rahim@gamil.com",
				Username:  "rahimm",
				Password:  "12345678",
			},
			wantErr: true,
		},
		{
			name: "REGISTER_PATIENT_USER_NAME_FAILURE",
			in: storage.User{
				FirstName: "Karim",
				LastName:  "Khanna",
				Email:     "rahimm@gamil.com",
				Username:  "rahim",
				Password:  "12345678",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Register(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestRegisterAppointment(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newuser := storage.User{
		FirstName: "Rahim",
		LastName:  "Hossain",
		Email:     "rahim@gmail.com",
		Username:  "rahim",
		Password:  "12345678",
		Role:      "user",
	}
	user, err := s.RegisterPatient(newuser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
	}
	newdoctor := storage.User{
		FirstName: "Pranto",
		LastName:  "Halder",
		Email:     "pranto@gmail.com",
		Username:  "pranto",
		Password:  "12345678",
		Role:      "doctor",
	}
	_, err = s.RegisterDoctorAdmin(newdoctor)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorAdmin() error = %v", err)
	}
	newdoctorType := storage.DoctorType{
		DoctorType: "Neurologists",
	}
	_, err = s.Registerdoctortype(newdoctorType)
	if err != nil {
		t.Fatalf("PostgresStorage.Registerdoctortype() error = %v", err)
	}
	newdoctorDetails := storage.Doctor{
		UserID:       2,
		DoctorTypeID: 1,
		Degree:       "MBBS",
		Gender:       "MALE",
	}
	doctordetails, err := s.RegisterDoctorDeatils(newdoctorDetails)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorDeatils() error = %v", err)
	}
	workdays := storage.Schedule{
		WorkDays: "Friday",
	}
	workday, err := json.Marshal(workdays.WorkDays)
	if err != nil {
		fmt.Printf("#%v", err)
	}
	newschedule := storage.Schedule{
		DoctorDetailsID: 1,
		WorkDays:        string(workday),
		Address:         "Khulna",
		Phone:           "01716504535",
	}
	schedule, err := s.RegisterDoctorSchedule(newschedule)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	tests := []struct {
		name    string
		in      storage.Appointment
		want    *storage.Appointment
		wantErr bool
	}{
		{
			name: "REGISTER_APPOINTMENT_SUCCESS",
			in: storage.Appointment{
				UserID:          1,
				DoctorDetailsID: 1,
				ScheduleID:      1,
			},
			want: &storage.Appointment{
				UserID:          1,
				DoctorDetailsID: 1,
				ScheduleID:      1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.DoctorDetailsID = doctordetails.ID
			tt.in.ScheduleID = schedule.ID
			tt.in.UserID = user.ID
			got, err := s.RegisterAppointment(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Appointment{}, "ID", "Is_Appointed", "TimeSlot"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestEditUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newuser := storage.User{
		FirstName: "Rahim",
		LastName:  "Hossain",
		Email:     "rahim@gmail.com",
		Username:  "rahim",
		Password:  "12345678",
		Role:      "user",
	}
	user, err := s.RegisterPatient(newuser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int
		want    *storage.User
		wantErr bool
	}{
		{
			name: "EDIT_USER_SUCCESS",
			in:   1,
			want: &storage.User{
				ID:        0,
				FirstName: "Rahim",
				LastName:  "Hossain",
				Email:     "rahim@gmail.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in = user.ID
			got, err := s.EditUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Username", "Password", "Role", "Is_active", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newuser := storage.User{
		FirstName: "Rahim",
		LastName:  "Hossain",
		Email:     "rahim@gmail.com",
		Username:  "rahim",
		Password:  "12345678",
		Role:      "user",
	}
	user, err := s.RegisterPatient(newuser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
	}
	tests := []struct {
		name    string
		in      storage.UpdateUser
		want    *storage.UpdateUser
		wantErr bool
	}{
		{
			name: "UPDATE_USER_SUCCESS",
			in: storage.UpdateUser{
				ID:        1,
				FirstName: "Rahim",
				LastName:  "Hossain",
				Email:     "rahim@gmail.com",
			},
			want: &storage.UpdateUser{
				ID:        1,
				FirstName: "Rahim",
				LastName:  "Hossain",
				Email:     "rahim@gmail.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.FirstName = user.FirstName
			tt.in.LastName = user.LastName
			tt.in.Email = user.Email
			got, err := s.UpdateUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestDeleteUserByID(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newuser := storage.User{
		FirstName: "Rahim",
		LastName:  "Hossain",
		Email:     "rahim@gmail.com",
		Username:  "rahim",
		Password:  "12345678",
		Role:      "user",
	}
	user, err := s.RegisterPatient(newuser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int
		wantErr bool
	}{
		{
			name:    "DELETE_USER_SUCCESS",
			in:      1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in = user.ID
			err := s.DeleteUserByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestListUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newUser := []storage.User{
		{
			FirstName: "Karim",
			LastName:  "Khan",
			Email:     "karim@gmail.com",
			Username:  "karin",
			Role:      "user",
			Password:  "12345678",
		},
		{
			FirstName: "Rahim",
			LastName:  "shake",
			Email:     "rahim@gmail.com",
			Username:  "rahim",
			Role:      "user",
			Password:  "12345678",
		}, {
			FirstName: "Pranto",
			LastName:  "Halder",
			Email:     "pranto@gmail.com",
			Username:  "pranto",
			Role:      "user",
			Password:  "12345678",
		}, {
			FirstName: "Shovon",
			LastName:  "Halder",
			Email:     "shovon@gmail.com",
			Username:  "shovon",
			Role:      "user",
			Password:  "12345678",
		},
	}
	for _, value := range newUser {
		_, err := s.RegisterPatient(value)
		if err != nil {
			t.Errorf("PostGressStorage.ListUser() error = %v", err)
			return
		}
	}
	tests := []struct {
		name    string
		in      storage.UserFilter
		want    []storage.User
		wantErr bool
	}{
		{
			name:    "SUCCESS",
			in:      storage.UserFilter{},
			want:    newUser,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Username", "Role", "Password", "Is_active", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestEditUserStatus(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newuser := storage.User{
		FirstName: "Rahim",
		LastName:  "Hossain",
		Email:     "rahim@gmail.com",
		Username:  "rahim",
		Password:  "12345678",
		Role:      "user",
	}
	user, err := s.RegisterPatient(newuser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int
		want    *storage.UpdateStatus
		wantErr bool
	}{
		{
			name: "SUCESS_STATUS_EDIT",
			in:   1,
			want: &storage.UpdateStatus{
				ID:        1,
				Is_active: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in = user.ID
			got, err := s.EditUserStatus(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestUpdateUserStatus(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newuser := storage.User{
		FirstName: "Rahim",
		LastName:  "Hossain",
		Email:     "rahim@gmail.com",
		Username:  "rahim",
		Password:  "12345678",
		Role:      "user",
	}
	_, err := s.RegisterPatient(newuser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
	}
	tests := []struct {
		name    string
		in      storage.UpdateStatus
		want    *storage.UpdateStatus
		wantErr bool
	}{
		{
			name: "USER_UPDATE_STATUS_SUCCESS",
			in: storage.UpdateStatus{
				ID:        1,
				Is_active: false,
			},
			want: &storage.UpdateStatus{
				ID:        1,
				Is_active: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UpdateUserStatus(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestShowDoctorListToUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newuser := []storage.User{
		{
			FirstName: "Rahim",
			LastName:  "Hossain",
			Email:     "rahim@gmail.com",
			Username:  "rahim",
			Password:  "12345678",
			Role:      "doctor",
		},
		{
			FirstName: "Karim",
			LastName:  "Khan",
			Email:     "karim@gmail.com",
			Username:  "karim",
			Password:  "12345678",
			Role:      "doctor",
		},
	}
	for _, value := range newuser {
		_, err := s.RegisterDoctorAdmin(value)
		if err != nil {
			t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
		}
	}

	doctortype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	_, err := s.Registerdoctortype(doctortype)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
	}
	doctordetails := storage.Doctor{
		UserID:       1,
		DoctorTypeID: 1,
		Degree:       "MBBS",
		Gender:       "Male",
	}
	_, err = s.RegisterDoctorDeatils(doctordetails)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
	}
	newshow := []storage.ShowDoctorToPatient{
		{
			FirstName:  "Rahim",
			LastName:   "Hossain",
			Degree:     "MBBS",
			DoctorType: "Nurologist",
			Gender:     "Male",
		},
	}
	tests := []struct {
		name    string
		in      int
		want    []storage.ShowDoctorToPatient
		wantErr bool
	}{
		{
			name:    "SHOW_DOCTOR_LIST_TO_USER_SUCCESS",
			in:      1,
			want:    newshow,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in = doctordetails.DoctorTypeID
			got, err := s.ShowDoctorListToUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.ShowDoctorListToUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.ShowDoctorToPatient{}, "ID"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

