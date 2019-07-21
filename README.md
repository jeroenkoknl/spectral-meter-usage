# Meter Usage Service
This application is part of the assignment for Spectral. This is a microservice
that exposes time based electricity consumption data over the gRPC protocol. 

## How to run

This application is written in Go. When not already done, please install the [Go tools](https://golang.org/doc/install#install) first.


After cloning this repo, run the following [dep](https://golang.github.io/dep/docs/installation.html) command to get all dependencies:

```
dep ensure
```

Next run the application by calling:

```
go run .
```

## API Description

The service is described in the following proto file:

```proto
syntax = "proto3";

package main;

service MeterUsage {
  rpc GetMeasurements (GetMeasurementsRequest) returns (stream Measurement) {}
}

message Measurement {
    int64 timestamp = 1;
    int32 metervalue = 2;
}

message GetMeasurementsRequest {}
```

The `MeterUsage` service has one method called `GeMeasurements`. It takes one parameter of the empty type `GetMeasurementsRequest`. This type is empty and is there because the input parameters are required in the proto description file and the actual method does not have input parameters. The method returns a `stream` of `Measurement` items. 

### Measurement type

Because the assignment was not giving more details about the data the service was providing except that it is 'time based electricity consumption data'. Because of this, I had to make some assumptions on the data type. I thought this type would at least have the following fields:

|Name|DataType|Description|
|----|--------|-----------|
|timestamp|int64|A UNIX timestamp of the measurement. This is a 64 bit integer to allow timestamps after the year 2032. A UNIX timestamp was chosen for maximum portability between platforms.|
|metervalue|int32|The actual measured value from the electricity meter, most likely in Ah.|

### gRPC stream

The reason the `GetMeasurements` returns a `stream` is to allow clients to start processing the results before all items are transferred.

## Go

This microservice is written in Go because of two reasons:

1. Go and gRPC are both developed by Google and I expected Go libraries available for gRPC would be well supported by Google. 
2. I already have experience in Go, even though it was already one and a half year ago since I had written my last line of Go.

## Docker

Even though I have not been able to fully automate the build and deployment process, I have created a `Dockerfile` that builds the application.

To create a Docker image you could run the following command: 

```
docker build -t spectral-meter-usage:0.1.0 .
```