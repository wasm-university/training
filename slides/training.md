---
theme: uncover
size: 16:9
paginate: true
---

# Petit cours de Wasm front/back par l'exemple

**Voxxed Luxembourg 2022**

https://github.com/wasm-university/training

---
# Merci 😃

**à vous & à l'équipe Voxxed Luxembourg**

---

![bg left:40% 80%](pictures/k33g.png)

#### Philippe Charrière

- TAM @GitLab
- 🦊 @k33g
- 🐦 @k33g_org
- 🍊🦸Gitpod Hero
- GDG Cloud IOT Lyon
- RdV des speakers

---

# Déroulement

- 👋 Vous pouvez intervenir à tout moment
- 10% Théorie 90% Démos (en gros)
- 🚀 Des démos que vous pourrez refaire :
  - https://github.com/wasm-university

---

# Objectif(s)

- Université “découverte” par l’exemple
- Rien de complexe
- Repartir avec le bagage nécessaire 🧳

🖐️ Ne posez pas de questions compliquées 😛

---

# WebAssembly ???

## WASM ???
### C'est parti ! 🚀

---

# Wasm Quoi/Pourquoi ?
> - prendre le schéma qui explique ce que c'est (+ajouter pourquoi)
> - 

---

# Histoire
> - prendre le schéma qui explique ce que c'est
> - 

---

## Wasm peut s’exécuter partout

JavaScript (navigateur)
JavaScript (Node.js)
GraalVM
Runtimes **WASI** (Wasmer, Wastime, Wasmedge, …): CLI & Libs
<!-- webassembly system interface -->
---

Wasm file ~= container image, **smaller**, safer, without an OS

---

# Hosts
> - prendre le schéma qui explique ce que c'est
> - 
<!-- la portabilité de wasm dépend de l'hôte -->

---

### 🖐️ Le module Wasm n’accède pas à l’OS

- Wasm c’est pour du compute (au départ)
- Pas d’accès aux fonctions systèmes de l’OS (hors host functions)
  - I/O
  - Sockets
- Pas d’accès à la mémoire hors allocation spécifique
<!-- vérifier cette partie -->

---

##### Toolchains par langage & hôte

<style scoped>
table {
    height: 80%;
    width: 100%;
    font-size: 20px;
    color: green;
}
th {
    color: blue;
}
</style>

Langage         | WASM (VM JS)                    | WASI                                     | Remarks
:---------------|:--------------------------------|:-----------------------------------------|:--------
C/C++           | EMScripten, LLVM (clang)        | LLVM, SDK C/C++ Wasmer                   |
Rust            | Wasm-pack + wasm-bindgen (glue) | rustup target add wasm32-wasi            |     
Go              | Intégré à la toolchain standard | Non ou alors utiliser TinyGo             |
Assemblyscript  | Intégré                         | Intégré                                  | Ne cible que du WASM
Swift           | SwiftWasm                       | SwiftWasm                                |
Kotlin          | Kotlin native (expérimental)    |                                          |
C#              | Blazor (solution complète)      | dotnet add package Wasi.Sdk --prerelease |
Ruby            | Artichoke                       | En cours (portage CRuby Wasm32-WASI)     |
Python          | Expérimental                    |                                          |

<!-- regarder prez de Sébastien pour Kotlin -->
###### *Liste non exhaustive*
---

# Title

---

# Title

---

# Title

---

# Title

---

# Title

---

# Title

---

# Title

---

# Title

---

# Title

---

# Title

---