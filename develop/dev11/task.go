package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

// Кастомный тип для корректного парсинга времени
type eventDate time.Time

// Способ анмаршалинга строки во время
func (d *eventDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = eventDate(t)
	return nil
}

// Способ маршалинга строки во время
func (d eventDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d))
}

// Событие
type event struct {
	EventID     int       `json:"-"`
	UserID      int       `json:"user_id"`
	Date        eventDate `json:"date"`
	Description string    `json:"description"`
}

// Хранилище
type storage struct {
	sync.RWMutex

	Events map[int]event
}

// GET ответ
type getResponse struct {
	Result []event `json:"result"`
}

// Middleware для логирования
func logging(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

// Создать событие
func (s *storage) createHandler(w http.ResponseWriter, r *http.Request) {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	var newEvent event
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newEvent)
	if err != nil {
		returnError(w, err.Error(), 400)
		return
	}

	newEvent.EventID = len(s.Events) + 1
	s.Events[newEvent.EventID] = newEvent

	returnOk(w, "Запись успешно создана")
}

// Обновить событие
func (s *storage) updateHandler(w http.ResponseWriter, r *http.Request) {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	type UpdateBody struct {
		EventID     int       `json:"event_id"`
		UserID      int       `json:"user_id,omitempty"`
		Date        eventDate `json:"date,omitempty"`
		Description string    `json:"description,omitempty"`
	}

	var updateInfo UpdateBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateInfo)
	if err != nil {
		returnError(w, err.Error(), 400)
		return
	}

	event, ok := s.Events[updateInfo.EventID]
	if !ok {
		returnError(w, err.Error(), 400)
		return
	}
	if event.Date != updateInfo.Date {
		event.Date = updateInfo.Date
	}
	if event.Description != updateInfo.Description {
		event.Description = updateInfo.Description
	}
	s.Events[updateInfo.EventID] = event

	returnOk(w, "Запись успешно обновлена")
}

// Удалить событие
func (s *storage) deleteHandler(w http.ResponseWriter, r *http.Request) {
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	type DeleteBody struct {
		EventID int `json:"event_id"`
	}

	var deleteInfo DeleteBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deleteInfo)
	if err != nil {
		returnError(w, err.Error(), 400)
		return
	}

	delete(s.Events, deleteInfo.EventID)

	returnOk(w, "Запись успешно удалена")
}

// Выдать все события
func (s *storage) allEventsHandler(w http.ResponseWriter, r *http.Request) {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	res := getResponse{}
	for _, e := range s.Events {
		res.Result = append(res.Result, e)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}

// Выдать события на указанный день
func (s *storage) forDayHandler(w http.ResponseWriter, r *http.Request) {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	userID := r.FormValue("user_id")
	initDate := r.FormValue("date")
	date, _ := time.Parse("2006-01-02", initDate)
	log.Print(initDate)
	log.Print(date)

	res := getResponse{}
	for _, e := range s.Events {
		if fmt.Sprint(e.UserID) == userID && time.Time(e.Date) == date {
			res.Result = append(res.Result, e)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}

// Выдать события на указанный день + 7 дней
func (s *storage) forWeekHandler(w http.ResponseWriter, r *http.Request) {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	userID := r.FormValue("user_id")
	initDate := r.FormValue("date")
	date, err := time.Parse("2006-01-02", initDate)
	if err != nil {
		returnError(w, err.Error(), 400)
		return
	}

	res := getResponse{}
	for _, e := range s.Events {
		subTime := int(date.AddDate(0, 0, 7).Sub(time.Time(e.Date)).Hours()) / 24
		if fmt.Sprint(e.UserID) == userID && subTime >= 0 && subTime <= 7 {
			res.Result = append(res.Result, e)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}

// Выдать события в указанном месяце
func (s *storage) forMonthHandler(w http.ResponseWriter, r *http.Request) {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	userID := r.FormValue("user_id")
	initDate := r.FormValue("date")
	date, err := time.Parse("2006-01-02", initDate)
	if err != nil {
		returnError(w, err.Error(), 400)
		return
	}

	res := getResponse{}
	for _, e := range s.Events {
		if fmt.Sprint(e.UserID) == userID && time.Time(e.Date).Month() == date.Month() {
			res.Result = append(res.Result, e)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}

// Возврат ошибки
func returnError(w http.ResponseWriter, err string, status int) {
	log.Printf("%s: %s", fmt.Sprint(status), err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err == "" {
		fmt.Fprintf(w, `{"error": "%s"}`, http.StatusText(status))
		return
	}
	fmt.Fprintf(w, `{"error": "%s: %s"}`, http.StatusText(status), err)
}

// Возврат 200
func returnOk(w http.ResponseWriter, result string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, `{"result": %s}`, result)
}

func main() {
	s := storage{Events: map[int]event{}}

	// Берем данные из config.ini
	inidata, err := ini.Load("config.ini")
	var ip string
	var port string
	if err != nil {
		log.Print("config.ini не найден. Сервер запускается на localhost:8080")
		ip = "localhost"
		port = ":8080"
	} else {
		section := inidata.Section("connection")
		ip = section.Key("ip").String()
		port = section.Key("port").String()
	}

	// Сервер
	router := http.NewServeMux()
	router.Handle("/create_event", logging(s.createHandler))
	router.Handle("/update_event", logging(s.updateHandler))
	router.Handle("/delete_event", logging(s.deleteHandler))
	router.Handle("/all_events", logging(s.allEventsHandler))
	router.Handle("/events_for_day", logging(s.forDayHandler))
	router.Handle("/events_for_week", logging(s.forWeekHandler))
	router.Handle("/events_for_month", logging(s.forMonthHandler))
	http.Handle("/", router)
	err = http.ListenAndServe(ip+":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
