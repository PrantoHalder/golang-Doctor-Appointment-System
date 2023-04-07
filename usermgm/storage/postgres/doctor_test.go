package postgres

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"main.go/usermgm/storage"
)

func TestRegisterDoctorDeatils(t *testing.T) {
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
		Role:      "doctor",
	}
	_, err := s.RegisterDoctorAdmin(newuser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
	}
	newdoctortype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	_, err = s.Registerdoctortype(newdoctortype)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
	}
	tests := []struct {
		name    string
		in      storage.DoctorDetails
		want    *storage.DoctorDetails
		wantErr bool
	}{
		{
			name: "REGISTER_DOCTOR_DETAILS_SUCCESS",
			in: storage.DoctorDetails{
				UserID:       1,
				DoctorTypeID: 1,
				Degree:       "MBBS",
				Gender:       "MALE",
			},
			want: &storage.DoctorDetails{
				UserID:       1,
				DoctorTypeID: 1,
				Degree:       "MBBS",
				Gender:       "MALE",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.RegisterDoctorDeatils(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.DoctorDetails{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestRegisterDoctorSchedule(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newdoctor := storage.User{
		FirstName: "Pranto",
		LastName:  "Halder",
		Email:     "pranto@gmail.com",
		Username:  "pranto",
		Password:  "12345678",
		Role:      "doctor",
	}
	_, err := s.RegisterDoctorAdmin(newdoctor)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	newdoctortype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	_, err = s.Registerdoctortype(newdoctortype)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	newdoctordetails := storage.DoctorDetails{
		UserID:       1,
		DoctorTypeID: 1,
		Degree:       "MBBS",
		Gender:       "Male",
	}
	_, err = s.RegisterDoctorDeatils(newdoctordetails)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
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
	tests := []struct {
		name    string
		in      storage.Schedule
		want    *storage.Schedule
		wantErr bool
	}{
		{
			name:    "REGISTER_DOCTOR_SCHEDULE",
			in:      newschedule,
			want:    &newschedule,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.RegisterDoctorSchedule(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.RegisterDoctorSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Schedule{}, "ID", "CreatedAt", "UpdatedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestListDoctor(t *testing.T) {
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
			Role:      "doctor",
			Password:  "12345678",
		},
		{
			FirstName: "Rahim",
			LastName:  "shake",
			Email:     "rahim@gmail.com",
			Username:  "rahim",
			Role:      "doctor",
			Password:  "12345678",
		}, {
			FirstName: "Pranto",
			LastName:  "Halder",
			Email:     "pranto@gmail.com",
			Username:  "pranto",
			Role:      "doctor",
			Password:  "12345678",
		}, {
			FirstName: "Shovon",
			LastName:  "Halder",
			Email:     "shovon@gmail.com",
			Username:  "shovon",
			Role:      "doctor",
			Password:  "12345678",
		},
	}
	for _, value := range newUser {
		_, err := s.RegisterDoctorAdmin(value)
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
			got, err := s.ListDoctor(tt.in)
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

func TestEditDoctorDetails(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newdoctor := storage.User{
		FirstName: "Pranto",
		LastName:  "Halder",
		Email:     "pranto@gmail.com",
		Username:  "pranto",
		Password:  "12345678",
		Role:      "doctor",
	}
	_, err := s.RegisterDoctorAdmin(newdoctor)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	newdoctortype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	_, err = s.Registerdoctortype(newdoctortype)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	doctordetails := storage.DoctorDetails{
		UserID:       1,
		DoctorTypeID: 1,
		Degree:       "MBBS",
		Gender:       "Male",
	}
	_, err = s.RegisterDoctorDeatils(doctordetails)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	tests := []struct {
		name    string
		in      int
		want    *storage.DoctorDetails
		wantErr bool
	}{
		{
			name: "",
			in:   1,
			want: &storage.DoctorDetails{
				ID:     1,
				Degree: "MBBS",
				Gender: "Male",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in = doctordetails.UserID
			got, err := s.EditDoctorDetails(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.EditDoctorDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.DoctorDetails{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateDoctorDetails(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newdoctor := storage.User{
		FirstName: "Pranto",
		LastName:  "Halder",
		Email:     "pranto@gmail.com",
		Username:  "pranto",
		Password:  "12345678",
		Role:      "doctor",
	}
	_, err := s.RegisterDoctorAdmin(newdoctor)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	newdoctortype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	_, err = s.Registerdoctortype(newdoctortype)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	doctordetails := storage.DoctorDetails{
		UserID:       1,
		DoctorTypeID: 1,
		Degree:       "MBBS",
		Gender:       "Male",
	}
	_, err = s.RegisterDoctorDeatils(doctordetails)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	tests := []struct {
		name    string
		in      storage.DoctorDetails
		want    *storage.DoctorDetails
		wantErr bool
	}{
		{
			name: "UPDATE_DOCTOR_DETAILS_SUCCESS",
			in: storage.DoctorDetails{
				ID:     1,
				Degree: "FRCS",
				Gender: "Female",
			},
			want: &storage.DoctorDetails{
				Degree: "FRCS",
				Gender: "Female",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UpdateDoctorDetails(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.DoctorDetails{}, "ID", "DoctorTypeID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestEditDoctorSchedule(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newdoctor := storage.User{
		FirstName: "Pranto",
		LastName:  "Halder",
		Email:     "pranto@gmail.com",
		Username:  "pranto",
		Password:  "12345678",
		Role:      "doctor",
	}
	_, err := s.RegisterDoctorAdmin(newdoctor)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	newdoctortype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	_, err = s.Registerdoctortype(newdoctortype)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	doctordetails := storage.DoctorDetails{
		UserID:       1,
		DoctorTypeID: 1,
		Degree:       "MBBS",
		Gender:       "Male",
	}
	_, err = s.RegisterDoctorDeatils(doctordetails)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
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
	_, err = s.RegisterDoctorSchedule(newschedule)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int
		want    *storage.Schedule
		wantErr bool
	}{
		{
			name:    "EDIT_DOCTOR_SCHEDULE_SUCCESS",
			in:      1,
			want:    &newschedule,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.EditDoctorSchedule(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.EditDoctorSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Schedule{}, "ID", "DoctorDetailsID", "CreatedAt", "UpdatedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateDoctorSchedule(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newdoctor := storage.User{
		FirstName: "Pranto",
		LastName:  "Halder",
		Email:     "pranto@gmail.com",
		Username:  "pranto",
		Password:  "12345678",
		Role:      "doctor",
	}
	_, err := s.RegisterDoctorAdmin(newdoctor)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	newdoctortype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	_, err = s.Registerdoctortype(newdoctortype)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	doctordetails := storage.DoctorDetails{
		UserID:       1,
		DoctorTypeID: 1,
		Degree:       "MBBS",
		Gender:       "Male",
	}
	_, err = s.RegisterDoctorDeatils(doctordetails)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
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
		StartAt:         time.Time{},
		EndAt:           time.Time{},
		WorkDays:        string(workday),
		Address:         "Khulna",
		Phone:           "01716504535",
	}
	_, err = s.RegisterDoctorSchedule(newschedule)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	tests := []struct {
		name    string
		in      storage.Schedule
		want    *storage.Schedule
		wantErr bool
	}{
		{
			name: "UPDATE_DOCTOR_SCHEDULE_SUCCESS",
			in: storage.Schedule{
				ID:       1,
				StartAt:  time.Time{},
				EndAt:    time.Time{},
				WorkDays: string(workday),
				Address:  "Khulna",
				Phone:    "01716504535",
			},
			want:    &newschedule,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UpdateDoctorSchedule(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.UpdateDoctorSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Schedule{}, "ID", "DoctorDetailsID", "CreatedAt", "UpdatedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestApproveEdit(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newPatient := storage.User{
		FirstName: "Shovon",
		LastName:  "Halder",
		Email:     "shovon@gmail.com",
		Username:  "shovon",
		Password:  "12345678",
		Role:      "user",
	}
	_, err := s.RegisterPatient(newPatient)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
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
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	newdoctortype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	_, err = s.Registerdoctortype(newdoctortype)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	doctordetails := storage.DoctorDetails{
		UserID:       2,
		DoctorTypeID: 1,
		Degree:       "MBBS",
		Gender:       "Male",
	}
	_, err = s.RegisterDoctorDeatils(doctordetails)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
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
		StartAt:         time.Time{},
		EndAt:           time.Time{},
		WorkDays:        string(workday),
		Address:         "Khulna",
		Phone:           "01716504535",
	}
	_, err = s.RegisterDoctorSchedule(newschedule)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	newappointment := storage.Appointment{
		UserID:          1,
		DoctorDetailsID: 1,
		ScheduleID:      1,
		TimeSlot:        "waiting",
	}
	_, err = s.RegisterAppointment(newappointment)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int
		want    *storage.Appointment
		wantErr bool
	}{
		{
			name: "APPROVE_APPOINTMENT_STATUS_SUCCESS",
			in:   1,
			want: &storage.Appointment{
				ID:           1,
				Is_Appointed: false,
				TimeSlot:     "waiting",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in = newappointment.DoctorDetailsID
			got, err := s.ApproveEdit(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.ApproveEdit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Schedule{}, "ID", "DoctorDetailsID", "CreatedAt", "UpdatedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestApproveUpdate(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newPatient := storage.User{
		FirstName: "Shovon",
		LastName:  "Halder",
		Email:     "shovon@gmail.com",
		Username:  "shovon",
		Password:  "12345678",
		Role:      "user",
	}
	_, err := s.RegisterPatient(newPatient)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
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
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	newdoctortype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	_, err = s.Registerdoctortype(newdoctortype)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
	}
	doctordetails := storage.DoctorDetails{
		UserID:       2,
		DoctorTypeID: 1,
		Degree:       "MBBS",
		Gender:       "Male",
	}
	_, err = s.RegisterDoctorDeatils(doctordetails)
	if err != nil {
		t.Errorf("PostGressStorage.ListUser() error = %v", err)
		return
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
		StartAt:         time.Time{},
		EndAt:           time.Time{},
		WorkDays:        string(workday),
		Address:         "Khulna",
		Phone:           "01716504535",
	}
	_, err = s.RegisterDoctorSchedule(newschedule)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	newappointment := storage.Appointment{
		UserID:          1,
		DoctorDetailsID: 1,
		ScheduleID:      1,
		TimeSlot:        "waiting",
	}
	_, err = s.RegisterAppointment(newappointment)
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
			name: "UPDATE_APPROVE_UPDATE_SUCCESS",
			in: storage.Appointment{
				ID:           1,
				Is_Appointed: true,
				TimeSlot:     "waiting",
			},
			want: &storage.Appointment{
				ID:           1,
				Is_Appointed: true,
				TimeSlot:     "waiting",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ApproveUpdate(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.ApproveUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Schedule{}, "ID", "DoctorDetailsID", "CreatedAt", "UpdatedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestListDoctorDetails(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	user := storage.User{
		FirstName: "Pranto",
		LastName:  "Halder",
		Email:     "pranto@gmail.com",
		Username:  "pranto",
		Password:  "12345678",
		Role:      "doctor",
	}
	doctor, err := s.RegisterDoctorAdmin(user)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	dtype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	DType, err := s.Registerdoctortype(dtype)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	ddetails := storage.DoctorDetails{
		UserID:       doctor.ID,
		DoctorTypeID: DType.ID,
		Degree:       "MBBS",
		Gender:       "Male",
	}
	_, err = s.RegisterDoctorDeatils(ddetails)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int
		want    *storage.DoctorDetailsList
		wantErr bool
	}{
		{
			name: "DOCTOR_DETAILS_LIST_SUCCESS",
			in:   doctor.ID,
			want: &storage.DoctorDetailsList{
				FirstName:  "Pranto",
				LastName:   "Halder",
				DoctorType: "Nurologist",
				Degree:     "MBBS",
				Gender:     "Male",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in = doctor.ID
			got, err := s.ListDoctorDetails(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.ListDoctorDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.DoctorDetailsList{}, "ID"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDoctorScheduleListQuery(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	user := storage.User{
		FirstName: "Pranto",
		LastName:  "Halder",
		Email:     "pranto@gmail.com",
		Username:  "pranto",
		Password:  "12345678",
		Role:      "doctor",
	}
	doctor, err := s.RegisterDoctorAdmin(user)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	dtype := storage.DoctorType{
		DoctorType: "Nurologist",
	}
	DType, err := s.Registerdoctortype(dtype)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	ddetails := storage.DoctorDetails{
		UserID:       doctor.ID,
		DoctorTypeID: DType.ID,
		Degree:       "MBBS",
		Gender:       "Male",
	}
	Dd, err := s.RegisterDoctorDeatils(ddetails)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	workdays := storage.Schedule{
		WorkDays: "Friday",
	}
	workday, err := json.Marshal(workdays.WorkDays)
	if err != nil {
		fmt.Printf("#%v", err)
	}
	newschedule := []storage.Schedule{
		{ID: 1,
			DoctorDetailsID: Dd.ID,
			StartAt:         time.Time{},
			EndAt:           time.Time{},
			WorkDays:        string(workday),
			Address:         "Khulna",
			Phone:           "01716504535",
		},
	}
	for _, value := range newschedule {
		_, err := s.RegisterDoctorSchedule(value)
		if err != nil {
			t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
		}
	}
	tests := []struct {
		name    string
		in      int
		want    []storage.Schedule
		wantErr bool
	}{
		{
			name:    "DOCTOR_SCHEDULE_LIST_SUCCESS",
			in:      Dd.ID,
			want:    newschedule,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in = Dd.ID
			got, err := s.DoctorScheduleList(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.DoctorScheduleListQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Schedule{}, "DoctorDetailsID", "CreatedAt", "UpdatedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestEditDoctorStatus(t *testing.T) {
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
		Role:      "doctor",
	}
	user, err := s.RegisterDoctorAdmin(newuser)
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
			got, err := s.EditDoctorStatus(tt.in)
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

func TestUpdateDoctorStatus(t *testing.T) {
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
		Role:      "doctor",
	}
	_, err := s.RegisterDoctorAdmin(newuser)
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
			got, err := s.UpdateDoctorStatus(tt.in)
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