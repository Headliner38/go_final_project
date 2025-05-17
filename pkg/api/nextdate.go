package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func afterNow(date, now time.Time) bool {
	return date.Truncate(24 * time.Hour).After(now.Truncate(24 * time.Hour))
}

// now — время, от которого ищется ближайшая дата;
// dstart — исходное время в формате 20060102, от которого начинается отсчёт повторений;
// repeat — правило повторения в описанном выше формате
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if len(repeat) == 0 {
		return "", fmt.Errorf("некорректный параметр repeat: длина %d", len(repeat))
	}

	parsedTime, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("ошибка выполнения time.Parse")
	}
	repeatParts := strings.Split(repeat, " ")
	if repeatParts[0] != `d` && repeatParts[0] != `y` {
		return "", fmt.Errorf("неподдерживаемый формат repeat")
	}

	switch repeatParts[0] {
	case "y":
		for {
			parsedTime = parsedTime.AddDate(1, 0, 0)
			if afterNow(parsedTime, now) {
				break
			}
		}
		return parsedTime.Format(dateFormat), nil

	case "d":
		if len(repeatParts) < 2 {
			return "", fmt.Errorf("не указан интервал дней")
		}
		interval, err := strconv.Atoi(repeatParts[1])
		if err != nil {
			return "", fmt.Errorf("интервал дней в repeat указан неверно")
		}
		if interval > 400 || interval <= 0 {
			return "", fmt.Errorf("интервал дней указан неверно, указанный интервал %d", interval)
		}
		for {
			parsedTime = parsedTime.AddDate(0, 0, interval)
			if afterNow(parsedTime, now) {
				break
			}
		}
		return parsedTime.Format(dateFormat), nil

	default:
		return "", fmt.Errorf("неподдерживаемый формат")
	}

}

func nextDayHandler(res http.ResponseWriter, req *http.Request) {
	nowS := req.FormValue("now")
	dateS := req.FormValue("date")
	repeatS := req.FormValue("repeat")

	var now time.Time
	var err error
	if len(nowS) == 0 {
		now = time.Now()
	} else {
		now, err = time.Parse(dateFormat, nowS)
		if err != nil {
			http.Error(res, fmt.Sprintf("неверный формат даты now: %q, ожидается формат %s", nowS, dateFormat), http.StatusBadRequest)
			return
		}
	}
	str, err := NextDate(now, dateS, repeatS)
	if err != nil {
		http.Error(res, "ошибка при выполнении NextDate", http.StatusBadRequest)
	}

	res.Write([]byte(str))

}
