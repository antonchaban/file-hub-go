package repository

import (
	"fmt"
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/jmoiron/sqlx"
)

type FilePostgres struct {
	db *sqlx.DB
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
