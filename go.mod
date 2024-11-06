module github.com/edgexfoundry/edgex-go

go 1.23

require (
	github.com/eclipse/paho.mqtt.golang v1.5.0
	github.com/edgexfoundry/go-mod-bootstrap/v4 v4.0.0-dev.3
	github.com/edgexfoundry/go-mod-configuration/v4 v4.0.0-dev.3
	github.com/edgexfoundry/go-mod-core-contracts/v4 v4.0.0-dev.3
	github.com/edgexfoundry/go-mod-messaging/v4 v4.0.0-dev.5
	github.com/edgexfoundry/go-mod-secrets/v4 v4.0.0-dev.2
	github.com/fxamacker/cbor/v2 v2.7.0
	github.com/go-co-op/gocron/v2 v2.12.3
	github.com/gomodule/redigo v1.9.2
	github.com/google/uuid v1.6.0
	github.com/jackc/pgerrcode v0.0.0-20240316143900-6e2875d9b438
	github.com/jackc/pgx/v5 v5.7.1
	github.com/labstack/echo/v4 v4.12.0
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	github.com/robfig/cron/v3 v3.0.1
	github.com/spf13/cast v1.7.0
	github.com/spiffe/go-spiffe/v2 v2.4.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.28.0
	gopkg.in/yaml.v3 v3.0.1
	modernc.org/sqlite v1.33.1
)

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/edgexfoundry/go-mod-registry/v4 v4.0.0-dev.2 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-jose/go-jose/v4 v4.0.4 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.22.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/nats-io/nats.go v1.37.0 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/zeebo/errs v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/grpc v1.67.1 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	modernc.org/gc/v3 v3.0.0-20240107210532-573471604cb6 // indirect
	modernc.org/libc v1.55.3 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.8.0 // indirect
	modernc.org/strutil v1.2.0 // indirect
	modernc.org/token v1.1.0 // indirect
)

replace github.com/edgexfoundry/go-mod-bootstrap/v4 => github.com/judehung/go-mod-bootstrap/v4 v4.0.0-20241105061404-aed124c70cea
