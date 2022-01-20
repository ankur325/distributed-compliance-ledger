# Run Light Client Proxy
Light Client Proxy can be used if there are no trusted Full Nodes (Validator or Observers) a client can connect to.

It can be a proxy for CLI or direct requests from code done via Tendermint RPC.

Please note, that CLI can use a Light Client proxy only for single-value query requests.
A Full Node (Validator or Observer) should be used for multi-value query requests and write requests.

See the following links for details about a Light Client:
- https://docs.tendermint.com/master/tendermint-core/light-client.html
- https://docs.tendermint.com/master/nodes/light-client.html

## Running Light Client Proxy
See https://docs.tendermint.com/master/nodes/light-client.html for details

### 1. Choose Semi-trusted or Non-trusted Nodes for Connection
Light Client must be connected to one Primary Full Node (Validator or Observer) and
a number of Witness Full Nodes (Validators or Observers).

### 2. Obtain Trusted Height & Hash
One way to obtain a semi-trusted hash & height is to query multiple full nodes and compare their hashes.

For the first node (for example Primary Node):
```bash
curl -s http(s)://<node host>:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```
Result Example:
```
{
  "height": "<height>",
  "hash": "<hash>"
}
```

For other nodes (for example Witness Nodes):
```bash
curl -s http(s)://<node host>:26657/commit?height=<height> | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```
Where `<height>` obtained from the first run.

If result is the same on all nodes, then obtained `<height>` and `<hash>` can be used as parameters for the light client at the next step.

### 3. Start a Light Client Proxy Server
```
dcld light <chain-id> -p tcp://<primary-host>:26657 -w tcp://<witness1-host>:26657,tcp://<witness2-host>:26657 --height <height> --hash <hash>
```
where `<height>` and `<hash>` obtained at the previous step.

Light Client Proxy is started at port `8888` by default. It can be changed with `--laddr` argument.

Light Client Proxy will write some information (about headers) to disk. Default location is `~/.tendermint-light`.
It can be changed with `--dir` argument.

### 4. Connect CLI to a Light Client Proxy Server
```
dcld config chain-id <chain-id>
dcld config node tcp://<light-client-proxy-host>:8888
```