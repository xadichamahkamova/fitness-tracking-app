package handler

import (
	"github.com/xadichamahkamova/fitness-tracking-app/storage"
	"github.com/xadichamahkamova/fitness-tracking-app/internal/email"
)

type HandlerST struct {
	Queries *storage.Queries
	Notification email.NotificationRepo
}