// database contains a Relation interface that defines a database relation that
// can be used to generate and execute queries to the database. Views and Tables
// can be implemented to make a difference between read-only and read-write
// resources. Only Tables can be used with the Insert/Update/Delete functions.
// But Select can be used with any type of Relation.

package database
