package postgres

import (
	"testing"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"main.go/usermgm/storage"
)

func TestLogin(t *testing.T) {
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
			Role:      "user",
		},
	}
	for _, value := range newuser {
		_, err := s.RegisterPatient(value)
		if err != nil {
			t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
		}
	}

	newdoctor := []storage.User{
		{
			FirstName: "Pranto",
			LastName:  "Halder",
			Email:     "pranto@gmail.com",
			Username:  "pranto",
			Password:  "12345678",
			Role:      "admin",
		},
	}
	for _, value := range newdoctor {
		_, err := s.RegisterDoctorAdmin(value)
		if err != nil {
			t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
		}
	}
	newadmin := []storage.User{
		{
			FirstName: "Shovon",
			LastName:  "Halder",
			Email:     "shovon@gmail.com",
			Username:  "shovon",
			Password:  "12345678",
			Role:      "doctor",
		},
	}
	for _, value := range newadmin {
		_, err := s.RegisterAdmin(value)
		if err != nil {
			t.Fatalf("PostgresStorage.RegisterPatient() error = %v", err)
		}
	}
	tests := []struct {
		name    string
		in      string
		want    *storage.User
		wantErr bool
	}{
		{
			name:    "ADMIN_LOGIN_SUCCESS",
			in:      "pranto",
			want:    &storage.User{
				FirstName: "Pranto",
				LastName:  "Halder",
				Email:     "pranto@gmail.com",
				Username:  "pranto",
				Role:      "admin",
				Is_active: true,
			},
			wantErr: false,
		},
		{
			name:    "DOCTOR_LOGIN_SUCCESS",
			in:      "shovon",
			want:    &storage.User{
				FirstName: "Shovon",
				LastName:  "Halder",
				Email:     "shovon@gmail.com",
				Username:  "shovon",
				Role:      "doctor",
				Is_active: true,
			},
			wantErr: false,
		},
		{
			name:    "PATIENT_LOGIN_SUCCESS",
			in:      "rahim",
			want:    &storage.User{
				FirstName: "Rahim",
				LastName:  "Hossain",
				Email:     "rahim@gmail.com",
				Username:  "rahim",
				Role:      "user",
				Is_active: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Login(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.AppinmentStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID","Password","CreatedAt","UpdatedAt","DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}