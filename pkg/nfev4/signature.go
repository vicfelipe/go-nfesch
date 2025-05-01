package nfev4

type SignedInfo struct {
	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod"`
	SignatureMethod        SignatureMethod        `xml:"SignatureMethod"`
	Reference              Reference              `xml:"Reference"`
}

type CanonicalizationMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type SignatureMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type Reference struct {
	URI          string       `xml:"URI,attr"`
	Transforms   Transforms   `xml:"Transforms"`
	DigestMethod DigestMethod `xml:"DigestMethod"`
	DigestValue  string       `xml:"DigestValue"`
}

type Transforms struct {
	Transform []Transform `xml:"Transform"`
}

type Transform struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type DigestMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

type KeyInfo struct {
	X509Data X509Data `xml:"X509Data"`
}

type X509Data struct {
	X509Certificate string `xml:"X509Certificate"`
}
