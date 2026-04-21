# Technical Documentation/Documentação Técnica


---

## 🇧🇷 Versão em Português

### Arquitetura do Sistema e Integração
O sistema é baseado em um modelo de comunicação por contrato. O arquivo de definição Interface Definition Language (IDL), denominado `telemetry.proto`, estabelece as regras de comunicação que todas as linguagens devem seguir.

* **Camada de Contrato (Protocol Buffers):** Define as mensagens `MetricRequest` e `MetricResponse`. Ao compilar este arquivo, geramos stubs específicos para cada linguagem, garantindo que o dado enviado pelo Go ou Python seja exatamente o dado esperado pelo servidor Rust.
* **Servidor Central (Rust):** Atua como o núcleo de processamento. Escolhido por sua segurança de memória e performance preditível, ele gerencia a recepção de buffers binários via HTTP/2.
* **SDK de Ingestão (Go):** Projetado para alta concorrência. Utiliza a eficiência do runtime de Go para gerenciar múltiplas conexões de sensores simultâneos com baixo overhead.
* **SDK de Análise (Python):** Provê uma interface de alto nível para manipulação de dados, ideal para integração com ferramentas de análise estatística e ciência de dados.

### Conceitos de Engenharia Aplicados

#### Comunicação Binária vs. Textual
Diferente de APIs REST tradicionais que utilizam JSON (texto), este sistema utiliza gRPC. Os dados são serializados em formato binário, o que reduz drasticamente o tamanho do payload e o custo computacional de parsing, resultando em menor latência.

#### Tipagem Forte e Segurança em Tempo de Compilação
A utilização de Protocol Buffers impede erros comuns de integração, como campos ausentes ou tipos de dados incompatíveis. Se o cliente Go tentar enviar uma string onde o servidor Rust espera um float, o código não será compilado ou falhará imediatamente na validação do contrato.

### Erros Tratados e Soluções Técnicas

1.  **Resolução de Namespaces em Go:** Durante o desenvolvimento do SDK em Go, foi identificado um erro de redeclaração de pacotes (*shadowing*). A solução consistiu na segregação da lógica de negócio em um pacote específico (`metrics/`) e da interface de linha de comando no arquivo raiz (`main.go`). Isso permitiu que o compilador resolvesse corretamente as dependências internas sem conflitos de escopo.
2.  **Manipulação de Caminhos Dinâmicos em Python:** No SDK Python, o erro `ModuleNotFoundError` era frequente devido à forma como o gRPC gera os arquivos de stub. Foi implementada uma solução utilizando `sys.path.append(os.path.dirname(os.path.abspath(__file__)))`. Isso garante que o interpretador Python localize os módulos gerados independentemente do diretório de execução, tornando o SDK portátil.
3.  **Resiliência de Conexão e Timeouts:** Em sistemas distribuídos, um servidor lento pode causar o travamento de todos os clientes. Implementamos o conceito de Contexto (*Context*) no Go e *Timeouts* no Python. Isso assegura que, se o servidor Rust não responder em um tempo determinado, o cliente encerre a tentativa de conexão, liberando recursos do sistema.

### Instruções de Operação

#### Fluxo de Execução
Para o funcionamento correto da malha de serviços, a ordem de inicialização é mandatória:
1.  **Servidor Rust:** Deve ser iniciado primeiro para abrir a porta de escuta 50051 e carregar o motor de processamento.
2.  **Clientes (Go e Python):** Uma vez que o servidor reporte o estado de "Listening", os SDKs podem ser executados para realizar a ingestão de telemetria.

#### Manutenção do Contrato
Qualquer alteração na estrutura de dados deve ser feita exclusivamente no arquivo `telemetry.proto`. Após a alteração, os comandos de geração de código (`protoc`) devem ser reexecutados para todas as linguagens para garantir que os SDKs permaneçam sincronizados com o motor central.

---

## 🇺🇸 English Version

### System Architecture and Integration
The system is based on a contract-first communication model. The Interface Definition Language (IDL) definition file, named `telemetry.proto`, establishes the communication rules that all languages must follow.

* **Contract Layer (Protocol Buffers):** Defines the `MetricRequest` and `MetricResponse` messages. By compiling this file, we generate language-specific stubs, ensuring that data sent via Go or Python matches exactly what the Rust server expects.
* **Central Server (Rust):** Acts as the processing core. Chosen for its memory safety and predictable performance, it manages the reception of binary buffers via HTTP/2.
* **Ingestion SDK (Go):** Designed for high concurrency. It leverages the efficiency of the Go runtime to manage multiple simultaneous sensor connections with low overhead.
* **Analysis SDK (Python):** Provides a high-level interface for data manipulation, ideal for integration with statistical analysis tools and data science.

### Applied Engineering Concepts

#### Binary vs. Textual Communication
Unlike traditional REST APIs that utilize JSON (text), this system uses gRPC. Data is serialized into a binary format, which drastically reduces payload size and computational parsing costs, resulting in lower latency.

#### Strong Typing and Compile-Time Safety
The use of Protocol Buffers prevents common integration errors, such as missing fields or incompatible data types. If a Go client attempts to send a string where the Rust server expects a float, the code will not compile or will fail immediately during contract validation.

### Handled Errors and Technical Solutions

1.  **Go Namespace Resolution:** During the development of the Go SDK, a package redeclaration error (*shadowing*) was identified. The solution involved segregating the business logic into a specific package (`metrics/`) and the command-line interface into the root file (`main.go`). This allowed the compiler to correctly resolve internal dependencies without scope conflicts.
2.  **Dynamic Path Handling in Python:** In the Python SDK, `ModuleNotFoundError` was frequent due to the way gRPC generates stub files. A solution was implemented using `sys.path.append(os.path.dirname(os.path.abspath(__file__)))`. This ensures the Python interpreter locates the generated modules regardless of the execution directory, making the SDK portable.
3.  **Connection Resilience and Timeouts:** In distributed systems, a slow server can cause all clients to hang. We implemented the concept of Context in Go and Timeouts in Python. This ensures that if the Rust server does not respond within a specific timeframe, the client terminates the connection attempt, releasing system resources.

### Operating Instructions

#### Execution Flow
For the service mesh to function correctly, the initialization order is mandatory:
1.  **Rust Server:** Must be started first to open listening port 50051 and load the processing engine.
2.  **Clients (Go and Python):** Once the server reports a "Listening" status, the SDKs can be executed to perform telemetry ingestion.

#### Contract Maintenance
Any changes to the data structure must be made exclusively in the `telemetry.proto` file. Following any changes, the code generation commands (`protoc`) must be re-executed for all languages to ensure the SDKs remain synchronized with the central engine.
