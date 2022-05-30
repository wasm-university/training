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

### Mode de fonctionnement des démos Web

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


---

![bg](#3AF1F2)
![fg](#000000)
# Wasm avec Go et NodeJS

---

# Title

---

# Title

---

# Title

---

# Title

---

![bg](#3AF1F2)
![fg](#000000)
# 🦀 Wasm avec Rust et NodeJS

---

# Title

---

# Title

---

# Title

---

# Title

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

---