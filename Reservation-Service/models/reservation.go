package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	Reservations map[string]*Reservation
)

type Reservation struct {
	ReservationId   string
	Score      int64
	PlayerName string
}

func init() {
	Reservations = make(map[string]*Reservation)
	Reservations["hjkhsbnmn123"] = &Reservation{"hjkhsbnmn123", 100, "astaxie"}
	Reservations["mjjkxsxsaa23"] = &Reservation{"mjjkxsxsaa23", 101, "someone"}
}

func AddOne(Reservation Reservation) (ReservationId string) {
	Reservation.ReservationId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Reservations[Reservation.ReservationId] = &Reservation
	return Reservation.ReservationId
}

func GetOne(ReservationId string) (Reservation *Reservation, err error) {
	if v, ok := Reservations[ReservationId]; ok {
		return v, nil
	}
	return nil, errors.New("ReservationId Not Exist")
}

func GetAll() map[string]*Reservation {
	return Reservations
}

func Update(ReservationId string, Score int64) (err error) {
	if v, ok := Reservations[ReservationId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ReservationId Not Exist")
}

func Delete(ReservationId string) {
	delete(Reservations, ReservationId)
}

