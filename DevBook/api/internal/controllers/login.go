package controllers

import (
	"api/internal/autenticacao"
	"api/internal/banco"
	"api/internal/modelo"
	"api/internal/repositorio"
	"api/internal/respostas"
	"api/internal/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var usuario modelo.Usuario
	if err = json.Unmarshal(requestBody, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Con()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositorio.NovoRepositoUsuarios(db)
	usuarioSalvoNoBanco, err := repo.BuscarPorEmail(usuario.Email)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	if err = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); err != nil {
		respostas.Erro(w, http.StatusUnauthorized, err)
		return
	}
	token, err := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
