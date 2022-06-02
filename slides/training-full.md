---
theme: uncover
size: 16:9
paginate: true
---
<style scoped>
  mark {
    background-color: #942EC1;
    color: #FFFFFF;
  }
</style>
# Petit cours de <mark>Wasm</mark> front/back par l'exemple

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
  - en utilisant <mark>Gitpod</mark>
  - ou en utilisant <mark>DevContainer</mark>

---

# Objectif(s)

- Université "découverte" par l’exemple
- Rien de complexe
- Repartir avec le bagage nécessaire 🧳

🖐️ Ne posez pas de questions compliquées 😛🙏
https://github.com/wasm-university/training/issues

---

# WebAssembly ???

## WASM ???
### C'est parti ! 🚀

---

# 🚧🚧🚧 Wasm Quoi/Pourquoi ?
> - prendre le schéma qui explique ce que c'est (+ajouter pourquoi)
> - 

---

# 🚧🚧🚧 Histoire
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

# 🚧🚧🚧 Hosts
> - prendre le schéma qui explique ce que c'est
> - 
<!-- la portabilité de wasm dépend de l'hôte -->

---
![bg](#C4D8F8)
# Wasm a "quelques" limitations

---

### 🖐️ Le module Wasm n’accède pas à l’OS

- Wasm c’est pour du compute (au départ)
- Pas d’accès aux fonctions systèmes de l’OS (hors host functions)
  - I/O
  - Sockets
- Pas d’accès à la mémoire hors allocation spécifique
<!-- vérifier cette partie -->

---

## C'est une bonne limitation
### <mark>Safe by default</mark>

---

### 📣 La Communication Wasm <=> Host  n’est pas triviale  
> (trop bas niveau ?)
  
#### 4 types de données pour les paramètres: 
  
  - 32 & 64 Bit Integer
  - 32 & 64 Bit Floating-Point Number

---

## String 😡

---

> Certains "hôtes" (et toolchains) ont déjà tout prévu (certains frameworks aussi pour WASI)

---

![bg](#B8F6C5)
# 🛠 ToolChains

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
Rust            | Wasm-pack + wasm-bindgen (glue) | rustup target add wasm32-wasi            | support navigateur 💖    
Go              | Intégré à la toolchain standard | Non ou alors utiliser TinyGo             | support navigateur 💖
Assemblyscript  | Intégré                         | Intégré                                  | Ne cible que du WASM
Swift           | SwiftWasm                       | SwiftWasm                                |
Kotlin          | Kotlin native (expérimental)    |                                          |
C#              | Blazor (solution complète)      | dotnet add package Wasi.Sdk --prerelease |
Ruby            | Artichoke                       | En cours (portage CRuby Wasm32-WASI)     |
Python          | Expérimental                    |                                          |

<!-- regarder prez de Sébastien pour Kotlin -->
###### *Liste non exhaustive*
---

![bg](#3A84F2)
![fg](#FFFFFF)
# Wasm & le Navigateur

---

- 1er contact: un peu de C
- Go
- Rust

---

### 🚧🚧🚧 Mode de fonctionnement des démos Web

> refaire un schema avec index.html et tout le touin touin

---

![bg](#3AF1F2)
![fg](#000000)

# 1er module Wasm en C

---

`main.c`
```c
#define WASM_EXPORT __attribute__((visibility("default")))

WASM_EXPORT 
float power(float number, int pow) {
 float res = number;
   for (int i = 0;i < pow - 1; i++) {
     res = res * number;
   }
 return res;
}

WASM_EXPORT 
char* greet()
{
    static char str[12] = "hello world!";
    return (char*)str;
}
```

---

```bash
clang --target=wasm32 \
  --no-standard-libraries -Wl,--export-all -Wl, \
  --no-entry -o main.wasm main.c
```

---

`index.html`
```javascript
WebAssembly.instantiateStreaming(fetch("main.wasm")) 
  .then(({ instance }) => {
    console.log("👋 main.wasm is loaded")
    
    const value = instance.exports.power(2, 2)

    console.log(`🤖 value: ${value}`)
    console.log(`👋 greet: ${instance.exports.greet()}`)

  })
  .catch(error => {
    console.log("😡 ouch", error)
  })
```

---
![bg](#000000)
![fg](#FFFFFF)
# Démo 🚀


<a href="https://github.com/wasm-university/training/tree/main/00-c-web" target="_blank">00-c-web</a>

---
![bg](#3AF1F2)
![fg](#000000)
# Wasm avec Go dans le navigateur

---

<style scoped>
  mark {
    background-color: #EFD217;
    color: #000000;
  }
</style>

# Go + JavaScript = 💖

```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

```html
<script src="wasm_exec.js"></script>
```
 
 > Disclaimer, I 💛 <mark>**JavaScript**</mark>
---

## Fonctions en Go: 
`[]js.Value` <mark>&</mark> `interface{}`

```go
func Hello(this js.Value, args []js.Value) interface{} {
  message := args[0].String() // get the parameters
  return "😃 Hello " + message
}
```

```go
js.Global().Set("Hello", js.FuncOf(Hello))
```

<!-- 
Et avec ça, on peut faire plein de choses ... 
Comme en JavaScript 😉
-->

---

## Utilisation de la fonction Go

```javascript
const go = new Go() // Go Wasm runtime

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject) 
  .then(result => { // Get the importObject from the go instance.
    // execute `main`
    go.run(result.instance)
    // instance object contains 
    // all the Exported WebAssembly functions	
    Hello("Bob Morane")
    //😃 Hello "Bob Morane
  })
  .catch(error => {
    console.log("😡 ouch", error)
  })
```

<!-- 
Il est temps de voir quelques exemples
-->

---
![bg](#000000)
![fg](#FFFFFF)
# Démos 🚀

<a href="https://github.com/wasm-university/training/tree/main/01-go-hello" target="_blank">01-go-hello</a>
<a href="https://github.com/wasm-university/training/tree/main/02-go-dom" target="_blank">02-go-dom</a>
<a href="https://github.com/wasm-university/training/tree/main/03-go-call-go-function" target="_blank">03-go-call-go-function</a>
<a href="https://github.com/wasm-university/training/tree/main/04-go-return-json" target="_blank">04-go-return-json</a>
<a href="https://github.com/wasm-university/training/tree/main/05-go-json-as-parameter" target="_blank">05-go-json-as-parameter</a>
<a href="https://github.com/wasm-university/training/tree/main/06-go-call-a-js-function" target="_blank">06-go-call-a-js-function</a>
<a href="https://github.com/wasm-university/training/tree/main/07-go-call-js-promise" target="_blank">07-go-call-js-promise</a>

---
![bg](#3AF1F2)
![fg](#000000)
# Wasm avec Rust dans le navigateur

## 🦀 + 🕸️ = 💖

https://rustwasm.github.io/

---

# Facile ?

---

# Avec Wasm Bindgen, OUI ‼️

https://github.com/rustwasm/wasm-bindgen
> Facilitating high-level interactions between Wasm modules and JavaScript

---

#### Créer un projet "Rust Wasm"

<mark>Créer un projet de type "library"</mark>

```bash
cargo new --lib hello
```

<mark>Mise à jour de `Cargo.toml`</mark>

```toml
[lib]
name = "hello"
path = "src/lib.rs"
crate-type =["cdylib"]

[dependencies]
wasm-bindgen = "0.2.50"
```

---

<mark>Modifier `main.rs`<mark>

```rust
use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub fn hello(s: String) -> String {
  let r = String::from("👋 hello ");
  
  return r + &s;
}
```

<mark>Compiler<mark>

```bash
cd hello
wasm-pack build --release --target web
```
> 🖐️ `--target web`
---

<mark>Utiliser<mark>

```html
<script type="module">
  import init, { hello } from './hello/pkg/hello.js'

  async function run() {
    await init()
    console.log(hello("Bob")) 
    console.log(hello("Jane")) 
    console.log(hello("John")) 
  }
  run();
</script>
```

---

![bg](#000000)
![fg](#FFFFFF)
# Démos 🚀

<a href="https://github.com/wasm-university/training/tree/main/08-rust-hello" target="_blank">08-rust-hello</a>
<a href="https://github.com/wasm-university/training/tree/main/09-rust-call-with-json" target="_blank">09-rust-call-with-json</a>
<a href="https://github.com/wasm-university/training/tree/main/10-rust-dom" target="_blank">10-rust-dom</a>


---
<style scoped>
  mark {
    background-color: #EFD217;
    color: #000000;
  }
  mark-green {
    background-color: #12984E;
    color: #000000;
  }
</style>

![bg](#3AF1F2)
![fg](#000000)
# Wasm & <mark-green>NodeJS</mark-green>

## VM <mark>JavaScript</mark>

---

![bg](#18CA8B)
![fg](#000000)
# C'est comme pour le navigateur ... 😍

🖐️ Attention, pour Rust :

```bash
wasm-pack build --release --target nodejs
```

---

![bg](#000000)
![fg](#FFFFFF)
# Démos 🚀

<a href="https://github.com/wasm-university/training/tree/main/11-nodejs-go-function" target="_blank">11-nodejs-go-function</a>
<a href="https://github.com/wasm-university/training/tree/main/12-nodejs-rust-function" target="_blank">12-nodejs-rust-function</a>

---
![bg](#3AF1F2)
![fg](#000000)

# Cas d'utilisation
### (Wasm dans le navigateur)
---

# Quelques applications

- Jeux Vidéos
- "Vraies" applications
- Traitement d’image en local (dans le navigateur), OCR
- Cartographie
- Machine Learning
- Chiffrement dans le navigateur
- ...

---

### https://web.autocad.com

![w:800](pictures/autocad.png)

---

#### https://beta.unity3d.com/jonas/AngryBots/

![w:800](pictures/unity.png)

<!--
https://blog.unity.com/technology/webassembly-is-here
-->
---

#### https://github.com/naptha/tesseract.js

![w:800](pictures/tesseract.png)

<!-- OCR ordonances Doctolib -->
---

# Plus besoin de l’AppStore ? 😬

<!-- l'avenir nous le dira -->

---
![bg](#3217EF)
![fg](#FFFFFF)

<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
</style>

# Libérez, délivrez Wasm
## ... du navigateur (de la VM JS)
# <mark>WASI</mark>
https://wasi.dev/

---

<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
</style>

### WASI: WebAssembly System Interface
#### WebAssembly comme <mark>"Portable Runtime"</mark>

WASI == Les fondations pour "sortir" Wasm du navigateur

#### Sous-groupe de spécifications WebAssembly

---

<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
</style>

### Comme la JVM (conceptuellement)
#### mais en mieux ?

- Sécurisé
- Polyglotte
- Rapide
- <mark>Léger</mark>

---

<style scoped>
  mark {
    background-color: #F7C00E;
    color: #000000;
  }
</style>

### Un module WebAssembly <mark>ne peut pas</mark>

- Accéder au système d’exploitation
- Accéder à la mémoire que le host ne lui a pas donnée
- Faire des requêtes sur le réseau
- Lire ou écrire dans des fichiers


---
<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
</style>

**WASI est une spécification pour pouvoir fournir un accès <mark>sécurisé et isolé</mark> au système sur lequel s’exécute <mark>l’hôte du module Wasm</mark>.**

---

# 🚧🚧🚧 Host Runtime

- Ici refaire un schéma comme pour le slide 55

---

## Les projets de runtimes WASI

- Pour exécuter du code Wasm à partir d’une CLI
- Pour exécuter du code Wasm à partir d’un autre langage (Rust, Go, C/C++) >> SDK

---

## Les 3 les plus reconnus du moment :

- **Wasmer**: https://wasmer.io/
- **Wasmtime**: https://wasmtime.dev/
- **WasmEdge**: https://wasmedge.org/
---

##### SDK WASI / Langage <mark>(<> CLI)</mark>

<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
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

Langage             | WASMER                   | WASMEDGE (+arm)           | WASMTIME (+arm)
:-------------------|:-------------------------|:--------------------------|:--------
  <mark>Rust</mark> |  x                       |  x                        |  x
  <mark>Go</mark>   |  x (<mark>TinyGo</mark>) |  x  (<mark>TinyGo</mark>) |  x (<mark>TinyGo</mark>)
  <mark>C</mark>    |  x                       |  x                        |  x
  C++               |  x                       |                           |
  Python            |  x                       |  x                        |  x
  Swift             |  x                       |  x                        |  ?
  Grain             |                          |  x                        |  ?
  .Net              |  x (C#)                  |                           |  x
  NodeJS            |  x                       |  x                        |
  Bash              |                          |                           |  x
  Java              |  x                       |                           |  x (outside Bytecode Alliance)
  Perl              |                          |                           |  x (outside Bytecode Alliance)
  Zig               |  x (not published)       |                           |  x (outside Bytecode Alliance)
  Ruby              |                          |                           |  x (outside Bytecode Alliance)
> - Wasmer supporte d'autres langages
---

# Bytecode Alliance
<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
</style>

https://bytecodealliance.org/

- WASM / WASI => les spécifications par le W3C, 
- La Bytecode Alliance s’occupe de l’implémentation

Avec Amazon, ARM, <mark>Cosmonic</mark>, Fastly, Google, Intel, <mark>Fermyon</mark>, <mark>Suborbital</mark>, Microsoft, Mozilla, Shopify, Siemens ...

---

![bg](#000000)
![fg](#FFFFFF)
# Démos 🚀
## Les CLI des runtimes WASI

<a href="https://github.com/wasm-university/training/tree/main/13-go-wasi-cli-app" target="_blank">13-go-wasi-cli-app</a>
<a href="https://github.com/wasm-university/training/tree/main/14-rust-wasi-cli-function" target="_blank">14-rust-wasi-cli-function</a>


---

## Utiliser un SDK (WasmEdge)

#### Faites votre propre "CLI Wasm" 🤓

#### Appeler des fonctions Wasm à partir de Go 🚀

La documentation de WasmEdge est 💖
https://wasmedge.org/book/en/embed/go.html
https://github.com/second-state/WasmEdge-go-examples

---

![bg](#000000)
![fg](#FFFFFF)
# Démos 🚀
## Utilisation du SDK WasmEdge

<a href="https://github.com/wasm-university/training/tree/main/15-go-wasmedge-cli" target="_blank">15-go-wasmedge-cli</a>
<a href="https://github.com/wasm-university/training/tree/main/16-go-wasmedge-function" target="_blank">16-go-wasmedge-function</a>

<!-- montrer le code -->

---
<style scoped>
  mark {
    background-color: #17EFE7;
    color: #000000;
  }
</style>
###  Ok, plutôt facile 😛
### Mais comment je fais si je veux passer une <mark>String</mark> à ma fonction ?


---


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

# Title

---

# Title


---

# Références

- https://wasmbyexample.dev
- WASI: https://wasi.dev/
- Fermyon: https://www.fermyon.com/

---