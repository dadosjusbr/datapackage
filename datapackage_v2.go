package datapackage

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/dadosjusbr/proto/coleta"
	"github.com/frictionlessdata/datapackage-go/datapackage"
	"github.com/frictionlessdata/datapackage-go/validator"
	"github.com/frictionlessdata/tableschema-go/csv"
)

const (
	contrachequeFileNameV2     = "contracheque.csv" // hardcoded in datapackage_descriptor.json
	contrachequeResourceNameV2 = "contracheque"     // hardcoded in datapackage_descriptor.json
)

// Essa versão passou a ser utilizada a partir de julho de 2023
func NewResultadoColetaCSV_V2(rc *coleta.ResultadoColeta) ResultadoColeta_CSV_V2 {
	var coletaCSV Coleta_CSV_V2
	var remuneracoes []Remuneracao_CSV_V2
	var folha []Contracheque_CSV_V2

	coletaCSV.ChaveColeta = rc.Coleta.ChaveColeta
	coletaCSV.Orgao = rc.Coleta.Orgao
	coletaCSV.Mes = rc.Coleta.Mes
	coletaCSV.Ano = rc.Coleta.Ano
	coletaCSV.TimestampColeta = rc.Coleta.TimestampColeta.AsTime()
	coletaCSV.RepositorioColetor = rc.Coleta.RepositorioColetor
	coletaCSV.VersaoColetor = rc.Coleta.VersaoColetor
	coletaCSV.RepositorioParser = rc.Coleta.RepositorioParser
	coletaCSV.VersaoParser = rc.Coleta.VersaoParser

	var metadados Metadados_CSV_V2
	metadados.Orgao = rc.Coleta.Orgao
	metadados.Mes = rc.Coleta.Mes
	metadados.Ano = rc.Coleta.Ano
	metadados.FormatoAberto = rc.Metadados.FormatoAberto
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

	for i, v := range rc.Folha.ContraCheque {
		var contraCheque Contracheque_CSV_V2
		contraCheque.IdContracheque = i + 1
		contraCheque.Orgao = rc.Coleta.Orgao
		contraCheque.Mes = rc.Coleta.Mes
		contraCheque.Ano = rc.Coleta.Ano
		contraCheque.Nome = v.Nome
		contraCheque.Matricula = v.Matricula
		contraCheque.Funcao = v.Funcao
		contraCheque.LocalTrabalho = v.LocalTrabalho
		var salario, beneficios, descontos float64
		for _, k := range v.Remuneracoes.Remuneracao {
			var remuneracao Remuneracao_CSV_V2
			remuneracao.IdContracheque = i + 1
			remuneracao.Orgao = rc.Coleta.Orgao
			remuneracao.Mes = rc.Coleta.Mes
			remuneracao.Ano = rc.Coleta.Ano
			remuneracao.Categoria = k.Categoria
			remuneracao.Item = k.Item
			remuneracao.Valor = k.Valor
			if k.TipoReceita == coleta.Remuneracao_B && k.Natureza == coleta.Remuneracao_R {
				salario += k.Valor
				remuneracao.Tipo = "R/B"
			} else if k.TipoReceita == coleta.Remuneracao_O && k.Natureza == coleta.Remuneracao_R {
				beneficios += k.Valor
				remuneracao.Tipo = "R/O"
			} else if k.Natureza == coleta.Remuneracao_D {
				descontos += math.Abs(k.Valor)
				remuneracao.Tipo = "D"
			}
			remuneracoes = append(remuneracoes, remuneracao)
		}
		contraCheque.Salario = salario
		contraCheque.Beneficios = beneficios
		contraCheque.Descontos = descontos
		contraCheque.Remuneracao = salario + beneficios - descontos

		if v.Ativo && strings.Contains(rc.Coleta.Orgao, "mp") {
			contraCheque.Situacao = "A"
		}

		folha = append(folha, contraCheque)
	}

	return ResultadoColeta_CSV_V2{
		Coleta:       append([]Coleta_CSV_V2{}, coletaCSV),
		Remuneracoes: remuneracoes,
		Folha:        folha,
		Metadados:    append([]Metadados_CSV_V2{}, metadados),
	}
}

