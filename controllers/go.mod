module icy-mountain/controllers

go 1.14

require (
	github.com/labstack/echo/v4 v4.6.3
	icy-mountain/database v0.0.0-00010101000000-000000000000
	icy-mountain/models v0.0.0-00010101000000-000000000000
)

replace icy-mountain/models => ../models

replace icy-mountain/database => ../database
