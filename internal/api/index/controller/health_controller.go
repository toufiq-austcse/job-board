package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/toufiq-austcse/go-api-boilerplate/config"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/api_response"
	"github.com/toufiq-austcse/go-api-boilerplate/pkg/db/providers/influxdb"
	"net/http"
)

// Index hosts godoc
// @Summary  Health Check
// @Tags     Index
// @Accept   json
// @Produce  json
// @Success  200
// @Router   / [get]
func Index() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": config.AppConfig.APP_NAME + " is Running",
		})
	}
}

func GetInFluxData() gin.HandlerFunc {

	return func(context *gin.Context) {
		client := influxdb.OpenInfluxDbConnection()
		queryAPI := client.QueryAPI("toufiq")

		result, err := queryAPI.Query(context, `
import "date"

t1=from(bucket: "demo")
  |> range(start: -50y)
  |> filter(fn: (r) => r["_measurement"] == "offer_live_classes")
  |>drop(columns:["_start","_stop"])
  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
  |> group(columns: ["_time"])
  |>count(column: "live_class_id")
  |> rename(columns: { "live_class_id": "count"})
  |> map(fn: (r) => ({ r with _time: date.truncate(t: r._time, unit: 1d) }))


t2=from(bucket: "demo")
  |> range(start: -50y)
  |> filter(fn: (r) => r["_measurement"] == "user_live_classes")
 |>filter(fn: (r)=>r.user_id=="1")
  |>drop(columns:["_start","_stop"])
  |> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
   |> group(columns: ["_time"])
  |>count(column: "live_class_id")
  |> rename(columns: { "live_class_id": "participated"})
   |> map(fn: (r) => ({ r with _time: date.truncate(t: r._time, unit: 1d) }))

join(tables: {t1: t1, t2: t2}, on: ["_time"])
  
`)

		if err != nil {
			fmt.Print("hello ")
			errRes := api_response.BuildErrorResponse(http.StatusInternalServerError, "Internal Server Error", err.Error(), nil)
			context.JSON(errRes.Code, errRes)
			return
		}

		res := []map[string]interface{}{}
		for result.Next() {
			if result.TableChanged() {
				//fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			//fmt.Printf("%v %v\n", result.Record().Values(), result.Record().Value())

			res = append(res, result.Record().Values())
		}
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
		context.JSON(http.StatusOK, res)
	}
}
