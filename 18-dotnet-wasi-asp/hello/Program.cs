using System.Runtime.InteropServices;

//var builder = WebApplication.CreateBuilder(args);
var builder = WebApplication.CreateBuilder(args).UseWasiConnectionListener();

var app = builder.Build();

app.MapGet("/", () => $"👋 Hello, World! 🌍 🖥️: {RuntimeInformation.OSArchitecture} ⏳: {DateTime.UtcNow.ToLongTimeString()} (UTC)");

app.Run();