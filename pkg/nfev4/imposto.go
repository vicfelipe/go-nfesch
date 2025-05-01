package nfev4

// IPI - Imposto sobre Produtos Industrializados
type IPI struct {
	CNPJProd string   `xml:"CNPJProd,omitempty"`
	CSelo    string   `xml:"cSelo,omitempty"`
	QSelo    int      `xml:"qSelo,omitempty"`
	CEnq     string   `xml:"cEnq"`
	IPITrib  *IPITrib `xml:"IPITrib,omitempty"`
	IPINT    *IPINT   `xml:"IPINT,omitempty"`
}

type IPITrib struct {
	CST   int     `xml:"CST"`
	VBC   float64 `xml:"vBC,omitempty"`
	PIPI  float64 `xml:"pIPI,omitempty"`
	QUnid float64 `xml:"qUnid,omitempty"`
	VUnid float64 `xml:"vUnid,omitempty"`
	VIPI  float64 `xml:"vIPI"`
}

type IPINT struct {
	CST int `xml:"CST"`
}

// II - Imposto de Importação
type II struct {
	VBC   float64 `xml:"vBC"`
	VDesp float64 `xml:"vDesp"`
	VIOF  float64 `xml:"vIOF"`
	VII   float64 `xml:"vII"`
}

// PIS - Programa de Integração Social
type PIS struct {
	PISAliq *PISAliq `xml:"PISAliq,omitempty"`
	PISQtde *PISQtde `xml:"PISQtde,omitempty"`
	PISNT   *PISNT   `xml:"PISNT,omitempty"`
	PISOutr *PISOutr `xml:"PISOutr,omitempty"`
}

type PISAliq struct {
	CST  int     `xml:"CST"`
	VBC  float64 `xml:"vBC"`
	PPIS float64 `xml:"pPIS"`
	VPIS float64 `xml:"vPIS"`
}

type PISQtde struct {
	CST       int     `xml:"CST"`
	QBCProd   float64 `xml:"qBCProd"`
	VAliqProd float64 `xml:"vAliqProd"`
	VPIS      float64 `xml:"vPIS"`
}

type PISNT struct {
	CST int `xml:"CST"`
}

type PISOutr struct {
	CST       int     `xml:"CST"`
	VBC       float64 `xml:"vBC,omitempty"`
	PPIS      float64 `xml:"pPIS,omitempty"`
	QBCProd   float64 `xml:"qBCProd,omitempty"`
	VAliqProd float64 `xml:"vAliqProd,omitempty"`
	VPIS      float64 `xml:"vPIS"`
}

// COFINS - Contribuição para o Financiamento da Seguridade Social
type COFINS struct {
	COFINSAliq *COFINSAliq `xml:"COFINSAliq,omitempty"`
	COFINSQtde *COFINSQtde `xml:"COFINSQtde,omitempty"`
	COFINSNT   *COFINSNT   `xml:"COFINSNT,omitempty"`
	COFINSOutr *COFINSOutr `xml:"COFINSOutr,omitempty"`
}

type COFINSAliq struct {
	CST     int     `xml:"CST"`
	VBC     float64 `xml:"vBC"`
	PCOFINS float64 `xml:"pCOFINS"`
	VCOFINS float64 `xml:"vCOFINS"`
}

type COFINSQtde struct {
	CST       int     `xml:"CST"`
	QBCProd   float64 `xml:"qBCProd"`
	VAliqProd float64 `xml:"vAliqProd"`
	VCOFINS   float64 `xml:"vCOFINS"`
}

type COFINSNT struct {
	CST int `xml:"CST"`
}

type COFINSOutr struct {
	CST       int     `xml:"CST"`
	VBC       float64 `xml:"vBC,omitempty"`
	PCOFINS   float64 `xml:"pCOFINS,omitempty"`
	QBCProd   float64 `xml:"qBCProd,omitempty"`
	VAliqProd float64 `xml:"vAliqProd,omitempty"`
	VCOFINS   float64 `xml:"vCOFINS"`
}

// ISSQN - Imposto sobre Serviços de Qualquer Natureza
type ISSQN struct {
	VBC         float64 `xml:"vBC"`
	VAliq       float64 `xml:"vAliq"`
	VISSQN      float64 `xml:"vISSQN"`
	CMunFG      string  `xml:"cMunFG"`
	CListServ   string  `xml:"cListServ"`
	VDeducao    float64 `xml:"vDeducao,omitempty"`
	VOutro      float64 `xml:"vOutro,omitempty"`
	VDescIncond float64 `xml:"vDescIncond,omitempty"`
	VDescCond   float64 `xml:"vDescCond,omitempty"`
	VISSRet     float64 `xml:"vISSRet,omitempty"`
	IndISS      int     `xml:"indISS"`
	ISSRetido   int     `xml:"ISSRetido"`
}

// Tipos de ICMS adicionais
type ICMS10 struct {
	Orig      int     `xml:"orig"`
	CST       int     `xml:"CST"`
	ModBC     int     `xml:"modBC"`
	VBC       float64 `xml:"vBC"`
	PICMS     float64 `xml:"pICMS"`
	VICMS     float64 `xml:"vICMS"`
	ModBCST   int     `xml:"modBCST"`
	PMVAST    float64 `xml:"pMVAST,omitempty"`
	PRedBCST  float64 `xml:"pRedBCST,omitempty"`
	VBCST     float64 `xml:"vBCST"`
	PICMSST   float64 `xml:"pICMSST"`
	VICMSST   float64 `xml:"vICMSST"`
	VFCP      float64 `xml:"vFCP,omitempty"`
	VBCFCP    float64 `xml:"vBCFCP,omitempty"`
	PFCP      float64 `xml:"pFCP,omitempty"`
	VFCPST    float64 `xml:"vFCPST,omitempty"`
	VFCPSTRet float64 `xml:"vFCPSTRet,omitempty"`
}

