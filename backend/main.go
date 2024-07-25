package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/BlackMocca/sqlx"
	"github.com/Blackmocca/example-geolocation/constants"
	"github.com/Blackmocca/example-geolocation/middleware"
	"github.com/Blackmocca/example-geolocation/route"
	"github.com/Blackmocca/example-geolocation/service/tracking/http"
	"github.com/Blackmocca/example-geolocation/service/tracking/repository"
	"github.com/Blackmocca/example-geolocation/service/tracking/usecase"
	"github.com/Blackmocca/example-geolocation/service/tracking/validator"
	"github.com/labstack/echo/v4"
	pq "github.com/lib/pq"
	"github.com/xeipuuv/gojsonschema"
)

func psqlConnect() *sqlx.DB {
	conn, err := pq.ParseURL(constants.ENV_DATABASE_URL)
	if err != nil {
		panic(err)
	}
	client, err := sqlx.Connect("postgres", conn)
	if err != nil {
		panic(err)
	}
	if err := client.Ping(); err != nil {
		panic(err)
	}
	return client
}

func initscript(client *sqlx.DB) {
	sql := `
	

	CREATE TABLE IF NOT EXISTS trackings (
		id SERIAL PRIMARY KEY,
		plate_number VARCHAR(20) NOT NULL,
		lat NUMERIC NOT NULL,
		lon NUMERIC NOT NULL,
		track_time TIMESTAMP NOT NULL
	); 

	CREATE INDEX IF NOT EXISTS track_idx ON trackings USING BRIN (track_time);
	`

	client.MustExec(sql)
}

func getJSONSchemaLoader(path string) *gojsonschema.Schema {
	path = filepath.Clean(path)
	bu, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	loader := gojsonschema.NewSchemaLoader()
	loader.Draft = gojsonschema.Draft7
	loader.AutoDetect = false
	schema, err := loader.Compile(gojsonschema.NewBytesLoader(bu))
	if err != nil {
		panic(err)
	}

	return schema
}

func main() {
	client := psqlConnect()
	defer client.Close()
	initscript(client)

	middl := middleware.InitMiddleware("")
	e := echo.New()
	e.Use(middl.InitContextIfNotExists, middl.InputForm)

	trackingRepo := repository.NewPsqlRepository(client)
	trackingUs := usecase.NewUsecase(trackingRepo)
	trackingHandler := http.NewTrackingHandler(trackingUs)

	trackingValidator := validator.Validation{
		SaveTrackingSchema: getJSONSchemaLoader("./assets/jsonschema/savetracking.json"),
	}

	/* register route */
	r := route.NewRoute(e, middl)
	r.Register(trackingHandler, trackingValidator)

	e.Start(":3000")
}
