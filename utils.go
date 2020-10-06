package lazyfit

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func findCourseIdFromName(name string) string {
	if name == "" {
		return ""
	}

	for k := range listCourses {
		if listCourses[k].nome == name {
			return listCourses[k].id
		}
	}

	return ""
}

func getDailyFilterParam() (string, string) {
	nextWeekDay := currentDay.AddDate(0, 0, 7)

	day := fmt.Sprintf("%d-%d-%d", nextWeekDay.Year(), nextWeekDay.Month(), nextWeekDay.Day())

	return day, day
}

func getWeeklyFilterParam() (string, string) {
	giorno := currentDay.Weekday()

	if giorno == 0 { // used because weekday starts with sunday
		giorno = 7
	}

	nextMonday := currentDay.AddDate(0, 0, 7-int(giorno)+1)
	nextSunday := nextMonday.AddDate(0, 0, 6)

	inizio := fmt.Sprintf("%d-%d-%d", nextMonday.Year(), nextMonday.Month(), nextMonday.Day())
	fine := fmt.Sprintf("%d-%d-%d", nextSunday.Year(), nextSunday.Month(), nextSunday.Day())

	return inizio, fine
}

//https://gist.github.com/johnlonganecker/0b1f857781a902a558f34f1b467d5df8
func ConvertStructToMapOfStrings(st interface{}) map[string]string {
	reqRules := make(map[string]string)

	v := reflect.ValueOf(st)
	t := reflect.TypeOf(st)

	for i := 0; i < v.NumField(); i++ {
		key := strings.ToLower(t.Field(i).Name)
		typ := v.FieldByName(t.Field(i).Name).Kind().String()
		structTag := t.Field(i).Tag.Get("json")
		jsonName := strings.TrimSpace(strings.Split(structTag, ",")[0])
		value := v.FieldByName(t.Field(i).Name)

		// if jsonName is not empty use it for the key
		if jsonName != "" && jsonName != "-" {
			key = jsonName
		}

		switch typ {
		case "string":
			if !(value.String() == "" && strings.Contains(structTag, "omitempty")) {
				reqRules[key] = value.String()
			}
		case "bool":
			if value.Interface().(bool) {
				reqRules[key] = "true"
			} else {
				reqRules[key] = "false"
			}
		default:
			reqRules[key] = "" //value.Interface()
		}
	}

	return reqRules
}

func SendHttpRequest(method, url string) []byte {
	var resp *http.Response
	var err error

	if method == "POST" {
		resp, err = http.Post(url, "", nil)
	} else {
		resp, err = http.Get(url)
	}

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return body
}
