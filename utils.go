package lazyfit

import (
	"fmt"
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

func getDailyFilterParam() (string,string) {
	giorno := currentDay.Weekday()

	if giorno == 0 { // used because weekday starts with sunday
		giorno = 7
	}

	nextMonday := currentDay.AddDate(0,0, 7 - int(giorno) + 1)

	inizio := fmt.Sprintf("%d-%d-%d", nextMonday.Year(), nextMonday.Month(), nextMonday.Day())

	return inizio, inizio
}

func getWeeklyFilterParam() (string, string) {
	giorno := currentDay.Weekday()

	if giorno == 0 { // used because weekday starts with sunday
		giorno = 7
	}

	nextMonday := currentDay.AddDate(0,0, 7 - int(giorno) + 1)
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
		if jsonName != ""  && jsonName != "-" {
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
			reqRules[key] = ""//value.Interface()
		}
	}

	return reqRules
}