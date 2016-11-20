package main

import (
    "log"
    "time"
    "fmt"
    "math/rand"
    "github.com/influxdata/influxdb/client/v2"
)

const (
    database = "quill_development"
)

func main() {
    // Make client
    c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: "http://localhost:8086",

    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    writePoints(c)
}

func writePoints(clnt client.Client) {
    sampleSize := 1000
    rand.Seed(42)

    bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
        Database:  "quill_development",
        Precision: "us",
    })

    for i := 0; i < sampleSize; i++ {
        regions := []string{"us-west1", "us-west2", "us-west3", "us-east1"}
        tags := map[string]string{
            "cpu":    "cpu-total",
            "host":   fmt.Sprintf("host%d", rand.Intn(1000)),
            "region": regions[rand.Intn(len(regions))],
        }

        idle := rand.Float64() * 100.0
        fields := map[string]interface{}{
            "idle": idle,
            "busy": 100.0 - idle,
        }
        data, _ := client.NewPoint(
            "cpu_usage",
            tags,
            fields,
            time.Now())
        bp.AddPoint(data)
    }

    err := clnt.Write(bp)
    if err != nil {
        log.Fatal(err)
    }
}
