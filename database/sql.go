package sql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb" //bblablalba
)

// //GetAll ...
// func (s *Str) GetAll() []*Item {
// 	rst := make([]*Item, 0)
// 	query := `SELECT A.DESCRICAO_PRODUTO,A.MARCA,A.COLECAO,A.SUBCOLECAO,A.CLASSIFICACAO,A.SUBCLASSIFICACAO,B.TAMANHOS,C.*,A.FOTO,A.DATA FROM LINX_TBFG..PRODUTOS A
// 	LEFT JOIN LINX_TBFG..[GRADES] B ON A.ID_GRADE = B.DESC_GRADE
// 	LEFT JOIN LINX_TBFG..[GRADE_PROPORCAO] C ON A.ID = C.ID_PRODUTO`
// 	rows, err := s.db.QueryContext(context.Background(), query, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	for rows.Next() {
// 		Item := Item{}
// 		if err := rows.Scan(&Item.Descricao, &Item.Marca, &Item.Colecao, &Item.Subcolecao, &Item.Classificacao, &Item.Subclassificacao); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		rst = append(rst, &Item)
// 	}

// 	return rst
// } // precisa terminar entender como trazer dois [] cor e proporção

// //GetAllColecao ...
// func (s *Str) GetAllColecao() []string {
// 	rst := make([]string, 0)

// 	query := `SELECT DESC_COLECAO FROM LINX_TBFG..COLECAO`
// 	rows, err := s.db.QueryContext(context.Background(), query, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}

// 	for rows.Next() {
// 		var colecao string
// 		if err := rows.Scan(&colecao); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		rst = append(rst, colecao)
// 	}

// 	return rst
// }

// //SubColecao ...
// func (s *Str) SubColecao() []string {
// 	rst := make([]string, 0)
// 	query := ``
// 	rows, err := s.db.QueryContext(context.Background(), query, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	for rows.Next() {
// 		var subcolecao string
// 		if err := rows.Scan(&subcolecao); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		rst = append(rst, subcolecao)
// 	}

// 	return rst
// }

// //GetAllGrade ...
// func (s *Str) GetAllGrade() []string {
// 	rst := make([]string, 0)
// 	query := `SELECT DESC_GRADE FROM LINX_TBFG..GRADES`
// 	rows, err := s.db.QueryContext(context.Background(), query, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	for rows.Next() {
// 		var grade string
// 		if err := rows.Scan(&grade); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		rst = append(rst, grade)
// 	}

// 	return rst
// }

// //GetAllClassificacao ...
// func (s *Str) GetAllClassificacao() []string {
// 	rst := make([]string, 0)
// 	query := `SELECT DESC_CLASSIFICACAO FROM LINX_TBFG..CLASSIFICACAO`
// 	rows, err := s.db.QueryContext(context.Background(), query, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	for rows.Next() {
// 		var classificacao string
// 		if err := rows.Scan(&classificacao); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		rst = append(rst, classificacao)
// 	}

// 	return rst
// }

// //SubClassificacao ...
// func (s *Str) SubClassificacao() []string {
// 	rst := make([]string, 0)
// 	query := `SELECT DESC_GRADE FROM LINX_TBFG..GRADES`
// 	rows, err := s.db.QueryContext(context.Background(), query, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	for rows.Next() {
// 		var subclassificacao string
// 		if err := rows.Scan(&subclassificacao); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		rst = append(rst, subclassificacao)
// 	}

// 	return rst
// }

// //GetItem ...
// func (s *Str) GetItem(cod *CodItem) []*Item {
// 	rst := make([]*Item, 0)
// 	query := fmt.Sprintf(`SELECT A.DESCRICAO_PRODUTO,A.MARCA,A.COLECAO,A.SUBCOLECAO,A.CLASSIFICACAO,A.SUBCLASSIFICACAO,B.TAMANHOS,C.*,A.FOTO,A.DATA FROM LINX_TBFG..PRODUTOS A
// 	LEFT JOIN LINX_TBFG..[GRADES] B ON A.ID_GRADE = B.DESC_GRADE
// 	LEFT JOIN LINX_TBFG..[GRADE_PROPORCAO] C ON A.ID = C.ID_PRODUTO
// 	WHERE A.PRODUTO = '`, cod, `'`)
// 	rows, err := s.db.QueryContext(context.Background(), query, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	for rows.Next() {
// 		Item := Item{}
// 		if err := rows.Scan("detalhes item"); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		rst = append(rst, &Item)
// 	}

// 	return rst
// }

// //NewItem ...
// func (s *Str) NewItem(newItem *Item) {
// }

// //Marca ...
// func (s *Str) Marca() []string {
// 	rst := make([]string, 0)
// 	query := `SELECT MARCA FROM LINX_TBFG..MARCA`
// 	rows, err := s.db.QueryContext(context.Background(), query, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	for rows.Next() {
// 		var marca string
// 		if err := rows.Scan(&marca); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		rst = append(rst, marca)
// 	}

// 	return rst
// }

//MakeSQL ...
func MakeSQL(host, port, username, password string) (*Str, error) {
	s := &Str{}
	s.url = &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     fmt.Sprintf("%s:%s", host, port),
		RawQuery: url.Values{}.Encode(),
	}
	return s, s.connect()
}

//Ping ...
func (s *Str) Ping() error {
	return s.db.Ping()
}

func (s *Str) connect() error {
	var err error
	if s.db, err = sql.Open("sqlserver", s.url.String()); err != nil {
		return err
	}
	return s.db.PingContext(context.Background())
}

// Str ...
type Str struct {
	url *url.URL
	db  *sql.DB
}
