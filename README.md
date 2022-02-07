# go-genproto

```bash
export BICONOM_API_PATH="$HOME/go/src/github.com/biconom/apis"

protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/agent.proto=biconom/type/agent biconom/type/agent.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/confirmation.proto=biconom/type/confirmation biconom/type/confirmation.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/contact.proto=biconom/type/contact biconom/type/contact.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/device.proto=biconom/type/device biconom/type/device.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/filter.proto=biconom/type/filter biconom/type/filter.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/profile.proto=biconom/type/profile biconom/type/profile.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/session.proto=biconom/type/session biconom/type/session.proto

protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/auth/v1/auth.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/confirmation/v1/confirmation.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/session/v1/session.proto
```
