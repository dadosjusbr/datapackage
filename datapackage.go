package datapackage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dadosjusbr/proto/coleta"
	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/tableschema-go/csv"
	"github.com/gocarina/gocsv"
)

const (
	coletaFileName       = "coleta.csv"                  // hardcoded in datapackage_descriptor.json
	folhaFileName        = "contra_cheque.csv"           // hardcoded in datapackage_descriptor.json
	remuneracaoFileName  = "remuneracao.csv"             // hardcoded in datapackage_descriptor.json
	metadadosFileName    = "metadados.csv"               // hardcoded in datapackage_descriptor.json
	descriptorFileName   = "datapackage_descriptor.json" // name of datapackage descriptor
	coletaResource       = "coleta"                      // hardcoded in datapackage_descriptor.json
	contrachequeResource = "contra_cheque"               // hardcoded in datapackage_descriptor.json
	remuneracaoResource  = "remuneracao"                 // hardcoded in datapackage_descriptor.json
	metadadosResource    = "metadados"                   // hardcoded in datapackage_descriptor.json
)

func NewResultadoColetaCSV(rc *coleta.ResultadoColeta) ResultadoColeta_CSV {
	var coleta Coleta_CSV
	var remuneracoes []Remuneracao_CSV
	var folha []ContraCheque_CSV

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
			remuneracoes = append(remuneracoes, remuneracao)
		}
		folha = append(folha, contraCheque)
	}

	return ResultadoColeta_CSV{
		Coleta:       append([]Coleta_CSV{}, coleta),
		Remuneracoes: remuneracoes,
		Folha:        folha,
		Metadados:    append([]Metadados_CSV{}, metadados),
	}
}

func Zip(outputPath string, descriptorPath string, rc *ResultadoColeta_CSV, cleanup bool) error {
	outDir := filepath.Dir(outputPath)
	coletaCSVPath := filepath.Join(outDir, coletaFileName)
	folhaCSVPath := filepath.Join(outDir, folhaFileName)
	remuneracaoCSVPath := filepath.Join(outDir, remuneracaoFileName)
	metadadosCSVPath := filepath.Join(outDir, metadadosFileName)

	defer func() {
		if cleanup {
			os.Remove(coletaCSVPath)
			os.Remove(folhaCSVPath)
			os.Remove(remuneracaoCSVPath)
			os.Remove(metadadosCSVPath)
		}
	}()

	// Creating coleta csv
	if err := toCSVFile(rc.Coleta, coletaCSVPath); err != nil {
		return fmt.Errorf("error creating Coleta CSV (%s):%q", coletaCSVPath, err)
	}

	// Creating contracheque csv
	if err := toCSVFile(rc.Folha, folhaCSVPath); err != nil {
		return fmt.Errorf("error creating Folha de pagamento CSV (%s):%q", folhaCSVPath, err)
	}

	// Creating remuneracao csv
	if err := toCSVFile(rc.Remuneracoes, remuneracaoCSVPath); err != nil {
		return fmt.Errorf("error creating Remuneração CSV (%s):%q", remuneracaoCSVPath, err)
	}

	// Creating metadata csv
	if err := toCSVFile(rc.Metadados, metadadosCSVPath); err != nil {
		return fmt.Errorf("error creating Metadados CSV (%s):%q", metadadosCSVPath, err)
	}

	c, err := ioutil.ReadFile(descriptorPath)
	if err != nil {
		return fmt.Errorf("error reading %s:%q", descriptorPath, err)
	}
	var desc map[string]interface{}
	if err := json.Unmarshal(c, &desc); err != nil {
		return fmt.Errorf("error unmarshaling datapackage descriptor (%s):%q", descriptorPath, err)
	}
	pkg, err := datapackage.New(desc, outDir)
	if err != nil {
		return fmt.Errorf("error create datapackage:%q", err)
	}
	if err := pkg.Zip(outputPath); err != nil {
		return fmt.Errorf("error zipping datapackage (%s):%q", outputPath, err)
	}
	return err
}

func Load(path string) (ResultadoColeta_CSV, error) {
	pkg, err := datapackage.Load(path)
	if err != nil {
		return ResultadoColeta_CSV{}, fmt.Errorf("error loading datapackage (%s):%q", path, err)
	}

	coleta := pkg.GetResource(coletaResource)
	if coleta == nil {
		return ResultadoColeta_CSV{}, fmt.Errorf("resource coleta not found in package %s", path)
	}
	var coleta_CSV []Coleta_CSV
	if err := coleta.Cast(&coleta_CSV, csv.LoadHeaders()); err != nil {
		return ResultadoColeta_CSV{}, fmt.Errorf("failed to cast Coleta_CSV: %s", err)
	}

	contracheque := pkg.GetResource(contrachequeResource)
	if contracheque == nil {
		return ResultadoColeta_CSV{}, fmt.Errorf("resource contra_cheque not found in package %s", path)
	}
	var contracheque_CSV []ContraCheque_CSV
	if err := contracheque.Cast(&contracheque_CSV, csv.LoadHeaders()); err != nil {
		return ResultadoColeta_CSV{}, fmt.Errorf("failed to cast ContraCheque_CSV: %s", err)
	}

	remuneracao := pkg.GetResource(remuneracaoResource)
	if remuneracao == nil {
		return ResultadoColeta_CSV{}, fmt.Errorf("resource remuneracao not found in package %s", path)
	}
	var remuneracao_CSV []Remuneracao_CSV
	if err := remuneracao.Cast(&remuneracao_CSV, csv.LoadHeaders()); err != nil {
		return ResultadoColeta_CSV{}, fmt.Errorf("failed to cast Remuneracao_CSV: %s", err)
	}

	metadados := pkg.GetResource(metadadosResource)
	if metadados == nil {
		return ResultadoColeta_CSV{}, fmt.Errorf("resource metadados not found in package %s", path)
	}
	var metadados_CSV []Metadados_CSV
	if err := metadados.Cast(&metadados_CSV, csv.LoadHeaders()); err != nil {
		return ResultadoColeta_CSV{}, fmt.Errorf("failed to cast Metadados_CSV: %s", err)
	}

	return ResultadoColeta_CSV{
		Coleta:       coleta_CSV,
		Remuneracoes: remuneracao_CSV,
		Folha:        contracheque_CSV,
		Metadados:    metadados_CSV,
	}, nil
}

// toCSVFile dumps the payroll into a file using the CSV format.
func toCSVFile(in interface{}, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating CSV file(%s):%q", path, err)
	}
	defer f.Close()
	return gocsv.MarshalFile(in, f)
}

// fromCSVFile gets from CSV to a certain struct.
func fromCSVFile(in interface{}, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	return gocsv.UnmarshalFile(f, in)
}
