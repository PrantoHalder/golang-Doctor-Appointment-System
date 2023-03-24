package postgres

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"main.go/usermgm/storage"
)

func TestRegisterAdmin(t *testing.T) {
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
			name: "REGISTER_ADMIN_SUCCESS",
			in: storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karin@gamil.com",
				Role:      "admin",
				Username:  "karin",
				Password:  "12345678",
			},
			want: &storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karin@gamil.com",
				Username:  "karin",
				Password:  "12345678",
				Role:      "admin",
				Is_active: true,
			},
			wantErr: false,
		},
		{
			name: "REGISTER_ADMIN_EMAIL_UNIQUE_FAILURE",
			in: storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karin@gamil.com",
				Role:      "admin",
				Username:  "karinn",
				Password:  "12345678",
			},
			wantErr: true,
		},
		{
			name: "REGISTER_ADMIN_USERNAME_UNIQUE_FAILURE",
			in: storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karinn@gamil.com",
				Role:      "admin",
				Username:  "karin",
				Password:  "12345678",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.RegisterAdmin(tt.in)
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

func TestRegisterDoctorAdmin(t *testing.T) {
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
			name: "REGISTER_ADMIN_DOCTOR_SUCCESS",
			in: storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karin@gamil.com",
				Role:      "doctor",
				Username:  "karin",
				Password:  "12345678",
			},
			want: &storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karin@gamil.com",
				Username:  "karin",
				Password:  "12345678",
				Role:      "doctor",
				Is_active: true,
			},
			wantErr: false,
		},
		{
			name: "REGISTER_ADMIN_DOCTOR_EMAIL_UNIQUE_FAILURE",
			in: storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karin@gamil.com",
				Role:      "doctor",
				Username:  "karinn",
				Password:  "12345678",
			},
			wantErr: true,
		},
		{
			name: "REGISTER_ADMIN_DOCTOR_USERNAME_UNIQUE_FAILURE",
			in: storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karinn@gamil.com",
				Role:      "doctor",
				Username:  "karin",
				Password:  "12345678",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.RegisterDoctorAdmin(tt.in)
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

func TestRegisterPatient(t *testing.T) {
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
			name: "REGISTER_ADMIN_PATIENT_SUCCESS",
			in: storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karin@gamil.com",
				Role:      "user",
				Username:  "karin",
				Password:  "12345678",
			},
			want: &storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karin@gamil.com",
				Username:  "karin",
				Password:  "12345678",
				Role:      "user",
				Is_active: true,
			},
			wantErr: false,
		},
		{
			name: "REGISTER_ADMIN_PATIENT_EMAIL_UNIQUE_FAILURE",
			in: storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karin@gamil.com",
				Role:      "user",
				Username:  "karinn",
				Password:  "12345678",
			},
			wantErr: true,
		},
		{
			name: "REGISTER_ADMIN_PATIENT_USERNAME_UNIQUE_FAILURE",
			in: storage.User{
				FirstName: "Karin",
				LastName:  "Khatun",
				Email:     "karinn@gamil.com",
				Role:      "user",
				Username:  "karin",
				Password:  "12345678",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.RegisterPatient(tt.in)
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

func TestEditAdmin(t *testing.T) {
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
		Role:      "admin",
	}
	user, err := s.RegisterAdmin(newuser)
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
			tt.in = user.ID
			got, err := s.EditAdmin(tt.in)
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

func TestUpdateAdmin(t *testing.T) {
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
		Role:      "admin",
	}
	user, err := s.RegisterAdmin(newuser)
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

func TestDeleteAdminByID(t *testing.T) {
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
		Role:      "admin",
	}
	user, err := s.RegisterAdmin(newuser)
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
			in:      user.ID,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteAdminByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEditDoctor(t *testing.T) {
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
			got, err := s.EditDoctor(tt.in)
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

func TestUpdateDoctor(t *testing.T) {
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
			got, err := s.UpdateDoctor(tt.in)
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

func TestDeleteDoctorByID(t *testing.T) {
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
		Role:      "admin",
	}
	user, err := s.RegisterDoctorAdmin(newuser)
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
			err := s.DeleteDoctorByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEditPatient(t *testing.T) {
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
			got, err := s.EditPatient(tt.in)
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

func TestUpdatePatient(t *testing.T) {
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
			got, err := s.UpdatePatient(tt.in)
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

func TestDeletePatientByID(t *testing.T) {
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
		Role:      "admin",
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
			err := s.DeletePatientByID(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestListAdmin(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newAdmin := []storage.User{
		{
			FirstName: "Karim",
			LastName:  "Khan",
			Email:     "karim@gmail.com",
			Username:  "karin",
			Role:      "admin",
			Password:  "12345678",
		},
		{
			FirstName: "Rahim",
			LastName:  "shake",
			Email:     "rahim@gmail.com",
			Username:  "rahim",
			Role:      "admin",
			Password:  "12345678",
		}, {
			FirstName: "Pranto",
			LastName:  "Halder",
			Email:     "pranto@gmail.com",
			Username:  "pranto",
			Role:      "admin",
			Password:  "12345678",
		}, {
			FirstName: "Shovon",
			LastName:  "Halder",
			Email:     "shovon@gmail.com",
			Username:  "shovon",
			Role:      "admin",
			Password:  "12345678",
		},
	}
	for _, value := range newAdmin {
		_, err := s.RegisterAdmin(value)
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
			want:    newAdmin,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.ListAdmin(tt.in)
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