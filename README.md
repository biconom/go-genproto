# go-genproto

```bash
export BICONOM_API_PATH="$HOME/go/src/github.com/biconom/apis"

protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/account.proto=biconom/type/account biconom/type/account.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/agent.proto=biconom/type/agent biconom/type/agent.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/condition.proto=biconom/type/condition biconom/type/condition.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/confirmation.proto=biconom/type/confirmation biconom/type/confirmation.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/contact.proto=biconom/type/contact biconom/type/contact.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/currency.proto=biconom/type/currency biconom/type/currency.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/device.proto=biconom/type/device biconom/type/device.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/limit.proto=biconom/type/limit biconom/type/limit.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/profile.proto=biconom/type/profile biconom/type/profile.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/rank.proto=biconom/type/rank biconom/type/rank.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/session.proto=biconom/type/session biconom/type/session.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/sort.proto=biconom/type/sort biconom/type/sort.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/staking.proto=biconom/type/staking biconom/type/staking.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/ticket.proto=biconom/type/ticket biconom/type/ticket.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/until.proto=biconom/type/until biconom/type/until.proto

protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/currency/v1/currency.proto

protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/bot/currency/v1/currency.proto

protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/auth/v1/auth.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/confirmation/v1/confirmation.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/currency/v1/currency.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/session/v1/session.proto
```

biconom.admin.currency.v1;