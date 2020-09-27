
export GOPATH=/Users/dharshekthvel/ac/code/FuzionMatching


# Dependency added
go get "github.com/gorilla/mux"
go get "github.com/gin-gonic/gin"
go get "github.com/segmentio/ksuid"
go get "github.com/streadway/amqp"
go get "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
go get "github.com/newrelic/go-agent"
go get "github.com/newrelic/go-agent/v3/newrelic"
go get "go.mongodb.org/mongo-driver"
go get "golang.org/x/exp/utf8string"
go get "golang.org/x/text/encoding/charmap"
go get "github.com/bradfitz/gomemcache/memcache"
go get "github.com/lib/pq"
go get "github.com/aws/aws-sdk-go"
go get "github.com/go-redis/redis"
go get "github.com/xlzd/gotp"
go get "github.com/segmentio/kafka-go"
go get "github.com/natefinch/lumberjack"
go get "firebase.google.com/go"

go run src/gateway/fuzion.go DEV REST

