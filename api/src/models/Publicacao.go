package models

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`   // max 50
	Conteudo  string    `json:"conteudo,omitempty"` // max 300
	AutorID   uint64    `json:"autor_id,omitempty"`
	AutorNick uint64    `json:"autor_nick,omitempty"`
	Curtidas  uint64    `json:"curtidas`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {
	if err := publicacao.validate(); err != nil {
		return err
	}
	publicacao.format()

	return nil
}

func (publicacao *Publicacao) validate() error {
	if publicacao.Titulo == "" {
		return errors.New("o título da publicação não pode estar vazio")
	}

	if publicacao.Conteudo == "" {
		return errors.New("o conteúdo da publicação não pode estar vazio")
	}

	if len(publicacao.Titulo) > 50 {
		return errors.New("o título da publicação não pode ter mais que 50 caracteres")
	}

	if len(publicacao.Conteudo) > 300 {
		return errors.New("o conteúdo da publicação não pode ter mais que 300 caracteres")
	}
	return nil
}

func (publicacao *Publicacao) format() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
