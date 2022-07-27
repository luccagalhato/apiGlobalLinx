package controller

import (
	"MANCHESTER/API-GLOBAL-LINX/config"
	"MANCHESTER/API-GLOBAL-LINX/server"
	"MANCHESTER/API-GLOBAL-LINX/sqlserver"

	"log"
	"net/http"
)

//NewController ...
func NewController(filePath string) (*Controller, error) {
	c := &Controller{}

	confChan := make(chan config.Config)

	firstRead := make(chan error)
	go func() {
		conf := <-confChan
		var err error
		l, err := sqlserver.NewSQL(&conf.Linx)
		if err != nil {
			firstRead <- err
			return
		}
		c.linx = (*linx)(l)
		if err := (*sqlserver.SQLStr)(c.linx).Connect(); err != nil {
			firstRead <- err
			return
		}
		// a, err := mysql.NewSQL(&conf.App)
		// if err != nil {
		// 	firstRead <- err
		// 	return
		// }

		// if err := (*mysql.SQLStr)(c.app).Connect(); err != nil {
		// 	firstRead <- err
		// 	return
		// }

		c.server = server.NewServer(conf, c.handlers())

		firstRead <- nil

		for conf := range confChan {
			if err := (*sqlserver.SQLStr)(c.linx).UpdateConfig(&conf.Linx); err != nil {
				log.Fatal(err)
			}
			// if err := (*mysql.SQLStr)(c.app).UpdateConfig(&conf.App); err != nil {
			// 	log.Fatal(err)
			// }

			c.server = server.NewServer(conf, c.handlers())
		}
	}()

	_, err := config.LoadYaml(filePath, func(conf config.Config) {
		confChan <- conf
	})
	if err != nil {
		return nil, err
	}
	if err := <-firstRead; err != nil {
		return nil, err
	}

	return c, nil
}

//Controller ...
type Controller struct {
	//conf config.Config
	linx *linx

	server *http.Server
}

type linx sqlserver.SQLStr

//ListenAndServe ...
func (c *Controller) ListenAndServe() error {
	return c.server.ListenAndServe()
}

// //Test ...
// func (c *Controller) Test() {
// 	json.NewEncoder(os.Stdout).Encode(c.cadastroPageData())
// }

func (c *Controller) handlers() map[string]server.Handler {
	return map[string]server.Handler{
		// "/cadastro_page": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.CadastroPage,
		// },
		// "/cadastro_produto": {
		// 	Method: http.MethodPost,
		// 	Fn:     c.CadastroProduto,
		// },
		// "/consulta_linx/{id}": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaLinx,
		// },
		// "/consulta_produtos": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaProdutos,
		// },
		// "/photos/{file}": {
		// 	Method: http.MethodGet,
		// 	Fn:     http.StripPrefix("/photos/", http.FileServer(http.Dir("photos/"))).ServeHTTP,
		// },
		// "/pedidos_page": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.PedidosPage,
		// },
		// "/cadastro_pedido": {
		// 	Method: http.MethodPost,
		// 	Fn:     c.CadastroPedido,
		// },
		// "/consulta_pedidos": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaPedidos,
		// },
		// "/consulta_produtos_ecom": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaProdutosEcom,
		// },
		// "/consulta_produtos_ecom_colecao": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaProdutosEcomColecao,
		// },
		// "/consulta_item_pedido": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaItemPedido,
		// },
		// "/consulta_produtos_look": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaProdutosEcom,
		// },
		// "/adicionar_item_pedido": {
		// 	Method: http.MethodPost,
		// 	Fn:     c.AdicionarItemPedido,
		// },

		// "/resumo_pedido": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ResumoPedido,
		// },
		// "/itemsCadastrado": {
		// 	http.MethodGet,
		// 	apihandler.Getall,
		// },
		// "/grade": {
		// 	http.MethodGet,
		// 	apihandler.GetAllGrade,
		// },
		// "/marca": {
		// 	http.MethodGet,
		// 	apihandler.GetMarca,
		// },
		// "/classificacao": {
		// 	http.MethodGet,
		// 	apihandler.GetAllClassificacao,
		// },
		// "/subclassificacao": {
		// 	http.MethodGet,
		// 	apihandler.GetSubClassificacao,
		// },
		// "/colecao": {
		// 	http.MethodGet,
		// 	apihandler.GetAllColecao,
		// },
		// "/subcolecao": {
		// 	http.MethodGet,
		// 	apihandler.GetSubColecao,
		// },
		// "/novoItem": {
		// 	http.MethodPost,
		// 	apihandler.NewItem,
		// },
		// "/item/{id}": {
		// 	http.MethodPost,
		// 	apihandler.GetItem,
		// },
	}
}

