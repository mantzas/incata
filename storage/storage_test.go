package storage_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/mantzas/incata/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Storage", func() {

	It("execute returns expected data", func() {

		dbInner, mock, _ := sqlmock.New()

		storage, _ := NewStorageFinalized(dbInner, MSSQL, "Event")

		mock.ExpectExec("123").WithArgs(1, 2, 3).WillReturnResult(sqlmock.NewResult(1, 1))

		storage.Exec("123", 1, 2, 3)

		err := mock.ExpectationsWereMet()
		Expect(err).NotTo(HaveOccurred(), "there were unfulfilled expections: %s", err)
	})

	It("query returns expected data", func() {

		dbInner, mock, _ := sqlmock.New()
		storage, _ := NewStorageFinalized(dbInner, PostgreSQL, "Event")

		var rows *sqlmock.Rows
		mock.ExpectQuery("123").WithArgs(1, 2, 3).WillReturnRows(rows)

		storage.Query("123", 1, 2, 3)
		err := mock.ExpectationsWereMet()
		Expect(err).NotTo(HaveOccurred(), "there were unfulfilled expections: %s", err)
	})

	It("close suceeds", func() {

		dbInner, mock, _ := sqlmock.New()
		storage, _ := NewStorageFinalized(dbInner, MSSQL, "Event")

		mock.ExpectClose()

		storage.Close()
		err := mock.ExpectationsWereMet()
		Expect(err).NotTo(HaveOccurred(), "there were unfulfilled expections: %s", err)
	})

	It("wrong db type returns error when creating new finalized storage", func() {

		dbInner, _, _ := sqlmock.New()
		_, err := NewStorageFinalized(dbInner, 3, "Event")
		Expect(err).To(HaveOccurred())
	})

	It("wrong db type returns error when creating new storage", func() {

		_, err := NewStorage(3, "123", "Event")
		Expect(err).To(HaveOccurred())
	})

	It("new storage returns error when opening", func() {

		_, err := NewStorage(MSSQL, "123", "Event")
		Expect(err).To(HaveOccurred())
	})

	DescribeTable("Convert string to DB type",
		func(text string, expectedDbType DbType, hasErrors bool) {

			actualDbType, err := ConvertToDbType(text)

			if hasErrors {
				Expect(err).To(HaveOccurred())
			} else {
				Expect(actualDbType).To(Equal(expectedDbType))
			}
		},
		Entry("mssql", "mssql", MSSQL, false),
		Entry("MSSQL", "MSSQL", MSSQL, false),
		Entry("MsSQL", "MsSQL", MSSQL, false),
		Entry("MsSql", "MsSql", MSSQL, false),
		Entry("postgresql", "postgresql", PostgreSQL, false),
		Entry("PostgreSQL", "PostgreSQL", PostgreSQL, false),
		Entry("POSTGRESQL", "POSTGRESQL", PostgreSQL, false),
		Entry("xxx", "xxx", PostgreSQL, true),
	)
})
