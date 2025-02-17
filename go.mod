module example.com/ships

go 1.23.6

require (
	example.com/decoder v0.0.0-00010101000000-000000000000
	example.com/mapper v0.0.0-00010101000000-000000000000
	example.com/producer v0.0.0-00010101000000-000000000000
	example.com/tcpclient v0.0.0-00010101000000-000000000000
	example.com/types v0.0.0-00010101000000-000000000000
	github.com/BertoldVdb/go-ais v0.4.0
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/adrianmo/go-nmea v1.10.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace example.com/types => ./types

replace example.com/producer => ./producer

replace example.com/tcpclient => ./tcpclient

replace example.com/mapper => ./mapper

replace example.com/decoder => ./decoder
