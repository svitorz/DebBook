package models

import (
	"api/src/security"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem"`
}

func (usuario *Usuario) Preparar(action string) error {
	if err := usuario.validar(action); err != nil {
		return err
	}

	if err := usuario.formatar(action); err != nil {
		return err
	}

	return nil
}

func (usuario *Usuario) validar(action string) error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("o email é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		fmt.Println(usuario.Email)
		return errors.New(" e-mail inserido é inválido")
	}

	if action == "cadastrar" && usuario.Senha == "" {
		return errors.New("o senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastrar" {
		senhaComHash, err := security.Hash(usuario.Senha)
		if err != nil {
			return err
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}
