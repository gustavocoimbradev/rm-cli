````markdown
# RM CLI

Ferramenta CLI desenvolvida em Go para consultar dados diretamente no banco de dados SQL Server do TOTVS RM (CorporeRM).

O projeto busca facilitar o acesso a informações do sistema por meio do terminal, permitindo a criação de comandos para consultas, listagens e relatórios.

## Funcionalidades

* Conecta diretamente ao banco de dados do TOTVS RM
* Executa consultas pelo terminal
* Permite buscar e listar informações do CorporeRM
* Exibe os resultados de forma organizada
* Suporta a criação de novos comandos e relatórios
* Utiliza consultas SQL parametrizadas
* Possui saída colorida no terminal

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
go run . student {{ra}}
```

Ou gere o executável:

```bash
go build -o rm-cli
```

Depois execute:

```bash
./rm-cli student {{ra}}
```

## Exemplo

```text
DADOS DO ALUNO

RA: 000000
Nome: João da Silva
CPF: 12345678900
Nascimento: 1990-01-01
E-mail: joao@example.com
Telefones: 35999999999

ENDEREÇO

Rua: Rua Principal
Número: 123
Complemento: -
Bairro: Centro
Cidade: Varginha
Estado: MG
CEP: 37000000
País: Brasil
```

## Tecnologias

* Go
* SQL Server
* TOTVS RM (CorporeRM)
* `database/sql`
* `go-mssqldb`
* `godotenv`

## Autor

Desenvolvido por [Gustavo Coimbra](https://github.com/gustavocoimbradev).

```
```
