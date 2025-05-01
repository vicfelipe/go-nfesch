package nfev4

type ISSQNTot struct {
	VServ       float64 `xml:"vServ"`
	VBC         float64 `xml:"vBC"`
	VISS        float64 `xml:"vISS"`
	VPIS        float64 `xml:"vPIS"`
	VCOFINS     float64 `xml:"vCOFINS"`
	DCompet     string  `xml:"dCompet"`
	VDeducao    float64 `xml:"vDeducao,omitempty"`
	VOutro      float64 `xml:"vOutro,omitempty"`
	VDescIncond float64 `xml:"vDescIncond,omitempty"`
	VDescCond   float64 `xml:"vDescCond,omitempty"`
	VISSRet     float64 `xml:"vISSRet,omitempty"`
	CRegTrib    int     `xml:"cRegTrib,omitempty"`
}

type RetTrib struct {
	VRetPIS    float64 `xml:"vRetPIS,omitempty"`
	VRetCOFINS float64 `xml:"vRetCOFINS,omitempty"`
	VRetCSLL   float64 `xml:"vRetCSLL,omitempty"`
	VBCIRRF    float64 `xml:"vBCIRRF,omitempty"`
	VIRRF      float64 `xml:"vIRRF,omitempty"`
	VBCRetPrev float64 `xml:"vBCRetPrev,omitempty"`
	VRetPrev   float64 `xml:"vRetPrev,omitempty"`
}
