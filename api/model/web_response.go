package model

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

const StatusInternalServerError = "Internal server error"
const StatusOK = "OK"
const StatusNotFound = "Status not found"
const StatusCreated = "Created"
const StatusDeleted = "Deleted"
const StatusRestored = "Restored"