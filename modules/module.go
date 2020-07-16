package modules

type Key string

type Import Module
type Imports map[Key]Module

type Export Module
type Exports map[Key]Module

type Module interface {
	GetImport(key Key) Import
	GetExport(key Key) Export
	GetProvider(key Key) Provider
	SetImport(key Key, i Import)
	SetExport(key Key, i Import)
	SetProvider(key Key, i Import)
}

type

type Controllers map[Key]Controller

type Controller interface {

}

type Providers map[Key]Provider

type Provider interface {

}

type Repository interface {

}