// //UpdateFiliais ...
// func (c *Controller) UpdateFiliais() {
// 	rows, err := c.linx.QueryRow("SELECT RTRIM(FILIAL), RTRIM(COD_FILIAL), RTRIM(CGC_CPF) FROM FILIAIS")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	filiais := make([]string, 0, 100)

// 	for rows.Next() {
// 		filial := linxFilial{}
// 		if err := rows.Scan(&filial.Filial, &filial.CodFilial, &filial.CNPJ); err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		filiais = append(filiais, filial.parse())
// 	}

// 	_, err = c.fin.QueryRow("INSERT IGNORE INTO FILIAIS_LINX (NOME_FANTASIA, COD_LINX, CNPJ) VALUES %s;", strings.Join(filiais, ", "))
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// //GetMatrizes ...
// func (c *Controller) GetMatrizes(tables ...string) ([][]string, error) {
// 	rows, err := c.fin.QueryRow("SELECT RTRIM(CNPJ) FROM MATRIZES")
// 	if err != nil {
// 		// log.Println(err)
// 		return nil, err
// 	}

// 	matrizes := make([][]string, len(tables))
// 	for i := 0; i < len(matrizes); i++ {
// 		matrizes[i] = make([]string, 0, 10)
// 	}

// 	var cnpj string
// 	for rows.Next() {
// 		if err := rows.Scan(&cnpj); err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		for i := 0; i < len(matrizes); i++ {
// 			matrizes[i] = append(matrizes[i], fmt.Sprintf("%s.CGC_CPF LIKE '%s%%'", tables[i], cnpj))
// 		}
// 	}
// 	return matrizes, nil
// }

