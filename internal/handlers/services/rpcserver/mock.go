package rpcserver

//go:generate mockgen -source=pjrpc_server_service.go -destination=pjrpc_server_service.go.mock.go -package=$GOPACKAGE -build_flags="-tags=mockgen"
