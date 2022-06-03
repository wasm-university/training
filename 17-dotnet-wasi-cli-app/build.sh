#!/bin/bash
cd hello
dotnet add package Wasi.Sdk --prerelease
dotnet build
