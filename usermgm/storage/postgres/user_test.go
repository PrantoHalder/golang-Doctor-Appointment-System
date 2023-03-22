package postgres

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"main.go/usermgm/storage"
)

func TestRegister(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	
	tests := []struct {
		name    string
		in      storage.Patient
		want    *storage.Patient
		wantErr bool
	}{
		{
			name: "REGISTER_PATIENT_SUCCESS",
			in: storage.Patient{
				FirstName: "Rahim",
				LastName:  "Khan",
				Email:     "rahim@gamil.com",
				Username:  "rahim",
				Password:  "12345678",
			},
			want: &storage.Patient{
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
			in: storage.Patient{
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
			in: storage.Patient{
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
				cmpopts.IgnoreFields(storage.Patient{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want,opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
