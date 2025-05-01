package nfev4

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Estrutura principal da NFe
type NFe struct {
	XMLName    xml.Name    `xml:"NFe"`
	InfNFe     InfNFe      `xml:"infNFe"`
	InfNFeSupl *InfNFeSupl `xml:"infNFeSupl,omitempty"` // Novo campo para QRCode
	Versao     string      `xml:"versao,attr"`
}

type InfNFe struct {
	Id        string     `xml:"Id,attr"`
	Ide       Ide        `xml:"ide"`
	Emit      Emit       `xml:"emit"`
	Dest      *Dest      `xml:"dest,omitempty"`
	Det       []Det      `xml:"det"`
	Total     Total      `xml:"total"`
	Transp    Transp     `xml:"transp"`
	Cobr      *Cobr      `xml:"cobr,omitempty"`
	Pag       Pag        `xml:"pag"`
	InfAdic   *InfAdic   `xml:"infAdic,omitempty"`
	Signature *Signature `xml:"Signature,omitempty"`
}

// Identificação da NFe
type Ide struct {
	CUF      int        `xml:"cUF"`
	CNF      string     `xml:"cNF"`
	NatOp    string     `xml:"natOp"`
	Mod      int        `xml:"mod"`
	Serie    int        `xml:"serie"`
	NNF      int        `xml:"nNF"`
	DhEmi    time.Time  `xml:"dhEmi"`
	DhSaiEnt *time.Time `xml:"dhSaiEnt,omitempty"`
	TpNF     int        `xml:"tpNF"`
	IdDest   int        `xml:"idDest"`
	CMunFG   string     `xml:"cMunFG"`
	TpImp    int        `xml:"tpImp"`
	TpEmis   int        `xml:"tpEmis"`
	CDV      int        `xml:"cDV"`
	TpAmb    int        `xml:"tpAmb"`
	FinNFe   int        `xml:"finNFe"`
	IndFinal int        `xml:"indFinal"`
	IndPres  int        `xml:"indPres"`
	ProcEmi  int        `xml:"procEmi"`
	VerProc  string     `xml:"verProc"`
	NFRef    []NFRef    `xml:"NFref,omitempty"`
}

// Emitente da NFe
type Emit struct {
	CNPJ      string    `xml:"CNPJ,omitempty"`
	CPF       string    `xml:"CPF,omitempty"`
	XNome     string    `xml:"xNome"`
	XFant     string    `xml:"xFant,omitempty"`
	EnderEmit EnderEmit `xml:"enderEmit"`
	IE        string    `xml:"IE"`
	IEST      string    `xml:"IEST,omitempty"`
	IM        string    `xml:"IM,omitempty"`
	CNAE      string    `xml:"CNAE,omitempty"`
	CRT       int       `xml:"CRT"`
}

// Destinatário da NFe
type Dest struct {
	CNPJ          string    `xml:"CNPJ,omitempty"`
	CPF           string    `xml:"CPF,omitempty"`
	IdEstrangeiro string    `xml:"idEstrangeiro,omitempty"`
	XNome         string    `xml:"xNome"`
	EnderDest     EnderDest `xml:"enderDest"`
	IndIEDest     int       `xml:"indIEDest"`
	IE            string    `xml:"IE,omitempty"`
	ISUF          string    `xml:"ISUF,omitempty"`
	IM            string    `xml:"IM,omitempty"`
	Email         string    `xml:"email,omitempty"`
}

// Endereço do Emitente
type EnderEmit struct {
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XCpl    string `xml:"xCpl,omitempty"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	XMun    string `xml:"xMun"`
	UF      string `xml:"UF"`
	CEP     string `xml:"CEP"`
	CPais   string `xml:"cPais,omitempty"`
	XPais   string `xml:"xPais,omitempty"`
	Fone    string `xml:"fone,omitempty"`
}

// Endereço do Destinatário
type EnderDest struct {
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XCpl    string `xml:"xCpl,omitempty"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	XMun    string `xml:"xMun"`
	UF      string `xml:"UF"`
	CEP     string `xml:"CEP"`
	CPais   string `xml:"cPais,omitempty"`
	XPais   string `xml:"xPais,omitempty"`
	Fone    string `xml:"fone,omitempty"`
}

// Detalhamento de Produtos/Serviços
type Det struct {
	NItem     int     `xml:"nItem,attr"`
	Prod      Prod    `xml:"prod"`
	Imposto   Imposto `xml:"imposto"`
	InfAdProd string  `xml:"infAdProd,omitempty"`
}