type ICMS20 struct {
	Orig       int     `xml:"orig"`
	CST        int     `xml:"CST"`
	ModBC      int     `xml:"modBC"`
	PRedBC     float64 `xml:"pRedBC"`
	VBC        float64 `xml:"vBC"`
	PICMS      float64 `xml:"pICMS"`
	VICMS      float64 `xml:"vICMS"`
	VICMSDeson float64 `xml:"vICMSDeson"`
	MotDesICMS int     `xml:"motDesICMS,omitempty"`
	VFCP       float64 `xml:"vFCP,omitempty"`
	VBCFCP     float64 `xml:"vBCFCP,omitempty"`
	PFCP       float64 `xml:"pFCP,omitempty"`
}

// ICMS40 - ICMS isenção
type ICMS40 struct {
	Orig       int     `xml:"orig"`
	CST        int     `xml:"CST"`
	VICMS      float64 `xml:"vICMS,omitempty"`
	MotDesICMS int     `xml:"motDesICMS,omitempty"`
}

// ICMS51 - ICMS diferido
type ICMS51 struct {
	Orig     int     `xml:"orig"`
	CST      int     `xml:"CST"`
	ModBC    int     `xml:"modBC,omitempty"`
	PRedBC   float64 `xml:"pRedBC,omitempty"`
	VBC      float64 `xml:"vBC,omitempty"`
	PICMS    float64 `xml:"pICMS,omitempty"`
	VICMSOp  float64 `xml:"vICMSOp,omitempty"`
	PDif     float64 `xml:"pDif,omitempty"`
	VICMSDif float64 `xml:"vICMSDif,omitempty"`
	VICMS    float64 `xml:"vICMS,omitempty"`
}

// ICMS60 - ICMS cobrado anteriormente por substituição tributária
type ICMS60 struct {
	Orig         int     `xml:"orig"`
	CST          int     `xml:"CST"`
	VBCSTRet     float64 `xml:"vBCSTRet"`
	VICMSSTRet   float64 `xml:"vICMSSTRet"`
	VICMSSTDeson float64 `xml:"vICMSSTDeson,omitempty"`
	MotDesICMS   int     `xml:"motDesICMS,omitempty"`
}

// ICMS70 - ICMS com redução da BC e cobrança do ICMS por ST
type ICMS70 struct {
	Orig       int     `xml:"orig"`
	CST        int     `xml:"CST"`
	ModBC      int     `xml:"modBC"`
	PRedBC     float64 `xml:"pRedBC"`
	VBC        float64 `xml:"vBC"`
	PICMS      float64 `xml:"pICMS"`
	VICMS      float64 `xml:"vICMS"`
	ModBCST    int     `xml:"modBCST"`
	PMVAST     float64 `xml:"pMVAST,omitempty"`
	PRedBCST   float64 `xml:"pRedBCST,omitempty"`
	VBCST      float64 `xml:"vBCST"`
	PICMSST    float64 `xml:"pICMSST"`
	VICMSST    float64 `xml:"vICMSST"`
	VICMSDeson float64 `xml:"vICMSDeson,omitempty"`
	MotDesICMS int     `xml:"motDesICMS,omitempty"`
}

// ICMS90 - ICMS outros
type ICMS90 struct {
	Orig       int     `xml:"orig"`
	CST        int     `xml:"CST"`
	ModBC      int     `xml:"modBC,omitempty"`
	VBC        float64 `xml:"vBC,omitempty"`
	PRedBC     float64 `xml:"pRedBC,omitempty"`
	PICMS      float64 `xml:"pICMS,omitempty"`
	VICMS      float64 `xml:"vICMS,omitempty"`
	ModBCST    int     `xml:"modBCST,omitempty"`
	PMVAST     float64 `xml:"pMVAST,omitempty"`
	PRedBCST   float64 `xml:"pRedBCST,omitempty"`
	VBCST      float64 `xml:"vBCST,omitempty"`
	PICMSST    float64 `xml:"pICMSST,omitempty"`
	VICMSST    float64 `xml:"vICMSST,omitempty"`
	VICMSDeson float64 `xml:"vICMSDeson,omitempty"`
	MotDesICMS int     `xml:"motDesICMS,omitempty"`
}

// PISST - PIS por ST
type PISST struct {
	VBC       float64 `xml:"vBC,omitempty"`
	PPIS      float64 `xml:"pPIS,omitempty"`
	QBCProd   float64 `xml:"qBCProd,omitempty"`
	VAliqProd float64 `xml:"vAliqProd,omitempty"`
	VPIS      float64 `xml:"vPIS"`
}

// COFINSST - COFINS por ST
type COFINSST struct {
	VBC       float64 `xml:"vBC,omitempty"`
	PCOFINS   float64 `xml:"pCOFINS,omitempty"`
	QBCProd   float64 `xml:"qBCProd,omitempty"`
	VAliqProd float64 `xml:"vAliqProd,omitempty"`
	VCOFINS   float64 `xml:"vCOFINS"`
}
