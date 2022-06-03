#include <string.h>

int mono_wasm_add_assembly(const char* name, const unsigned char* data, unsigned int size);

extern const unsigned char hello_dll_F832AC0D[];
extern const int hello_dll_F832AC0D_len;
extern const unsigned char System_Console_dll_0566C3A3[];
extern const int System_Console_dll_0566C3A3_len;
extern const unsigned char System_Memory_dll_7E7749D6[];
extern const int System_Memory_dll_7E7749D6_len;
extern const unsigned char System_Private_CoreLib_dll_F5990610[];
extern const int System_Private_CoreLib_dll_F5990610_len;
extern const unsigned char System_Private_Runtime_InteropServices_JavaScript_dll_1CFCEC65[];
extern const int System_Private_Runtime_InteropServices_JavaScript_dll_1CFCEC65_len;
extern const unsigned char System_Threading_dll_D30B628A[];
extern const int System_Threading_dll_D30B628A_len;
extern const unsigned char System_Collections_dll_C69F445F[];
extern const int System_Collections_dll_C69F445F_len;
extern const unsigned char System_Runtime_InteropServices_dll_BC50FF4B[];
extern const int System_Runtime_InteropServices_dll_BC50FF4B_len;
extern const unsigned char System_Runtime_dll_3402885C[];
extern const int System_Runtime_dll_3402885C_len;
extern const unsigned char System_Private_Uri_dll_4BA7B4A5[];
extern const int System_Private_Uri_dll_4BA7B4A5_len;

const unsigned char* dotnet_wasi_getbundledfile(const char* name, int* out_length) {
  return NULL;
}

void dotnet_wasi_registerbundledassemblies() {
  mono_wasm_add_assembly ("hello.dll", hello_dll_F832AC0D, hello_dll_F832AC0D_len);
  mono_wasm_add_assembly ("System.Console.dll", System_Console_dll_0566C3A3, System_Console_dll_0566C3A3_len);
  mono_wasm_add_assembly ("System.Memory.dll", System_Memory_dll_7E7749D6, System_Memory_dll_7E7749D6_len);
  mono_wasm_add_assembly ("System.Private.CoreLib.dll", System_Private_CoreLib_dll_F5990610, System_Private_CoreLib_dll_F5990610_len);
  mono_wasm_add_assembly ("System.Private.Runtime.InteropServices.JavaScript.dll", System_Private_Runtime_InteropServices_JavaScript_dll_1CFCEC65, System_Private_Runtime_InteropServices_JavaScript_dll_1CFCEC65_len);
  mono_wasm_add_assembly ("System.Threading.dll", System_Threading_dll_D30B628A, System_Threading_dll_D30B628A_len);
  mono_wasm_add_assembly ("System.Collections.dll", System_Collections_dll_C69F445F, System_Collections_dll_C69F445F_len);
  mono_wasm_add_assembly ("System.Runtime.InteropServices.dll", System_Runtime_InteropServices_dll_BC50FF4B, System_Runtime_InteropServices_dll_BC50FF4B_len);
  mono_wasm_add_assembly ("System.Runtime.dll", System_Runtime_dll_3402885C, System_Runtime_dll_3402885C_len);
  mono_wasm_add_assembly ("System.Private.Uri.dll", System_Private_Uri_dll_4BA7B4A5, System_Private_Uri_dll_4BA7B4A5_len);
}

