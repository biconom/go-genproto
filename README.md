# go-genproto

```bash
export BICONOM_API_PATH="$HOME/go/src/github.com/biconom/apis"

protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/account.proto=biconom/type/account biconom/type/account.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/activity.proto=biconom/type/activity biconom/type/activity.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/agent.proto=biconom/type/agent biconom/type/agent.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/bonus.proto=biconom/type/bonus biconom/type/bonus.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/bot_key.proto=biconom/type/bot_key biconom/type/bot_key.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/condition.proto=biconom/type/condition biconom/type/condition.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/confirmation.proto=biconom/type/confirmation biconom/type/confirmation.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/contact.proto=biconom/type/contact biconom/type/contact.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/currency.proto=biconom/type/currency biconom/type/currency.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/device.proto=biconom/type/device biconom/type/device.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/limit.proto=biconom/type/limit biconom/type/limit.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/network.proto=biconom/type/network biconom/type/network.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/payment_system.proto=biconom/type/payment_system biconom/type/payment_system.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/position.proto=biconom/type/position biconom/type/position.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/profile.proto=biconom/type/profile biconom/type/profile.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/quests.proto=biconom/type/quests biconom/type/quests.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/rank_system.proto=biconom/type/rank_system biconom/type/rank_system.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/session.proto=biconom/type/session biconom/type/session.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/sort.proto=biconom/type/sort biconom/type/sort.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/staking.proto=biconom/type/staking biconom/type/staking.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/ticket.proto=biconom/type/ticket biconom/type/ticket.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/tree.proto=biconom/type/tree biconom/type/tree.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/until.proto=biconom/type/until biconom/type/until.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/voucher.proto=biconom/type/voucher biconom/type/voucher.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=Mbiconom/type/wallet.proto=biconom/type/wallet biconom/type/wallet.proto


#auth
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/auth/v1/auth.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/auth/v1/auth.proto
#confirmation
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/confirmation/v1/confirmation_self.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/confirmation/v1/confirmation_self.proto
#currency
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/currency/v1/currency.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/currency/v1/currency_pair.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/currency/v1/currency_pair_rate_source.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/currency/v1/currency_pair_rate_source_bot_key.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/bot/currency/v1/currency.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/bot/currency/v1/currency_pair.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/bot/currency/v1/currency_pair_rate_source.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/bot/currency/v1/currency_pair_rate_source_bot_key.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/currency/v1/currency.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/currency/v1/currency_pair.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/currency/v1/currency_pair_rate_source.proto
#rank
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/rank/v1/rank.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/rank/v1/rank_account.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/rank/v1/rank_network_statistics.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/rank/v1/rank_statistics.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/rank/v1/rank_system.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/admin/rank/v1/rank_system_account.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/rank/v1/rank.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/rank/v1/rank_account.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/rank/v1/rank_network_statistics.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/rank/v1/rank_statistics.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/rank/v1/rank_system.proto
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/rank/v1/rank_system_account.proto
```

```bash
protoc -I $BICONOM_API_PATH --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative biconom/client/session/v1/session.proto
```
