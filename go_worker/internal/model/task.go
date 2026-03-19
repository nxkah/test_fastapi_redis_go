package model

import "encoding/json"

type Task struct {
	TaskID  string          `json:"task_id"`
	Type    string          `json:"type"`
	UserID  string          `json:"user_id,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

type Result struct {
	TaskID string `json:"task_id"`
	Status string `json:"status"`
	Data   string `json:"data"`
}
