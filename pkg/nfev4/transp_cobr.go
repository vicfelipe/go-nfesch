package nfev4

// Transporte
type Transporta struct {
	CNPJ   string `xml:"CNPJ,omitempty"`
	CPF    string `xml:"CPF,omitempty"`
	XNome  string `xml:"xNome"`
	IE     string `xml:"IE,omitempty"`
	XEnder string `xml:"xEnder"`
	XMun   string `xml:"xMun"`
	UF     string `xml:"UF"`
}

type RetTransp struct {
	VServ    float64 `xml:"vServ"`
	VBCRet   float64 `xml:"vBCRet"`
	PICMSRet float64 `xml:"pICMSRet"`
	VICMSRet float64 `xml:"vICMSRet"`
	CFOP     string  `xml:"CFOP"`
	CMunFG   string  `xml:"cMunFG"`
}

type VeicTransp struct {
	Placa string `xml:"placa"`
	UF    string `xml:"UF"`
	RNTC  string `xml:"RNTC,omitempty"`
}

type Reboque struct {
	Placa string `xml:"placa"`
	UF    string `xml:"UF"`
	RNTC  string `xml:"RNTC,omitempty"`
	Vagao string `xml:"vagao,omitempty"`
	Balsa string `xml:"balsa,omitempty"`
}

type Vol struct {
	QVol   int      `xml:"qVol,omitempty"`
	Esp    string   `xml:"esp,omitempty"`
	Marca  string   `xml:"marca,omitempty"`
	NVol   string   `xml:"nVol,omitempty"`
	PesoL  float64  `xml:"pesoL,omitempty"`
	PesoB  float64  `xml:"pesoB,omitempty"`
	Lacres []Lacres `xml:"lacres,omitempty"`
}

type Lacres struct {
	NLacre string `xml:"nLacre"`
}

// Cobrança
type Cobr struct {
	Fat *Fat  `xml:"fat,omitempty"`
	Dup []Dup `xml:"dup,omitempty"`
}

type Fat struct {
	NFat  string  `xml:"nFat,omitempty"`
	VOrig float64 `xml:"vOrig,omitempty"`
	VDesc float64 `xml:"vDesc,omitempty"`
	VLiq  float64 `xml:"vLiq,omitempty"`
}

type Dup struct {
	NDup  string  `xml:"nDup,omitempty"`
	DVenc string  `xml:"dVenc,omitempty"`
	VDup  float64 `xml:"vDup,omitempty"`
}

// Pagamento
type Card struct {
	CNPJ      string `xml:"CNPJ,omitempty"`
	TBand     string `xml:"tBand,omitempty"`
	CAut      string `xml:"cAut,omitempty"`
	TPIntegra int    `xml:"tpIntegra,omitempty"`
}
