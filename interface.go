package main

type Storage interface {
	Write() error
	Read() error
}

type SQLStorage interface {
	Write() error
	WriteData() error
	Read() error
}
