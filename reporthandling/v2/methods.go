package v2

import "encoding/json"

func (postureReport *PostureReport) ToBytes() ([]byte, error) {
	return json.Marshal(*postureReport)
}

func (postureReport *PostureReport) ToString() string {
	if bpr, err := postureReport.ToBytes(); err == nil {
		return string(bpr)
	}

	return "{}"
}
