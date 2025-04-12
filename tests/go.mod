module gorm.io/gen/tests

go 1.22.2

toolchain go1.22.12

require (
	gorm.io/driver/mysql v1.5.6
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gen v0.3.19
	gorm.io/gorm v1.25.9
	gorm.io/plugin/dbresolver v1.5.0
)

require gorm.io/hints v1.1.1 // indirect

replace gorm.io/gen => ../
