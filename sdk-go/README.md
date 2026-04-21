### 🇧🇷 Versão em Português

**Visão Geral**
Este módulo funciona como o SDK de Telemetria, desenvolvido em Go. Seu papel principal é atuar como um cliente de alta performance que se comunica com um motor de processamento em Rust via gRPC. Ele abstrai a complexidade da serialização Protobuf e do gerenciamento de rede, oferecendo uma interface simples para ingestão de dados.

**Por que rodar este projeto?**
Em um sistema de métricas distribuídas, o Go oferece o equilíbrio ideal entre velocidade de desenvolvimento e eficiência de execução. Este SDK permite que sensores ou serviços externos enviem dados de telemetria (ID e Valor) para o motor central sem a necessidade de implementar o protocolo gRPC do zero.

**Como funciona**
1. **Contrato Primeiro:** Utiliza uma definição `.proto` para garantir que o servidor Rust e o cliente Go falem a mesma língua.
2. **Abstração:** O pacote `metrics` esconde a lógica de conexão gRPC (`Dial`, `Context`, `Insecure Credentials`).
3. **Transmissão:** Os dados são enviados como uma mensagem `MetricRequest`, otimizada para baixa latência.

**Erros Tratados e Corrigidos**
* **Sobreposição de Pacotes:** Corrigido o erro crítico de "main redeclared" separando a lógica do SDK em um pacote `metrics/` e a interface de linha de comando em `main.go`.
* **Visibilidade/Exportação:** Resolvidos problemas onde as structs geradas não eram acessíveis, garantindo o uso de letras maiúsculas no `.proto` e arquivos Go (nomes exportados).
* **Resolução de Dependências:** Ajustes no `go.mod` para garantir que imports locais (`sdk-go/gen`) funcionem corretamente sem necessidade de hospedagem externa.
* **Encerramento Seguro:** Implementação do fechamento de conexão para evitar vazamentos de memória e sockets TCP órfãos.

**Como Rodar**
1. **Pré-requisitos:** Certifique-se de ter o Go 1.21+ instalado.
2. **Dependências:** Execute `go mod tidy` para baixar as bibliotecas gRPC e Protobuf.
3. **Executando o Cliente:** Execute `go run main.go`.
   * *Nota: O servidor Rust deve estar rodando em localhost:50051 para uma conexão bem-sucedida.*
  
---


### 🇺🇸 English Version

**Overview**
This module serves as the Telemetry SDK, built in Go. Its primary role is to act as a high-performance client that communicates with a Rust-based processing engine via gRPC. It abstracts the complexity of Protobuf serialization and network handling, providing a simple interface for data ingestion.

**Why Run This?**
In a distributed metrics system, Go provides the ideal balance between development speed and execution efficiency. This SDK allows external sensors or services to send telemetry data (ID and Value) to the central motor without needing to implement the gRPC protocol from scratch.

**How It Works**
1. **Contract First:** It uses a `.proto` definition to ensure the Rust server and Go client speak the same language.
2. **Abstraction:** The `metrics` package hides the gRPC connection logic (`Dial`, `Context`, `Insecure Credentials`).
3. **Transmission:** Data is sent as a `MetricRequest` message, optimized for low latency.

**Errors Handled & Fixed**
* **Package Shadowing:** Fixed a critical "main redeclared" error by separating the SDK logic into a dedicated `metrics/` package and the CLI into `main.go`.
* **Visibility/Exporting:** Resolved issues where generated structs weren't accessible by ensuring proper capitalization in the `.proto` and Go files (Exported names).
* **Dependency Resolution:** Fixed `go.mod` conflicts to ensure local imports (`sdk-go/gen`) work correctly without external hosting.
* **Graceful Shutdown:** Implemented connection closing to prevent memory leaks and dangling TCP sockets.

**How to Run**
1. **Pre-requisites:** Ensure you have Go 1.21+ installed.
2. **Dependencies:** Run `go mod tidy` to download gRPC and Protobuf libraries.
3. **Running the Client:** Execute `go run main.go`.
   * *Note: The Rust server must be running on localhost:50051 for a successful connection.*