func ZipV2(outputPath string, rc ResultadoColeta_CSV_V2, cleanup bool) error {
	outDir := filepath.Dir(outputPath)
	coletaCSVPath := filepath.Join(outDir, coletaFileName)
	folhaCSVPath := filepath.Join(outDir, contrachequeFileNameV2)
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

	// Creating datapackage
	desc, err := descriptorMapV2()
	if err != nil {
		return err
	}
	pkg, err := datapackage.New(desc, outDir)
	if err != nil {
		return fmt.Errorf("error create datapackage:%q", err)
	}
	if err := pkg.Zip(outputPath); err != nil {
		return fmt.Errorf("error zipping datapackage (%s):%q", outputPath, err)
	}
	return nil
}

func descriptorMapV2() (map[string]interface{}, error) {
	b, err := json.Marshal(dadosjusbrDescriptorV2)
	if err != nil {
		return nil, fmt.Errorf("error marshalling dadosjusbr descriptor:%w", err)
	}
	var desc map[string]interface{}
	if err := json.Unmarshal(b, &desc); err != nil {
		return nil, fmt.Errorf("error converting datapackage descriptor into map:%w", err)
	}
	return desc, nil
}

func LoadV2(path string) (ResultadoColeta_CSV_V2, error) {
	pkg, err := datapackage.Load(path, validator.InMemoryLoader())
	if err != nil {
		return ResultadoColeta_CSV_V2{}, fmt.Errorf("error loading datapackage (%s):%q", path, err)
	}

	coleta := pkg.GetResource(coletaResourceName)
	if coleta == nil {
		return ResultadoColeta_CSV_V2{}, fmt.Errorf("resource coleta not found in package %s", path)
	}
	var coleta_CSV []Coleta_CSV_V2
	if err := coleta.Cast(&coleta_CSV, csv.LoadHeaders()); err != nil {
		return ResultadoColeta_CSV_V2{}, fmt.Errorf("failed to cast Coleta_CSV: %s", err)
	}

	contracheque := pkg.GetResource(contrachequeResourceNameV2)
	if contracheque == nil {
		return ResultadoColeta_CSV_V2{}, fmt.Errorf("resource contra_cheque not found in package %s", path)
	}
	var contracheque_CSV []Contracheque_CSV_V2
	if err := contracheque.Cast(&contracheque_CSV, csv.LoadHeaders()); err != nil {
		return ResultadoColeta_CSV_V2{}, fmt.Errorf("failed to cast ContraCheque_CSV: %s", err)
	}

	remuneracao := pkg.GetResource(remuneracaoResourceName)
	if remuneracao == nil {
		return ResultadoColeta_CSV_V2{}, fmt.Errorf("resource remuneracao not found in package %s", path)
	}
	var remuneracao_CSV []Remuneracao_CSV_V2
	if err := remuneracao.Cast(&remuneracao_CSV, csv.LoadHeaders()); err != nil {
		return ResultadoColeta_CSV_V2{}, fmt.Errorf("failed to cast Remuneracao_CSV: %s", err)
	}

	metadados := pkg.GetResource(metadadosResourceName)
	if metadados == nil {
		return ResultadoColeta_CSV_V2{}, fmt.Errorf("resource metadados not found in package %s", path)
	}
	var metadados_CSV []Metadados_CSV_V2
	if err := metadados.Cast(&metadados_CSV, csv.LoadHeaders()); err != nil {
		return ResultadoColeta_CSV_V2{}, fmt.Errorf("failed to cast Metadados_CSV: %s", err)
	}

	return ResultadoColeta_CSV_V2{
		Coleta:       coleta_CSV,
		Remuneracoes: remuneracao_CSV,
		Folha:        contracheque_CSV,
		Metadados:    metadados_CSV,
	}, nil
}
