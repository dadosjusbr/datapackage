package datapackage

// Versão compilável do descritor JSON dos dados consolidados e publicados pelo projeto DadosJusBR.
// Além de facilitar a manutenção, permite a criação de uma biblioteca que lida com a criação
// e carregamento de pacotes de dados.
var dadosjusbrDescriptor = dadosjusbrDescriptorStruct{
	Name:         "remuneracoes-jusbr",
	Title:        "Remunerações do Sistema de Justiça Brasileiro",
	Description:  "Remunerações do Sistema de Justiça Brasileiro libertadas por DadosJusBR",
	Profile:      "tabular-data-package",
	Homepage:     "https://dadosjusbr.org",
	Version:      "2.0.0",
	Contributors: []Contributor{{Title: "DadosJusBR", Role: "author"}},
	Licenses:     []License{{Name: "CC-BY-4.0", Title: "Creative Commons Attribution 4.0", Path: "https://creativecommons.org/licenses/by/4.0/"}},
	Keywords:     []string{"justice-system", "payments", "remunerações", "sistema-de-justiça"},
	Resources:    []Resource{coletaResource, contraChequeResource, remuneracaoResource, metadadosResource},
}

var coletaResource = Resource{
	Name:        coletaResourceName,
	Description: "Descreve a coleta e o coletor de um determinado órgão, mês e ano.",
	Path:        coletaFileName,
	Profile:     "tabular-data-resource",
	Schema: Schema{
		PrimaryKey: "chave_coleta",
		Fields: []Field{
			{
				Name:            "chave_coleta",
				Type:            "string",
				Format:          "default",
				Title:           "Chave coleta",
				Description:     "The Unique Key of collection",
				DescriptionPTBR: "A chave única da coleta",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "orgao",
				Type:            "string",
				Format:          "default",
				Title:           "Órgão",
				Description:     "The ID of agency",
				DescriptionPTBR: "A sigla da agência",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "mes",
				Type:            "integer",
				Format:          "default",
				Title:           "Mês",
				Description:     "Month that data collection refers to",
				DescriptionPTBR: "O mês que a os dados foram coletados se referem",
				Constraints:     Constraints{Required: true, Maximum: "12", Minimum: "1"},
			},
			{
				Name:            "ano",
				Type:            "integer",
				Format:          "default",
				Title:           "Ano",
				Description:     "Year that data collection refers to",
				DescriptionPTBR: "O ano que a os dados foram coletados se referem",
				Constraints:     Constraints{Required: true, Minimum: "2018"},
			},
			{
				Name:            "timestamp_coleta",
				Type:            "datetime",
				Format:          "%Y-%m-%dT%H:%M:%S.%fZ",
				Title:           "Timestamp que a coleta aconteceu",
				Description:     "Timestamp mark of data crawled",
				DescriptionPTBR: "Contém a marca temporal em que o dado foi coletado",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "repositorio_coletor",
				Type:            "string",
				Format:          "default",
				Title:           "Repositório Coletor",
				Description:     "The name of the repository that performed the collection",
				DescriptionPTBR: "Repositório que realizou a coleta",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "versao_coletor",
				Type:            "string",
				Format:          "default",
				Title:           "Versão do coletor",
				Description:     "Version of the collector that performed the collection",
				DescriptionPTBR: "Versão do coletor que realizou a coleta",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "repositorio_parser",
				Type:            "string",
				Format:          "default",
				Title:           "Repositório Parser",
				Description:     "The name of the repository that processed the data",
				DescriptionPTBR: "Repositório que processou os dados",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "versao_parser",
				Type:            "string",
				Format:          "default",
				Title:           "Versão do Parser",
				Description:     "Version of the parser that processed the data",
				DescriptionPTBR: "Versão do parser que processou os dados",
				Constraints:     Constraints{Required: true},
			},
		},
	},
}