type Prod struct {
	CProd    string  `xml:"cProd"`
	CEAN     string  `xml:"cEAN,omitempty"`
	XProd    string  `xml:"xProd"`
	NCM      string  `xml:"NCM"`
	CEST     string  `xml:"CEST,omitempty"`
	CFOP     string  `xml:"CFOP"`
	UCom     string  `xml:"uCom"`
	QCom     float64 `xml:"qCom"`
	VUnCom   float64 `xml:"vUnCom"`
	VProd    float64 `xml:"vProd"`
	CEANTrib string  `xml:"cEANTrib,omitempty"`
	UTrib    string  `xml:"uTrib"`
	QTrib    float64 `xml:"qTrib"`
	VUnTrib  float64 `xml:"vUnTrib"`
	VFrete   float64 `xml:"vFrete,omitempty"`
	VSeg     float64 `xml:"vSeg,omitempty"`
	VDesc    float64 `xml:"vDesc,omitempty"`
	VOutro   float64 `xml:"vOutro,omitempty"`
	IndTot   int     `xml:"indTot"`
}

// Tributos
type Imposto struct {
	ICMS   *ICMS   `xml:"ICMS,omitempty"`
	IPI    *IPI    `xml:"IPI,omitempty"`
	II     *II     `xml:"II,omitempty"`
	PIS    *PIS    `xml:"PIS,omitempty"`
	COFINS *COFINS `xml:"COFINS,omitempty"`
	ISSQN  *ISSQN  `xml:"ISSQN,omitempty"`
}

// ICMS
type ICMS struct {
	ICMS00 *ICMS00 `xml:"ICMS00,omitempty"`
	ICMS10 *ICMS10 `xml:"ICMS10,omitempty"`
	ICMS20 *ICMS20 `xml:"ICMS20,omitempty"`
	// Outros tipos de ICMS podem ser adicionados aqui
}

type ICMS00 struct {
	Orig  int     `xml:"orig"`
	CST   int     `xml:"CST"`
	ModBC int     `xml:"modBC"`
	VBC   float64 `xml:"vBC"`
	PICMS float64 `xml:"pICMS"`
	VICMS float64 `xml:"vICMS"`
	PFCP  float64 `xml:"pFCP,omitempty"`
	VFCP  float64 `xml:"vFCP,omitempty"`
}

// Estruturas para outros tipos de ICMS (ICMS10, ICMS20, etc.) devem ser adicionadas

// Totalizadores
type Total struct {
	ICMSTot  ICMSTot   `xml:"ICMSTot"`
	ISSQNTot *ISSQNTot `xml:"ISSQNtot,omitempty"`
	RetTrib  *RetTrib  `xml:"retTrib,omitempty"`
}

type ICMSTot struct {
	VBC        float64 `xml:"vBC"`
	VICMS      float64 `xml:"vICMS"`
	VICMSDeson float64 `xml:"vICMSDeson"`
	VFCP       float64 `xml:"vFCP"`
	VBCST      float64 `xml:"vBCST"`
	VST        float64 `xml:"vST"`
	VFCPST     float64 `xml:"vFCPST"`
	VFCPSTRet  float64 `xml:"vFCPSTRet"`
	VProd      float64 `xml:"vProd"`
	VFrete     float64 `xml:"vFrete"`
	VSeg       float64 `xml:"vSeg"`
	VDesc      float64 `xml:"vDesc"`
	VII        float64 `xml:"vII"`
	VIPI       float64 `xml:"vIPI"`
	VIPIDevol  float64 `xml:"vIPIDevol"`
	VPIS       float64 `xml:"vPIS"`
	VCOFINS    float64 `xml:"vCOFINS"`
	VOutro     float64 `xml:"vOutro"`
	VNF        float64 `xml:"vNF"`
	VTotTrib   float64 `xml:"vTotTrib,omitempty"`
}

// Transporte
type Transp struct {
	ModFrete   int         `xml:"modFrete"`
	Transporta *Transporta `xml:"transporta,omitempty"`
	RetTransp  *RetTransp  `xml:"retTransp,omitempty"`
	VeicTransp *VeicTransp `xml:"veicTransp,omitempty"`
	Reboque    []Reboque   `xml:"reboque,omitempty"`
	Vol        []Vol       `xml:"vol,omitempty"`
}

// Pagamento
type Pag struct {
	DetPag []DetPag `xml:"detPag"`
	VTroco float64  `xml:"vTroco,omitempty"`
}

type DetPag struct {
	TPag int     `xml:"tPag"`
	VPag float64 `xml:"vPag"`
	Card *Card   `xml:"card,omitempty"`
}

// Informações adicionais
type InfAdic struct {
	InfAdFisco string     `xml:"infAdFisco,omitempty"`
	InfCpl     string     `xml:"infCpl,omitempty"`
	ObsCont    []ObsCont  `xml:"obsCont,omitempty"`
	ObsFisco   []ObsFisco `xml:"obsFisco,omitempty"`
	ProcRef    []ProcRef  `xml:"procRef,omitempty"`
}

// Assinatura digital
type Signature struct {
	XMLName        xml.Name `xml:"Signature"`
	Xmlns          string   `xml:"xmlns,attr"`
	SignedInfo     SignedInfo
	SignatureValue string `xml:"SignatureValue"`
	KeyInfo        KeyInfo
}

type infNFeForDigest struct {
	XMLName xml.Name `xml:"infNFe"`
	*InfNFe
}

