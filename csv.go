package datapackage

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

// Essa versão deixou de ser utilizada em julho de 2023
type ResultadoColeta_CSV struct {
	Coleta       []Coleta_CSV
	Remuneracoes []Remuneracao_CSV
	Folha        []ContraCheque_CSV
	Metadados    []Metadados_CSV
}

// Essa versão deixou de ser utilizada em julho de 2023
type Coleta_CSV struct {
	ChaveColeta        string    `csv:"chave_coleta" tableheader:"chave_coleta"`
	Orgao              string    `csv:"orgao" tableheader:"orgao"`
	Mes                int32     `csv:"mes" tableheader:"mes"`
	Ano                int32     `csv:"ano" tableheader:"ano"`
	TimestampColeta    time.Time `csv:"timestamp_coleta" tableheader:"timestamp_coleta"`
	RepositorioColetor string    `csv:"repositorio_coletor" tableheader:"repositorio_coletor"`
	VersaoColetor      string    `csv:"versao_coletor" tableheader:"versao_coletor"`
	DirColetor         string    `csv:"dir_coletor" tableheader:"dir_coletor"`
}

// Essa versão deixou de ser utilizada em julho de 2023
type ContraCheque_CSV struct {
	IdContraCheque string `csv:"id_contra_cheque" tableheader:"id_contra_cheque"`
	ChaveColeta    string `csv:"chave_coleta" tableheader:"chave_coleta"`
	Nome           string `csv:"nome" tableheader:"nome"`
	Matricula      string `csv:"matricula" tableheader:"matricula"`
	Funcao         string `csv:"funcao" tableheader:"funcao"`
	LocalTrabalho  string `csv:"local_trabalho" tableheader:"local_trabalho"`
	Tipo           string `csv:"tipo" tableheader:"tipo"`
	Ativo          bool   `csv:"ativo" tableheader:"ativo"`
}

// Essa versão deixou de ser utilizada em julho de 2023
type Metadados_CSV struct {
	ChaveColeta                string  `csv:"chave_coleta" tableheader:"chave_coleta"`
	FormatoAberto              bool    `csv:"formato_aberto" tableheader:"formato_aberto"`                             // Os dados são disponibilizados em formato aberto?
	Acesso                     string  `csv:"acesso" tableheader:"acesso"`                                             // Conseguimos prever/construir uma URL com base no órgão/mês/ano que leve ao download do dado?
	Extensao                   string  `csv:"extensao" tableheader:"extensao"`                                         // Extensao do arquivo de dados, ex: CSV, JSON, XLS, etc
	EstritamenteTabular        bool    `csv:"estritamente_tabular" tableheader:"estritamente_tabular"`                 // Órgãos que disponibilizam dados limpos (tidy data)
	FormatoConsistente         bool    `csv:"formato_consistente" tableheader:"formato_consistente"`                   // Órgão alterou a forma de expor seus dados entre o mês em questão e o mês anterior?
	TemMatricula               bool    `csv:"tem_matricula" tableheader:"tem_matricula"`                               // Órgão disponibiliza matrícula do servidor?
	TemLotacao                 bool    `csv:"tem_lotacao" tableheader:"tem_lotacao"`                                   // Órgão disponibiliza lotação do servidor?
	TemCargo                   bool    `csv:"tem_cargo" tableheader:"tem_cargo"`                                       // Órgão disponibiliza a função do servidor?
	DetalhamentoReceitaBase    string  `csv:"detalhamento_receita_base" tableheader:"detalhamento_receita_base"`       // Contra-cheque
	DetalhamentoOutrasReceitas string  `csv:"detalhamento_outras_receitas" tableheader:"detalhamento_outras_receitas"` // Inclui indenizações, direitos eventuais, diárias, etc
	DetalhamentoDescontos      string  `csv:"detalhamento_descontos" tableheader:"detalhamento_descontos"`             // Inclui imposto de renda, retenção por teto e contribuição previdenciária
	IndiceCompletude           float32 `csv:"indice_completude" tableheader:"indice_completude"`                       // Componente do índice de transparência resultante da análise dos metadados relacionados a disponibilidade dos dados.
	IndiceFacilidade           float32 `csv:"indice_facilidade" tableheader:"indice_facilidade"`                       // Componente do índice de transparência resultante da análise dos metadados relacionados a dificuldade para acessar os dados que estão disponíveis.
	IndiceTransparencia        float32 `csv:"indice_transparencia" tableheader:"indice_transparencia"`                 // Nota final, calculada utilizada os componentes de disponibilidade e dificuldade.
}

// Essa versão deixou de ser utilizada em julho de 2023
type Remuneracao_CSV struct {
	IdContraCheque string  `csv:"id_contra_cheque" tableheader:"id_contra_cheque"`
	ChaveColeta    string  `csv:"chave_coleta" tableheader:"chave_coleta"`
	Natureza       string  `csv:"natureza" tableheader:"natureza"`
	Categoria      string  `csv:"categoria" tableheader:"categoria"`
	Item           string  `csv:"item" tableheader:"item"`
	Valor          float64 `csv:"valor" tableheader:"valor"`
}

// toCSVFile dumps the payroll into a file using the CSV format.
func toCSVFile(in interface{}, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating CSV file(%s):%q", path, err)
	}
	defer f.Close()

	// Cria um novo escritor CSV com o separador de colunas personalizado
	csvWriter := csv.NewWriter(f)
	csvWriter.Comma = ';'
	csvWriter.UseCRLF = true // Para garantir que os fins de linha estejam no formato correto

	return gocsv.MarshalCSV(in, csvWriter)
}

// fromCSVFile gets from CSV to a certain struct.
func fromCSVFile(in interface{}, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	// Cria um leitor CSV com o separador de colunas personalizado
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'

	return gocsv.UnmarshalCSV(csvReader, in)
}
