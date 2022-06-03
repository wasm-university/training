using System.Runtime.InteropServices;

Console.WriteLine("👋 Hello, World! 🌍");
Console.WriteLine($"OS Architecture: {RuntimeInformation.OSArchitecture}");
Console.WriteLine($"Current time (UTC): {DateTime.UtcNow.ToLongTimeString()}");