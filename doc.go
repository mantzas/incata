/*
Package incata is a source eventing data access library. The name combines incremental (inc) and data (ata).
Details about event sourcing can be read on Martin Fowlers site(http://martinfowler.com/eaaDev/EventSourcing.html).

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

In order to use the appender or retriever we have to provide the following:

1. a serializer or deserializer which implements the Serializer or Deserializer interface or the .  A JSONMarshaller is provided.

2. a writer or reader which implements the Writer or Reader interface. A SQLWriter and SQLReader is provided.

3. a appender and retriever which implement the Appender and Retriever interface. Appender and Retriever are provided.

The supported relational DB's are MS Sql Server and PostgreSQL.


Check out the examples in the examples folder for setting up the default marshaller and reader/writers.


MS SQL Server Setup

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

PostgreSQL Setup

PostgreSQL Driver used:

    "github.com/lib/pq"

DB Table setup (Provide a table name)

      CREATE TABLE {TableName}
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
       ON {TableName}
       USING btree
       ("SourceId");

For a full guide visit https://github.com/mantzas/incata
*/
package incata
