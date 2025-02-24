module example.com/producer

go 1.23.6

replace example.com/types => ../types

require github.com/confluentinc/confluent-kafka-go/v2 v2.8.0