var contraChequeResource = Resource{
	Name:        contrachequeResourceName,
	Description: "Descreve os contracheques associados a uma determinada coleta (mês, ano, órgão)",
	Path:        contrachequeFileName,
	Profile:     "tabular-data-resource",
	Schema: Schema{
		PrimaryKey: "id_contra_cheque",
		ForeignKeys: []ForeignKey{
			{"orgao", FKRef{coletaResourceName, "orgao"}},
			{"mes", FKRef{coletaResourceName, "mes"}},
			{"ano", FKRef{coletaResourceName, "ano"}}},
		Fields: []Field{
			{
				Name:            "id_contra_cheque",
				Type:            "number",
				Format:          "default",
				Title:           "Identificador da folha de pagamento",
				Description:     "Payroll identifier",
				DescriptionPTBR: "Identificador da folha de pagamento",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "orgao",
				Type:            "string",
				Format:          "default",
				Title:           "Órgão",
				Description:     "The ID of agency",
				DescriptionPTBR: "A sigla da agência",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "mes",
				Type:            "integer",
				Format:          "default",
				Title:           "Mês",
				Description:     "Month that data collection refers to",
				DescriptionPTBR: "O mês que a os dados foram coletados se referem",
				Constraints:     Constraints{Required: true, Maximum: "12", Minimum: "1"},
			},
			{
				Name:            "ano",
				Type:            "integer",
				Format:          "default",
				Title:           "Ano",
				Description:     "Year that data collection refers to",
				DescriptionPTBR: "O ano que a os dados foram coletados se referem",
				Constraints:     Constraints{Required: true, Minimum: "2018"},
			},
			{
				Name:            "nome",
				Type:            "string",
				Format:          "default",
				Title:           "Nome",
				Description:     "Public servant name",
				DescriptionPTBR: "Nome do servidor público",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "matricula",
				Type:            "string",
				Format:          "default",
				Title:           "Matrícula",
				Description:     "Public servant work id",
				DescriptionPTBR: "Matrícula do servidor público",
			},
			{
				Name:            "funcao",
				Type:            "string",
				Format:          "default",
				Title:           "Função",
				Description:     "Public servant role",
				DescriptionPTBR: "Função do servidor público",
			},
			{
				Name:            "local_trabalho",
				Type:            "string",
				Format:          "default",
				Title:           "Local de trabalho",
				Description:     "Public servant workplace",
				DescriptionPTBR: "O local onde o membro está lotado",
			},
			{
				Name:            "salario",
				Type:            "number",
				Format:          "default",
				Title:           "Salário",
				Description:     "Public servant payroll salary",
				DescriptionPTBR: "Salário base no contracheque do membro",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "beneficios",
				Type:            "number",
				Format:          "default",
				Title:           "Benefícios",
				Description:     "Public servant payroll benefits",
				DescriptionPTBR: "Benefícios no contracheque do membro",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "descontos",
				Type:            "number",
				Format:          "default",
				Title:           "Descontos",
				Description:     "Public servant payroll discounts",
				DescriptionPTBR: "Descontos no contracheque do membro",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "remuneracao",
				Type:            "number",
				Format:          "default",
				Title:           "Remuneração",
				Description:     "Public servant remuneration",
				DescriptionPTBR: "Remuneração do membro",
				Constraints:     Constraints{Required: true},
			},
		},
	},
}

var remuneracaoResource = Resource{
	Name:        remuneracaoResourceName,
	Description: "Detalha as remunerações associadas aos contracheques de uma determinada coleta (mês, ano, órgão)",
	Path:        remuneracaoFileName,
	Profile:     "tabular-data-resource",
	Schema: Schema{
		ForeignKeys: []ForeignKey{
			{"id_contra_cheque", FKRef{contrachequeResourceName, "id_contra_cheque"}},
			{"orgao", FKRef{coletaResourceName, "orgao"}},
			{"mes", FKRef{coletaResourceName, "mes"}},
			{"ano", FKRef{coletaResourceName, "ano"}},
		},
		Fields: []Field{
			{
				Name:            "id_contra_cheque",
				Type:            "number",
				Format:          "default",
				Title:           "Identificador da folha de pagamento",
				Description:     "Payroll identifier",
				DescriptionPTBR: "Identificador da folha de pagamento",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "orgao",
				Type:            "string",
				Format:          "default",
				Title:           "Órgão",
				Description:     "The ID of agency",
				DescriptionPTBR: "A sigla da agência",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "mes",
				Type:            "integer",
				Format:          "default",
				Title:           "Mês",
				Description:     "Month that data collection refers to",
				DescriptionPTBR: "O mês que a os dados foram coletados se referem",
				Constraints:     Constraints{Required: true, Maximum: "12", Minimum: "1"},
			},
			{
				Name:            "ano",
				Type:            "integer",
				Format:          "default",
				Title:           "Ano",
				Description:     "Year that data collection refers to",
				DescriptionPTBR: "O ano que a os dados foram coletados se referem",
				Constraints:     Constraints{Required: true, Minimum: "2018"},
			},
			{
				Name:            "tipo",
				Type:            "string",
				Format:          "default",
				Title:           "Tipo",
				Description:     "Describes whether it is an income (base or other remuneration) or a discount",
				DescriptionPTBR: "Descreve se é um rendimento (base ou outras remunerações) ou um desconto",
				Constraints:     Constraints{Required: true, Enum: []string{"R/B", "R/O", "D"}},
			},
			{
				Name:            "categoria",
				Type:            "string",
				Format:          "default",
				Title:           "Categoria",
				Description:     "Category of the remuneration",
				DescriptionPTBR: "Categoria da remuneração",
			},
			{
				Name:            "item",
				Type:            "string",
				Format:          "default",
				Title:           "Ítem de remuneração",
				Description:     "Description of the remuneration item",
				DescriptionPTBR: "Descrição do ítem de remuneração",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "valor",
				Type:            "number",
				Format:          "default",
				Title:           "Valor",
				Description:     "Value associated to the remuneration item",
				DescriptionPTBR: "Valor associado ao item de remuneração",
				Constraints:     Constraints{Required: true},
			},
		},
	},
}

