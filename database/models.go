package sql

//Item ...
type Item struct {
	Descricao        string
	Marca            string
	Colecao          string
	Subcolecao       string
	Classificacao    string
	Subclassificacao string
	Cor_cod          string
	Grade            string
	ID_proporcao     string
	Foto             string
	Data             string
}

type CodItem struct {
}

type Colecoes struct {
	Desc_colecao string
}