func (nfe *NFe) GenerateAccessKey() string {
	ide := nfe.InfNFe.Ide
	emit := nfe.InfNFe.Emit

	// Formato: UF + AAMM + CNPJ + MOD + SERIE + NUMERO + TPEMIS + CODIGO NUMERICO + DV
	key := fmt.Sprintf("%02d", ide.CUF) +
		ide.DhEmi.Format("06") + ide.DhEmi.Format("01") +
		emit.CNPJ +
		fmt.Sprintf("%02d", ide.Mod) +
		fmt.Sprintf("%03d", ide.Serie) +
		fmt.Sprintf("%09d", ide.NNF) +
		fmt.Sprintf("%01d", ide.TpEmis) +
		fmt.Sprintf("%08d", ide.CNF) +
		fmt.Sprintf("%01d", ide.CDV)

	return key
}

// GenerateQRCode gera o conteúdo do QRCode conforme padrão SEFAZ
// Gerar QRCode para ambiente de homologação
// qrCode, err := nfe.GenerateQRCode(2, "", "") // Sem CSC para emissão normal
//
//	if err != nil {
//		fmt.Printf("Erro ao gerar QRCode: %v\n", err)
//		return
//	}
func (nfe *NFe) GenerateQRCode(ambiente int, cscID string, cscToken string) (string, error) {
	// Validações básicas
	if nfe.InfNFe.Id == "" {
		return "", fmt.Errorf("ID da NFe não pode ser vazio")
	}

	if ambiente != 1 && ambiente != 2 {
		return "", fmt.Errorf("ambiente deve ser 1 (produção) ou 2 (homologação)")
	}

	// Extrair chave de acesso (remover "NFe" do início)
	chaveAcesso := strings.TrimPrefix(nfe.InfNFe.Id, "NFe")
	if len(chaveAcesso) != 44 {
		return "", fmt.Errorf("chave de acesso inválida")
	}

	// Obter dígito verificador (último caractere da chave)
	dv := chaveAcesso[len(chaveAcesso)-1:]

	// Montar parâmetros do QRCode
	params := url.Values{}
	params.Add("p", chaveAcesso)
	params.Add("chNFe", chaveAcesso)
	params.Add("nVersao", "100")
	params.Add("tpAmb", fmt.Sprintf("%d", ambiente))

	// Adicionar CSC quando disponível (contingência)
	if cscToken != "" && cscID != "" {
		digest, err := nfe.calculateDigestValue()
		if err != nil {
			return "", fmt.Errorf("falha ao calcular digest value: %w", err)
		}

		params.Add("cDest", nfe.InfNFe.Dest.CNPJ)
		params.Add("dhEmi", nfe.InfNFe.Ide.DhEmi.Format(time.RFC3339))
		params.Add("vNF", fmt.Sprintf("%.2f", nfe.InfNFe.Total.ICMSTot.VNF))
		params.Add("vICMS", fmt.Sprintf("%.2f", nfe.InfNFe.Total.ICMSTot.VICMS))
		params.Add("digVal", digest)
		params.Add("cIdToken", cscID)
	}

	// Montar URL base conforme ambiente
	var baseURL string
	if ambiente == 1 {
		baseURL = "https://www.sefazvirtual.fazenda.gov.br/NFeConsulta/qrcode"
	} else {
		baseURL = "https://hom.nfe.fazenda.gov.br/NFeConsulta/qrcode"
	}

	// Montar URL completa
	qrCodeURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Adicionar dígito verificador no final quando em contingência
	if cscToken != "" {
		qrCodeURL += "|" + string(dv)
	}

	return qrCodeURL, nil
}

// calculateDigestValue calcula o digest value conforme padrão SEFAZ
func (nfe *NFe) calculateDigestValue() (string, error) {
	// 1. Criar uma estrutura temporária apenas com infNFe
	inf := infNFeForDigest{
		InfNFe: &nfe.InfNFe,
	}

	// 2. Serializar para XML
	var buf bytes.Buffer
	encoder := xml.NewEncoder(&buf)

	// Configurações importantes para o padrão SEFAZ
	encoder.Indent("", "")

	if err := encoder.Encode(inf); err != nil {
		return "", fmt.Errorf("falha ao serializar infNFe: %w", err)
	}

	// 3. Normalizar o XML (conforme exigido pelo padrão)
	xmlContent := buf.String()
	xmlContent = strings.ReplaceAll(xmlContent, "\n", "")
	xmlContent = strings.ReplaceAll(xmlContent, "\r", "")
	xmlContent = strings.ReplaceAll(xmlContent, "\t", "")
	xmlContent = strings.Join(strings.Fields(xmlContent), " ")

	// 4. Calcular SHA-1
	hasher := sha1.New()
	hasher.Write([]byte(xmlContent))
	sha1Hash := hasher.Sum(nil)

	// 5. Codificar em base64
	digestValue := base64.StdEncoding.EncodeToString(sha1Hash)

	return digestValue, nil
}
