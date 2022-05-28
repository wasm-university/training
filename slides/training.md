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
### Safe by default

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

---
![bg](#3AF1F2)
![fg](#000000)
# Wasm avec Go dans le navigateur

---

# Title

---

# Title

---

# Title

---

![bg](#3A84F2)
![fg](#FFFFFF)

# Libérez, délivrez Wasm
## ... du navigateur
### WASI
---

# Title

---

# Title

---