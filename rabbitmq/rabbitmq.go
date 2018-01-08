// Handle the connection to the remote AMQP server to receive messages

package rabbitmq

import (
  "github.com/streadway/amqp"
  "log"
  "time"
)

type RabbitMQ struct {
  Url                 string  `yaml:"url"`
  Exchange            string  `yaml:"exchange"`
  ConnectionName      string  `yaml:"connectionName"`
  HeartBeat           int     `yaml:"heartBeat"`
  Product             string  `yaml:"product"`
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

// Connect to amqp server as necessary
func (s *RabbitMQ) Connect( ) {
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
  log.Println( "Publishing to ", s.exchange(), routingKey )

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
