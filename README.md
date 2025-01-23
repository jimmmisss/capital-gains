# Captial Gains 

Este projeto é uma aplicação Go que implementa uma lógica de cálculo de impostos para ganho de capital. Ele define uma struct `CalculadoraImpostos` que gerencia cálculos de impostos com base em diferentes operações, como compra e venda.

## Estrutura do Projeto

```
capital-gains
├── cmd
│   └── main.go
├── internal
│   └── domain
│       └── models.go
│       └── tax.go
│       └── tax_test.go
│   └── infra
│       └── input.go
│       └── input_test.go
│   └── service
│       └── processor.go
│       └── processor_test.go
├── input.txt         
├── Dockerfile          
├── docker-compose.yaml  
└── README.md           
```

## Primeiros Passos

Para construir e executar a aplicação usando Docker, siga estes passos:

1. **Clone o repositório:**

   ```sh
   git clone <repository-url>
   cd capital-gains
   ```

2. **Construa a imagem Docker:**

   ```sh
   docker-compose build
   ```

3. **Execute a aplicação:**

   ```sh
   docker-compose run --rm -T capital-gains < input.txt
   ```

## Uso

Uma vez que a aplicação seja executada com o comando acima, você pode verá no terminal o resultado para os cálculos baseado nas regras de negócios

## Testes

Para executar os testes deste projeto, use o seguinte comando:

```sh
go test ./...
go test ./... -v
```

## Justificativa da Arquitetura

Esta aplicação foi desenvolvida utilizando uma arquitetura modular, organizada em camadas internas (`internal`) para promover encapsulamento e separação de responsabilidades. A estrutura do projeto segue boas práticas de design para projetos em Go, garantindo manutenibilidade, extensibilidade e testabilidade.  

### Escolhas Arquiteturais

1. **Separação de Domínios**:
   - O código fonte foi dividido em módulos específicos como `domain`, `infra` e `service`. Cada módulo é responsável por uma parte distinta da lógica da aplicação:
     - **Domain**: Contém os modelos principais (`models.go`) e a lógica de negócio (`tax.go`) que centralizam as regras fiscais e cálculos de impostos.
     - **Infra**: Gerencia a entrada de dados (`input.go`), oferecendo uma camada que abstrai interações externas.
     - **Service**: Processa a lógica de alto nível, conectando infraestruturas e lógica de domínio para executar a operação principal (`processor.go`).

2. **Teste e Validação**:
   - Todos os módulos possuem seus respectivos testes unitários, como `tax_test.go` e `processor_test.go`, garantindo que cada componente seja validado individualmente.

3. **Execução via CLI**:
   - A aplicação é projetada para rodar no terminal com suporte à leitura de arquivos de entrada (`input.txt`). Essa abordagem simplifica a interface do usuário para o objetivo inicial e facilita automações ou integrações futuras.

4. **Dockerização**:
   - O projeto está encapsulado em um ambiente Docker, permitindo sua execução consistente independentemente do sistema operacional ou das dependências locais do desenvolvedor. O uso do `docker-compose.yaml` simplifica a construção e execução da aplicação.

5. **Configuração de Entrada Personalizável**:
   - O uso de volumes no Docker (`input.txt`) permite fácil customização dos dados de entrada sem necessidade de alterar o código-fonte.

6. **Simples e Funcional**:
   - A escolha de uma estrutura relativamente simples reflete o foco em resolver o problema central (cálculo de impostos para ganhos de capital) com o menor número de dependências externas e complexidade.

### Benefícios da Arquitetura

- **Escalabilidade**: A organização em módulos permite adicionar novas funcionalidades ou adaptar a lógica de negócio sem impacto significativo no restante do sistema.
- **Portabilidade**: O uso de Docker assegura que a aplicação pode ser executada de forma idêntica em qualquer ambiente.
- **Manutenibilidade**: A separação de camadas e os testes fornecem uma base sólida para evolução do projeto com menor risco de introdução de falhas.
- **Clareza e Organização**: A estrutura modular e documentação clara garantem que novos desenvolvedores consigam se ambientar rapidamente ao projeto.
