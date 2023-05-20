package repository

import (
	"database/sql"
	"errors"
	"fmt"
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestFilePostgres_CreateFile(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewFilePostgres(db)

	type args struct {
		folderId int
		file     fhub.File
	}
	type mockBehavior func(args args, id int)

	tests := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		id           int
		wantErr      bool
	}{
		{
			name: "Ok",
			args: args{
				folderId: 1,
				file: fhub.File{
					FileName: "file",
					FileSize: 300,
					FilePath: "path",
				},
			},
			id: 0,
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				fmt.Println(rows)
				mock.ExpectQuery("INSERT INTO files").
					WithArgs(args.file.FileName, args.file.FileSize, args.file.FilePath).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO folders_files").WithArgs(args.folderId, id).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "Empty Fields",
			args: args{
				folderId: 1,
				file: fhub.File{
					FileName: "",
					FilePath: "/dir/file",
				},
			},
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id).RowError(0, errors.New("insert error"))
				mock.ExpectQuery("INSERT INTO files").
					WithArgs(args.file.FileName, args.file.FilePath).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "Failed 2nd Insert",
			args: args{
				folderId: 1,
				file: fhub.File{
					FileName: "title",
					FilePath: "/dir/file",
				},
			},
			mockBehavior: func(args args, id int) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO files").
					WithArgs(args.file.FileName, args.file.FilePath).WillReturnRows(rows)

				mock.ExpectExec("INSERT INTO folders_files").WithArgs(id, args.folderId).
					WillReturnError(errors.New("insert error"))

				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.id)

			got, err := r.CreateFile(tt.args.folderId, tt.args.file)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.id, got)
			}
		})
	}
}

func TestFilePostgres_GetAllFiles(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewFilePostgres(db)

	type args struct {
		folderId int
		userId   int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    []fhub.File
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "file_name", "file_date", "file_size", "file_path"}).
					AddRow(1, "title1", "", 0, "path/one").
					AddRow(2, "title2", "", 0, "path/two").
					AddRow(3, "title3", "", 0, "path/three")

				mock.ExpectQuery(
					"SELECT (.+) FROM files fi "+
						"INNER JOIN folders_files ff on (.+) "+
						"INNER JOIN users_folders uf on (.+) WHERE (.+)").
					WithArgs(1, 1).WillReturnRows(rows)
			},
			input: args{
				folderId: 1,
				userId:   1,
			},
		},
		{
			name: "No Records",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "file_name", "file_path"})

				mock.ExpectQuery("SELECT (.+) FROM files fi "+
					"INNER JOIN folders_files ff on (.+) "+
					"INNER JOIN users_folders uf on (.+) WHERE (.+)").
					WithArgs(1, 1).WillReturnRows(rows)
			},
			input: args{
				folderId: 1,
				userId:   1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetAllFiles(tt.input.userId, tt.input.folderId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestFilePostgres_GetFileById(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewFilePostgres(db)

	type args struct {
		fileId int
		userId int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    fhub.File
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "file_name", "file_date", "file_size", "file_path"}).
					AddRow(1, "title1", "", 0, "path/one")

				mock.ExpectQuery("SELECT (.+) FROM files fi "+
					"INNER JOIN folders_files ff on (.+) "+
					"INNER JOIN users_folders uf on (.+) WHERE (.+)").
					WithArgs(1, 1).WillReturnRows(rows)
			},
			input: args{
				fileId: 1,
				userId: 1,
			},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "file_name", "file_date", "file_size", "file_path"})

				mock.ExpectQuery("SELECT (.+) FROM files fi "+
					"INNER JOIN folders_files ff on (.+) "+
					"INNER JOIN users_folders uf on (.+) WHERE (.+)").
					WithArgs(404, 1).WillReturnRows(rows)
			},
			input: args{
				fileId: 404,
				userId: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetFileById(tt.input.userId, tt.input.fileId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestFilePostgres_DeleteFile(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := NewFilePostgres(db)

	type args struct {
		fileId int
		userId int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mock.ExpectExec("DELETE FROM files fi USING folders_files ff, users_folders uf WHERE (.+)").
					WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: true,
			input: args{
				fileId: 1,
				userId: 1,
			},
		},
		{
			name: "Not Found",
			mock: func() {
				mock.ExpectExec("DELETE FROM files fi USING folders_files ff, users_folders uf WHERE (.+)").
					WithArgs(1, 404).WillReturnError(sql.ErrNoRows)
			},
			input: args{
				fileId: 404,
				userId: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.DeleteFile(tt.input.userId, tt.input.fileId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
