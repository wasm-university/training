

spin templates list
spin templates install --git https://github.com/fermyon/spin

```
+---------------------------------------------------+
| Name         Description                          |
+===================================================+
| http-go      HTTP request handler using (Tiny)Go  |
| http-rust    HTTP request handler using Rust      |
| redis-go     Redis message handler using (Tiny)Go |
| redis-rust   Redis message handler using Rust     |
+---------------------------------------------------+
```

mkdir hello
cd hello
spin new http-go --name hello --value variable=hello

https://github.com/fermyon/spin/tree/main/examples/http-tinygo