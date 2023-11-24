package models

import (
	"errors"
	"fmt"
	"webhook/pkg/utils"
)

type Payload struct {
	Event      string           `json:"event"`
	EventType  string           `json:"event_type"`
	AppID      string           `json:"app_id"`
	UserID     string           `json:"user_id"`
	MessageID  string           `json:"message_id"`
	PageTitle  string           `json:"page_title"`
	PageURL    string           `json:"page_url"`
	Language   string           `json:"browser_language"`
	ScreenSize string           `json:"screen_size"`
	Attributes map[string]types `json:"attributes"`
	Traits     map[string]types `json:"traits"`
}

type types struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

func ValidateAndDecodePayload(payload map[string]any, attrLimit, uattrLimit uint) (Payload, error) {
	var data Payload

	// validate event
	if event, ok := payload["ev"]; ok {
		data.Event = utils.ConvertToString(event)
	} else {
		return data, errors.New("event is missing")
	}

	// validate event type
	if eventType, ok := payload["et"]; ok {
		data.EventType = utils.ConvertToString(eventType)
	} else {
		return data, errors.New("event type is missing")
	}

	//validate Application ID
	if appID, ok := payload["id"]; ok {
		data.AppID = utils.ConvertToString(appID)
	} else {
		return data, errors.New("application id is missing")
	}

	// validate User ID
	if userID, ok := payload["uid"]; ok {
		data.UserID = utils.ConvertToString(userID)
	} else {
		return data, errors.New("user id is missing")
	}

	// validate Message ID
	if msgID, ok := payload["mid"]; ok {
		data.MessageID = utils.ConvertToString(msgID)
	} else {
		return data, errors.New("message id is missing")
	}

	// validate Page Title
	if pageTitle, ok := payload["t"]; ok {
		data.PageTitle = utils.ConvertToString(pageTitle)
	} else {
		return data, errors.New("page title is missing")
	}

	// Validate page URL
	if pageURL, ok := payload["p"]; ok {
		data.PageURL = utils.ConvertToString(pageURL)
	} else {
		return data, errors.New("page url is missing")
	}

	// Validate browser language
	if lang, ok := payload["l"]; ok {
		data.Language = utils.ConvertToString(lang)
	} else {
		return data, errors.New("page url is missing")
	}

	// Validate Screen size
	if screenSize, ok := payload["sc"]; ok {
		data.ScreenSize = utils.ConvertToString(screenSize)
	} else {
		return data, errors.New("screen size is missing")
	}

	// Validate Attributes with predefined number of times or lesser
	data.Attributes = make(map[string]types)
	for i := 1; i <= int(attrLimit); i++ {
		var (
			key, val, valType             = "atrk", "atrv", "atrt"
			dataKey, dataVal, dataValType string
		)
		if atKey, ok := payload[fmt.Sprint(key, i)]; ok {
			dataKey = utils.ConvertToString(atKey)
		} else {
			if i > 1 {
				break
			}
			return data, fmt.Errorf("attribute %v key is missing", i)
		}

		if atVal, ok := payload[fmt.Sprint(val, i)]; ok {
			dataVal = utils.ConvertToString(atVal)
		} else {
			return data, fmt.Errorf("attribute %v value is missing", i)
		}

		if atValType, ok := payload[fmt.Sprint(valType, i)]; ok {
			dataValType = utils.ConvertToString(atValType)
		} else {
			return data, fmt.Errorf("attribute %v's values data type is missing", i)
		}

		data.Attributes[dataKey] = types{
			Value: dataVal,
			Type:  dataValType,
		}
	}

	// Validate traits with predefined max number of times or lesser
	data.Traits = make(map[string]types)
	for i := 1; i <= int(uattrLimit); i++ {
		var (
			key, val, valType             = "uatrk", "uatrv", "uatrt"
			dataKey, dataVal, dataValType string
		)
		if atKey, ok := payload[fmt.Sprint(key, i)]; ok {
			dataKey = utils.ConvertToString(atKey)
		} else {
			if i > 1 {
				break
			}
			return data, fmt.Errorf("trait %v's key is missing", i)
		}

		if atVal, ok := payload[fmt.Sprint(val, i)]; ok {
			dataVal = utils.ConvertToString(atVal)
		} else {
			return data, fmt.Errorf("trait %v's value is missing", i)
		}

		if atValType, ok := payload[fmt.Sprint(valType, i)]; ok {
			dataValType = utils.ConvertToString(atValType)
		} else {
			return data, fmt.Errorf("trait %v's value data type is missing", i)
		}

		data.Traits[dataKey] = types{
			Value: dataVal,
			Type:  dataValType,
		}
	}

	return data, nil
}
