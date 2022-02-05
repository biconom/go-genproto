# go-genproto

```bash
export BICONOM_API_PATH="$HOME/go/src/github.com/biconom/apis"

mkdir -p ./type/agent
protoc -I $BICONOM_API_PATH -I $BICONOM_API_PATH/biconom/type --go_out ./type/agent --go_opt paths=source_relative agent.proto

mkdir -p ./type/confirmation
protoc -I $BICONOM_API_PATH -I $BICONOM_API_PATH/biconom/type --go_out ./type/confirmation --go_opt paths=source_relative confirmation.proto

mkdir -p ./type/contact
protoc -I $BICONOM_API_PATH -I $BICONOM_API_PATH/biconom/type --go_out ./type/contact --go_opt paths=source_relative contact.proto

mkdir -p ./type/device
protoc -I $BICONOM_API_PATH -I $BICONOM_API_PATH/biconom/type --go_out ./type/device --go_opt paths=source_relative device.proto

mkdir -p ./type/filter
protoc -I $BICONOM_API_PATH -I $BICONOM_API_PATH/biconom/type --go_out ./type/filter --go_opt paths=source_relative filter.proto

mkdir -p ./type/profile
protoc -I $BICONOM_API_PATH -I $BICONOM_API_PATH/biconom/type --go_out ./type/profile --go_opt paths=source_relative profile.proto

mkdir -p ./type/session
protoc -I $BICONOM_API_PATH -I $BICONOM_API_PATH/biconom/type --go_out ./type/session --go_opt paths=source_relative session.proto

mkdir -p ./client/auth/v1
protoc -I $BICONOM_API_PATH -I $BICONOM_API_PATH/biconom --go_out ./ --go_opt paths=source_relative --go-grpc_out ./ --go-grpc_opt paths=source_relative client/auth/v1/auth.proto

mkdir -p ./client/confirmation/v1
protoc -I $BICONOM_API_PATH -I $BICONOM_API_PATH/biconom --go_out ./ --go_opt paths=source_relative --go-grpc_out ./ --go-grpc_opt paths=source_relative client/confirmation/v1/confirmation.proto

mkdir -p ./client/session/v1
protoc -I $BICONOM_API_PATH -I $BICONOM_API_PATH/biconom --go_out ./ --go_opt paths=source_relative --go-grpc_out ./ --go-grpc_opt paths=source_relative client/session/v1/session.proto
```

