package models

import "time"

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autor_id,omitempty"`
	AutorNick uint64    `json:"autor_nick,omitempty"`
	Curtidas  uint64    `json:"curtidas`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}
