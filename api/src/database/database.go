package database

import (
	"api/src/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	if err := CriarBancoSeNecessario(db); err != nil {
		log.Fatal(err)
	}

	if err := CriarTabelas(db); err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func CriarBancoSeNecessario(db *sql.DB) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS devbook")
	if err != nil {
		return fmt.Errorf("erro ao criar banco: %w", err)
	}

	_, err = db.Exec("USE devbook")
	if err != nil {
		return fmt.Errorf("erro ao usar banco: %w", err)
	}

	return nil
}

func CriarTabelas(db *sql.DB) error {
	comandos := []string{
		`CREATE TABLE IF NOT EXISTS USUARIOS (
			id INT AUTO_INCREMENT PRIMARY KEY,
			nome VARCHAR(50) NOT NULL,
			nick VARCHAR(50) NOT NULL UNIQUE,
			email VARCHAR(50) NOT NULL UNIQUE,
			senha VARCHAR(255) NOT NULL,
			criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		) ENGINE=INNODB;`,

		`CREATE TABLE IF NOT EXISTS SEGUIDORES (
			usuario_id INT NOT NULL,
			seguidor_id INT NOT NULL,
			PRIMARY KEY (usuario_id, seguidor_id),
			FOREIGN KEY (usuario_id) REFERENCES USUARIOS(id) ON DELETE CASCADE,
			FOREIGN KEY (seguidor_id) REFERENCES USUARIOS(id) ON DELETE CASCADE
		) ENGINE=INNODB;`,

		`CREATE TABLE IF NOT EXISTS PUBLICACOES (
			id INT AUTO_INCREMENT PRIMARY KEY,
			titulo VARCHAR(50) NOT NULL,
			conteudo VARCHAR(300) NOT NULL,
			autor_id INT NOT NULL,
			criadaEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (autor_id) REFERENCES USUARIOS(id) ON DELETE CASCADE
		) ENGINE=INNODB;`,

		`CREATE TABLE IF NOT EXISTS likes (
			user_id INT NOT NULL,
			post_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (user_id, post_id),
			FOREIGN KEY (user_id) REFERENCES USUARIOS(id) ON DELETE CASCADE,
			FOREIGN KEY (post_id) REFERENCES PUBLICACOES(id) ON DELETE CASCADE
		) ENGINE=INNODB;`,
	}

	for _, comando := range comandos {
		if _, err := db.Exec(comando); err != nil {
			return fmt.Errorf("erro ao criar tabela: %w", err)
		}
	}

	log.Println("Tabelas criadas ou já existentes.")
	return nil
}
