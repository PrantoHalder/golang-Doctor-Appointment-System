package postgres

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"main.go/usermgm/storage"
)

func TestPostGressStorage_RegisterDoctorDeatils(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		u storage.Doctor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *storage.Doctor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PostGressStorage{
				DB: tt.fields.DB,
			}
			got, err := s.RegisterDoctorDeatils(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostGressStorage.RegisterDoctorDeatils() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostGressStorage.RegisterDoctorDeatils() = %v, want %v", got, tt.want)
			}
		})
	}
}
