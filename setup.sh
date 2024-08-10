#!/bin/bash

# Instalar dependências
go get -u github.com/go-git/go-git/v5
go get -u github.com/go-git/go-git/v5/plumbing

# Configurar variáveis de ambiente
export GOPATH=$PWD
export PATH=$GOPATH/bin:$PATH

# Instalar ferramentas necessárias
go install github.com/go-git/go-git/v5/cmd/git
