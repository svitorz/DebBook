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

func (p publicacoes) FindById(publicacaoID uint64) (models.Publicacao, error) {
	rows, err := p.db.Query(`SELECT p.*, u.nick FROM PUBLICACOES p INNER JOIN USUARIOS u ON u.id=p.autor_id WHERE p.id = ?`, publicacaoID)
	if err != nil {
		return models.Publicacao{}, err
	}

	defer rows.Close()

	var post models.Publicacao

	if rows.Next() {
		if err = rows.Scan(
			&post.ID,
			&post.Titulo,
			&post.Conteudo,
			&post.AutorID,
			&post.Curtidas,
			&post.CriadaEm,
			&post.AutorNick,
		); err != nil {
			return models.Publicacao{}, err
		}
	}
	return post, nil
}

func (p publicacoes) Index(userId uint64) ([]models.Publicacao, error) {
	rows, err := p.db.Query(`SELECT DISTINCT p.*, u.nick 
				FROM PUBLICACOES p 
				INNER JOIN USUARIOS u ON u.id = p.autor_id 
				LEFT JOIN SEGUIDORES s ON p.autor_id = s.usuario_id 
				WHERE u.id = ? OR s.seguidor_id = ?
				ORDER BY p.id DESC;`, userId, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var publicacoes []models.Publicacao

	for rows.Next() {
		var publicacao models.Publicacao

		if err = rows.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}

func (p publicacoes) Update(publicacaoId uint64, publicacao models.Publicacao) error {
	statement, err := p.db.Prepare("UPDATE PUBLICACOES SET titulo = ?, conteudo = ? where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoId); err != nil {
		return err
	}

	return nil
}

func (p publicacoes) Destroy(publicacaoId uint64) error {
	statement, err := p.db.Prepare("DELETE FROM PUBLICACOES where id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publicacaoId); err != nil {
		return err
	}

	return nil
}

func (p publicacoes) FindByUserId(userId uint64) (models.Publicacao, error) {
	rows, err := p.db.Query(`SELECT p.*, u.nick FROM PUBLICACOES p INNER JOIN USUARIOS u ON u.id=p.autor_id WHERE p.autor_id = ?`, userId)
	if err != nil {
		return models.Publicacao{}, err
	}

	defer rows.Close()

	var post models.Publicacao

	if rows.Next() {
		if err = rows.Scan(
			&post.ID,
			&post.Titulo,
			&post.Conteudo,
			&post.AutorID,
			&post.Curtidas,
			&post.CriadaEm,
			&post.AutorNick,
		); err != nil {
			return models.Publicacao{}, err
		}
	}
	return post, nil
}

func (p publicacoes) LikePost(userId uint64, postId uint64) error {
	stmt, err := p.db.Prepare(`INSERT INTO likes (user_id, post_id) VALUES (?, ?)`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(userId, postId); err != nil {
		return err
	}
	return nil
}

func (p publicacoes) UnlikePost(userId uint64, postId uint64) error {
	stmt, err := p.db.Prepare(`DELETE FROM likes WHERE user_id = ? AND post_id = ?`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(userId, postId); err != nil {
		return err
	}
	return nil
}
