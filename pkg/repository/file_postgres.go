package repository

import (
	"fmt"
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/jmoiron/sqlx"
	"strings"
)

type FilePostgres struct {
	db *sqlx.DB
}

type File interface {
	CreateFile(folderId int, file fhub.File) (int, error)
	GetAllFiles(userId, folderId int) ([]fhub.File, error)
	GetFileById(userId, fileId int) (fhub.File, error)
	DeleteFile(userId, fileId int) error
	UpdateFile(userId, fileId int, input fhub.UpdateFileInput) error
}

func NewFilePostgres(db *sqlx.DB) *FilePostgres {
	return &FilePostgres{db: db}
}

func (r *FilePostgres) CreateFile(folderId int, file fhub.File) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var fileId int
	createFileQuery := fmt.Sprintf(`insert into %s (file_name, file_size, file_path) values ($1, $2, $3) returning id`, filesTable)
	row := tx.QueryRow(createFileQuery, file.FileName, file.FileSize, file.FilePath)
	err = row.Scan(&fileId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createFileFolderQuery := fmt.Sprintf(`insert into %s (file_id, folder_id) values ($1, $2)`, foldersFilesTable)
	_, err = tx.Exec(createFileFolderQuery, fileId, folderId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return fileId, tx.Commit()
}

func (r *FilePostgres) GetAllFiles(userId, folderId int) ([]fhub.File, error) {
	var files []fhub.File

	query := fmt.Sprintf(`select fit.id, fit.file_name, fit.file_date, fit.file_size, fit.file_path from %s fit inner join %s fft on fft.file_id = fit.id
					inner join %s uft on uft.folder_id = fft.folder_id where fft.folder_id = $1 and uft.user_id = $2`,
		filesTable, foldersFilesTable, usersFoldersTable)
	if err := r.db.Select(&files, query, folderId, userId); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *FilePostgres) GetFileById(userId, fileId int) (fhub.File, error) {
	var file fhub.File

	query := fmt.Sprintf(`select fit.id, fit.file_name, fit.file_date, fit.file_size, fit.file_path from %s fit inner join %s fft on fft.file_id = fit.id
					inner join %s uft on uft.folder_id = fft.folder_id where fit.id = $1 and uft.user_id = $2`,
		filesTable, foldersFilesTable, usersFoldersTable)
	if err := r.db.Get(&file, query, fileId, userId); err != nil {
		return file, err
	}

	return file, nil
}

func (r *FilePostgres) DeleteFile(userId, fileId int) error {
	query := fmt.Sprintf(`DELETE FROM %s fit USING %s fft, %s uft 
									WHERE fit.id = fft.file_id AND fft.folder_id = uft.folder_id AND uft.user_id = $1 AND fit.id = $2`,
		filesTable, foldersFilesTable, usersFoldersTable)
	_, err := r.db.Exec(query, userId, fileId)
	return err
}

func (r *FilePostgres) UpdateFile(userId, fileId int, input fhub.UpdateFileInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.FileName != nil {
		setValues = append(setValues, fmt.Sprintf("file_name=$%d", argId))
		args = append(args, *input.FileName)
		argId++
	}

	if input.FilePath != nil {
		setValues = append(setValues, fmt.Sprintf("file_path=$%d", argId))
		args = append(args, *input.FilePath)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s fit SET %s FROM %s fft, %s uft
									WHERE fit.id = fft.file_id AND fft.folder_id = uft.folder_id AND uft.user_id = $%d AND fit.id = $%d`,
		filesTable, setQuery, foldersFilesTable, usersFoldersTable, argId, argId+1)
	args = append(args, userId, fileId)

	_, err := r.db.Exec(query, args...)
	return err
}
