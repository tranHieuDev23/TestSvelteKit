package rpcclient

//go:generate mockgen -source=pjrpc_client_service.go -destination=pjrpc_client_service.go.mock.go -package=$GOPACKAGE -build_flags="-tags=mockgen"
