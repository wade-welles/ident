package main

type Counter struct {
	Cluster string `json:"cluster"`
	Seq     uint16 `json:"seq"`
}
