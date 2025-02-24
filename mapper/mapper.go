package mapper

import (
	"fmt"
	"strings"

	"github.com/BertoldVdb/go-ais"

	"example.com/types"
)

func TransformPositionReport(packet *ais.PositionReport) *types.PositionReportEvent {
	return &types.PositionReportEvent{
		MMSI:       packet.UserID,
		Status:     packet.NavigationalStatus,
		StatusText: "todo",
		Speed:      float64(packet.Sog),
		Latitude:   float64(packet.Latitude),
		Longitude:  float64(packet.Longitude),
		Course:     float64(packet.Cog),
		Heading:    packet.TrueHeading,
		Turn:       uint16(packet.RateOfTurn),
		Accuracy:   packet.PositionAccuracy,
	}
}

func TransformShipStaticData(packet *ais.ShipStaticData) *types.ShipStaticDataEvent {
	return &types.ShipStaticDataEvent{
		MMSI:         packet.UserID,
		ShipName:     strings.TrimSpace(packet.Name),
		ShipType:     packet.Type,
		ShipTypeText: "todo",
		CallSign:     strings.TrimSpace(packet.CallSign),
		ToBow:        packet.Dimension.A,
		ToStern:      packet.Dimension.B,
		ToPort:       uint16(packet.Dimension.C),
		ToStarboard:  uint16(packet.Dimension.D),
		ETA: fmt.Sprintf("%02d/%02dT%02d:%02d",
			packet.Eta.Month,
			packet.Eta.Day,
			packet.Eta.Hour,
			packet.Eta.Minute,
		),
		Destination: packet.Destination,
	}
}
