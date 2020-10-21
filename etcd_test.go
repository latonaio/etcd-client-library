package etcd

import (
	"log"
	"testing"
)

var (
	testAddress      = "localhost"
	testWrongAddress = "128.0.0.1"
	testPort         = 12379
	testWrongPort    = 123
	testkey          = "testkey"
	testvalue        = "testvalue"
	testkey1         = "testkey1"
	testvalue1       = "testvalue1"
	testkey2         = "testkey2"
	testvalue2       = "testvalue2"
	testkey3         = "testkey3"
	testvalue3       = "testvalue3"
	testkey4         = "testkey3"
	testvalue4       = "testvalue3"
)

func getCon() (*Etcd, error) {
	etcd := NewEtcd(testAddress, testPort)
	if err := etcd.Connection(); err != nil {
		log.Fatal(err)
	}
	return etcd, nil
}

func TestPutAndGet(t *testing.T) {
	key := "/test/devices/device-1"
	value := "test"

	etcd, _ := getCon()
	if err := etcd.Put(key, value); err != nil {
		t.Fatal(err)
	}

	resp, err := etcd.Get(key)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Key != key || resp.Value != value {
		t.Fatalf("not match key:%s, value:%s", resp.Key, resp.Value)
	}
}
