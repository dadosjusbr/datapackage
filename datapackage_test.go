package datapackage

import (
	"os"
	"testing"
	"time"

	"github.com/dadosjusbr/proto/coleta"
	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// variáveis usadas pelos testes.
var (
	pt, _      = time.Parse(time.RFC3339Nano, "2021-11-28T14:51:05.35811Z")
	coletaTest = Coleta_CSV_V2{
		ChaveColeta:        "tjal/02/2020",
		Orgao:              "tjal",
		Mes:                2,
		Ano:                2020,
		TimestampColeta:    pt,
		RepositorioColetor: "https://github.com/dadosjusbr/coletor-cnj",
		VersaoColetor:      "unspecified",
		RepositorioParser:  "https://github.com/dadosjusbr/parser-cnj",
		VersaoParser:       "unspecified",
	}
	contraChequeTest = ContraCheque_CSV_V2{
		IdContraCheque: 1,
		Orgao:          "tjal",
		Mes:            2,
		Ano:            2020,
		Nome:           "ADALBERTO CORREIA DE LIMA",
		Salario:        35462.22,
		Remuneracao:    35462.22,
	}
	metadadosTest = Metadados_CSV_V2{
		Orgao:                      "tjal",
		Mes:                        2,
		Ano:                        2020,
		FormatoAberto:              false,
		Acesso:                     "NECESSITA_SIMULACAO_USUARIO",
		Extensao:                   "XLS",
		EstritamenteTabular:        true,
		FormatoConsistente:         true,
		TemMatricula:               false,
		TemLotacao:                 false,
		TemCargo:                   false,
		DetalhamentoReceitaBase:    "DETALHADO",
		DetalhamentoOutrasReceitas: "DETALHADO",
		DetalhamentoDescontos:      "DETALHADO",
		IndiceCompletude:           0.61538464,
		IndiceFacilidade:           0.5,
		IndiceTransparencia:        0.8,
	}
	remuneracaoTest = Remuneracao_CSV_V2{
		IdContraCheque: 1,
		Orgao:          "tjal",
		Mes:            2,
		Ano:            2020,
		Tipo:           "R/B",
		Categoria:      "contracheque",
		Item:           "Subsídio",
		Valor:          35462.22,
	}
	resultadoColeta = ResultadoColeta_CSV_V2{
		Coleta:       []Coleta_CSV_V2{coletaTest},
		Folha:        []ContraCheque_CSV_V2{contraChequeTest},
		Metadados:    []Metadados_CSV_V2{metadadosTest},
		Remuneracoes: []Remuneracao_CSV_V2{remuneracaoTest},
	}
)

func TestNewResultadoColetaCSV(t *testing.T) {
	in := coleta.ResultadoColeta{
		Coleta: &coleta.Coleta{
			ChaveColeta:        "tjal/02/2020",
			Mes:                2,
			Ano:                2020,
			Orgao:              "tjal",
			RepositorioColetor: "https://github.com/dadosjusbr/coletor-cnj",
			VersaoColetor:      "unspecified",
			RepositorioParser:  "https://github.com/dadosjusbr/parser-cnj",
			VersaoParser:       "unspecified",
			TimestampColeta:    timestamppb.New(pt),
		},
		Folha: &coleta.FolhaDePagamento{
			ContraCheque: []*coleta.ContraCheque{{
				IdContraCheque: "tjal/02/2020/1",
				ChaveColeta:    "tjal/02/2020",
				Nome:           "ADALBERTO CORREIA DE LIMA",
				Tipo:           coleta.ContraCheque_MEMBRO,
				Ativo:          true,
				Remuneracoes: &coleta.Remuneracoes{
					Remuneracao: []*coleta.Remuneracao{{
						Natureza:  coleta.Remuneracao_R,
						Valor:     35462.22,
						Item:      "Subsídio",
						Categoria: "contracheque",
					}},
				},
			}},
		},
		Metadados: &coleta.Metadados{
			FormatoAberto:       false,
			Acesso:              coleta.Metadados_NECESSITA_SIMULACAO_USUARIO,
			Extensao:            coleta.Metadados_XLS,
			EstritamenteTabular: true,
			FormatoConsistente:  true,
			TemMatricula:        false,
			TemLotacao:          false,
			TemCargo:            false,
			ReceitaBase:         coleta.Metadados_DETALHADO,
			OutrasReceitas:      coleta.Metadados_DETALHADO,
			Despesas:            coleta.Metadados_DETALHADO,
			IndiceCompletude:    0.61538464,
			IndiceFacilidade:    0.5,
			IndiceTransparencia: 0.8,
		},
	}
	assert.Equal(t, resultadoColeta, NewResultadoColetaCSV(&in))
}

func TestLoad_Success(t *testing.T) {
	rc, err := Load("test_datapackage_load.zip")
	assert.NoError(t, err, "want no erro on Load")

	t.Run("CheckColeta", func(t *testing.T) {
		assert.Equal(t, 1, len(rc.Coleta))
		assert.Equal(t, coletaTest, rc.Coleta[0])
	})

	t.Run("CheckContraCheque", func(t *testing.T) {
		assert.Equal(t, 214, len(rc.Folha))
		assert.Equal(t, contraChequeTest, rc.Folha[0])
	})

	t.Run("CheckMetadados", func(t *testing.T) {
		assert.Equal(t, 1, len(rc.Metadados))
		assert.Equal(t, metadadosTest, rc.Metadados[0])
	})

	t.Run("CheckRemuneracao", func(t *testing.T) {
		assert.Equal(t, 5354, len(rc.Remuneracoes))
		assert.Equal(t, remuneracaoTest, rc.Remuneracoes[0])
	})
}

// func TestLoad_Remote(t *testing.T) {
// 	rc, err := Load("https://github.com/dadosjusbr/datapackage/raw/main/test_datapackage_load.zip")
// 	assert.NoError(t, err, "want no erro on Load")

// 	t.Run("CheckColeta", func(t *testing.T) {
// 		assert.Equal(t, 1, len(rc.Coleta))
// 		assert.Equal(t, coletaTest, rc.Coleta[0])
// 	})

// 	t.Run("CheckContraCheque", func(t *testing.T) {
// 		assert.Equal(t, 214, len(rc.Folha))
// 		assert.Equal(t, contraChequeTest, rc.Folha[0])
// 	})

// 	t.Run("CheckMetadados", func(t *testing.T) {
// 		assert.Equal(t, 1, len(rc.Metadados))
// 		assert.Equal(t, metadadosTest, rc.Metadados[0])
// 	})

// 	t.Run("CheckRemuneracao", func(t *testing.T) {
// 		assert.Equal(t, 5354, len(rc.Remuneracoes))
// 		assert.Equal(t, remuneracaoTest, rc.Remuneracoes[0])
// 	})
// }

func TestLoad_Error(t *testing.T) {
	testData := []struct {
		desc string
		path string
	}{
		{"FileNotFound", "fileNotFound"},
		{"MissingColeta", "test_datapackage_missing_coleta.zip"},
		{"MissingContracheque", "test_datapackage_missing_contracheque.zip"},
		{"MissingMetadados", "test_datapackage_missing_metadados.zip"},
		{"MissingRemuneracao", "test_datapackage_missing_remuneracao.zip"},
	}
	for _, d := range testData {
		t.Run(d.desc, func(t *testing.T) {
			_, err := Load(d.path)
			assert.Error(t, err)
		})
	}
}
func TestZip_Success(t *testing.T) {
	assert.NoError(t, Zip("datapackage_criado.zip", resultadoColeta, false), "want no err during Zip")
	defer func() {
		os.Remove(coletaFileName)
		os.Remove(contrachequeFileName)
		os.Remove(remuneracaoFileName)
		os.Remove(metadadosFileName)
		os.Remove("datapackage_criado.zip")
	}()

	t.Run("Contents", func(t *testing.T) {
		t.Run("CheckColeta", func(t *testing.T) {
			var got []Coleta_CSV_V2
			assert.NoError(t, fromCSVFile(&got, coletaFileName), "want no err during retrieving coleta csv")
			assert.Equal(t, 1, len(got))
			assert.Equal(t, coletaTest, got[0])
		})
		t.Run("ContraCheque", func(t *testing.T) {
			var got []ContraCheque_CSV_V2
			assert.NoError(t, fromCSVFile(&got, contrachequeFileName), "want no err during retrieving folha csv")
			assert.Equal(t, 1, len(got))
			assert.Equal(t, contraChequeTest, got[0])
		})

		t.Run("Metadados", func(t *testing.T) {
			var got []Metadados_CSV_V2
			assert.NoError(t, fromCSVFile(&got, metadadosFileName), "want no err during retrieving metadados csv")
			assert.Equal(t, 1, len(got))
			assert.Equal(t, metadadosTest, got[0])
		})

		t.Run("Remuneracoes", func(t *testing.T) {
			var got []Remuneracao_CSV_V2
			assert.NoError(t, fromCSVFile(&got, remuneracaoFileName), "want no err during retrieving remuneracoes csv")
			assert.Equal(t, 1, len(got))
			assert.Equal(t, remuneracaoTest, got[0])
		})
	})

	t.Run("CheckSchema", func(t *testing.T) {
		pkg, err := datapackage.Load("datapackage_criado.zip")
		assert.NoError(t, err)
		resNames := []struct {
			resName   string
			numFields int
		}{
			{coletaResourceName, len(coletaResource.Schema.Fields)},
			{contrachequeResourceName, len(contraChequeResource.Schema.Fields)},
			{remuneracaoResourceName, len(remuneracaoResource.Schema.Fields)},
			{metadadosResourceName, len(metadadosResource.Schema.Fields)},
		}
		for _, data := range resNames {
			t.Run(data.resName, func(t *testing.T) {
				res := pkg.GetResource(data.resName)
				assert.NotNil(t, res)
				sch, err := res.GetSchema()
				assert.NoError(t, err)
				assert.NoError(t, sch.Validate())
				assert.Equal(t, data.resName, res.Name())
				assert.Equal(t, data.numFields, len(sch.Fields))
				assert.Greater(t, len(sch.Fields), 0)
			})
		}
	})
}
