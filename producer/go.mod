module example.com/producer

go 1.23.6

replace example.com/types => ../types

require example.com/types v0.0.0-00010101000000-000000000000

require github.com/BertoldVdb/go-ais v0.4.0 // indirect
