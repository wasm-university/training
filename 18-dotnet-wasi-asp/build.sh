#!/bin/bash
cd hello
dotnet add package Wasi.Sdk --prerelease
dotnet add package Wasi.AspNetCore.Server.Native --prerelease
dotnet build
ls -lh bin/Debug/net7.0/*.wasm
