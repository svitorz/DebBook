package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (u usuarios) Show(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick&%

	rows, err := u.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM USUARIOS WHERE nome LIKE ? OR nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.Usuario

	for rows.Next() {
		var user models.Usuario

		if err = rows.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
