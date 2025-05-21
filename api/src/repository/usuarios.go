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

func (u usuarios) ShowAll(nomeOuNick string) ([]models.Usuario, error) {
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

func (u usuarios) Show(ID uint64) (models.Usuario, error) {
	linhas, err := u.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM USUARIOS WHERE id = ?",
		ID,
	)
	if err != nil {
		return models.Usuario{}, nil
	}

	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (u usuarios) Update(ID uint64, usuario models.Usuario) error {
	statement, err := u.db.Prepare("UPDATE USUARIOS SET nome = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}

	return nil
}

func (u usuarios) Destroy(ID uint64) error {
	statement, err := u.db.Prepare("DELETE FROM USUARIOS WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (u usuarios) FindByEmail(email string) (models.Usuario, error) {
	row, err := u.db.Query("SELECT id, senha FROM USUARIOS WHERE email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}

	defer row.Close()

	var usuario models.Usuario

	if row.Next() {
		if err := row.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil
}

func (u usuarios) Follow(usuarioID, seguidorId uint64) error {
	statement, err := u.db.Prepare("INSERT IGNORE INTO SEGUIDORES(usuario_id, seguidor_id) VALUES (?,?)")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(usuarioID, seguidorId); err != nil {
		return err
	}

	return nil
}

func (u usuarios) Unfollow(usuarioID, seguidorId uint64) error {
	statement, err := u.db.Prepare("DELETE FROM SEGUIDORES WHERE usuario_id = ? AND seguidor_id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(usuarioID, seguidorId); err != nil {
		return err
	}

	return nil
}

func (u usuarios) GetFollowers(usuarioId uint64) ([]models.Usuario, error) {
	rows, err := u.db.Query(
		"SELECT u.id, u.nome, u.nick, u.email, u.criadoEm FROM USUARIOS u INNER JOIN SEGUIDORES s ON s.seguidor_id = u.id WHERE s.usuario_id = ?",
		usuarioId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var usuario models.Usuario

		if err = rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (u usuarios) GetFollowing(usuarioId uint64) ([]models.Usuario, error) {
	rows, err := u.db.Query(
		"SELECT u.id, u.nome, u.nick, u.email, u.criadoEm FROM USUARIOS u INNER JOIN SEGUIDORES s ON s.seguidor_id = u.id WHERE s.seguidor_id = ?",
		usuarioId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var usuarios []models.Usuario
	for rows.Next() {
		var usuario models.Usuario

		if err = rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
