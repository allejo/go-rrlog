package main

import (
	"bytes"
	"encoding/binary"
)

type TeamUpdateData struct {
	count uint8

	Type string `json:"type"`
	Teams []TeamData `json:"teams"`
}

type TeamData struct {
	Team uint16	`json:"team"`
	Size uint16 `json:"size"`
	Wins uint16 `json:"wins"`
	Losses uint16 `json:"losses"`
}

func handleMsgTeamUpdate(data []byte) (unpacked TeamUpdateData) {
	buf := bytes.NewBuffer(data)

	unpacked.Type = "MsgTeamUpdate"

	binary.Read(buf, binary.BigEndian, &unpacked.count)

	var i uint8
	for i = 0; i < unpacked.count; i++ {
		var data TeamData

		binary.Read(buf, binary.BigEndian, &data.Team)
		binary.Read(buf, binary.BigEndian, &data.Size)
		binary.Read(buf, binary.BigEndian, &data.Wins)
		binary.Read(buf, binary.BigEndian, &data.Losses)

		unpacked.Teams = append(unpacked.Teams, data)
	}

	return
}