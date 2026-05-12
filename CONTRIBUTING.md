
# Contributing to Real-Time Metrics / Contribuindo para o Real-Time Metrics

## 🇧🇷 Português

Obrigado por dedicar seu tempo para contribuir! Este projeto é um ecossistema de baixa latência e alta densidade de dados. Para manter a excelência técnica, estabelecemos as seguintes diretrizes:

###  Nosso Foco: Performance Bruta
Qualquer contribuição deve respeitar o "Hot-Path" do sistema. Mudanças que introduzam alocações desnecessárias na Heap, bloqueios de thread ou aumentem a complexidade ciclomática sem ganho de performance serão revisadas com rigor.

###  Processo de Submissão
1. **Issues Primeiro:** Antes de qualquer PR grande, abra uma Issue para discutir a abordagem técnica.
2. **Qualidade de Código:** - **C++:** Use `clang-format`. Foque em semântica de movimento (move semantics) e evite cópias.
   - **Rust:** O código deve passar no `cargo clippy` sem avisos.
   - **Go/Python:** Siga rigorosamente os linters padrão.
3. **Benchmarks:** PRs que afetam o core de processamento devem incluir resultados de benchmark (usando Criterion em Rust ou benchmarks nativos em C++).

### 🛠️ Configuração de Ambiente
- Certifique-se de ter o `LLVM`, `Rust (Stable)`, `Go` e `Protocol Buffers` instalados.
- Use o script `./scripts/generate_proto.sh` após alterar qualquer arquivo `.proto`.

---

## 🇺🇸 English

Thank you for your interest in contributing! This project is a low-latency, high-data-density ecosystem. To maintain technical excellence, we have established the following guidelines:

### Our Focus: Raw Performance
Any contribution must respect the system's "Hot-Path." Changes introducing unnecessary Heap allocations, thread locking, or increased cyclomatic complexity without a clear performance gain will be strictly reviewed.

###  Submission Process
1. **Issues First:** Before any major Pull Request, open an Issue to discuss the technical approach.
2. **Code Quality:** - **C++:** Must use `clang-format`. Focus on move semantics and avoid copies.
   - **Rust:** Code must pass `cargo clippy` without warnings.
   - **Go/Python:** Strictly follow standard linters.
3. **Benchmarks:** PRs affecting the processing core must include benchmark results (using Criterion for Rust or native C++ benchmarks).

### 🛠️ Environment Setup
- Ensure you have `LLVM`, `Rust (Stable)`, `Go`, and `Protocol Buffers` installed.
- Run the `./scripts/generate_proto.sh` script after modifying any `.proto` files.

---

### 📋 Convenções de Commit / Commit Conventions
Usamos **Conventional Commits**:
- `feat:` Nova funcionalidade.
- `fix:` Correção de bug.
- `perf:` Mudança focada exclusivamente em performance.
- `docs:` Alterações na documentação.
