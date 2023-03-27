package postgres

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"main.go/usermgm/storage"
)

func TestRegisterdoctortype(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.DoctorType
		want    *storage.DoctorType
		wantErr bool
	}{
		{
			name: "REGISTER_DOCTOR_TYPE_SUCCESS",
			in: storage.DoctorType{
				DoctorType: "Neurologists",
			},
			want: &storage.DoctorType{
				DoctorType: "Neurologists",
			},
			wantErr: false,
		},
		{
			name: "REGISTER_DOCTOR_TYPE_UNIQUE_FAILURE",
			in: storage.DoctorType{
				DoctorType: "Neurologists",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Registerdoctortype(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.DoctorType{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestEditDoctorType(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newdoctortype := storage.DoctorType{
		ID:         1,
		DoctorType: "Nurologist",
	}
	_, err := s.Registerdoctortype(newdoctortype)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int
		want    *storage.DoctorType
		wantErr bool
	}{
		{
			name: "",
			in:   1,
			want: &storage.DoctorType{
				ID:         1,
				DoctorType: "Nurologist",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in = newdoctortype.ID
			got, err := s.EditDoctorType(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.EditDoctorType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.DoctorType{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateDoctorType(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newdoctortype := storage.DoctorType{
		ID:         1,
		DoctorType: "Nurologist",
	}
	_, err := s.Registerdoctortype(newdoctortype)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterDoctorSchedule() error = %v", err)
	}
	tests := []struct {
		name    string
		in      storage.DoctorType
		want    *storage.DoctorType
		wantErr bool
	}{
		{
			name:    "UPDATE_DOCTOR_TYPE_SUCCESS",
			in:      storage.DoctorType{
				ID:         1,
				DoctorType: "Child Spacialist",
			},
			want:    &storage.DoctorType{
				DoctorType: "Child Spacialist",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UpdateDoctorType(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.UpdateDoctorType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.DoctorType{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostGressStorage.Register() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}