// //UpdateFaturamentos ...
// func (c *Controller) UpdateFaturamentos() error {
// 	m, err := c.GetMatrizes("H", "B")
// 	if err != nil {
// 		return err
// 	}
// 	rows, err := c.linx.QueryRow(`Select 'PDV' AS PLATAFORMA, RTRIM(A.CODIGO_FILIAL) AS CODIGO_FILIAL, RTRIM(B.SERIE_NF_SAIDA) AS SERIE_NF, RTRIM(B.NUMERO_FISCAL_VENDA) AS NF_SAIDA
// 		, RTRIM(D.CLIENTE_VAREJO) AS CLIENTE, RTRIM(D.CPF_CGC) AS CNPJ, RTRIM(E.NATUREZA_SAIDA) AS NATUREZA_SAIDA, RTRIM(F.DESC_NATUREZA) AS DESC_NATUREZA_SAIDA, RTRIM(G.FORMA_PGTO) AS DESC_COND_PGTO
// 		 ,RTRIM(C.EMISSAO) AS EMISSAO, RTRIM(C.DATA_SAIDA) AS DATA_SAIDA, C.VALOR_TOTAL, C.QTDE_TOTAL
// 		 , CAST(COALESCE(C.OBS,'') AS VARCHAR(255)) AS OBS
// 		 , C.NOTA_CANCELADA, RTRIM(C.MOTIVO_CANCELAMENTO_NFE) AS MOTIVO_CANCELAMENTO_NFE, RTRIM(C.DATA_CANCELAMENTO) AS DATA_CANCELAMENTO, RTRIM(C.CHAVE_NFE) AS CHAVE_NFE, RTRIM(A.TICKET) AS TICKET,  RTRIM(B.NUMERO_FISCAL_TROCA) AS NUMERO_FISCAL_TROCA
// 		FROM LOJA_VENDA A
// 		LEFT JOIN LOJA_VENDA_PGTO B ON A.LANCAMENTO_CAIXA=B.LANCAMENTO_CAIXA AND A.TERMINAL=B.TERMINAL AND A.CODIGO_FILIAL=B.CODIGO_FILIAL
// 	LEFT JOIN LOJA_NOTA_FISCAL C ON B.CODIGO_FILIAL=C.CODIGO_FILIAL AND B.SERIE_NF_SAIDA=C.SERIE_NF AND B.NUMERO_FISCAL_VENDA=C.NF_NUMERO
// 	LEFT JOIN CLIENTES_VAREJO D ON A.CODIGO_CLIENTE=D.CODIGO_CLIENTE
// 	LEFT JOIN LOJAS_NATUREZA_OPERACAO E ON C.NATUREZA_OPERACAO_CODIGO=E.NATUREZA_OPERACAO_CODIGO
// 	LEFT JOIN NATUREZAS_SAIDAS F ON E.NATUREZA_SAIDA=F.NATUREZA_SAIDA
// 	LEFT JOIN LOJA_FORMAS_PGTO G ON B.COD_FORMA_PGTO=G.COD_FORMA_PGTO
// 	LEFT JOIN FILIAIS H ON A.CODIGO_FILIAL=H.COD_FILIAL
// 	WHERE C.RECEBIMENTO=0 AND (%s)
// 	UNION SELECT 'ERP' AS PLATAFORMA, RTRIM(B.COD_FILIAL) AS CODIGO_FILIAL, RTRIM(SERIE_NF) AS SERIE_NF, RTRIM(NF_SAIDA) AS NF_SAIDA
// 	 , RTRIM(A.NOME_CLIFOR) AS CLIENTE, RTRIM(E.CGC_CPF) AS CNPJ, RTRIM(A.NATUREZA_SAIDA) AS NATUREZA_SAIDA, RTRIM(D.DESC_NATUREZA) AS DESC_NATUREZA_SAIDA, RTRIM(C.DESC_COND_PGTO) AS DESC_COND_PGTO
// 	 ,RTRIM(EMISSAO) AS EMISSAO, RTRIM(DATA_SAIDA) AS DATA_SAIDA, VALOR_TOTAL, QTDE_TOTAL
// 	 , CAST(COALESCE(A.OBS,'') AS VARCHAR(255)) AS OBS
// 	 , NOTA_CANCELADA, RTRIM(MOTIVO_CANCELAMENTO_NFE) AS MOTIVO_CANCELAMENTO_NFE, RTRIM(DATA_CANCELAMENTO) AS DATA_CANCELAMENTO, RTRIM(A.CHAVE_NFE) AS CHAVE_NFE,  NULL AS TICKET,  NULL AS NUMERO_FISCAL_TROCA
// 		FROM [LINX_TBFG].[dbo].[Faturamento] A
// 		LEFT JOIN FILIAIS B ON A.FILIAL=B.FILIAL
// 		LEFT JOIN COND_ATAC_PGTOS C ON A.CONDICAO_PGTO=C.CONDICAO_PGTO
// 		LEFT JOIN NATUREZAS_SAIDAS D ON A.NATUREZA_SAIDA=D.NATUREZA_SAIDA
// 		LEFT JOIN CADASTRO_CLI_FOR E ON A.NOME_CLIFOR=E.NOME_CLIFOR
// 		WHERE %s`, strings.Join(m[0], " OR "), strings.Join(m[1], " OR "))
// 	if err != nil {
// 		return err
// 	}

// 	rsts := make([]linxFat, 0, 100)
// 	for rows.Next() {
// 		var rst linxFat
// 		if err := rows.Scan(&rst.Plataforma, &rst.CodFilial, &rst.NFSerie, &rst.NF, &rst.NomeCliente, &rst.CNPJ, &rst.Natureza, &rst.NaturezaDesc, &rst.DescCondPgto, &rst.Emissao, &rst.Data, &rst.ValorTotal, &rst.QtdeTotal, &rst.Obs, &rst.NotaCancelada, &rst.MotivoCancelamento, &rst.DataCancelamento, &rst.ChaveNFE, &rst.Ticket, &rst.NFDevolucao); err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		rsts = append(rsts, rst)
// 	}

// 	values := make([]string, len(rsts))

// 	for i := 0; i < len(rsts); i++ {
// 		values[i] = rsts[i].parse()
// 	}

