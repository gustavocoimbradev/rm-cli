# RM CLI

Ferramenta CLI desenvolvida em Go para consultar dados diretamente no banco de dados SQL Server do TOTVS RM (CorporeRM).

O projeto busca facilitar o acesso a informações do sistema por meio do terminal, permitindo a criação de comandos para consultas, listagens e geração de relatórios.

## Funcionalidades

* Conecta diretamente ao banco de dados do TOTVS RM
* Executa consultas pelo terminal
* Permite buscar e listar informações do CorporeRM
* Exibe os resultados de forma organizada
* Suporta a criação de novos comandos e relatórios
* Utiliza consultas SQL parametrizadas

## Configuração

Crie um arquivo `.env` na raiz do projeto:

```env
DB_HOST=127.0.0.1
DB_PORT=1433
DB_DATABASE=CorporeRM
DB_USERNAME=usuario
DB_PASSWORD=senha
````

## Como usar

Execute diretamente com Go:

```bash
go run .
```

Ou gere o executável:

```bash
go build -o rm-cli
```

Depois execute:

```bash
./rm-cli
```

A ferramenta exibirá no terminal os comandos disponíveis.

## Tecnologias

* Go
* SQL Server
* TOTVS RM (CorporeRM)
* database/sql
* go-mssqldb
* godotenv

## Autor

Desenvolvido por [Gustavo Coimbra](https://github.com/gustavocoimbradev).
