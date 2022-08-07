protoc -I ../internal/transport/grpc/proto/ \
auth.proto \
area.proto  \
--go-grpc_out=../internal/transport --go_out=../internal/transport