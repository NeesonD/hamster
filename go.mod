module hamster

go 1.16

require (
	fyne.io/fyne/v2 v2.3.1
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/gocarina/gocsv v0.0.0-20230226133904-70c27cb2918a
	github.com/gocolly/colly/v2 v2.1.0
	github.com/json-iterator/go v1.1.12
	github.com/mattn/go-sqlite3 v1.14.16
	github.com/mitchellh/mapstructure v1.4.3
	github.com/panjf2000/ants/v2 v2.7.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.10.1
	github.com/tickstep/aliyunpan-api v0.1.3
)

replace github.com/tickstep/aliyunpan-api v0.1.3 => github.com/NeesonD/aliyunpan-api v0.0.0-20230323092347-5d196238b60d
