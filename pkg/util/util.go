package util

import (
	"bytes"
	"encoding/xml"
	"io"

	qrcode "github.com/skip2/go-qrcode"
	"pkg/nfev4"
)

// ParseNFeFromXML - Agora lida com ou sem declaração XML
func ParseNFeFromXML(xmlData io.Reader) (*nfev4.NFe, error) {
	// Ler todo o conteúdo para verificar se tem declaração XML
	content, err := io.ReadAll(xmlData)
	if err != nil {
		return nil, err
	}

	// Adicionar declaração XML se não existir
	if !bytes.HasPrefix(content, []byte("<?xml")) {
		content = append([]byte(`<?xml version="1.0" encoding="UTF-8"?>`+"\n"), content...)
	}

	// Parsear o conteúdo
	var nfe nfev4.NFe
	decoder := xml.NewDecoder(bytes.NewReader(content))
	decoder.Strict = false

	err = decoder.Decode(&nfe)
	if err != nil {
		return nil, err
	}

	return &nfe, nil
}

func GenerateQRCodeImage(content string, size int) ([]byte, error) {
	return qrcode.Encode(content, qrcode.Medium, size)
}
