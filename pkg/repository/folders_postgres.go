package repository

import (
	"fmt"
	todo "github.com/antonchaban/file-hub-go"
	"github.com/jmoiron/sqlx"
)

type FolderPostgres struct {
	db *sqlx.DB
}

func NewFolderPostgres(db *sqlx.DB) *FolderPostgres {
	return &FolderPostgres{db: db}
}

func (r *FolderPostgres) CreateFolder(userId int, folder todo.Folder) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createFolderQuery := fmt.Sprintf("INSERT INTO %s (folder_name) values ($1) RETURNING id", foldersTable)
	row := tx.QueryRow(createFolderQuery, folder.FolderName)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersFolderQuery := fmt.Sprintf("INSERT INTO %s (user_id, folder_id) values ($1, $2)", usersFoldersTable)
	_, err = tx.Exec(createUsersFolderQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *FolderPostgres) GetAllFolders(userId int) ([]todo.Folder, error) {
	var folders []todo.Folder
	query := fmt.Sprintf("select ft.id, ft.folder_name, ft.folderdate from %s ft inner join %s uft on ft.id = uft.folder_id where uft.user_id = $1",
		foldersTable, usersFoldersTable)
	err := r.db.Select(&folders, query, userId)
	return folders, err
}

func (r *FolderPostgres) GetById(userId, id int) (todo.Folder, error) {
	var folder todo.Folder
	query := fmt.Sprintf(
		"select ft.id, ft.folder_name, ft.folderdate from %s ft inner join %s uft on ft.id = uft.folder_id "+
			"where uft.user_id = $1 and ft.id = $2",
		foldersTable, usersFoldersTable)
	err := r.db.Get(&folder, query, userId, id)
	return folder, err

}
