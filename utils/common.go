package utils

type (
	ResponseSchema struct {
		Status   ResponseStatus `json:"status"`
		Metadata interface{}    `json:"metadata,omitempty"`
		Data     interface{}    `json:"data"`
	}

	ResponseStatus struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)
