package repository

import (
	"api/src/models"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (u usuarios) Store(usuario models.Usuario) (uint64, error) {
	statement, err := u.db.Prepare(
		"insert into USUARIOS (nome, nick, email, senha) values (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
