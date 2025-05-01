package nfev4

// InfNFeSupl - Informações suplementares
type InfNFeSupl struct {
	QRCode   string `xml:"qrCode"`
	URLChave string `xml:"urlChave,omitempty"`
}

// InfSolicNFF - Informações de solicitação de NFF
type InfSolicNFF struct {
	XSolic string `xml:"xSolic"`
}
