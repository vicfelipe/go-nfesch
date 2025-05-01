package nfev4

// ObsCont - Observações do contribuinte
type ObsCont struct {
	XCampo string `xml:"xCampo,attr"`
	XTexto string `xml:"xTexto"`
}

// ObsFisco - Observações do fisco
type ObsFisco struct {
	XCampo string `xml:"xCampo,attr"`
	XTexto string `xml:"xTexto"`
}

// ProcRef - Processo referenciado
type ProcRef struct {
	NProc   string `xml:"nProc"`
	IndProc int    `xml:"indProc"`
}
