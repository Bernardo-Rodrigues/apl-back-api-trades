# Tech Challenge - Solução

Este repositório contém a solução desenvolvida para o desafio técnico proposto. Para mais detalhes, veja [challenge.md](challenge.md).
O objetivo foi construir um sistema para gerar relatórios financeiros a partir de dados de transações e preços de ativos.

## Tecnologias Utilizadas
- **Golang**: Linguagem principal utilizada para a implementação da lógica do sistema.
- **CSV**: Formato para exportação dos relatórios gerados.
- **gRPC**: Utilizado para a comunicação eficiente entre a API e os clientes.
- **Testes**: Utilização de `testify` e `mockery` para garantir o funcionamento correto do código.
- **Docker**: Facilita a execução e distribuição da aplicação em ambientes isolados e reproduzíveis.

### **Ferramenta de Apoio**: [Make](https://www.gnu.org/software/make/)
- O Arquivo makefile contém um conjunto de diretivas usadas pela ferramenta de automação de compilação ``make`` para apoiar no desenvolvimento:
#### Baixar o Chocolatey (se necessário no Windows):
```
https://chocolatey.org/install 
```
#### Instalar o make (se necessário no Windows):
```
choco install make
```

## Arquitetura e Design
A arquitetura adotada segue princípios de **Clean Architecture** e **DDD (Domain-Driven Design)**, aplicados aqui mais como uma demonstração de boas práticas do que por necessidade do projeto. A solução está dividida em camadas bem definidas, com regras de negócio implementadas de forma clara e independente de frameworks e detalhes de infraestrutura.

## Estrutura do Código
- `main.go`: Contém o código de inicialização e execução do servidor.
- `domain/`: Contém a lógica de domínio, incluindo a definição dos modelos de dados, regras de negócios.
- `controller/`: Camada responsável pela comunicação entre a interface e a lógica de domínio.
- `use-case/`: Implementação dos casos de uso do sistema, como a geração do relatório.
- `infra/`: Contém as implementações de infraestrutura, como repositórios, conexões com bancos de dados, clients de APIs externas e serviços auxiliares.

## Fluxo da Solução
### Geração do Relatório:
- O sistema gera relatórios financeiros baseados em transações (compras e vendas de ativos) e preços de ativos durante um intervalo de tempo especificado.
- O arquivo CSV gerado contém o saldo total de patrimônio e a rentabilidade acumulada a cada intervalo de tempo, permitindo análises detalhadas.

### Input:
- O sistema recebe dados de transações e preços em formato de arquivos CSV através de uma API gRPC.
- A entrada inclui informações como data, quantidade e preço dos ativos, além do intervalo de tempo e o saldo inicial.


Abaixo está a descrição e o formato esperado para cada um dos parâmetros de entrada no sistema.

| Parâmetro        | Descrição                                                                                                                                                                         |
|------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `start_date`     | Data de início do período a ser analisado. **Formato esperado**: `YYYY-MM-DD HH:MM:SS`                                                                                            |
| `end_date`       | Data de término do período a ser analisado. **Formato esperado**: `YYYY-MM-DD HH:MM:SS`                                                                                           |
| `interval_minutes` | Intervalo de tempo em minutos para agrupar as transações. **Exemplo**: 15 (15 minutos)                                                                                            |
| `initial_balance` | O saldo inicial, representando o patrimônio inicial para cálculos. **Exemplo**: 100_000                                                                                           |
| `trades_file`    | Arquivo contendo os dados das transações realizadas. **Formato esperado**: Arquivo CSV (bytes)                                                                                    |
| `assets_files`   | Mapa contendo arquivos CSV com dados de preços dos ativos, a chave sendo o nome do ativo, assim como aparece no arquivo de transações. **Formato esperado**: `map<string, bytes>` |

## Exemplo de Entrada

1. `start_date`: "2021-03-01 10:00:00"
2. `end_date`: "2021-03-31 17:59:00"
3. `interval_minutes`: 15
4. `initial_balance`: 100_000
5. `trades_file`: Arquivo CSV com transações de compra e venda.
6. `assets_files`: Mapa de arquivos CSV para ativos como "BTC", "ETH".


### Execução:
- Leitura e processamento dos arquivos CSV são feitos em paralelo para melhorar a performance.
- Os dados são carregados em memória de forma otimizada, armazenando apenas as informações necessárias para os cálculos e geração do relatório.
- A camada de **use-case** processa os dados, aplicando as regras de negócio.
- O resultado final é retornado via gRPC, garantindo eficiência na comunicação.

### Saída:
- O sistema gera um arquivo CSV com o relatório financeiro detalhado.
- O arquivo contém os campos: **timestamp, Patrimônio Total e Rentabilidade Acumulada** para cada intervalo.


