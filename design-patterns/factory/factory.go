package main

import "fmt"

// enums
type storageType string

const (
	DiskStorage     storageType = "disk_storage"
	InMemoryStorage storageType = "inmemory_storage"
	RedisStorage    storageType = "redis_storage"
	TempStorage     storageType = "temp_storage"
)

type Storage interface {
	getStorage() storageType
}

type Disk struct{}

func (d *Disk) getStorage() storageType {
	return DiskStorage
}

type InMemory struct{}

func (d *InMemory) getStorage() storageType {
	return InMemoryStorage
}

type Redis struct{}

func (d *Redis) getStorage() storageType {
	return RedisStorage
}

type Temp struct{}

func (d *Temp) getStorage() storageType {
	return TempStorage
}
func NewStorage(st storageType) Storage {
	switch st {
	case DiskStorage:
		return &Disk{}
	case InMemoryStorage:
		return &InMemory{}
	case RedisStorage:
		return &Redis{}
	default:
		return &Temp{}
	}
}

func main() {
	var st storageType = TempStorage
	var storage Storage = NewStorage(st)
	fmt.Println(storage.getStorage())
}
