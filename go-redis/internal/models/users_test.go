package models

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestUserModel_Insert(t *testing.T) {
	DB := newTestDB(t)

	type args struct {
		name string
		age  int
	}
	tests := []struct {
		name    string
		dbPool  *pgxpool.Pool
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Insert first user",
			dbPool:  DB,
			args:    args{"test name", 30},
			want:    1,
			wantErr: false,
		},
		{
			name:    "Insert second user",
			dbPool:  DB,
			args:    args{"test name 2", 30},
			want:    2,
			wantErr: false,
		},
		{
			name:    "Insert second user again",
			dbPool:  DB,
			args:    args{"test name 2", 30},
			want:    2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UserModel{
				DB: tt.dbPool,
			}
			got, err := m.Insert(tt.args.name, tt.args.age)
			if tt.wantErr {
				if err == nil {
					t.Errorf("UserModel.Insert() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				return
			}

			if got != tt.want {
				t.Errorf("UserModel.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestUserModel_Get(t *testing.T) {
	type fields struct {
		db *pgx.Conn
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UserModel{
				db: tt.fields.db,
			}
			got, err := m.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserModel.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
