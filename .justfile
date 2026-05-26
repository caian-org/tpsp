set shell := ["/bin/bash", "-c"]

# --------------------------------------------------------------------------------------------------

_help:
    @just --list

# --------------------------------------------------------------------------------------------------

# build tpsp into bin/tpsp
build:
    @mkdir -p bin
    @go build -o bin/tpsp ./cmd/tpsp
    @echo "built tpsp"

# remove generated binary
clean:
    @rm -f bin/tpsp
