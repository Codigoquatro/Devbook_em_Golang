package modelo

import (
	"api/internal/seguranca"
	"errors"
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
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}
	if err := usuario.formatar(etapa); err != nil {
		return err
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}
	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("O email inserido é inválido!")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (usario *Usuario) formatar(etapa string) error {
	usario.Nome = strings.TrimSpace(usario.Nome)
	usario.Nick = strings.TrimSpace(usario.Nick)
	usario.Email = strings.TrimSpace(usario.Email)

	if etapa == "cadastro" {
		senhaComHash, err := seguranca.Hash(usario.Senha)
		if err != nil {
			return err
		}

		usario.Senha = string(senhaComHash)
	}
	return nil
}
