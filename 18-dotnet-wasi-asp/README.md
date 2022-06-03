# How to

> create the application
```bash
dotnet new console -o hello
```

> change the content of `hello/hello.csproj``
```xml
<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <OutputType>Exe</OutputType>
    <TargetFramework>net7.0</TargetFramework>
    <ImplicitUsings>enable</ImplicitUsings>
    <Nullable>enable</Nullable>
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="Wasi.Sdk" Version="0.1.1" />
  </ItemGroup>

</Project>
```

> change the source code of `hello/Program.cs`
```csharp
using System.Runtime.InteropServices;

Console.WriteLine("👋 Hello, World! 🌍");
Console.WriteLine($"OS Architecture: {RuntimeInformation.OSArchitecture}");
Console.WriteLine($"Current time (UTC): {DateTime.UtcNow.ToLongTimeString()}");
```

> build
```bash
cd hello
dotnet add package Wasi.Sdk --prerelease
dotnet build
```

> rub
```bash
wasmtime bin/Debug/net7.0/hello.wasm
wasmer bin/Debug/net7.0/hello.wasm
```

