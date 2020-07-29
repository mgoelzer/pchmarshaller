Decodes Filecoin payment channel creation receipts to give the address of the payment channel.

## Using

If your payment channel create message receipt looks like this:

```json
{ "exit_code": 0, "ret": "gkMA8TdVAuEnBlturfm30tDuzwd59jqHcDfr", "gas_used": 8897444 }
```

then extract the payment channel address like this:

```bash
echo "gkMA8TdVAuEnBlturfm30tDuzwd59jqHcDfr" | openssl enc -d -a | ./pchmarshaller
```

Alternatively, if you know the signed message CID for the message that created the payment channel:

```bash
$ curl [...]
{"jsonrpc":"2.0","result":{"/":"bafy2bzaceakdfzn37uqqsosyu5nvg5jk2rwcor7zcthnqotn6vidxgacehnbk"},"id":999}
$ lotus state search-msg bafy2bzaceakdfzn37uqqsosyu5nvg5jk2rwcor7zcthnqotn6vidxgacehnbk | grep "Return" | sed -e 's/Return://' | ./pchmarshaller
```

## Build

`go build -o pchmarshaller main.go`

