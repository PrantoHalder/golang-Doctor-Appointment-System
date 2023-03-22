package postgres

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"main.go/usermgm/storage"
)

func TestPostGressStorage_Login(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *storage.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostGressStorage{
				DB: tt.fields.DB,
			}
			got, err := s.Login(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostGressStorage.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
