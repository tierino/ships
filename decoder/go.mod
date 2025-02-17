module example.com/decoder

go 1.23.6

replace example.com/types => ../types

require (
	example.com/types v0.0.0-00010101000000-000000000000
	github.com/BertoldVdb/go-ais v0.4.0
	github.com/adrianmo/go-nmea v1.10.0
)
