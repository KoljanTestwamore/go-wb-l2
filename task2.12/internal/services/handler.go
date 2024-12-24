package services

import "time"

type Handler interface {
	HandleCreateEvent(id uint, e time.Time) (uint, time.Time)
	HandleUpdateEvent(id uint, eventId uint, e time.Time)
	HandleDeleteEvent(id uint, eventId uint)
	HandleGetEventsForPeriod(id uint, e time.Time, period time.Duration) map[uint]time.Time
}

type CalendarHandler struct {
	// ключ - id пользователя
	// массив - ивенты пользователя
	// 1) по хорошему тут использовать uuid
	// 2) по хорошему тут использовать b-tree для дат
	// Так как это курс не по алгоритмам, я считаю, что
	// подобные допущения уместны.
	// P.S. В реальном проекте тут была бы бд, она
	// избавит от необходимости делать деревья
	data map[uint]map[uint]time.Time
}

func NewCalendarHandler() *CalendarHandler {
	return &CalendarHandler{
		data: make(map[uint]map[uint]time.Time),
	}
}

func (ch *CalendarHandler) HandleCreateEvent(id uint, e time.Time) (uint, time.Time) {
	//                   :-)
	eid := uint(len(ch.data[id]))
	ch.data[id][eid] = e

	return eid, e 
}

func (ch *CalendarHandler) HandleUpdateEvent(id uint, eventId uint, e time.Time) {
	ch.data[id][eventId] = e
}

func (ch *CalendarHandler) HandleDeleteEvent(id uint, eventId uint) {
	delete(ch.data[id], eventId)
}

func (ch *CalendarHandler) HandleGetEventsForPeriod(id uint, e time.Time, period time.Duration) (res map[uint]time.Time) {
	userEvents := ch.data[id]
	res = make(map[uint]time.Time)

	for eid, event := range userEvents {
		if event.After(event) && event.Before(e.Add(period)) {
			res[eid] = event
		}
	}

	return res
}