package model

type PoiSearch struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Geohash   string  `json:"geohash"`
}
