# TPSP: Transporte Público de São Paulo

`tpsp` (acrônimo para "Transporte Público de São Paulo") é uma pequena aplicação
de linha de comando que exibe o estado atual das linhas do [Metro], [CPTM],
[ViaMobilidade] e [ViaQuatro].

**AVISO: Este projeto não possui relações com o Estado de São Paulo, a CPTM, o
Metro ou qualquer outro órgão governamental.**

[Metro]: http://www.metro.sp.gov.br
[CPTM]: https://www.cptm.sp.gov.br
[ViaMobilidade]: https://www.viamobilidade.com.br
[ViaQuatro]: https://www.viaquatro.com.br


## Requerimentos

- Go 1.26 ou superior (apenas para compilação)
- just (apenas para compilação)
- Docker (opcional, para uso via imagem de container)


## Compilação

```sh
just build
```

O binário será gerado em `bin/tpsp`.


## Docker

A imagem oficial do projeto é publicada no GitHub Container Registry:

```sh
docker pull ghcr.io/caian-org/tpsp:latest
```

Também é possível usar uma versão específica:

```sh
docker pull ghcr.io/caian-org/tpsp:v2.0.3
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


## Licença

Na medida do possível sob a lei, [Caian Ertl][me] renunciou a __todos os
direitos autorais e direitos relacionados ou adjacentes a este trabalho__. No
espírito da _liberdade de informação_, encorajo você a forkar, modificar,
alterar, compartilhar ou fazer o que quiser com este projeto! [`^ C ^ V`][kopimi]

[me]: https://github.com/caian-org
[kopimi]: https://kopimi.com
