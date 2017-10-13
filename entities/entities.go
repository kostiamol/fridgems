package entities

import "time"

type Server struct {
	Host string
	Port string
}

type DevMeta struct {
	Type string `json:"type"`
	Name string `json:"name"`
	MAC  string `json:"mac"`
}

type FridgeData struct {
	TempCam1 map[int64]float32
	TempCam2 map[int64]float32
}

type FridgeRequest struct {
	Time   int64      `json:"time"`
	Meta   DevMeta   `json:"meta"`
	Data   FridgeData `json:"data"`
}

type FridgeConfig struct {
	TurnedOn    bool   `json:"turnedOn"`
	CollectFreq int64  `json:"collectFreq"`
	SendFreq    int64  `json:"sendFreq"`
}

func (fc *FridgeConfig) IsEmpty() bool {
	if fc.CollectFreq == 0 && fc.SendFreq == 0 && fc.TurnedOn == false {
		return true
	}
	return false
}

type FridgeGenerData struct {
	Time int64
	Data float32
}

type CollectFridgeData struct {
	CBot chan FridgeGenerData
	CTop chan FridgeGenerData
	ReqChan chan FridgeRequest
}

type RoutinesController struct {
	StopChan chan struct{}
}

func (c *RoutinesController) Wait() {
	<-c.StopChan
	<-time.NewTimer(time.Second * 3).C
}

func (c *RoutinesController) Terminate() {
	select {
	case <-c.StopChan:
	default:
		close(c.StopChan)
	}
}
