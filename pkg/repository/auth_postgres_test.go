package repository

import (
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestAuthPostgres_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthPostgres(db)

	tests := []struct {
		name    string
		mock    func()
		input   fhub.User
		want    int
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO users").
					WithArgs("Test", "test", "password").WillReturnRows(rows)
			},
			input: fhub.User{
				Username: "test",
				Password: "password",
			},
			want: 0,
		},
		{
			name: "Empty Fields",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO users").
					WithArgs("Test", "test", "").WillReturnRows(rows)
			},
			input: fhub.User{
				Username: "test",
				Password: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.CreateUser(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestAuthPostgres_GetUser(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewAuthPostgres(db)

	type args struct {
		username string
		password string
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		want    fhub.User
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "username", "password"}).
					AddRow(1, "Test", "test", "password")
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("test", "password").WillReturnRows(rows)
			},
			input: args{"test", "password"},
			want: fhub.User{
				Id:       1,
				Username: "test",
				Password: "password",
			},
			wantErr: true,
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "username", "password"})
				mock.ExpectQuery("SELECT (.+) FROM users").
					WithArgs("not", "found").WillReturnRows(rows)
			},
			input:   args{"not", "found"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetUser(tt.input.username, tt.input.password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
