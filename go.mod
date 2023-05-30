module github.com/barebitcoin/btc-buf

go 1.18

replace github.com/btcsuite/btcd => github.com/barebitcoin/btcd v0.23.5-0.20230530095723-ea49e0152fc5

require (
	github.com/btcsuite/btcd v0.23.4
	github.com/btcsuite/btcd/btcutil v1.1.3
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/jessevdk/go-flags v1.5.0
	github.com/rs/zerolog v1.26.1
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd // indirect
	github.com/btcsuite/websocket v0.0.0-20150119174127-31079b680792 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)
