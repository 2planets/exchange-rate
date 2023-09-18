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

# configuration
* config.json - Configuration for host and rate.json file location
* rate.json - Exchange rate information