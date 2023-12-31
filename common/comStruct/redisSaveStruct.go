package comStruct

import "encoding/json"

type TrainStationRelationRedis map[string]string

func (t TrainStationRelationRedis) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t TrainStationRelationRedis) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &t)
}
