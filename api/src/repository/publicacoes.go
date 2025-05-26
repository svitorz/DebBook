package repository

import (
	"api/src/models"
	"database/sql"
)

type publicacoes struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *publicacoes {
	return &publicacoes{db}
}

func (p publicacoes) Store(publicacao models.Publicacao) (uint64, error) {
	stmt, err := p.db.Prepare("INSERT INTO PUBLICACOES(titulo, conteudo, autor_id) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}
