package postgres

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"main.go/usermgm/storage"
)

func TestPostGressStorage_RegisterAdmin(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		u storage.Admin
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *storage.Admin
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostGressStorage{
				DB: tt.fields.DB,
			}
			got, err := s.RegisterAdmin(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.RegisterAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostGressStorage.RegisterAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}
