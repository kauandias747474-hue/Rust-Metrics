# 🚀 Rust-Metrics: The Deterministic Telemetry Engine

![Rust](https://img.shields.io/badge/rust-%23E32A1C.svg?style=for-the-badge&logo=rust&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-4285F4?style=for-the-badge&logo=google-cloud&logoColor=white)
![Zero-Copy](https://img.shields.io/badge/Zero--Copy-rkyv-orange?style=for-the-badge)
![Status](https://img.shields.io/badge/status-active-brightgreen?style=for-the-badge)

---

## 🇧🇷 Português

### 🎯 Por que Rust-Metrics?
O nome **Rust-Metrics** nasceu da necessidade de eliminar a incerteza. Em sistemas de telemetria de alta escala, o maior inimigo é o **Garbage Collector (GC)**. Linguagens como Go e Java sofrem pausas imprevisíveis que resultam em perda de pacotes e latência instável. 

O **Rust-Metrics** utiliza Rust para garantir processamento **determinístico**: sem pausas, sem perda de dados, apenas performance bruta em nível de hardware.

### 🏗️ Arquitetura Poliglota (Novas Integrações)
Para tornar o projeto uma solução completa, integramos o motor Rust com as melhores ferramentas de cada domínio:

1.  **Go (Control Plane):** Adicionado em `sdk-go/` para atuar como orquestrador. O Go é o padrão ouro para infraestrutura cloud e Kubernetes. Ele gerencia o ciclo de vida do motor Rust.
2.  **Python (Analysis Plane):** Adicionado em `sdk-python/`. Enquanto o Rust processa, o Python utiliza bibliotecas como **Polars** e **Pandas** para análise estatística e detecção de anomalias em tempo real.
3.  **TypeScript (Visual Plane):** Adicionado em `dashboard-ts/`. Um frontend moderno que consome dados via **gRPC-Web**, garantindo que os gráficos reflitam o que está acontecendo no motor com o mínimo de atraso visual.

---

## 🇺🇸 English

### 🎯 Why Rust-Metrics?
**Rust-Metrics** was engineered to kill uncertainty. In high-throughput telemetry, the **Garbage Collector (GC)** is the enemy. Languages like Go and Java have unpredictable pauses that cause packet loss and jitter.

**Rust-Metrics** leverages Rust to provide **deterministic** processing: no pauses, no data loss, just raw hardware-level performance.

### 🏗️ Polyglot Architecture (New Tooling)
To turn this engine into a full-scale ecosystem, we've integrated Rust with the best-in-class tools for every domain:

1.  **Go (Control Plane):** Located in `sdk-go/`. Go is the industry standard for cloud infrastructure. It acts as the orchestrator, managing node health and configuration.
2.  **Python (Analysis Plane):** Located in `sdk-python/`. While Rust handles the ingestion, Python uses **Polars** and **ML libraries** to perform real-time statistical analysis and anomaly detection.
3.  **TypeScript (Visual Plane):** Located in `dashboard-ts/`. A modern UI that consumes data via **gRPC-Web**, ensuring dashboards stay synced with the engine with microsecond-level precision.

---

## 🛠️ Arsenal de Ferramentas / Toolchain

| Camada / Layer | Ferramenta / Tool | Propósito / Purpose |
| :--- | :--- | :--- |
| **Engine Core** | `Rust` | Processamento determinístico de baixa latência. |
| **Runtime** | `Tokio` | I/O assíncrono para lidar com milhões de conexões. |
| **Serialization** | `gRPC / Protobuf` | Contrato universal entre Rust, Go, Python e TS. |
| **Zero-Copy** | `rkyv` | Desserialização sem alocação na Heap (Velocidade máxima). |
| **Orchestration** | `Go` | Gerenciamento de containers e API de controle. |
| **Analytics** | `Python / Polars` | Processamento de dados e inteligência artificial. |
| **Frontend** | `TS / React` | Interface visual de alta fidelidade e tempo real. |

----

**🚀 Como Iniciar / Quick Start**
Contratos: make gen-proto (Gera o código para todas as linguagens).

Motor: cargo run --release (Inicia o servidor gRPC em Rust).

Análise: python sdk-python/client.py (Inicia o consumo de dados).
---

**📜 Licença / License:** 
*MIT License. Free for all.*


## 📂 Estrutura de Pastas / Project Structure

```text
├── src/                # Motor Principal em Rust
├── proto/              # Contratos gRPC (Fonte da Verdade)
├── sdk-go/             # Orquestrador e Ferramentas Cloud (Go)
├── sdk-python/         # Análise de Dados e ML (Python)
├── dashboard-ts/       # Interface Web (TypeScript/React)
├── benches/            # Relatórios de Performance (Criterion)
└── .github/workflows/  # CI/CD de Performance e Automação