var metadadosResource = Resource{
	Name:        metadadosResourceName,
	Description: "Metadados associados a uma determinada coleta (mês, ano, órgão), incluindo o índice de transparência DadosJusBR.",
	Path:        metadadosFileName,
	Profile:     "tabular-data-resource",
	Schema: Schema{
		ForeignKeys: []ForeignKey{
			{"orgao", FKRef{coletaResourceName, "orgao"}},
			{"mes", FKRef{coletaResourceName, "mes"}},
			{"ano", FKRef{coletaResourceName, "ano"}},
		},
		Fields: []Field{
			{
				Name:            "orgao",
				Type:            "string",
				Format:          "default",
				Title:           "Órgão",
				Description:     "The ID of agency",
				DescriptionPTBR: "A sigla da agência",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "mes",
				Type:            "integer",
				Format:          "default",
				Title:           "Mês",
				Description:     "Month that data collection refers to",
				DescriptionPTBR: "O mês que a os dados foram coletados se referem",
				Constraints:     Constraints{Required: true, Maximum: "12", Minimum: "1"},
			},
			{
				Name:            "ano",
				Type:            "integer",
				Format:          "default",
				Title:           "Ano",
				Description:     "Year that data collection refers to",
				DescriptionPTBR: "O ano que a os dados foram coletados se referem",
				Constraints:     Constraints{Required: true, Minimum: "2018"},
			},
			{
				Name:            "formato_aberto",
				Type:            "boolean",
				Title:           "Formato Aberto",
				Description:     "Is the data available in an open format?",
				DescriptionPTBR: "Os dados são disponibilizados em formato aberto?",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "acesso",
				Type:            "string",
				Format:          "default",
				Title:           "Acesso",
				Description:     "Can we build a URL that leads to the data download based on agency/month/year?",
				DescriptionPTBR: "Conseguimos prever/construir uma URL com base no órgão/mês/ano que leve ao download do dado?",
				Constraints:     Constraints{Required: true, Enum: []string{"ACESSO_DIRETO", "RASPAGEM_DIFICULTADA", "NECESSITA_SIMULACAO_USUARIO"}},
			},
			{
				Name:            "extensao",
				Type:            "string",
				Format:          "default",
				Title:           "Extensão",
				Description:     "Extension of the original data file.",
				DescriptionPTBR: "Extensao do arquivo de dados, ex: CSV, JSON, XLS, etc",
				Constraints:     Constraints{Required: true, Enum: []string{"PDF", "ODS", "XLS", "JSON", "CSV", "HTML"}},
			},
			{
				Name:            "estritamente_tabular",
				Type:            "boolean",
				Title:           "Estritamente Tabular",
				Description:     "Is the available data tidy?",
				DescriptionPTBR: "Órgãos que disponibilizam dados limpos (tidy data)",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "formato_consistente",
				Type:            "boolean",
				Title:           "Formato Consistente",
				Description:     "Has the data changed since last month?",
				DescriptionPTBR: "Órgão alterou a forma de expor seus dados entre o mês em questão e o mês anterior?",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "tem_matricula",
				Type:            "boolean",
				Title:           "Tem Matrícula",
				Description:     "Does the agency publicize the employee id?",
				DescriptionPTBR: "Órgão disponibiliza matrícula do servidor?",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "tem_lotacao",
				Type:            "boolean",
				Title:           "Tem Lotação",
				Description:     "Does the agency publicize the employee workplace?",
				DescriptionPTBR: "Órgão disponibiliza lotação do servidor?",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "tem_cargo",
				Type:            "boolean",
				Title:           "Tem Cargo",
				Description:     "Does the agency publicize the employee role?",
				DescriptionPTBR: "Órgão disponibiliza a função do servidor?",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "detalhamento_receita_base",
				Type:            "string",
				Format:          "default",
				Title:           "Detalhamento Receita Base",
				Description:     "Detail level of the base remuneration (wage).",
				DescriptionPTBR: "Quão detalhado é a publicação da receita base.",
				Constraints:     Constraints{Required: true, Enum: []string{"AUSENCIA", "SUMARIZADO", "DETALHADO"}},
			},
			{
				Name:            "detalhamento_outras_receitas",
				Type:            "string",
				Format:          "default",
				Title:           "Detalhamento Outras Despesas",
				Description:     "Detail level of other remunerations.",
				DescriptionPTBR: "Quão detalhado é a publicação das demais receitas.",
				Constraints:     Constraints{Required: true, Enum: []string{"AUSENCIA", "SUMARIZADO", "DETALHADO"}},
			},
			{
				Name:            "detalhamento_descontos",
				Type:            "string",
				Format:          "default",
				Title:           "Detalhamento Descontos",
				Description:     "Detail level of the base discounts.",
				DescriptionPTBR: "Quão detalhado é a publicação dos descontos.",
				Constraints:     Constraints{Required: true, Enum: []string{"AUSENCIA", "SUMARIZADO", "DETALHADO"}},
			},
			{
				Name:            "indice_completude",
				Type:            "number",
				Format:          "default",
				Title:           "Componente completude do índice de transparência DadosJusBR",
				Description:     "DadosJusBR Transparency Index Component which results from metadata related to data availability",
				DescriptionPTBR: "Componente do índice de transparência resultante da análise dos metadados relacionados a disponibilidade dos dados",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "indice_facilidade",
				Type:            "number",
				Format:          "default",
				Title:           "Componente facilidade do índice de transparência DadosJusBR",
				Description:     "DadosJusBR Transparency Index Component which results from metadata related to how easy is to access the data",
				DescriptionPTBR: "Componente do índice de transparência resultante da análise dos metadados relacionados a dificuldade para acessar os dados que estão disponíveis",
				Constraints:     Constraints{Required: true},
			},
			{
				Name:            "indice_transparencia",
				Type:            "number",
				Format:          "default",
				Title:           "Índice de Transparência DadosJusBR",
				Description:     "DadosJusBR Transparency Score, calculated from the availability and access components",
				DescriptionPTBR: "Nota final, calculada utilizada os componentes de disponibilidade e dificuldade",
				Constraints:     Constraints{Required: true},
			},
		},
	},
}

