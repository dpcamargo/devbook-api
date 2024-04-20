package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db: db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		`INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil

}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, err := repositorio.db.Query(`
	SELECT p.*, u.nick FROM
	publicacoes p INNER JOIN usuarios u
	on u.id = p.autor_id WHERE p.id = ?`,
		publicacaoID)
	if err != nil {
		return modelos.Publicacao{}, err
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if err = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return modelos.Publicacao{}, err
		}
	}
	return publicacao, nil
}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
	SELECT DISTINCT p.*, u.nick FROM publicacoes p
	INNER JOIN usuarios u on u.id = p.autor_id
	INNER JOIN seguidores s on p.autor_id = s.usuario_id
	WHERE u.id = ? OR s.seguidor_id = ?
	ORDER BY 1 DESC`, usuarioID, usuarioID)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao
		if err = linhas.Scan(
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

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, err := repositorio.db.Prepare(`
	UPDATE publicacoes SET titulo = ?, conteudo = ? WHERE id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); err != nil {
		return err
	}
	return nil
}

func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, err := repositorio.db.Prepare(`DELETE FROM publicacoes WHERE id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacaoID); err != nil {
		return err
	}

	return nil
}
