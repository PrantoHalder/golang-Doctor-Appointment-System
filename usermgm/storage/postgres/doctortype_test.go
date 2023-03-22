package postgres

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"main.go/usermgm/storage"
)

func TestPostGressStorage_Registerdoctortype(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		u storage.DoctorType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *storage.DoctorType
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostGressStorage{
				DB: tt.fields.DB,
			}
			got, err := s.Registerdoctortype(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.Registerdoctortype() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostGressStorage.Registerdoctortype() = %v, want %v", got, tt.want)
			}
		})
	}
}
