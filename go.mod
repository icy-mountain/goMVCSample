module mvc_sample.com

go 1.14

require (
	github.com/labstack/echo/v4 v4.6.3
	gorm.io/driver/mysql v1.2.3
	gorm.io/driver/postgres v1.2.3 // indirect
	gorm.io/gorm v1.22.5
	icy-mountain/controllers v0.0.0-00010101000000-000000000000
	icy-mountain/database v0.0.0-00010101000000-000000000000
	icy-mountain/models v0.0.0-00010101000000-000000000000
)

replace icy-mountain/controllers => ./controllers

replace icy-mountain/models => ./models

replace icy-mountain/database => ./database
