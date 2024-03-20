package datapackage

import (
	"strconv"
	"strings"
	"time"
)

// *_V2: Essa versão passou a ser utilizada a partir de julho de 2023
type ResultadoColeta_CSV_V2 struct {
	Coleta       []Coleta_CSV_V2
	Remuneracoes []Remuneracao_CSV_V2
	Folha        []Contracheque_CSV_V2
	Metadados    []Metadados_CSV_V2
}

// *_V2: Essa versão passou a ser utilizada a partir de julho de 2023
type Coleta_CSV_V2 struct {
	ChaveColeta        string    `csv:"chave_coleta" tableheader:"chave_coleta"`
	Orgao              string    `csv:"orgao" tableheader:"orgao"`
	Mes                int32     `csv:"mes" tableheader:"mes"`
	Ano                int32     `csv:"ano" tableheader:"ano"`
	TimestampColeta    time.Time `csv:"timestamp_coleta" tableheader:"timestamp_coleta"`
	RepositorioColetor string    `csv:"repositorio_coletor" tableheader:"repositorio_coletor"`
	VersaoColetor      string    `csv:"versao_coletor" tableheader:"versao_coletor"`
	RepositorioParser  string    `csv:"repositorio_parser" tableheader:"repositorio_parser"`
	VersaoParser       string    `csv:"versao_parser" tableheader:"versao_parser"`
}

// *_V2: Essa versão passou a ser utilizada a partir de julho de 2023
type Contracheque_CSV_V2 struct {
	IdContracheque int           `csv:"id_contracheque" tableheader:"id_contracheque"`
	Orgao          string        `csv:"orgao" tableheader:"orgao"`
	Mes            int32         `csv:"mes" tableheader:"mes"`
	Ano            int32         `csv:"ano" tableheader:"ano"`
	Nome           string        `csv:"nome" tableheader:"nome"`
	Matricula      string        `csv:"matricula" tableheader:"matricula"`
	Funcao         string        `csv:"funcao" tableheader:"funcao"`
	LocalTrabalho  string        `csv:"local_trabalho" tableheader:"local_trabalho"`
	Salario        CustomFloat32 `csv:"salario" tableheader:"salario"`
	Beneficios     CustomFloat32 `csv:"beneficios" tableheader:"beneficios"`
	Descontos      CustomFloat32 `csv:"descontos" tableheader:"descontos"`
	Remuneracao    CustomFloat32 `csv:"remuneracao" tableheader:"remuneracao"`
	Situacao       string        `csv:"situacao" tableheader:"situacao"`
}

// *_V2: Essa versão passou a ser utilizada a partir de julho de 2023
type Metadados_CSV_V2 struct {
	Orgao                      string        `csv:"orgao" tableheader:"orgao"`
	Mes                        int32         `csv:"mes" tableheader:"mes"`
	Ano                        int32         `csv:"ano" tableheader:"ano"`
	FormatoAberto              bool          `csv:"formato_aberto" tableheader:"formato_aberto"`                             // Os dados são disponibilizados em formato aberto?
	Acesso                     string        `csv:"acesso" tableheader:"acesso"`                                             // Conseguimos prever/construir uma URL com base no órgão/mês/ano que leve ao download do dado?
	Extensao                   string        `csv:"extensao" tableheader:"extensao"`                                         // Extensao do arquivo de dados, ex: CSV, JSON, XLS, etc
	EstritamenteTabular        bool          `csv:"estritamente_tabular" tableheader:"estritamente_tabular"`                 // Órgãos que disponibilizam dados limpos (tidy data)
	FormatoConsistente         bool          `csv:"formato_consistente" tableheader:"formato_consistente"`                   // Órgão alterou a forma de expor seus dados entre o mês em questão e o mês anterior?
	TemMatricula               bool          `csv:"tem_matricula" tableheader:"tem_matricula"`                               // Órgão disponibiliza matrícula do servidor?
	TemLotacao                 bool          `csv:"tem_lotacao" tableheader:"tem_lotacao"`                                   // Órgão disponibiliza lotação do servidor?
	TemCargo                   bool          `csv:"tem_cargo" tableheader:"tem_cargo"`                                       // Órgão disponibiliza a função do servidor?
	DetalhamentoReceitaBase    string        `csv:"detalhamento_receita_base" tableheader:"detalhamento_receita_base"`       // Contra-cheque
	DetalhamentoOutrasReceitas string        `csv:"detalhamento_outras_receitas" tableheader:"detalhamento_outras_receitas"` // Inclui indenizações, direitos eventuais, diárias, etc
	DetalhamentoDescontos      string        `csv:"detalhamento_descontos" tableheader:"detalhamento_descontos"`             // Inclui imposto de renda, retenção por teto e contribuição previdenciária
	IndiceCompletude           CustomFloat32 `csv:"indice_completude" tableheader:"indice_completude"`                       // Componente do índice de transparência resultante da análise dos metadados relacionados a disponibilidade dos dados.
	IndiceFacilidade           CustomFloat32 `csv:"indice_facilidade" tableheader:"indice_facilidade"`                       // Componente do índice de transparência resultante da análise dos metadados relacionados a dificuldade para acessar os dados que estão disponíveis.
	IndiceTransparencia        CustomFloat32 `csv:"indice_transparencia" tableheader:"indice_transparencia"`                 // Nota final, calculada utilizada os componentes de disponibilidade e dificuldade.
}

// *_V2: Essa versão passou a ser utilizada a partir de julho de 2023
type Remuneracao_CSV_V2 struct {
	IdContracheque int           `csv:"id_contracheque" tableheader:"id_contracheque"`
	Orgao          string        `csv:"orgao" tableheader:"orgao"`
	Mes            int32         `csv:"mes" tableheader:"mes"`
	Ano            int32         `csv:"ano" tableheader:"ano"`
	Tipo           string        `csv:"tipo" tableheader:"tipo"`
	Categoria      string        `csv:"categoria" tableheader:"categoria"`
	Item           string        `csv:"item" tableheader:"item"`
	Valor          CustomFloat32 `csv:"valor" tableheader:"valor"`
}

type CustomFloat32 float32

func (c CustomFloat32) MarshalCSV() (string, error) {
	// Substitui o ponto separador decimal por vírgula
	return strings.ReplaceAll(strconv.FormatFloat(float64(c), 'f', -1, 32), ".", ","), nil
}

func (c *CustomFloat32) UnmarshalCSV(csv string) error {
	// Substitui a vírgula por ponto para garantir que o valor seja interpretado corretamente
	csv = strings.ReplaceAll(csv, ",", ".")

	// Converte a string para float64
	floatVal, err := strconv.ParseFloat(csv, 32)
	if err != nil {
		return err
	}

	// Converte float64 para float32 e atribui ao CustomFloat32
	*c = CustomFloat32(float32(floatVal))

	return nil
}
