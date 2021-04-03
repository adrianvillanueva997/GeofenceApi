package middleware

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
)

func InitMetrics(app *fiber.App) {
	// labels := map[string]string{"custom_label1":"custom_value1", "custom_label2":"custom_value2"}
	// fiberprometheus.NewWithLabels(labels, namespace, subsystem )
	prometheus := fiberprometheus.New("GeofenceAPI")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)
}
