package nfev4

// NFRef - Referência a documentos fiscais
type NFRef struct {
	RefNFe string  `xml:"refNFe,omitempty"`
	RefNF  *RefNF  `xml:"refNF,omitempty"`
	RefNFP *RefNFP `xml:"refNFP,omitempty"`
	RefCTe string  `xml:"refCTe,omitempty"`
	RefECF *RefECF `xml:"refECF,omitempty"`
}

type RefNF struct {
	CUF   int    `xml:"cUF"`
	AAMM  string `xml:"AAMM"`
	CNPJ  string `xml:"CNPJ"`
	Mod   string `xml:"mod"`
	Serie int    `xml:"serie"`
	NNF   int    `xml:"nNF"`
}

type RefNFP struct {
	CUF   int    `xml:"cUF"`
	AAMM  string `xml:"AAMM"`
	CNPJ  string `xml:"CNPJ,omitempty"`
	CPF   string `xml:"CPF,omitempty"`
	IE    string `xml:"IE"`
	Mod   string `xml:"mod"`
	Serie int    `xml:"serie"`
	NNF   int    `xml:"nNF"`
}

type RefECF struct {
	Mod  string `xml:"mod"`
	NECF string `xml:"nECF"`
	NCOO string `xml:"nCOO"`
}
