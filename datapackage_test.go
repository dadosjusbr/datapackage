package datapackage

import (
	"testing"
	"time"
)

func TestLoad_Success(t *testing.T) {
	rc, err := Load("test_datapackage.zip")
	if err != nil {
		t.Errorf("want no errr, got: %v", err)
	}
	t.Run("CheckColeta", func(t *testing.T) {
		if len(rc.Coleta) != 1 {
			t.Errorf("want 1 coleta, got: %d", len(rc.Coleta))
		}
		pt, err := time.Parse(time.RFC3339Nano, "2021-11-28T14:51:05.35811Z")
		if err != nil {
			t.Errorf("want no errr, got: %v", err)
		}
		want := Coleta_CSV{
			ChaveColeta:        "tjal/02/2020",
			Orgao:              "tjal",
			Mes:                2,
			Ano:                2020,
			TimestampColeta:    pt,
			RepositorioColetor: "https://github.com/dadosjusbr/coletor-cnj",
			VersaoColetor:      "unspecified",
		}
		if rc.Coleta[0] != want {
			t.Errorf("want %+v, got: %+v", want, rc.Remuneracoes[0])
		}
	})

	t.Run("CheckContraCheque", func(t *testing.T) {
		if len(rc.Folha) != 214 { // quantidade de contra-cheques
			t.Errorf("want 1 folha, got: %d", len(rc.Folha))
		}
		want := ContraCheque_CSV{
			ChaveColeta:    "tjal/02/2020",
			IdContraCheque: "tjal/02/2020/1",
			Nome:           "ADALBERTO CORREIA DE LIMA",
			Tipo:           "MEMBRO",
			Ativo:          true,
		}
		if rc.Folha[0] != want {
			t.Errorf("want %+v, got: %+v", want, rc.Remuneracoes[0])
		}
	})

	t.Run("CheckMetadados", func(t *testing.T) {
		if len(rc.Metadados) != 1 {
			t.Errorf("want 1 metadados, got: %d", len(rc.Metadados))
		}
		want := Metadados_CSV{
			ChaveColeta:                "tjal/02/2020",
			NaoRequerLogin:             true,
			NaoRequerCaptcha:           true,
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
		}
		if rc.Metadados[0] != want {
			t.Errorf("want %+v, got: %+v", want, rc.Metadados[0])
		}
	})

	t.Run("CheckRemuneracao", func(t *testing.T) {
		if len(rc.Remuneracoes) != 5354 { // quantidade de entradas totais dos contra-cheques
			t.Errorf("want 5355 remunerações, got: %d", len(rc.Remuneracoes))
		}
		want := Remuneracao_CSV{
			ChaveColeta:    "tjal/02/2020",
			IdContraCheque: "tjal/02/2020/1",
			Natureza:       "R",
			Categoria:      "contracheque",
			Item:           "Subsídio",
			Valor:          35462.22,
		}
		if rc.Remuneracoes[0] != want {
			t.Errorf("want %+v, got: %+v", want, rc.Remuneracoes[0])
		}
	})
}
