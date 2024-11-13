# Query Package

This Go package provides functionality for executing database queries and mapping the results to a 2D slice. It is designed to work with the `database/sql` package and allows you to retrieve rows from a database and store them in a format that is easy to process.

## Overview

The main component of the package is the `Recordset` struct, which represents a set of rows fetched from a database query. The `Map` method maps each row of data into a 2D slice of `interface{}` type. This allows for a flexible, dynamic way of handling query results where the column types are unknown in advance.

## Features

- Fetch rows from a database using the `sql.Rows` object.
- Dynamically scan and store column data into a 2D slice (`[][]interface{}`).
- Handle any number of columns and rows.
- Error handling for column retrieval and row scanning.

## Installation

To use this package in your project, import it and use the `NewRecordset` function to create an instance of the `Recordset` struct.

```bash
go get github.com/choconutella/gocho-query