// 	_, err = c.fin.QueryRow(`INSERT INTO FATURAMENTOS (PLATAFORMA, COD_LINX_FILIAL, NF_SERIE, NF_SAIDA, NOME_CLIFOR, CNPJ_CPF, NATUREZA_SAIDA, DESC_NATUREZA, DESC_COND_PGTO, EMISSAO, DATA_SAIDA, VALOR_TOTAL, OBS, NF_CANCELADA, MOTIVO_CANCELAMENTO, DATA_CANCELAMENTO, NF_CHAVE, TICKET, NF_DEVOLUCAO)
// 	VALUES %s AS NEW ON DUPLICATE KEY UPDATE
// 	PLATAFORMA = NEW.PLATAFORMA,
// 	COD_LINX_FILIAL = NEW.COD_LINX_FILIAL,
// 	NF_SERIE = NEW.NF_SERIE,
// 	NF_SAIDA = NEW.NF_SAIDA,
// 	NOME_CLIFOR = NEW.NOME_CLIFOR,
// 	CNPJ_CPF = NEW.CNPJ_CPF,
// 	NATUREZA_SAIDA = NEW.NATUREZA_SAIDA,
// 	DESC_NATUREZA = NEW.DESC_NATUREZA,
// 	DESC_COND_PGTO = NEW.DESC_COND_PGTO,
// 	EMISSAO = NEW.EMISSAO,
// 	DATA_SAIDA = NEW.DATA_SAIDA,
// 	VALOR_TOTAL = NEW.VALOR_TOTAL,
// 	OBS = NEW.OBS,
// 	NF_CANCELADA = NEW.NF_CANCELADA,
// 	MOTIVO_CANCELAMENTO = NEW.MOTIVO_CANCELAMENTO,
// 	DATA_CANCELAMENTO = NEW.DATA_CANCELAMENTO,
// 	NF_CHAVE = NEW.NF_CHAVE,
// 	TICKET = NEW.TICKET,
// 	NF_DEVOLUCAO = NEW.NF_DEVOLUCAO;`, strings.Join(values, ", "))

// 	return err
// }

// //UpdateEntradas ...
// func (c *Controller) UpdateEntradas() error {
// 	m, err := c.GetMatrizes("G", "B")
// 	if err != nil {
// 		return err
// 	}
// 	rows, err := c.linx.QueryRow(`Select 'PDV' AS PLATAFORMA, RTRIM(A.CODIGO_FILIAL) AS CODIGO_FILIAL, RTRIM(A.SERIE_NF) AS SERIE_NF, RTRIM(A.NF_NUMERO) AS NF_ENTRADA
// 	, RTRIM(D.CLIENTE_VAREJO) AS CLIENTE, RTRIM(D.CPF_CGC) AS CNPJ, RTRIM(E.NATUREZA_ENTRADA) AS NATUREZA_ENTRADA, RTRIM(F.DESC_NATUREZA) AS DESC_NATUREZA_ENTRADA
// 	 ,RTRIM(A.EMISSAO) AS EMISSAO, RTRIM(A.DATA_SAIDA) AS DATA_ENTRADA, A.VALOR_TOTAL, A.QTDE_TOTAL
// 	 , CAST(COALESCE(A.OBS,'') AS VARCHAR(255)) AS OBS, 1 AS EMISSAO_PROPRIA
// 	 , A.NOTA_CANCELADA, RTRIM(A.MOTIVO_CANCELAMENTO_NFE) AS MOTIVO_CANCELAMENTO_NFE, RTRIM(A.DATA_CANCELAMENTO) AS DATA_CANCELAMENTO, RTRIM(A.CHAVE_NFE) AS CHAVE_NFE
// 	FROM LOJA_NOTA_FISCAL A
// LEFT JOIN CLIENTES_VAREJO D ON A.CODIGO_CLIENTE=D.CODIGO_CLIENTE
// LEFT JOIN LOJAS_NATUREZA_OPERACAO E ON A.NATUREZA_OPERACAO_CODIGO=E.NATUREZA_OPERACAO_CODIGO
// LEFT JOIN NATUREZAS_ENTRADAS F ON E.NATUREZA_ENTRADA=F.NATUREZA
// LEFT JOIN FILIAIS G ON A.CODIGO_FILIAL=G.COD_FILIAL
// WHERE A.RECEBIMENTO=1 AND (%s)
// UNION SELECT 'ERP' AS PLATAFORMA, RTRIM(B.COD_FILIAL) AS CODIGO_FILIAL, RTRIM(A.SERIE_NF_ENTRADA) AS SERIE_NF, RTRIM(A.NF_ENTRADA) AS NF_ENTRADA
//  , RTRIM(A.NOME_CLIFOR) AS CLIENTE, RTRIM(E.CGC_CPF) AS CNPJ, RTRIM(A.NATUREZA) AS NATUREZA_ENTRADA, RTRIM(D.DESC_NATUREZA) AS DESC_NATUREZA_ENTRADA
//  ,RTRIM(EMISSAO) AS EMISSAO, RTRIM(A.RECEBIMENTO) AS DATA_ENTRADA, A.VALOR_TOTAL, A.QTDE_TOTAL
//  , CAST(COALESCE(A.OBS,'') AS VARCHAR(255)) AS OBS, A.NF_ENTRADA_PROPRIA AS EMISSAO_PROPRIA
//  , A.NOTA_CANCELADA, RTRIM(A.MOTIVO_CANCELAMENTO_NFE) AS MOTIVO_CANCELAMENTO_NFE, RTRIM(A.DATA_CANCELAMENTO) AS DATA_CANCELAMENTO, RTRIM(A.CHAVE_NFE) AS CHAVE_NFE
// 	FROM ENTRADAS A
// 	LEFT JOIN FILIAIS B ON A.FILIAL=B.FILIAL
// 	LEFT JOIN COND_ATAC_PGTOS C ON A.CONDICAO_PGTO=C.CONDICAO_PGTO
// 	LEFT JOIN NATUREZAS_ENTRADAS D ON A.NATUREZA=D.NATUREZA
// 	LEFT JOIN CADASTRO_CLI_FOR E ON A.NOME_CLIFOR=E.NOME_CLIFOR
// 	WHERE A.NATUREZA LIKE '250.%%' AND (%s)`, strings.Join(m[0], " OR "), strings.Join(m[1], " OR "))
// 	if err != nil {
// 		return err
// 	}

