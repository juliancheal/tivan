package tivan

import (
  "log"
  "github.com/influxdata/influxdb/client/v2"
)

const (
  database = "example"
)

func NewTivanClient() {
  c, err := client.NewHTTPClient(client.HTTPConfig{
      Addr: "http://localhost:8086",
  })

  if err != nil {
      log.Fatalln("Error: ", err)
  }
  clnt client.Client
}

func data_point(measurement string, tags map, fields map, timestamp string) {
  data, err := client.NewPoint(
      measurement,
      tags,
      fields,
      timestamp)
  if err != nil {
      log.Fatal(err)
  }
}

func collector(arrayOfDataPoints []int) {
  bp, err := client.NewBatchPoints(client.BatchPointsConfig{
      Database:  "",
      Precision: "us",
  })
  if err != nil {
      log.Fatal(err)
  }

  for i := 0; i < sizeOfArrayOfStuff; i++ {
    bp.AddPoint(sizeOfArrayOfStuff[i])
  }

  err = clnt.Write(bp)
  if err != nil {
      log.Fatal(err)
  }
}
