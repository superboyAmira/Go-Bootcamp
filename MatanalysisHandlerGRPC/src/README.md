# Backend Service with gRPC conn & ORM

## Stack

* go 1.22
* gRPC/Protobuf
* GORM
* docker + docker-compose
* PostgreSQL

## Short description

We have a server and a client side. When connected, the server part generates the session UUID, random EV and SD (the boundaries of randomization are specified in the code), and the connection time. 

```
SessionId: UUIDcon,
Frequency: normDist.Rand(),
TimeUtc: timestamppb.Now(),
```

Based on this data, the server response is generated, providing a data stream with the session UUID, a random number based on EV and STD according to the law of normal distribution, and the time of sending. This stream is received by the client.

This stream is read by the client until the moment that the client application determines by sending SIGINT to the terminal (or the connection time, determined by the contexts, expires) From the array of data sent from the server, the client calculates the approximate expectation and STD of the array (with each new number from the stream it is updated)

```
type ReceivedData struct {
	count int64
	seq   []float64

	Mean float64
	STD  float64
}
```

After processing the server array is completed, the client switches to anomaly analysis mode. The anomaly is a number that is not included in the boundaries of the ```Mean +- (k * STD)``` equation, these data were calculated in the client receiving mode, as described above. The k coefficient is set by the client at startup via a command line parameter.

All anomalies are sent to the database, automatically creating a table and entering all information about the found anomaly in turn.

```
type AnomalyModel struct {
	SessionId string
	Frequency float64
	Mean      float64
	STD       float64
	Time      time.Time
}
```

.xml consist all settings to a server and database.