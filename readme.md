# Decomplicando o Go

Anotações feitas durante o curso 

## Day-1

Todo código em go começa no arquivo main.go

A func main() é rodada automaticamente como ponto de entradas

fmt é um pacote da standard library do go.

para rodar o pacote no terminar usamos ```go run main.go ```

para saber mais comando do go no terminal digitar ``go --help``

go usa camelCase. Quando começo o nome da função com letra minuscula, significa que ela é privada dentro do pacote ex ``func inicioFuncao(){}``, se começo com letra maiuscula ela é publica dentro do código``func InicioFuncao(){}``

uma string pode ser vazia, mas não nula. Toda string é imutavel. Quando 'mudamos' uma string é como se criassemos uma nova

o comando ``gofmt -w main.go`` formata seus arquivos go seguindo as melhores praticas (https://go.dev/blog/gofmt)