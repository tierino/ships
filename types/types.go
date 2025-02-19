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
	// todo
}
