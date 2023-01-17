package repositorio

import (
	"api/internal/modelo"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositoUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repo Usuarios) Criar(usuario modelo.Usuario) (uint64, error) {
	statement, err := repo.db.Prepare(
		"INSERT INTO usuarios(nome,nick,email,senha) VALUES (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}
	UltimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(UltimoIDInserido), nil
}

func (repo Usuarios) Buscar(nomeOuNick string) ([]modelo.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, err := repo.db.Query(
		"SELECT id,nome,nick,email,criadoEm FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []modelo.Usuario

	for linhas.Next() {
		var usuario modelo.Usuario
		if err = linhas.Scan(
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

func (repo Usuarios) BuscarPorID(ID uint64) (modelo.Usuario, error) {
	linhas, err := repo.db.Query(
		"SELECT id,nome,nick,email,criadoEm FROM usuarios WHERE id = ?",
		ID,
	)
	if err != nil {
		return modelo.Usuario{}, err
	}
	defer linhas.Close()

	var usuario modelo.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return modelo.Usuario{}, err
		}
	}
	return usuario, nil
}

func (repo Usuarios) AtualizarUsuario(ID uint64, usuario modelo.Usuario) error {
	statement, err := repo.db.Prepare(
		"UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		ID,
	); err != nil {
		return err
	}
	return nil

}

func (repo Usuarios) DeletarUsuario(ID uint64) error {

	statement, err := repo.db.Prepare(
		"DELETE FROM usuarios WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

func (repo Usuarios) BuscarPorEmail(email string) (modelo.Usuario, error) {
	linha, err := repo.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if err != nil {
		return modelo.Usuario{}, err
	}
	defer linha.Close()

	var usuario modelo.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return modelo.Usuario{}, err
		}
	}
	return usuario, nil
}
