package datapackage
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/dadosjusbr/proto/coleta"
	"github.com/dadosjusbr/coletores/status"
	"github.com/frictionlessdata/datapackage-go/datapackage"
)

// Adicionar os scores a struct de Metadados existente 
func NewColetaCSV(rc *coleta.ResultadoColeta) *ResultadoColeta_CSV {
    var coleta Coleta_CSV
	var remuneracoes []*Remuneracao_CSV
	var folha []*ContraCheque_CSV

	coleta.ChaveColeta = rc.Coleta.ChaveColeta
	coleta.Orgao = rc.Coleta.Orgao
	coleta.Mes = rc.Coleta.Mes
	coleta.Ano = rc.Coleta.Ano
	coleta.TimestampColeta = rc.Coleta.TimestampColeta.AsTime()
	coleta.RepositorioColetor = rc.Coleta.RepositorioColetor
	coleta.VersaoColetor = rc.Coleta.VersaoColetor
	coleta.DirColetor = rc.Coleta.DirColetor

	var metadados Metadados_CSV
	metadados.ChaveColeta = rc.Coleta.ChaveColeta
	metadados.NaoRequerLogin = rc.Metadados.NaoRequerLogin
	metadados.NaoRequerCaptcha = rc.Metadados.NaoRequerCaptcha
	metadados.Acesso = rc.Metadados.Acesso.String()
	metadados.Extensao = rc.Metadados.Extensao.String()
	metadados.EstritamenteTabular = rc.Metadados.EstritamenteTabular
	metadados.FormatoConsistente = rc.Metadados.FormatoConsistente
	metadados.TemMatricula = rc.Metadados.TemMatricula
	metadados.TemLotacao = rc.Metadados.TemLotacao
	metadados.TemCargo = rc.Metadados.TemCargo
	metadados.DetalhamentoReceitaBase = rc.Metadados.ReceitaBase.String()
	metadados.DetalhamentoOutrasReceitas = rc.Metadados.OutrasReceitas.String()
	metadados.DetalhamentoDescontos = rc.Metadados.Despesas.String()
	metadados.IndiceCompletude = rc.Metadados.IndiceCompletude
	metadados.IndiceFacilidade = rc.Metadados.IndiceFacilidade
	metadados.IndiceTransparencia = rc.Metadados.IndiceTransparencia

	for _, v := range rc.Folha.ContraCheque {
		var contraCheque ContraCheque_CSV
		contraCheque.IdContraCheque = v.IdContraCheque
		contraCheque.ChaveColeta = v.ChaveColeta
		contraCheque.Nome = v.Nome
		contraCheque.Matricula = v.Matricula
		contraCheque.Funcao = v.Funcao
		contraCheque.Ativo = v.Ativo
		contraCheque.LocalTrabalho = v.LocalTrabalho
		contraCheque.Tipo = v.Tipo.String()
		for _, k := range v.Remuneracoes.Remuneracao {
			var remuneracao Remuneracao_CSV
			remuneracao.IdContraCheque = v.IdContraCheque
			remuneracao.ChaveColeta = v.ChaveColeta
			remuneracao.Natureza = k.Natureza.String()
			remuneracao.Categoria = k.Categoria
			remuneracao.Item = k.Item
			remuneracao.Valor = k.Valor
			remuneracoes = append(remuneracoes, &remuneracao)
		}
		folha = append(folha, &contraCheque)
	}

	return &ResultadoColeta_CSV{
		Coleta:       append([]*Coleta_CSV{}, &coleta),
		Remuneracoes: remuneracoes,
		Folha:        folha,
		Metadados:    append([]*Metadados_CSV{}, &metadados),
	}
}

// Atualiza o datapackage_descriptor e compacta novamente
func Zip(packageFileName, outputPath, zipFileName string) {
	c, err := ioutil.ReadFile(packageFileName)
	if err != nil {
		err = status.NewError(status.InvalidParameters, fmt.Errorf("error reading datapackge_descriptor.json:%q", err))
		status.ExitFromError(err)
	}

	var desc map[string]interface{}
	if err := json.Unmarshal(c, &desc); err != nil {
		err = status.NewError(status.InvalidParameters, fmt.Errorf("error unmarshaling datapackage descriptor:%q", err))
		status.ExitFromError(err)
	}

	pkg, err := datapackage.New(desc, ".")
	if err != nil {
		err = status.NewError(status.InvalidParameters, fmt.Errorf("error create datapackage:%q", err))
		status.ExitFromError(err)
	}

	// Packing CSV and package descriptor.
	zipName := filepath.Join(outputPath, fmt.Sprintf("%s.zip", zipFileName))
	if err := pkg.Zip(zipName); err != nil {
		err = status.NewError(status.SystemError, fmt.Errorf("error zipping datapackage (%s):%q", zipName, err))
		status.ExitFromError(err)
	}
}

func Load(path string) {
    pkg, err := datapackage.Load(path)

    if err != nil {
		err = status.NewError(status.SystemError, fmt.Errorf("error loading datapackage (%s):%q", path, err))
		status.ExitFromError(err)
	}

    fmt.Printf("Data package \"%s\" successfully created.\n", pkg.Descriptor()["name"])
}