// 	rsts := make([]linxEnt, 0, 100)
// 	for rows.Next() {
// 		var rst linxEnt
// 		if err := rows.Scan(&rst.Plataforma, &rst.CodFilial, &rst.NFSerie, &rst.NF, &rst.NomeCliente, &rst.CNPJ, &rst.Natureza, &rst.NaturezaDesc, &rst.Emissao, &rst.Data, &rst.ValorTotal, &rst.QtdeTotal, &rst.Obs, &rst.EmissaoPropria, &rst.NotaCancelada, &rst.MotivoCancelamento, &rst.DataCancelamento, &rst.ChaveNFE); err != nil {
// 			log.Println(err)

// 			// var t [18]interface{}
// 			// rows.Scan(&t[0], &t[1], &t[2], &t[3], &t[4], &t[5], &t[6], &t[7], &t[8], &t[9], &t[10], &t[11], &t[12], &t[13], &t[14], &t[15], &t[16], &t[17])
// 			// fmt.Printf("%+v", t)
// 			continue
// 		}
// 		rsts = append(rsts, rst)
// 	}

// 	values := make([]string, len(rsts))

// 	for i := 0; i < len(rsts); i++ {
// 		values[i] = rsts[i].parse()
// 	}

// 	_, err = c.fin.QueryRow(`INSERT INTO ENTRADAS (PLATAFORMA, COD_LINX_FILIAL, NF_SERIE, NF_ENTRADA, NOME_CLIFOR, CNPJ_CPF, NATUREZA_ENTRADA, DESC_NATUREZA, EMISSAO, DATA_ENTRADA, VALOR_TOTAL, OBS, EMISSAO_PROPRIA, NF_CANCELADA, MOTIVO_CANCELAMENTO, DATA_CANCELAMENTO, NF_CHAVE)
// 	VALUES %s AS NEW ON DUPLICATE KEY UPDATE
// 	PLATAFORMA = NEW.PLATAFORMA,
// 	COD_LINX_FILIAL = NEW.COD_LINX_FILIAL,
// 	NF_SERIE = NEW.NF_SERIE,
// 	NF_ENTRADA = NEW.NF_ENTRADA,
// 	NOME_CLIFOR = NEW.NOME_CLIFOR,
// 	CNPJ_CPF = NEW.CNPJ_CPF,
// 	NATUREZA_ENTRADA = NEW.NATUREZA_ENTRADA,
// 	DESC_NATUREZA = NEW.DESC_NATUREZA,
// 	EMISSAO = NEW.EMISSAO,
// 	DATA_ENTRADA = NEW.DATA_ENTRADA,
// 	VALOR_TOTAL = NEW.VALOR_TOTAL,
// 	OBS = NEW.OBS,
// 	EMISSAO_PROPRIA = NEW.EMISSAO_PROPRIA,
// 	NF_CANCELADA = NEW.NF_CANCELADA,
// 	MOTIVO_CANCELAMENTO = NEW.MOTIVO_CANCELAMENTO,
// 	DATA_CANCELAMENTO = NEW.DATA_CANCELAMENTO,
// 	NF_CHAVE = NEW.NF_CHAVE;`, strings.Join(values, ", "))

// 	return err
// }
