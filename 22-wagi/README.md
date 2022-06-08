# Wagi

## Install

```bash
wget https://github.com/deislabs/wagi/releases/download/v0.8.1/wagi-v0.8.1-linux-amd64.tar.gz
tar -zxf wagi-v0.8.1-linux-amd64.tar.gz 
sudo cp wagi /usr/local/bin/wagi
```

## Demo

```bash
wagi -c modules.toml 
```

> GET
```bash
curl http://localhost:3000/hello
```

> GET query string
```bash
curl http://localhost:3000/hey?message=hello
```

> POST
```bash
curl -d "👋 hello people 😃" http://localhost:3000/hey 
```