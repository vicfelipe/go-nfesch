package validator

import (
	"errors"

	"pkg/nfev4"
)

// validateNFe - Validações básicas da estrutura NFe
func ValidateNFe(nfe *nfev4.NFe) error {
	if nfe.Versao != "4.00" {
		return errors.New("versão da NFe deve ser 4.00")
	}

	if nfe.InfNFe.Id == "" {
		return errors.New("ID da NFe é obrigatório")
	}

	if nfe.InfNFe.Ide.Mod != 55 {
		return errors.New("modelo da NFe deve ser 55")
	}

	// Adicione mais validações conforme necessário
	return nil
}
