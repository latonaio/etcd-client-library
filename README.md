# Etcd Client Library

## Description
Etcdに対して接続、データの取得、書き込み、削除を行うためのGo言語で書かれたライブラリです。


## Usage
・Etcdに接続する

```go
// 接続先のEtcdのアドレスをリストに格納する。
address = []string{"localhost:12379"}
ec := EtcdClient{address: address}
cli, err := ec.NewConnection()

// 接続を終了する
err := cli.CloseConnection()
```

・データの取得

```go
// Keyに対してValueを取得する。
res, err := cli.GET_value("Key")

// 特定の範囲のKeyを取得する。
res, err := cli.GET_key_range("key1", "key2")
```

・データの書き込み

```go
// 特定のKey Valueの値を書き込む。
err := cli.PUT_single("key", "value")

// Maps（連想配列）に格納されたKey Valueを書き込む。
keyvalue := map[string]string{"key1": "value1", "key2": "value2"}
err := cli.PUT_multiple(keyvalue)
```

・データの削除

```go
// 特定のキーのデータを削除する。
err := cli.DELETE_key("key")

// 指定された範囲のKeyを削除する。
err := cli2.DELETE_key_range("key1", "key2")
```