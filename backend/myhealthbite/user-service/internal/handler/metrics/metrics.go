package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	UserRegistrationCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "user_registration_total",
		Help: "Total number of successful user registrations",
	})

	UserLoginCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "user_login_total",
		Help: "Total number of successful user logins",
	})

	UserErrorCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "user_error_total",
		Help: "Total number of user-related errors",
	})
)

func init() {
	prometheus.MustRegister(UserRegistrationCounter)
	prometheus.MustRegister(UserLoginCounter)
	prometheus.MustRegister(UserErrorCounter)
}
