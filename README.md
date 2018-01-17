# tether

#
Get a Trillian log running:

```bash
go get github.com/google/trillian
cd !$
go build ./server/trillian_log_server
go build ./server/trillian_log_signer --http_endpoint=localhost:8093

# In one terminal:
./trillian_log_server --logtostderr ...

# In another terminal:
./trillian_log_signer --logtostderr --force_master --http_endpoint=localhost:8092 --batch_size=1000 --sequencer_guard_window=0 --sequencer_interval=200ms
```

Create a Log in Trillian:
```bash
go build ./cmd/createtree/
./createtree --admin_server=localhost:8090
<LOGID printed here>
```

Build and run geth.
Note that it will take time for geth to actually start downloading blocks, watch for status updates on its console.

```bash
# In yet another terminal:
make geth
build/bin/geth --cache=512 --verbosity 3 --rpc --fast console

```

Build and run the tether Follower:

```bash
# Yes, another terminal:
go run ./cmd/follower/main.go --geth=http://127.0.0.1:8545 --trillian_log=localhost:8090 --log_id LOGID --logtostderr
```

Watch as your diskspace gets eaten.


