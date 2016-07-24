# incata [![alt text](https://godoc.org/github.com/mantzas/incata?status.png)](https://godoc.org/github.com/mantzas/incata)&nbsp;[![build status](https://img.shields.io/travis/mantzas/incata.svg)](http://travis-ci.org/mantzas/incata)&nbsp;[![Coverage Status](https://coveralls.io/repos/github/mantzas/incata/badge.svg?branch=master)](https://coveralls.io/github/mantzas/incata?branch=master)&nbsp;[![Go Report Card](https://goreportcard.com/badge/github.com/mantzas/incata)](https://goreportcard.com/report/github.com/mantzas/incata)

Event Sourcing Data Access Library

Package incata is a source eventing data access library. The name combines incremental (inc) and data (ata).
Details about event sourcing can be read on Martin Fowlers [site](http://martinfowler.com/eaaDev/EventSourcing.html).

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

In order to use the appender or retriever we have to provide the following

- A serializer or deserializer which implements the Serializer or Deserializer interface or the .  A JSONMarshaller is provided.
- A writer or reader which implements the Writer or Reader interface. A SQLWriter and SQLReader is provided.
- A appender and retriever which implement the Appender and Retriever interface. Appender and Retriever are provided.

The supported relational DB's are MS Sql Server and PostgreSQL.

## Check out the examples in the examples folder for setting up the default marshaller and reader/writers

### MS SQL Server Setup

 SQL Server Driver used:

    "github.com/denisenkom/go-mssqldb"

 DB Table setup (Provide a table name)

        CREATE TABLE {TableName} (
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
         ON {TableName} (SourceId)
         ON [PRIMARY]
         GO

### PostgreSQL Setup

PostgreSQL Driver used:

    "github.com/lib/pq"

DB Table setup (Provide a table name)

      CREATE TABLE {table_name}
      (
        id bigserial NOT NULL,
        source_id uuid NOT NULL,
        created timestamp with time zone NOT NULL,
        event_type character varying(250) NOT NULL,
        version integer NOT NULL,
        payload text NOT NULL,
        CONSTRAINT pk_{table_name}_id PRIMARY KEY (id)
      )
      WITH (
        OIDS=FALSE
      );

      CREATE INDEX ix_order_event_source_id
        ON {table_name}
        USING btree
        (source_id);
