# exchange-rate

This app that provides currency exchange services

# start the app
```
./app
```

# API
```
curl "http://localhost:8080/apis/v1/exchange-rates?source=USD&target=JPY&amount=\$1,525"
```

# Test
```
go test ./...
```

# Set Config.json
* **Environment Variable Configuration (Preferred)**

    Configuration settings are provided through an environment variable named `CONFIG` (prioritized).

* **Command Line Parameter Configuration**

    Configuration settings can also be supplied during runtime using the `-config`` command line parameter.

# Configuration
* config.json - Configuration for host and rate.json file location
* rate.json - Exchange rate information