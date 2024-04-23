# OKEX API V5 

include RESTFul and Websocket API 

# Environments

```shell
export OK_ACCESS_KEY={your-access-key}
export OK_ACCESS_SECRET={your-access-secret}
export OK_ACCESS_PASSPHRASE={your-passphrase}
```

# Balance

```shell
make && ./okex balance
```

# Buy Co-currency

## Buy with market price

```shell
# default use USDT
make && ./okex buy PEPE 2000000
```

## Buy with limit price

```shell
# default use USDT
make && ./okex buy --order-type limit --price 0.00000693 PEPE 2000000
```

# Sell Co-currency

## Sell with market order

```shell
# default use USDT
make && ./okex sell PEPE 2000000
```

## Sell with limit order

```shell
# default use USDT
make && ./okex sell --order-type limit --price 0.0000073 PEPE 2000000
```

