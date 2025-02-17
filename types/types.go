package types

import "github.com/BertoldVdb/go-ais"

type AISMessageType string

const (
	PositionReportType           AISMessageType = "PositionReport"
	StaticDataReportType         AISMessageType = "StaticDataReport"
	ShipStaticDataType           AISMessageType = "ShipStaticData"
	AidsToNavigationReport       AISMessageType = "AidsToNavigationReport"
	StandardClassBPositionReport AISMessageType = "StandardClassBPositionReport"
)

type AISMessage struct {
	Type                         AISMessageType
	PositionReport               *ais.PositionReport
	StaticDataReport             *ais.StaticDataReport
	ShipStaticData               *ais.ShipStaticData
	AidsToNavigationReport       *ais.AidsToNavigationReport
	StandardClassBPositionReport *ais.StandardClassBPositionReport
}

type PositionReportEvent struct {
	MMSI         uint32
	ShipType     uint8
	ShipTypeText string
	Status       uint8
	StatusText   string
	Speed        float64
	Latitude     float64
	Longitude    float64
	Course       float64
	Heading      uint16
	Turn         uint16
	Accuracy     bool
}

type StaticDataReportEvent struct {
	MMSI uint32
}
