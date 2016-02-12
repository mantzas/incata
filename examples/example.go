package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"

	_ "github.com/denisenkom/go-mssqldb" // import sql driver
	_ "github.com/lib/pq"                // import postgrsql driver
	"github.com/mantzas/incata"
	"github.com/satori/go.uuid"
)

type payload struct {
	Description string `json:"description"`
}

func main() {
	eventCount := flag.Int("events", 1, "count of events stored")
	dbTypeInput := flag.String("db", "MSSQL", "db type: MSSQL or PostgreSQL")
	dbName := flag.String("dbName", "", "db name")
	connection := flag.String("conn", "", `MSSQL:      Server={server};Database={db};User Id={user};Password={password};
        PostgreSQL: postgres://{user}:{passwd}@{host}/{db}?sslmode=disable`)
	flag.Parse()

	if *eventCount == 0 || *dbTypeInput == "" || *connection == "" || *dbName == "" {
		flag.Usage()
		return
	}

	fmt.Printf("Event Count: %d", *eventCount)
	fmt.Println()
	fmt.Printf("Max Procs: %d", runtime.NumCPU())
	fmt.Println()

	dbType, err := incata.ConvertToDbType(*dbTypeInput)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(*connection)

	storage, err := incata.NewStorage(dbType, *connection, *dbName)

	if err != nil {
		panic(err)
	}

	sr := incata.NewJSONMarshaller()
	wr := incata.NewSQLWriter(storage, sr)

	incata.SetupAppender(wr)

	sourceID := uuid.NewV4()

	fmt.Println("Starting")
	var wg sync.WaitGroup
	wg.Add(*eventCount)
	startTime := time.Now()

	for i := 1; i <= *eventCount; i++ {

		go func(index int) {
			defer wg.Done()
			event := incata.NewEvent(sourceID, payload{Description: string(index)}, "TestEvent", 1)
			ar, err := incata.NewAppender()
			if err != nil {
				log.Print(err)
			}
			err = ar.Append(*event)
		}(i)
	}

	wg.Wait()

	timePassed := time.Since(startTime)
	timePerEvent := timePassed.Nanoseconds() / int64(*eventCount)
	timeDurationPerEvent, _ := time.ParseDuration(strconv.FormatInt(timePerEvent, 10) + "ns")

	fmt.Printf("Finished appending events in %s. %s per event", timePassed, timeDurationPerEvent)
	fmt.Println()

	reader := incata.NewSQLReader(storage, sr)
	incata.SetupRetriever(reader)

	retriever, err := incata.NewRetriever()
	if err != nil {
		panic(err)
	}

	events, err := retriever.Retrieve(sourceID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Read %d events, expected %d", len(events), *eventCount)
	fmt.Println()

	for index, event := range events {
		if !uuid.Equal(sourceID, event.SourceID) {
			fmt.Printf("Expected %s but was %s on item %d. %s", sourceID, event.SourceID, index, event.Created)
			fmt.Println()
		}
	}
}
