[![CI][ci-shield]][ci-url]
[![Release][rel-shield]][rel-url]
[![GitHub tag][tag-shield]][tag-url]

# `tpsp`

`tpsp` (acrônimo para "Transporte Público de São Paulo") é uma pequena
aplicação de linha de comando que exibe o estado atual das linhas do
[Metro][metro], [CPTM][cptm], [ViaMobilidade][viamobilidade] e
[ViaQuatro][viaquatro].

> **AVISO**: Este projeto não possui relações com o Estado de São Paulo,
> a CPTM, o Metro ou qualquer outro órgão governamental.

[ci-shield]: https://img.shields.io/github/actions/workflow/status/caian-org/tpsp/ci.yml?label=ci&logo=github&style=flat-square
[ci-url]: https://github.com/caian-org/tpsp/actions/workflows/ci.yml
[rel-shield]: https://img.shields.io/github/actions/workflow/status/caian-org/tpsp/release.yml?label=release&logo=github&style=flat-square
[rel-url]: https://github.com/caian-org/tpsp/actions/workflows/release.yml
[tag-shield]: https://img.shields.io/github/tag/caian-org/tpsp.svg?logo=git&logoColor=FFF&style=flat-square
[tag-url]: https://github.com/caian-org/tpsp/releases

[metro]: http://www.metro.sp.gov.br
[cptm]: https://www.cptm.sp.gov.br
[viamobilidade]: https://www.viamobilidade.com.br
[viaquatro]: https://www.viaquatro.com.br


## Requerimentos

- Go 1.26 ou superior (apenas para compilação)
- [just][just] (apenas para compilação)
- Docker (opcional, para uso via imagem de container)

[just]: https://github.com/casey/just


## Compilação

```sh
just build
```

O binário será gerado em `bin/tpsp`.


## Uso

```
tpsp [service] [flags]

Services:
    metro          Exibe apenas linhas do Metro
    cptm           Exibe apenas linhas da CPTM
    viamobilidade  Exibe apenas linhas da ViaMobilidade
    viaquatro      Exibe apenas linhas da ViaQuatro

    Se nenhum serviço for especificado, todas as linhas são exibidas.

Flags:
    -j, --json     Exibe a saída em formato JSON
    -v, --version  Exibe a versão do programa
    -h, --help     Exibe a ajuda
    --copyright    Exibe informações de copyright
```

### Exemplos

```sh
# Exibe o estado de todas as linhas
tpsp

# Exibe apenas as linhas do Metro
tpsp metro

# Exibe as linhas da CPTM em formato JSON
tpsp cptm --json
```

Saída de exemplo:

```
Linha              Status
---------------------------------
Azul               Operação Normal
Verde              Operação Normal
Vermelha           Operação Normal
Amarela            Operação Normal
Lilás              Operação Normal
Prata              Operação Normal
```


### Saída em JSON

```json
{
    "code": 200,
    "data": [
        {
            "line": "Azul",
            "status": "Operação Normal"
        },
        {
            "line": "Verde",
            "status": "Operação Normal"
        }
    ],
    "message": "success"
}
```


## Docker

A imagem oficial do projeto é publicada no GitHub Container Registry:

```sh
docker pull ghcr.io/caian-org/tpsp:latest
```

Também é possível usar uma versão específica:

```sh
docker pull ghcr.io/caian-org/tpsp:v2.0.4
```

Exemplos:

```sh
# Exibe o estado de todas as linhas
docker run --rm ghcr.io/caian-org/tpsp:latest

# Exibe apenas as linhas do Metro
docker run --rm ghcr.io/caian-org/tpsp:latest metro

# Exibe as linhas da CPTM em formato JSON
docker run --rm ghcr.io/caian-org/tpsp:latest cptm --json
```


## Licença

Na medida do permitido por lei, [Caian Ertl][me] renunciou a __todos os
direitos autorais e direitos conexos a este trabalho__. No espírito de
_liberdade de informação_, você é encorajado a clonar, modificar, distribuir,
compartilhar ou fazer o que quiser com este projeto! [`^C ^V`][kopimi]

[![Licença][cc-shield]][cc-url]

[me]: https://github.com/upsetbit
[cc-shield]: https://forthebadge.com/images/badges/cc-0.svg
[cc-url]: http://creativecommons.org/publicdomain/zero/1.0
[kopimi]: https://kopimi.com
