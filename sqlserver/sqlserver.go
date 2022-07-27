package sqlserver

import (
	config "MANCHESTER/API-GLOBAL-LINX/config"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb" //bblablalba
)

//SQLStr ...
type SQLStr struct {
	conf *config.SQL
	url  *url.URL
	db   *sql.DB
}

//NewSQL ...
func NewSQL(conf *config.SQL) (*SQLStr, error) {
	s := &SQLStr{}
	return s, s.UpdateConfig(conf)
}

//UpdateConfig ...
func (s *SQLStr) UpdateConfig(conf *config.SQL) error {
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			return err
		}
	}
	s.conf = conf
	s.url = &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(conf.Username, conf.Password),
		Host:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		RawQuery: url.Values{"database": {conf.Db}}.Encode(),
	}
	return s.Connect()
}

//Connect ...
func (s *SQLStr) Connect() error {
	var err error
	s.db, err = sql.Open("sqlserver", s.url.String())
	if err != nil {
		return err
	}
	return s.db.Ping()
}

//QueryRow ...
func (s *SQLStr) QueryRow(format string, parameters ...interface{}) (*sql.Rows, error) {
	q := fmt.Sprintf(format, parameters...)
	// fmt.Println(q)
	rows, err := s.db.Query(q)
	if err == nil {
		return rows, nil
	}
	if err := s.db.Ping(); err != nil {
		return nil, err
	}
	return s.db.Query(q)
}

// func (s *sqlStr) getStatementConcat(from, to time.Time) (map[string]int, error) {
// 	rows, err := s.db.Query(fmt.Sprintf("SELECT CONCAT(BANCO,AGENCIA,CONTA,DATE_FORMAT(DATA, '%%Y-%%m-%%d'),`DESC`,VALOR) FROM EXTRATO_BANCARIO WHERE DATA>='%s' AND DATA<='%s'", from.Format("2006-01-02"), to.Format("2006-01-02")))
// 	if err != nil {
// 		return nil, err
// 	}
// 	var value string
// 	m := make(map[string]int)
// 	for rows.Next() {
// 		rows.Scan(&value)
// 		m[value]++
// 	}
// 	return m, nil
// }

// func (s *sqlStr) insertCards(m map[int]trans) error {

// 	codes := make([]string, 0, len(m))
// 	for k := range m {
// 		codes = append(codes, strconv.Itoa(k))
// 	}

// 	rows, err := s.db.Query(fmt.Sprintf("SELECT NSU FROM VENDAS_CARTAO WHERE NSU IN ('%s')", strings.Join(codes, "', '")))
// 	if err != nil {
// 		return err
// 	}

// 	updates := map[int]trans{}
// 	for rows.Next() {
// 		var nsu int
// 		rows.Scan(&nsu)
// 		if trans, ok := m[nsu]; ok {
// 			updates[nsu] = trans
// 			delete(m, nsu)
// 		}
// 	}

// 	log.Printf("%d inserted, %d updated", len(m), len(updates))

// 	queryFormat := "INSERT INTO VENDAS_CARTAO (NSU,DATA,VALOR_BRUTO,PARCELAS,SE_CREDITO,BANCEIRA,MDR,DESCONTO,VALOR_LIQ,LOTE,AUTORIZACAO,LOJA_COD,LOJA_NOME,CNPJ,CARTAO_ID,CARTEIRA_ID,CAPTURA_TIPO,TERMINAL_TIPO,TERMINAL_COD,TID,PEDIDO_COD,TAXA_EMBARQUE,CANCELAMENTO_LOJA,CANCELAMENTO_DATA,CANCELAMENTO_VALOR) VALUES %s"
// 	values := make([]string, 0, len(m))
// 	for _, v := range m {
// 		values = append(values, v.sqlMarshal())
// 	}

// 	if len(m) > 0 {
// 		row := s.db.QueryRow(fmt.Sprintf(queryFormat, strings.Join(values, ",\n")))
// 		if err := row.Err(); err != nil {
// 			return err
// 		}
// 	}

// 	updateQuery := "UPDATE VENDAS_CARTAO SET %s WHERE NSU=%d"
// 	// updateValues := make([]string, 0, len(m))
// 	for k, v := range updates {
// 		// updateValues = append(values, v.sqlMarshal())
// 		row := s.db.QueryRow(fmt.Sprintf(updateQuery, v.sqlUpdateMarshal(), k))
// 		if err := row.Err(); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
