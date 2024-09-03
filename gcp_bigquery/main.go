package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// SampleData represents the structure of the data to be inserted into BigQuery.
type SampleData struct {
	Name  string  `bigquery:"name"`  // Name of the data
	Value float32 `bigquery:"value"` // Value of the data
	Time  string  `bigquery:"time"`  // Time when the data was recorded
}

const credentialsFile = "./key.json"

// BatchPushToBigQuery is a function that inserts data into BigQuery in batches.
func BatchFetchToBigQuery() {
	ctx := context.Background()

	// Path to your service account key file

	// Create a BigQuery client with credentials
	client, err := bigquery.NewClient(ctx, "gcp-learning-414814", option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatalf("Failed to create BigQuery client: %v", err)
	}
	defer client.Close()

	// Example: Running a simple query
	query := client.Query(`
        SELECT *
        FROM ` + "`gcp-learning-414814.rnd_query.sample_table`" + `
        LIMIT 10
    `)

	// Execute the query
	it, err := query.Read(ctx)
	if err != nil {
		log.Fatalf("Failed to run query: %v", err)
	}

	// Iterate through the query results
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate through query results: %v", err)
		}
		fmt.Println(values)
	}
}

func BatchPushToBigQuery() {
	ctx := context.Background()
	// Create a new BigQuery client using the provided credentials file
	client, err := bigquery.NewClient(ctx, "gcp-learning-414814", option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatalf("Failed to create BigQuery client: %v", err)
	}
	dataset := client.Dataset("rnd_query") // Specify the dataset name
	table := dataset.Table("sample_table") // Specify the table name
	// Retrieve the schema for the specified table
	meta, err := table.Metadata(ctx)
	if err != nil {
		log.Fatalf("Failed to get table meta data: %v", err)
	}
	// Print the schema of the table
	for _, field := range meta.Schema {
		fmt.Printf("Field Name: %s, Field Type: %s\\n", field.Name, field.Type)
	}
	// Create sample data to be inserted into the table
	data := SampleData{
		Name:  "anik",
		Value: 587.123,
		Time:  time.Now().Format(time.RFC3339),
	}
	data2 := SampleData{
		Name:  "shojib",
		Value: 587.123,
		Time:  time.Now().Format(time.RFC3339),
	}
	inserter := table.Inserter() // Create an inserter for the table
	items := []*SampleData{
		&data,
		&data2,
	}
	// Insert the data into the table
	if err := inserter.Put(ctx, items); err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}
	fmt.Println("Data inserted successfully!")
}

func UpdateInBigQuery() {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, "gcp-learning-414814", option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatalf("Failed to create BigQuery client: %v", err)
	}
	defer client.Close()

	// Query to update the table using a JOIN with another table
	query := `
		UPDATE ` + "`gcp-learning-414814.rnd_query.sample_table`" + `
		SET value = 12312
		WHERE name = "shojib" AND TIMESTAMP(time) < TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 90 MINUTE)`

	// Run the update query
	q := client.Query(query)
	job, err := q.Run(ctx)
	if err != nil {
		log.Fatalf("Failed to run update query: %v", err)
	}

	// Wait for the query job to complete
	status, err := job.Wait(ctx)
	if err != nil {
		log.Fatalf("Failed to wait for query job: %v", err)
	}
	if err := status.Err(); err != nil {
		log.Fatalf("Query job failed: %v", err)
	}

	log.Println("Rows updated successfully!")
}

func DeleteInBigQuery() {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, "gcp-learning-414814", option.WithCredentialsFile(credentialsFile))
	if err != nil {
		log.Fatalf("Failed to create BigQuery client: %v", err)
	}
	defer client.Close()

	// Query to filter out the rows that should be deleted
	query := `
		DELETE FROM ` + "`gcp-learning-414814.rnd_query.sample_table`" + `
		WHERE name = "anik"
		AND TIMESTAMP(time) < TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 90 MINUTE)
	`

	// Run the delete (create/replace) query
	q := client.Query(query)
	job, err := q.Run(ctx)
	if err != nil {
		log.Fatalf("Failed to run delete query: %v", err)
	}

	// Wait for the query job to complete
	status, err := job.Wait(ctx)
	if err != nil {
		log.Fatalf("Failed to wait for query job: %v", err)
	}
	if err := status.Err(); err != nil {
		log.Fatalf("Query job failed: %v", err)
	}

	log.Println("Rows deleted successfully!")
}

func main() {
	// BatchPushToBigQuery()
	// BatchFetchToBigQuery()
	UpdateInBigQuery()
	DeleteInBigQuery()
}
