# eth-tx-parser

## Introduction
This is a simple parser with which you can use to subscribe ETH addresses and poll for transaction new information.

## Usage
Compile and start up application:
`go build; ./eth-tx-parser`

This will start a poller for all subscribed addresses

## Future work
Currently subscribed addresses are hard-coded. It would be nice to have an endpoint with which to subscribe addresses