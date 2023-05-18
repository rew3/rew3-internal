package date

import "time"

type DateTime struct {
	time.Time
}

const dateTimeFormat = "2006-01-02T15:04:05.000Z"

func (d *DateTime) UnmarshalJSON(b []byte) error {
	// Parse the JSON string into a time.Time value
	parsedTime, err := time.Parse(`"`+dateTimeFormat+`"`, string(b))
	if err != nil {
		return err
	}

	// Set the value of the DateTime field to the parsed time
	d.Time = parsedTime
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	// Format the DateTime value as a JSON string
	formattedTime := "\"" + d.Time.Format(dateTimeFormat) + "\""
	return []byte(formattedTime), nil
}
