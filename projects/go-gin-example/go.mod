module github.com/fcf/go-gin-example

go 1.19

require (
	github.com/astaxie/beego v1.12.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.9.0
	github.com/go-ini/ini v1.67.0
	github.com/go-sql-driver/mysql v1.7.1
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
)

require (
	github.com/bytedance/sonic v1.8.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/shiena/ansicolor v0.0.0-20230509054315-a9deabde6e02 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/fcf/go-gin-example/conf => ./conf
	github.com/fcf/go-gin-example/docs => ./docs
	github.com/fcf/go-gin-example/docs/sql => ./docs/sql
	github.com/fcf/go-gin-example/middleware => ./middleware
	github.com/fcf/go-gin-example/models => ./models
	github.com/fcf/go-gin-example/pkg/e => ./pkg/e
	github.com/fcf/go-gin-example/pkg/setting => ./pkg/setting
	github.com/fcf/go-gin-example/pkg/util => ./pkg/util
	github.com/fcf/go-gin-example/routers => ./routers
	github.com/fcf/go-gin-example/routers/api => ./routers/api
	github.com/fcf/go-gin-example/runtime => ./runtime
	github.com/fcf/go-gin-example/pkg/logging => ./pkg/logging
)
