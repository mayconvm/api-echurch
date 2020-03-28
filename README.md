# E-Church API #

API para o aplicativo da E-Church

## Iniciando a aplicação

* Configurações manuais
    * Adicione o projeto na pasta do GOPATH
        go get -u bitbucket.org/eChurchIPB/api-echurch
        ln -s $GOPATH/src/bitbucket.org/eChurchIPB/api-echurch api

* Gerenciador de dependências
    * Pelo GO
        go get -u github.com/golang/dep/cmd/dep
        no arquivo .bashrc adicionar a "export PATH=$PATH:~/.go/bin/"
    
    * Ou instalar por script
        curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

    * Instalando dependências
        * dep ensure

## Todo

Arquitetura
    [x] Definir qual será o framework
    [] Definir qual o banco de dados

Aplicação
    [x] Configurar o webservice
    [] Conectar com o banco de dados
    [] Adicionar CORS
    [] Adicionar Authenticação por JWT

Infra-estrutura
    [] Configurar o docker para realizar o build da aplicação
    [] Configurar o docker para executar a aplicação
    [] Configurar o docker para o banco de dados