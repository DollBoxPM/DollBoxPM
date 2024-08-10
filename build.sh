#!/bin/bash

# Compilar o código
go build -o dollbox src/main.go

# Verificar se a compilação foi bem-sucedida
if [ $? -eq 0 ]; then
  echo "Compilation sucess!"
else
  echo "Compilation error!"
  exit 1
fi
