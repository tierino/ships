package types

type AISMessageType string

const (
	PositionReportLabel               AISMessageType = "PositionReport"
	StaticDataReportLabel             AISMessageType = "StaticDataReport"
	ShipStaticDataLabel               AISMessageType = "ShipStaticData"
	AidsToNavigationReportLabel       AISMessageType = "AidsToNavigationReport"
	StandardClassBPositionReportLabel AISMessageType = "StandardClassBPositionReport"
)

type PositionReportEvent struct {
	MMSI       uint32  `json:"mmsi"`
	Status     uint8   `json:"status"`
	StatusText string  `json:"status_text"`
	Speed      float64 `json:"speed"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lon"`
	Course     float64 `json:"course"`
	Heading    uint16  `json:"heading"`
	Turn       uint16  `json:"turn"`
	Accuracy   bool    `json:"accuracy"`
}

func (p *PositionReportEvent) GetType() AISMessageType {
	return PositionReportLabel
}

type ShipStaticDataEvent struct {
	MMSI         uint32 `json:"mmsi"`
	ShipName     string `json:"ship_name"`
	ShipType     uint8  `json:"ship_type"`
	ShipTypeText string `json:"ship_type_text"`
	CallSign     string `json:"call_sign"`
	ToBow        uint16 `json:"to_bow"`
	ToStern      uint16 `json:"to_stern"`
	ToPort       uint16 `json:"to_port"`
	ToStarboard  uint16 `json:"to_starboard"`
	ETA          string `json:"eta"`
	Destination  string `json:"dest"`
}

func (p *ShipStaticDataEvent) GetType() AISMessageType {
	return ShipStaticDataLabel
}