### Como Converter a Resposta Base64 em Arquivo
A resposta do gRPC retorna o arquivo em **Base64**. Para convertê-lo de volta para um arquivo CSV, siga estas etapas:

1. Copie o conteúdo Base64 retornado pela API.
2. Acesse [Base64 to File](https://www.base64decode.org/) ou use o seguinte comando no terminal:

   ```bash
   echo "SEU_BASE64_AQUI" | base64 --decode > relatorio.csv
   ```

Agora você pode abrir o `relatorio.csv` em qualquer editor de planilhas!

## Auditoria:
Cada operação realizada no sistema (geração de relatório) é registrada com informações de auditoria como a data, IP do cliente e dados da execução.

#### Como Ver os Logs de Auditoria?
Os logs de auditoria são armazenados em arquivos de log no diretório `/logs/audit.log`. Para visualizar os logs basta acessar o diretório na raiz do projeto.
## Como Rodar a Solução
### 1. Usando Golang
#### **Pré-requisitos:**
- **Golang**: Certifique-se de ter o Go instalado. Se não tiver, siga as instruções de instalação: [Golang - Install](https://go.dev/doc/install).
- **Dependências**: As dependências são gerenciadas via `go mod`. Execute o comando abaixo para instalar as dependências:

```bash
go mod tidy
```

#### **Rodar a aplicação:**
Para iniciar o servidor gRPC que expõe a API de geração de relatórios, execute o comando:

```bash
go run main.go
```

### 2. Usando Docker
Para rodar a solução utilizando **Docker**, basta executar:

```bash
docker-compose up -d --build
```

Ou se tiver o **Make**, execute:

```bash
make up
```

Isso irá expor o serviço na porta `50051`.

## Como Acessar o gRPCUI
O **gRPCUI** é uma interface gráfica para interagir com a API gRPC. Para acessá-lo, execute:

```bash
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
grpcui -plaintext localhost:50051
```

Após isso, abra o navegador e acesse a url presente no console:

```
gRPC Web UI available at ...
```

Ou se tiver iniciado a aplicação pelo docker, basta acessar:

```bash
http://localhost:8080/
```

## Testes
Os testes foram implementados para garantir que a lógica do sistema está funcionando corretamente. Para rodar os testes, use o comando:

```bash
go test ./...
```

## Resposta à Pergunta

Considerando um patrimônio inicial de R$ 100.000,00 e os arquivos de exemplo `march_2021_trades.csv`, `march_2021_pricesA.csv` e `march_2021_pricesB.csv`, a análise mostra que foi mais vantajoso para Alice ter operado ao longo do dia.

### Cálculos:

- **Retorno de Alice**: 0,448% → Valor final: **R$ 100.448,00**
- **Ativo A**: -0,208% → Valor final: **R$ 99.792,00**
- **Ativo B**: -0,503% → Valor final: **R$ 99.497,00**

### Conclusão:

Se Alice tivesse comprado 100% em **ativo A** ou **ativo B** no início do período, teria tido um retorno negativo. Já com as **trades** realizadas, ela obteve um retorno positivo de **R$ 448,00**.

Portanto, a melhor estratégia foi operar durante o dia, ao invés de comprar 100% em um único ativo.

## Pontos de Melhoria

- **Explorar mais edge cases nos testes unitários**: Incluir casos de borda como dados faltantes, valores extremos ou inconsistentes.
- **Criar testes de integração utilizando BDD**: Implementar testes de integração que validem o comportamento da aplicação com base em exemplos de cenários de negócios.
- **Utilizar uma base para salvar os dados da auditoria**: Criar uma base de dados em memória ou persistente para armazenar os dados de auditoria, garantindo a integridade e o histórico.
- **Adicionar mais validações aos dados de entrada**: Implementar verificações para garantir que os dados dos arquivos CSV estejam no formato correto, com campos não nulos e valores válidos.
- **Tratamento de erros mais elaborado**: Melhorar o tratamento de erros, garantindo que falhas na leitura dos arquivos ou nas operações de cálculo sejam tratadas de forma amigável e que o sistema forneça mensagens de erro claras para o usuário.

**Obs:** Melhorias já implementadas na v2 da solução, presente na branch develop. A branch main se mantém com a v1 que é a implementação feita dentro do prazo do desafio.

## Considerações Finais
A solução foi desenvolvida de forma incremental, seguindo a metodologia Git Flow para organizar e gerenciar o ciclo de desenvolvimento, com foco em legibilidade, modularidade e manutenção. Os testes garantem a qualidade do código, e a auditoria adiciona uma camada importante de rastreabilidade.
