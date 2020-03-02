module github.com/uhppoted/uhppoted

go 1.14

require (
	github.com/uhppoted/uhppote-cli v0.0.0-20200228202133-c3f7228e2f9e
	github.com/uhppoted/uhppote-core v0.0.0-20200228192138-00c62a4d6ea3
	github.com/uhppoted/uhppote-simulator v0.0.0-20200302182046-56cd930d95f1
	github.com/uhppoted/uhppoted-api v0.0.0-20200302181311-56c5fea77afc
	github.com/uhppoted/uhppoted-mqtt v0.0.0-20200228205150-28235df3168b
	github.com/uhppoted/uhppoted-rest v0.0.0-20200302184005-f30d02a22101
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
)

replace (
	github.com/uhppoted/uhppote-cli => ./uhppote-cli
	github.com/uhppoted/uhppote-core => ./uhppote-core
	github.com/uhppoted/uhppote-simulator => ./uhppote-simulator
	github.com/uhppoted/uhppoted-api => ./uhppoted-api
	github.com/uhppoted/uhppoted-rest => ./uhppoted-rest
)
