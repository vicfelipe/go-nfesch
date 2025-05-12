# GO-NFESCH

GO-NFESCH é um projeto em Go que implementa o Schema 4.0 da Nota Fiscal Eletrônica (NFe), conforme especificado pela Receita Federal do Brasil. Este projeto fornece ferramentas para manipulação, validação e geração de informações relacionadas à NFe, incluindo suporte para QR Codes e cálculos de valores fiscais.

## Funcionalidades Implementadas

- **Schema Layout NFe 4.0**: Estruturas completas para representar o layout da NFe 4.0.
- **Parse NFe from XML**: Leitura e interpretação de arquivos XML de NFe.
- **GenerateQRCodeImage**: Geração de imagens de QR Code para a NFe.
- **GenerateQRCode**: Geração do conteúdo do QR Code conforme o padrão SEFAZ.
- **calculateDigestValue**: Cálculo do valor de digest (hash) para assinatura digital.
- **ValidateNFe**: Validação básica da estrutura da NFe.

## Estrutura do Projeto

O projeto está organizado da seguinte forma:

- **`pkg/nfev4`**: Contém as estruturas e funcionalidades relacionadas ao layout da NFe 4.0.
- **`pkg/util`**: Funções utilitárias, como parsing de XML e geração de QR Codes.
- **`pkg/validator`**: Validações básicas da estrutura da NFe.
- **`configs/sefaz.yaml`**: Configurações específicas para diferentes estados brasileiros.
- **`data/nf_teste.xml`**: Exemplo de arquivo XML de uma NFe.

## Como Rodar

Certifique-se de ter o Go instalado em sua máquina. Para executar o projeto, utilize o comando:

```bash
go run main.go
```

O programa irá abrir o arquivo `data/nf_teste.xml`, realizar o parsing da NFe e exibir informações como emitente, número da nota e valor total.

## Dependências

O projeto utiliza a biblioteca [go-qrcode](https://github.com/skip2/go-qrcode) para geração de QR Codes. As dependências estão especificadas no arquivo `go.mod`.

## Configurações

O arquivo `configs/sefaz.yaml` contém URLs e configurações específicas para os webservices e QR Codes de cada estado brasileiro.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).

## Referência

O Schema 4.0 da Nota Fiscal Eletrônica pode ser encontrado no portal oficial da Receita Federal: [NFe - Receita Federal](https://www.nfe.fazenda.gov.br/portal/listaConteudo.aspx?tipoConteudo=BMPFMBoln3w=&AspxAutoDetectCookieSupport=1).

## Autor

Este projeto foi desenvolvido por [Victor Alencastro](https://www.victorcode.dev).