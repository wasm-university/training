#!/bin/bash
cd hello
dotnet add package Wasi.Sdk --prerelease
dotnet build
ls -lh bin/Debug/net7.0/*.wasm
