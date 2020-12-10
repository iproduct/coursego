package iot

import (
	"sync/atomic"
	"time"
)

type IotEventType int

const(
	Ping IotEventType = iota
	Distance
	Button
	Temperature
	Humidity
	Light
	Electricity
)

func (iet IotEventType) String() string {
	return [...]string{"Ping", "Distance", "Button", "Temperature"}[iet]
}

var nextID uint64 // holds next ID to be given

type IotEvent struct {
	ID             uint64 		`json:"id"`
	Type 		   IotEventType	`json:"type"`
	Timestamp      time.Time	`json:"time"`
	Readings[]     int			`json:"readings,omitempty"`
}

func NewEvent(id uint64, kind IotEventType, time time.Time, readings ...int) *IotEvent {
	ev := new(IotEvent)
	ev.ID, ev.Type, ev.Timestamp, ev.Readings = id, kind, time, readings
	return ev
}

func NewPingEvent() *IotEvent {
	return NewEvent(atomic.AddUint64(&nextID, 1), Ping, time.Now())
}

func NewDistanceEvent(distance int) *IotEvent {
	return NewEvent(atomic.AddUint64(&nextID, 1), Distance, time.Now(), distance)
}

func NewButtonEvent(state int) *IotEvent {
	return NewEvent(atomic.AddUint64(&nextID, 1), Button, time.Now(), state)
}

func NewTemperatureEvent(temperature int) *IotEvent {
	return NewEvent(atomic.AddUint64(&nextID, 1), Temperature, time.Now(), temperature)
}

func NewHumidityEvent(humidity int) *IotEvent {
	return NewEvent(atomic.AddUint64(&nextID, 1), Humidity, time.Now(), humidity)
}

func NewLightEvent(light int) *IotEvent {
	return NewEvent(atomic.AddUint64(&nextID, 1), Light, time.Now(), light)
}

func NewElectricityEvent(electricity int) *IotEvent {
	return NewEvent(atomic.AddUint64(&nextID, 1), Electricity, time.Now(), electricity)
}