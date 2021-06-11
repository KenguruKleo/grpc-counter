module main.go

go 1.16

replace services/counter => ./models/protos/services/counter

require (
	google.golang.org/grpc v1.38.0
	services/counter v0.0.0-00010101000000-000000000000
)
