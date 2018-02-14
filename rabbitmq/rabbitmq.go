// A simple amqp library for connecting to RabbitMQ
//
// This is a wrapper around the github.com/streadway/amqp library.
//
package rabbitmq

import (
  "github.com/streadway/amqp"
  "log"
  "time"
)

type RabbitMQ struct {
  // The amqp url to connect to
  Url                 string  `yaml:"url"`
  // The schange for publishing, defaults to amq.topic
  Exchange            string  `yaml:"exchange"`
  // The name of the connection that appears in the management plugin
  ConnectionName      string  `yaml:"connectionName"`
  // The heartBeat in seconds. Defaults to 10
  HeartBeat           int     `yaml:"heartBeat"`
  // The product name in the management plugin (optional)
  Product             string  `yaml:"product"`
  // The product version in the management plugin (optional)
  Version             string  `yaml:"version"`
  // ===== Internal
  connection     *amqp.Connection  `yaml:"-"`  // amqp connection
  channel        *amqp.Channel     `yaml:"-"`  // amqp channel
}

// called by main() ensure mandatory config is present
func (s *RabbitMQ) url( ) string {
  if s.Url == "" {
    log.Fatal( "amqp.url is mandatory" )
  }
  return s.Url
}

func (s *RabbitMQ) exchange() string {
  if s.Exchange == "" {
    return "amq.topic"
  }
  return s.Exchange
}

// Connect connects to the RabbitMQ instace thats been configured.
func (s *RabbitMQ) Connect( ) {
  if s.connection != nil {
    return
  }
  
  log.Println( "Connecting to amqp" )

  var heartBeat = s.HeartBeat
  if heartBeat == 0 {
    heartBeat = 10
  }

  var product = s.Product
  if product == "" {
    product = "Area51 GO"
  }

  var version = s.Version
  if version == "" {
    version = "0.2β"
  }

  // Use the user provided client name
  connection, err := amqp.DialConfig( s.url(), amqp.Config{
    Heartbeat:  time.Duration( heartBeat ) * time.Second,
    Properties: amqp.Table{
      "product": product,
      "version": version,
      "connection_name": s.ConnectionName,
    },
    Locale: "en_US",
  } )

  if err != nil {
    log.Fatal( err )
  }

  s.connection = connection

  // Most operations happen on a channel.  If any error is returned on a
  // channel, the channel will no longer be valid, throw it away and try with
  // a different channel.  If you use many channels, it's useful for the
  // server to
  channel, err := s.connection.Channel()

  if err != nil {
    log.Fatal( err )
  }

  s.channel = channel

  log.Println( "AMQP Connected" )

  err = s.channel.ExchangeDeclare( s.exchange(), "topic", true, false, false, false, nil)

  if err != nil {
    log.Fatal( err )
  }

}

// Publish a message
func (s *RabbitMQ) Publish( routingKey string, msg []byte ) {

  err := s.channel.Publish(
    s.exchange(),
    routingKey,
    false,
    false,
    amqp.Publishing{
      Body: msg,
  })

  if err != nil {
    log.Fatal( err )
  }

}

// QueueDeclare declares a queue
func (r *RabbitMQ) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
  return r.channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}

// QueueBind binds a queue to an exchange & routingKey
func (r *RabbitMQ) QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error {
  return r.channel.QueueBind(name, key, exchange, noWait, args)
}

// Consume adds a consumer to the queue and returns a GO channel
func (r *RabbitMQ) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
  return r.channel.Consume( queue, consumer, autoAck, exclusive, noLocal, noWait, args )
}
