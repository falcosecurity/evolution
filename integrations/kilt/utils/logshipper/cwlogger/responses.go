package cwlogger

type putLogEventsSuccessResponse struct {
	NextSequenceToken string `json:"nextSequenceToken"`
}

type putLogEventsErrorResponse struct {
	Code                  string  `json:"__type"`
	Message               string  `json:"message"`
	ExpectedSequenceToken *string `json:"expectedSequenceToken"`
}

