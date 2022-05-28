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
