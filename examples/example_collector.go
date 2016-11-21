package main

import (
  "github.com/juliancheal/tivan"
)

tivan = NewTivanClient()
// For each thing collect data
// store data in tivan.data_point() -> which stores data in array
// collector then creates a batch of the data_points and writes ti influxdb
tivan.collector()
