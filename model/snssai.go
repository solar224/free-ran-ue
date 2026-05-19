package model

type SnssaiIE struct {
	Sst string `yaml:"sst" valid:"required"`
	Sd  string `yaml:"sd"`
}
