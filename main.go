package main

import (
	"fmt"
	"os"

	"pkg/util"
)

func main() {
	// Abrir arquivo XML
	file, err := os.Open("data/nf_teste.xml")
	if err != nil {
		fmt.Printf("Erro ao abrir arquivo: %v", err)
	}
	defer file.Close()

	// Parsear NFe
	nfeParsed, err := util.ParseNFeFromXML(file)
	if err != nil {
		fmt.Printf("Erro ao parsear NFe: %v", err)
	}

	fmt.Printf("NFe parseada com sucesso:\n")
	fmt.Printf("Emitente: %s\n", nfeParsed.InfNFe.Emit.XNome)
	fmt.Printf("Número: %d\n", nfeParsed.InfNFe.Ide.NNF)
	fmt.Printf("Valor Total: %.2f\n", nfeParsed.InfNFe.Total.ICMSTot.VNF)

	for _, det := range nfeParsed.InfNFe.Det {
		fmt.Printf("Produto %d: %s - R$ %.2f\n", det.NItem, det.Prod.XProd, det.Prod.VProd)
	}
}
