package deltaexchapigo

import (
	"fmt"
	"reflect"
	"time"

	"github.com/golang/glog"
)

type Time struct {
	time.Time
}

// API endpoints
const (
	URIAssets        string = "/v2/assets"
	URIHistoryCandle string = "/v2/history/candles"
)

func structToMap(obj interface{}, tagName string) map[string]interface{} {
	var values reflect.Value
	switch obj.(type) {
	case TSPriceParam:
		{
			con := obj.(TSPriceParam)
			values = reflect.ValueOf(&con).Elem()
		}

	}

	tags := reflect.TypeOf(obj)
	params := make(map[string]interface{})
	for i := 0; i < values.NumField(); i++ {
		params[tags.Field(i).Tag.Get(tagName)] = values.Field(i).Interface()
	}

	return params
}

func GetTime(timeString string) string {
	layout := "02-01-2006 15:04:05"
	time_location, _ := time.LoadLocation("Asia/Kolkata")
	t, err := time.ParseInLocation(layout, timeString, time_location)
	// t, err := time.Parse(layout, timeString)
	if err != nil {
		glog.Fatal(err)
	}
	return fmt.Sprintf("%d", t.Unix())
}

func GetDate(timestampStr string) (string, error) {
	// Define the layout of the timestamp string
	layout := "02-01-2006"

	// Parse the timestamp string into a time.Time object
	timestamp, err := time.Parse(layout, timestampStr)
	if err != nil {
		return "", fmt.Errorf("error parsing timestamp: %v", err)
	}

	// Convert the time.Time object to epoch time (in seconds)
	epochTime := timestamp.Unix()

	// Convert the epoch time to a string
	epochTimeStr := fmt.Sprintf("%d", epochTime)

	return epochTimeStr, nil
}

func GetTodayAndLastWeekEpoch() (int64, int64) {
	// Get today's date
	today := time.Now()

	// Subtract 7 days to get the date from one week ago
	lastWeek := today.AddDate(0, 0, -7)

	// Convert dates to epoch time (in seconds)
	todayEpoch := today.Unix()
	lastWeekEpoch := lastWeek.Unix()

	return todayEpoch, lastWeekEpoch
}
