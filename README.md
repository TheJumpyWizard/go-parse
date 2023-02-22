# Go Parse (WIP)
###### This is a SQL parser written in Go that can parse and analyze SQL statements. It supports basic SELECT and INSERT statements, as well as WHERE, GROUP BY, HAVING, ORDER BY, LIMIT, and OFFSET clauses. It also includes basic semantic analysis checks to ensure that the input SQL statement is well-formed and conforms to the rules of SQL.

### Getting Started
To build and run the SQL parser, you will need Go version 1.16 or later installed on your machine. To check your Go version, run the following command:
```
$ go version
```
If you don't have Go installed, you can download it from the official Go website.

To get started with the SQL parser, clone the repository and navigate to the project directory:

```
shell
$ git clone https://github.com/your-username/sql-parser-go.git
$ cd sql-parser-go
```
Next, you can build and run the program using the following commands:

```
shell
$ go build -o sql-parser cmd/main.go
$ ./sql-parser
```
This will start the SQL parser in interactive mode, where you can enter SQL statements and see the parsed result. To exit the program, type exit.

### Usage
The SQL parser supports the following SQL statements:

* SELECT Statement
    * The SELECT statement is used to select data from a database table.

``` SELECT column1, column2, ... FROM table_name WHERE condition GROUP BY column1, column2, ... HAVING condition ORDER BY column1 ASC/DESC, column2 ASC/DESC, ... LIMIT offset, count ```

* INSERT Statement
    * The INSERT statement is used to insert new data into a database table.

``` INSERT INTO table_name (column1, column2, ...) VALUES (value1, value2, ...), (value1, value2, ...), ... ```