# incata ![alt text](https://travis-ci.org/mantzas/incata.svg?branch=master "Build Status")&nbsp;[![alt text](https://godoc.org/github.com/mantzas/incata?status.png)](https://godoc.org/github.com/mantzas/incata)&nbsp;[![Report card](http://goreportcard.com/badge/mantzas/incata)](http://goreportcard.com/report/mantzas/incata)
Event Sourcing Data Access Library

Package incata is a source eventing data access library. The name combines incremental (inc) and data (ata).
Details about event sourcing can be read on Martin Fowlers site(http://martinfowler.com/eaaDev/EventSourcing.html).

We have to provide two components in order to setup a appender, the serializer and the writer.
In order to create a custom serializer just implement the Serializer interface.
The same goes for the writer which needs to implement the Writer interface.

Currently we support two relational DB's, MS Sql Server and Postgresql.

The stored Event has the following structure:

    type Event struct {
      Id        int64
      SourceID  uuid.UUID
      Created   time.Time
      Payload   interface{}
      EventType string
      Version   int
    }

The payload is the actual data that we like to store in our DB.
Since the serializer can be anything the data type is set to interface{}.
This means that our db table column for the Payload have to match the serializer's result data type.

In order to use the appender we have to do the following once on application start up (Set up):

1. We create a serializer. Here we have a JSONSerializer

        sr := serializer.NewJSONSerializer()

2. We create the DB table to hold the events. For MS Sql Server we should have the following statement.
Payload is set to NVARCHAR(MAX) since we are serializing using the JSON format.

        CREATE TABLE Event (
          Id BIGINT IDENTITY
          ,SourceId UNIQUEIDENTIFIER NOT NULL
          ,Created DATETIME2 NOT NULL
          ,EventType NVARCHAR(250) NOT NULL
          ,Version INT NOT NULL
          ,Payload NVARCHAR(MAX) NOT NULL
          ,CONSTRAINT PK_Event PRIMARY KEY CLUSTERED (Id)
        ) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]

         GO

         CREATE INDEX IX_Event_SourceId
         ON Linear.dbo.Event (SourceId)
         ON [PRIMARY]
         GO

3. We create a db object providing the connection string and the insert statement. We have to import the MS SQL server package!

         _ "github.com/denisenkom/go-mssqldb"

         conn := fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;", *dbServer, *dbName, *dbUser, *dbPassword)

         db, err := relational.NewDb(relational.MsSQL, "mssql", conn,
         		"INSERT INTO Event (SourceId, Created, EventType, Version, Payload)  VALUES (?, ?, ?, ?, ?)")`

4. We create a writer

        wr := writer.NewSQLWriter(db, sr)

5. We set up the appender.

        incata.SetupAppender(wr)


That's it. Every time we need to append a event to the store we (Usage):

  1. Create a event

          event := model.NewEvent(sourceID, Payload{Description: string(index)}, "TestEvent", 1)

  2. Get a new appender and append the event

          ar, err := incata.NewAppender()
          if err != nil {
            return
          }
          err = ar.Append(*event)

In order to support Postgresql we have to alter step 2. in Set up with

      CREATE TABLE linearevents
      (
       "Id" serial NOT NULL,
       "SourceId" uuid,
       "Created" timestamp without time zone,
       "EventType" character varying(250),
       "Version" integer,
       "Payload" text,
       CONSTRAINT "PK_Event" PRIMARY KEY ("Id")
      )
      WITH (
       OIDS=FALSE
      );

      CREATE INDEX "event_idx_sourceId"
       ON linearevents
       USING btree
       ("SourceId");

and step 3. with

    _ "github.com/lib/pq"

    conn := fmt.Sprintf("postgres://postgres:%s@%s/%s", *dbUser, *dbPassword, *dbName)

    db, err = relational.NewDb(relational.Postgresql, "postgres", conn,
    	`INSERT INTO linearevents ("SourceId", "Created", "EventType", "Version", "Payload") VALUES ($1, $2, $3, $4, $5)`)

The example below uses MS SQL Server and JSON serialization.

     package main

     import (
     	"flag"
     	"fmt"
     	"runtime"
     	"strconv"
     	"sync"
     	"time"

     	_ "github.com/denisenkom/go-mssqldb"
     	"github.com/mantzas/incata"
     	"github.com/mantzas/incata/model"
     	"github.com/mantzas/incata/relational"
     	"github.com/mantzas/incata/serializer"
     	"github.com/mantzas/incata/writer"
     	"github.com/twinj/uuid"
     )

     type Payload struct {
     	Description string `json:"description"`
     }

     func main() {
     	eventCount := flag.Int("eventCount", 50, "count of events to be stored")
     	dbServer := flag.String("dbServer", "localhost", "the db hostname")
     	dbName := flag.String("dbName", "name", "name of the db")
     	dbUser := flag.String("dbUser", "user", "db user")
     	dbPassword := flag.String("dbPassword", "password", "db password")
     	flag.Parse()
     	fmt.Printf("Event Count: %d", *eventCount)
     	fmt.Println()
     	fmt.Printf("Max Procs: %d", runtime.NumCPU())
     	fmt.Println()

     	conn := fmt.Sprintf("Server=%s;Database=%s;User Id=%s;Password=%s;", *dbServer, *dbName, *dbUser, *dbPassword)

     	fmt.Println(conn)

     	db, err := relational.NewDb(relational.MsSQL, "mssql", conn,
     		"INSERT INTO Event (SourceId, Created, EventType, Version, Payload)  VALUES (?, ?, ?, ?, ?)")
     	if err != nil {
     		panic(err.Error())
     	}

     	sr := serializer.NewJSONSerializer()

     	wr := writer.NewSQLWriter(db, sr)

     	incata.SetupAppender(wr)

     	sourceID := uuid.NewV4()

     	fmt.Println("Starting")
     	var wg sync.WaitGroup
     	wg.Add(*eventCount)
     	startTime := time.Now()

     	for i := 1; i <= *eventCount; i++ {

     		go func(index int) {
     			defer wg.Done()
     			event := model.NewEvent(sourceID, Payload{Description: string(index)}, "TestEvent", 1)
     			ar, err := incata.NewAppender()
     			if err != nil {
     				return
     			}
     			err = ar.Append(*event)
     		}(i)
     	}

     	wg.Wait()

     	timePassed := time.Since(startTime)
     	timePerEvent := timePassed.Nanoseconds() / int64(*eventCount)
     	timeDurationPerEvent, _ := time.ParseDuration(strconv.FormatInt(timePerEvent, 10) + "ns")

     	fmt.Printf("Finished in %s. %s per event", timePassed, timeDurationPerEvent)
     }
