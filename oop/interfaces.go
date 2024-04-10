package main

import "fmt"

type Storage interface {
    Save(filename string, data []byte) error
}

type LocalDiskStorage struct{}

func (lds LocalDiskStorage) Save(filename string, data []byte) error {
    fmt.Printf("Saving data to local disk: %s\n", filename)
    // Implementation for saving to local disk
    return nil
}

type CloudStorage struct{}

func (cs CloudStorage) Save(filename string, data []byte) error {
    fmt.Printf("Saving data to cloud storage: %s\n", filename)
    // Implementation for saving to cloud storage
    return nil
}

func ProcessStorage(s Storage) error {
	
	s.Save("hi",[]byte("hu"))
	return nil
}

func main() {
    localDisk := LocalDiskStorage{}
    cloud := CloudStorage{}

    ProcessStorage(localDisk)
    ProcessStorage(cloud)
}