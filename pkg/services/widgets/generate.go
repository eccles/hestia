package widgetsService

// Generate boilerplate for grpcserver.go and service.go
// First argument value=1 enables GRPC
//go:generate ../../../scripts/grpcserver.sh 1 widgets

//go:generate mockery --all --dry-run=false
