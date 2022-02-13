module icy-mountain/models

go 1.14

replace icy-mountain/database => ../database

require (
	github.com/jinzhu/gorm v1.9.16
	gorm.io/gorm v1.22.4
	icy-mountain/database v0.0.0-00010101000000-000000000000
)
