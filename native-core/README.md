#  Native Core (C++ Mathematical Engine) / Núcleo Nativo (Motor Matemático C++)

---

## 🇺🇸 English Version

### **Overview**
The `native-core` is the mathematical backbone of the system. Built in C++, it is designed to perform CPU-intensive calculations that require maximum hardware proximity. It is compiled as a shared library (`.dll` or `.so`) to be dynamically invoked by the Rust server.

### **Why C++?**
While Go handles the sensor network and Rust handles system safety, C++ is utilized for its unrivaled performance in heavy mathematical operations (SIMD, vectorization) and its mature ecosystem for scientific computing.

### **FFI Integration (Foreign Function Interface)**
* **Symbol Export:** Uses `extern "C"` to prevent name mangling, allowing Rust to "find" and call the C++ functions.
* **Memory Management:** Implements strict pointer handling to ensure data passed from Rust is processed without memory leaks.

### **Key Features**
* **Optimized StdDev:** High-speed Standard Deviation calculation for large batches of telemetry.
* **Zero-Copy Strategy:** Designed to process data in-place whenever possible to minimize memory overhead.

---

## 🇧🇷 Versão em Português

### **Visão Geral**
O `native-core` é a espinha dorsal matemática do sistema. Desenvolvido em C++, ele é projetado para realizar cálculos intensivos de CPU que exigem proximidade máxima com o hardware. Ele é compilado como uma biblioteca compartilhada (`.dll` ou `.so`) para ser invocado dinamicamente pelo servidor Rust.

### **Por que C++?**
Enquanto o Go gerencia a rede de sensores e o Rust cuida da segurança do sistema, o C++ é utilizado por sua performance inigualável em operações matemáticas pesadas (SIMD, vetorização) e seu ecossistema maduro para computação científica.

### **Integração via FFI (Foreign Function Interface)**
* **Exportação de Símbolos:** Utiliza `extern "C"` para evitar o *name mangling*, permitindo que o Rust localize e chame as funções C++.
* **Gerenciamento de Memória:** Implementa manipulação rigorosa de ponteiros para garantir que os dados vindos do Rust sejam processados sem vazamentos de memória.

### **Principais Funcionalidades**
* **Desvio Padrão Otimizado:** Cálculo de alta velocidade para grandes lotes de telemetria enviados pelo Agente Go.
* **Estratégia Zero-Copy:** Projetado para processar dados no local (*in-place*) sempre que possível para minimizar o overhead de memória.
