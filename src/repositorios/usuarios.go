package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um novo repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoIDInserido), nil
}

func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	statement, err := repositorio.db.Prepare(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE nome LIKE ? or nick LIKE ?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	linhas, err := statement.Query(nomeOuNick, nomeOuNick)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if err := linhas.Scan(
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

func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	statement, err := repositorio.db.Prepare(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?")
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer statement.Close()

	var usuario modelos.Usuario
	err = statement.QueryRow(fmt.Sprintf("%d", ID)).Scan(
		&usuario.ID,
		&usuario.Nome,
		&usuario.Nick,
		&usuario.Email,
		&usuario.CriadoEm,
	)
	if err != nil {
		return modelos.Usuario{}, err
	}

	return usuario, nil
}

func (repositorio Usuarios) Atualizar(usuarioID uint64, usuario modelos.Usuario) error {
	statement, err := repositorio.db.Prepare(
		"UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(
		usuario.Nome,
		usuario.Nick,
		usuario.Email,
		usuarioID,
	); err != nil {
		return err
	}
	return nil
}

func (repositorio Usuarios) Deletar(usuarioID uint64) error {
	statement, err := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioID); err != nil {
		return err
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	statement, err := repositorio.db.Prepare("SELECT id, senha FROM usuarios where email = ?")
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer statement.Close()

	usuarioBanco := statement.QueryRow(email)

	var usuario modelos.Usuario

	if err := usuarioBanco.Scan(&usuario.ID, &usuario.Senha); err != nil {
		return modelos.Usuario{}, err
	}

	return usuario, nil
}
