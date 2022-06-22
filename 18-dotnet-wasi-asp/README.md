# How to

> create the application
```bash
dotnet new web -o hello
```

> change the content of `hello/hello.csproj``
```xml
<Project Sdk="Microsoft.NET.Sdk.Web">

  <PropertyGroup>
    <TargetFramework>net7.0</TargetFramework>
    <Nullable>enable</Nullable>
    <ImplicitUsings>enable</ImplicitUsings>
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="Wasi.AspNetCore.Server.Native" Version="0.1.1" />
    <PackageReference Include="Wasi.Sdk" Version="0.1.1" />
  </ItemGroup>

</Project>
```

> change the source code of `hello/Program.cs`
```csharp
using System.Runtime.InteropServices;

//var builder = WebApplication.CreateBuilder(args);
var builder = WebApplication.CreateBuilder(args).UseWasiConnectionListener();

var app = builder.Build();

app.MapGet("/", () => $"👋 Hello, World! 🌍 🖥️: {RuntimeInformation.OSArchitecture} ⏳: {DateTime.UtcNow.ToLongTimeString()} (UTC)");

app.Run();
```

> Change `Properties/launchSettings.json`
```json
{
  "profiles": {
    "AspNetCoreOnCustomHost": {
      "commandName": "Project",
      "launchBrowser": true,
      "environmentVariables": {
        "ASPNETCORE_ENVIRONMENT": "Development"
      },
      "applicationUrl": "http://localhost:8080"
    }
  }
}
```



> build
```bash
cd hello
dotnet add package Wasi.Sdk --prerelease
dotnet add package Wasi.AspNetCore.Server.Native --prerelease
dotnet build
```

> run
```bash
cd hello
wasmtime bin/Debug/net7.0/hello.wasm --tcplisten localhost:8080
```

