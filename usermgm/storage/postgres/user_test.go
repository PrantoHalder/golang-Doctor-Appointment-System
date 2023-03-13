package postgres

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"main.go/usermgm/storage"
)

func TestPostGressStorage_Register(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "CREATE_USER_SUCCESS",
			in: storage.User{
				FirstName: "first",
				LastName:  "last",
				Email:     "first@example.com",
				Username:  "user",
				Password:  "123456",
			},
			want: &storage.User{
				FirstName: "first",
				LastName:  "last",
				Email:     "first@example.com",
				Username:  "user",
				Status:    true,
			},
		},
		{
			name: "CREATE_USER_SUCCESS",
			in: storage.User{
				FirstName: "Rahim",
				LastName:  "Khan",
				Email:     "rahim@example.com",
				Username:  "rahim",
				Password:  "123456",
			},
			want: &storage.User{
				FirstName: "Rahim",
				LastName:  "Khan",
				Email:     "rahim@example.com",
				Username:  "rahim",
				Status:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.Register(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.CreateUser() error = got %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Role", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestPostGressStorage_GetStatusbyUsernameQueryOFUsers(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.Register(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.CreateUser() error = got %v, wantErr %v", err, tt.wantErr)
				return
			}

			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Role", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
