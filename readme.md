# Rspamd Statistics

## Example

### Generate keys

```shell
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 365 -subj '/CN=localhost'
```

### Make requests

```shell
echo '[1,3,9999922,22223332,1844674407370955162]' | json2msgpack | curl --data-binary @- -X POST -H "Content-Type: application/msgpack" -k -v https://127.0.0.1:8000/storage
```

Methods:

* POST - Find
* PUT - Add
* DELETE - Delete