type dadosjusbrDescriptorStruct struct {
	Profile      string        `json:"profile,omitempty"`
	Resources    []Resource    `json:"resources,omitempty"`
	Keywords     []string      `json:"keywords,omitempty"`
	Name         string        `json:"name,omitempty"`
	Title        string        `json:"title,omitempty"`
	Description  string        `json:"description,omitempty"`
	Homepage     string        `json:"homepage,omitempty"`
	Version      string        `json:"version,omitempty"`
	Contributors []Contributor `json:"contributors,omitempty"`
	Licenses     []License     `json:"licenses,omitempty"`
}

type Constraints struct {
	Required bool     `json:"required,omitempty"`
	Minimum  string   `json:"minimum,omitempty"`
	Maximum  string   `json:"maximum,omitempty"`
	Enum     []string `json:"enum,omitempty"`
}

type Field struct {
	Name            string      `json:"name,omitempty"`
	Type            string      `json:"type,omitempty"`
	Format          string      `json:"format,omitempty"`
	Title           string      `json:"title,omitempty"`
	Description     string      `json:"description,omitempty"`
	DescriptionPTBR string      `json:"description-ptbr,omitempty"`
	Constraints     Constraints `json:"constraints,omitempty"`
	BareNumber      bool        `json:"bareNumber,omitempty"`
}

type Schema struct {
	Fields      []Field      `json:"fields,omitempty"`
	PrimaryKey  string       `json:"primaryKey,omitempty"`
	ForeignKeys []ForeignKey `json:"foreignKeys,omitempty"`
}

type Resource struct {
	Name        string `json:"name,omitempty"`
	Path        string `json:"path,omitempty"`
	Profile     string `json:"profile,omitempty"`
	Description string `json:"description,omitempty"`
	Schema      Schema `json:"schema,omitempty"`
}

type Contributor struct {
	Title string `json:"title,omitempty"`
	Role  string `json:"role,omitempty"`
}

type License struct {
	Name  string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
	Path  string `json:"path,omitempty"`
}

type ForeignKey struct {
	Fields    string `json:"fields,omitempty"`
	Reference FKRef  `json:"reference,omitempty"`
}

type FKRef struct {
	Resource string `json:"resource,omitempty"`
	Fields   string `json:"fields,omitempty"`
}
