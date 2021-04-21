# Rota de Viagem #

Um turista deseja viajar pelo mundo pagando o menor preço possível independentemente do número de conexões necessárias.
Vamos construir um programa que facilite ao nosso turista, escolher a melhor rota para sua viagem.

Para isso precisamos inserir as rotas através de um arquivo de entrada.

## Input Example ##
```csv
GRU,BRC,10
BRC,SCL,5
GRU,CDG,75
GRU,SCL,20
GRU,ORL,56
ORL,CDG,5
SCL,ORL,20
```

## Explicando ## 
Caso desejemos viajar de **GRU** para **CDG** existem as seguintes rotas:

1. GRU - BRC - SCL - ORL - CDG ao custo de **$40**
2. GRU - ORL - CDG ao custo de **$64**
3. GRU - CDG ao custo de **$75**
4. GRU - SCL - ORL - CDG ao custo de **$45**

O melhor preço é da rota **1** logo, o output da consulta deve ser **GRU - BRC - SCL - ORL - CDG**.

### Execução do programa ###
A inicializacao do teste se dará por linha de comando onde o primeiro argumento é o arquivo com a lista de rotas inicial.

```shell
$ mysolution input-routes.csv
```

Duas interfaces de consulta devem ser implementadas:
- Interface de console deverá receber um input com a rota no formato "DE-PARA" e imprimir a melhor rota e seu respectivo valor.
  Exemplo:
  
```shell
please enter the route: GRU-CDG
best route: GRU - BRC - SCL - ORL - CDG > $40
please enter the route: BRC-ORL
best route: BRC - SCL - ORL > $25
```

- Interface Rest
    A interface Rest deverá suportar:
    - Registro de novas rotas. Essas novas rotas devem ser persistidas no arquivo csv utilizado como input(input-routes.csv),
    - Consulta de melhor rota entre dois pontos.

Também será necessária a implementação de 2 endpoints Rest, um para registro de rotas e outro para consula de melhor rota.

## Recomendações ##
Para uma melhor fluides da nossa conversa, atente-se aos seguintes pontos:

* Envie apenas o código fonte,
* Estruture sua aplicação seguindo as boas práticas de desenvolvimento,
* Evite o uso de frameworks ou bibliotecas externas à linguagem. Utilize apenas o que for necessário para a exposição do serviço,
* Implemente testes unitários seguindo as boas praticas de mercado,
* Documentação
  Em um arquivo Texto ou Markdown descreva:
  * Como executar a aplicação,
  * Estrutura dos arquivos/pacotes,
  * Explique as decisões de design adotadas para a solução,
  * Descreva sua APÌ Rest de forma simplificada.

## Procedimento para utilizar a interface de linha de comando
modos de execução
### Para executar no modo padrão
```
go run cmd/mysolution-cli.go ./input-routes.csv
```

### O programa também funcionário em modo silencioso da seguinte maneira
```
go run cmd/mysolution-cli.go -s --origin GRU --destination CDG ./input-routes.csv
```

### Para acessar o help do utilitário
```
go run cmd/mysolution-cli.go --help
```


## Procedimentos para cosumir a interface REST da aplicação
Esse API segue a espeficicação definida em arquivo ***travels-api_openapi3.yaml***

Primeiramente deverá ser iniciado o serviço REST via comando:

```
go run main.go ./input-routes.csv
```

Basta escolher os seu HTTP Client preferido para interagir com a API...
### Consulta de rota para viagem mais econômica
exemplo: Origem aeroporto de Grarulhos [Brasil] (GRU) e Destino aeroporto Charles de Gaulle [França] (CDG)

```shell
curl --location --request GET 'http://localhost:8080/travels?origin=GRU&destination=CDG'
```

### Adicionar uma nova rota de Viagem
exemplo: Origem aeroporto de Grarulhos [Brasil] (GRU), Destino aeroporto de Los Angeles [EUA] (LAX) e custo de 99.77 (Dolars)

```shell
curl --location --request POST 'http://localhost:8080/travels' \
--header 'Content-Type: application/json' \
--data-raw '{
  "origin": "GRU",
  "destination": "LAX",
  "cost": 99.77
}'
```


