package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type GrainFeed struct {
	SC_Commodity_ID           int
	SC_Commodity_Desc         string
	Amount                    float64
	SC_Geography_ID           int
	SC_GeographyIndented_Desc string
	Year_ID                   int
}

func main() {
	ctx := context.Background()
	c, err := bigquery.NewClient(ctx, "dl-gcp-cngo-sbox-dataenv-b1")
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer c.Close()

	DeleteUDFRoutine(&ctx, c)

}

func QueryingExample(ctx *context.Context, c *bigquery.Client) {

	query := c.Query(
		"SELECT SC_Commodity_ID, SC_Commodity_Desc, Amount, SC_Geography_ID, SC_GeographyIndented_Desc, Year_ID " +
			"FROM `dl-gcp-cngo-sbox-dataenv-b1.dl_gcp_cngo_sbox_dataenv_b1_data.feed-grains` " +
			"WHERE SC_Frequency_Desc = 'Monthly' AND Year_ID=2009 " +
			"ORDER BY Year_ID desc")

	rows, err := query.Read(*ctx)

	if err != nil {
		log.Fatal(err)
	}

	for {
		var values []bigquery.Value
		err := rows.Next(&values)

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(values)
	}
}

func QueryingExampleExported(ctx *context.Context, c *bigquery.Client) {

	query := c.Query(
		"SELECT SC_Commodity_ID, SC_Commodity_Desc, Amount, SC_Geography_ID, SC_GeographyIndented_Desc, Year_ID " +
			"FROM `dl-gcp-cngo-sbox-dataenv-b1.dl_gcp_cngo_sbox_dataenv_b1_data.feed-grains` " +
			"WHERE SC_Frequency_Desc = 'Monthly' AND Year_ID=2009 " +
			"ORDER BY Year_ID desc")

	rows, err := query.Read(*ctx)

	if err != nil {
		log.Fatal(err)
	}

	for {
		var gf GrainFeed
		err := rows.Next(&gf)

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(gf)
	}
}

func BrowseFeedGrainTable(ctx *context.Context, c *bigquery.Client) {
	myDataset := c.Dataset("dl_gcp_cngo_sbox_dataenv_b1_data")
	myGrainTable := myDataset.Table("feed-grains")

	rows := myGrainTable.Read(*ctx)
	delimiter := 0
	for {
		var gf GrainFeed
		err := rows.Next(&gf)
		if delimiter > 10 {
			break
		}

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(gf)
		delimiter++
	}
}

func CreateDataset(ctx *context.Context, c *bigquery.Client) {
	meta := &bigquery.DatasetMetadata{
		Location: "US",
	}
	if err := c.Dataset("go_course_data").Create(*ctx, meta); err != nil {
		log.Fatal(err)
	}
}

func CreateTable(ctx *context.Context, c *bigquery.Client) error {
	tableRef := c.Dataset("go_course_data").Table("players")

	//The second argument will be nil for our metadata
	if err := tableRef.Create(*ctx, nil); err != nil {
		return err
	}
	return nil
}

func JobExample(ctx *context.Context, c *bigquery.Client) {

	query := c.Query(
		"SELECT SC_Commodity_ID, SC_Commodity_Desc, Amount, SC_Geography_ID, SC_GeographyIndented_Desc, Year_ID " +
			"FROM `dl-gcp-cngo-sbox-dataenv-b1.dl_gcp_cngo_sbox_dataenv_b1_data.feed-grains` " +
			"WHERE SC_Frequency_Desc = 'Monthly' AND Year_ID=2009 " +
			"ORDER BY Year_ID desc")

	myJob, err := query.Run(*ctx)

	if err != nil {
		log.Fatal(err)
	}

	myJobID := myJob.ID()
	fmt.Printf("The ID of my job: %v\n", myJobID)

	rows, err := myJob.Read(*ctx)

	if err != nil {
		log.Fatal(err)
	}

	for {
		var gf GrainFeed
		err := rows.Next(&gf)

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(gf)

	}
}

func ImportJSONExampleWithExplicitSchema(ctx *context.Context, c *bigquery.Client) {

	gf := bigquery.NewGCSReference("gs://dl-gcp-cngo-sbox-dataenv-b1-data/battingupload.json")
	gf.SourceFormat = bigquery.JSON

	gf.Schema = bigquery.Schema{
		{Name: "playerID", Type: bigquery.StringFieldType},
		{Name: "yearID", Type: bigquery.IntegerFieldType},
		{Name: "stint", Type: bigquery.IntegerFieldType},
		{Name: "teamID", Type: bigquery.StringFieldType},
		{Name: "leagueID", Type: bigquery.StringFieldType},
		{Name: "G", Type: bigquery.IntegerFieldType},
		{Name: "AB", Type: bigquery.IntegerFieldType},
		{Name: "R", Type: bigquery.IntegerFieldType},
		{Name: "H", Type: bigquery.IntegerFieldType},
		{Name: "Doubles", Type: bigquery.IntegerFieldType},
		{Name: "Triples", Type: bigquery.IntegerFieldType},
		{Name: "HR", Type: bigquery.IntegerFieldType},
		{Name: "RBI", Type: bigquery.IntegerFieldType},
		{Name: "SB", Type: bigquery.IntegerFieldType},
		{Name: "CS", Type: bigquery.IntegerFieldType},
		{Name: "BB", Type: bigquery.IntegerFieldType},
		{Name: "SO", Type: bigquery.IntegerFieldType},
		{Name: "IBB", Type: bigquery.IntegerFieldType},
		{Name: "HBP", Type: bigquery.IntegerFieldType},
		{Name: "SH", Type: bigquery.IntegerFieldType},
		{Name: "SF", Type: bigquery.IntegerFieldType},
		{Name: "GIDP", Type: bigquery.IntegerFieldType},
	}

	loader := c.Dataset("go_course_data").Table("batting").LoaderFrom(gf)
	loader.WriteDisposition = bigquery.WriteEmpty

	myJob, err := loader.Run(*ctx)
	if err != nil {
		log.Fatal(err)
	}
	status, err := myJob.Wait(*ctx)
	if err != nil {
		log.Fatal(err)
	}
	if status.Err() != nil {
		log.Fatalf("job completed with error: %v", status.Err())
	}
}

func ImportCSVExampleWithExplicitSchema(ctx *context.Context, c *bigquery.Client) {
	gf := bigquery.NewGCSReference("gs://dl-gcp-cngo-sbox-dataenv-b1-data/Batting.csv")
	gf.SkipLeadingRows = 1
	gf.Schema = bigquery.Schema{
		{Name: "playerID", Type: bigquery.StringFieldType},
		{Name: "yearID", Type: bigquery.IntegerFieldType},
		{Name: "stint", Type: bigquery.IntegerFieldType},
		{Name: "teamID", Type: bigquery.StringFieldType},
		{Name: "leagueID", Type: bigquery.StringFieldType},
		{Name: "G", Type: bigquery.IntegerFieldType},
		{Name: "AB", Type: bigquery.IntegerFieldType},
		{Name: "R", Type: bigquery.IntegerFieldType},
		{Name: "H", Type: bigquery.IntegerFieldType},
		{Name: "Doubles", Type: bigquery.IntegerFieldType},
		{Name: "Triples", Type: bigquery.IntegerFieldType},
		{Name: "HR", Type: bigquery.IntegerFieldType},
		{Name: "RBI", Type: bigquery.IntegerFieldType},
		{Name: "SB", Type: bigquery.IntegerFieldType},
		{Name: "CS", Type: bigquery.IntegerFieldType},
		{Name: "BB", Type: bigquery.IntegerFieldType},
		{Name: "SO", Type: bigquery.IntegerFieldType},
		{Name: "IBB", Type: bigquery.IntegerFieldType},
		{Name: "HBP", Type: bigquery.IntegerFieldType},
		{Name: "SH", Type: bigquery.IntegerFieldType},
		{Name: "SF", Type: bigquery.IntegerFieldType},
		{Name: "GIDP", Type: bigquery.IntegerFieldType},
	}
	loader := c.Dataset("go_course_data").Table("batting").LoaderFrom(gf)
	loader.WriteDisposition = bigquery.WriteEmpty

	myJob, err := loader.Run(*ctx)
	if err != nil {
		log.Fatal(err)
	}
	status, err := myJob.Wait(*ctx)
	if err != nil {
		log.Fatal(err)
	}
	if status.Err() != nil {
		log.Fatalf("job completed with error: %v", status.Err())
	}
}

func CreateUDFRoutine(ctx *context.Context, c *bigquery.Client) {
	//`dl-gcp-cngo-sbox-dataenv-b1.dl_gcp_cngo_sbox_dataenv_b1_data.totalBases`
	metaData := &bigquery.RoutineMetadata{
		Type:     "SCALAR_FUNCTION",
		Language: "SQL",
		Body:     "(n +  x + (y * 2) + (z * 3))",
		Arguments: []*bigquery.RoutineArgument{
			{Name: "n", DataType: &bigquery.StandardSQLDataType{TypeKind: "INT64"}},
			{Name: "x", DataType: &bigquery.StandardSQLDataType{TypeKind: "INT64"}},
			{Name: "y", DataType: &bigquery.StandardSQLDataType{TypeKind: "INT64"}},
			{Name: "z", DataType: &bigquery.StandardSQLDataType{TypeKind: "INT64"}},
		},
	}

	routineRef := c.Dataset("go_course_data").Routine("totalBases")

	if err := routineRef.Create(*ctx, metaData); err != nil {
		log.Fatal(err)
	}
}

func DeleteUDFRoutine(ctx *context.Context, c *bigquery.Client) {
	err := c.Dataset("dl_gcp_cngo_sbox_dataenv_b1_data").Routine("totalBases").Delete(*ctx)

	if err != nil {
		log.Fatal(err)
	}

}

func UpdateUDFRoutine(ctx *context.Context, c *bigquery.Client) {
	routineRef := c.Dataset("go_course_data").Routine("totalBases")

	// We need to fetch the existing metadata
	meta, err := routineRef.Metadata(*ctx)
	if err != nil {
		log.Fatal(err)
	}

	// We must supply all arguments or we will get an error
	update := &bigquery.RoutineMetadataToUpdate{
		Type:        meta.Type,
		Language:    meta.Language,
		Arguments:   meta.Arguments,
		Description: "UDF to calculate total bases",
		ReturnType:  meta.ReturnType,
		Body:        meta.Body,
	}

	if _, err := routineRef.Update(*ctx, update, meta.ETag); err != nil {
		log.Fatal(err)
	}
}

// DO NOT RUN THIS, WE DON'T REALLY WANT TO COPY 400K DATA ROWS
// func CopyDatasetExample(ctx *context.Context, c *bigquery.Client) {
// 	desiredDataset := c.Dataset(""dl_gcp_cngo_sbox_dataenv_b1_data"")
// 	desiredSetInProject := c.DatasetInProject("dl-gcp-cngo-sbox-dataenv-b1", "dl_gcp_cngo_sbox_dataenv_b1_data")
// 	myGrainFeedTable := desiredSetInProject.Table("feed-grains")
// 	destinationTable := desiredSetInProject.Table("some-other-table")
// 	copier := destinationTable.CopierFrom(myGrainFeedTable)

// 	copier.WriteDisposition = bigquery.WriteTruncate

// 	job, err := copier.Run(*ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(desiredDataset, desiredSetInProject, job)

// 	gf := bigquery.NewGCSReference("gs://some-bucket/some-object")
// 	gf.AllowJaggedRows = true
// 	loader := myGrainFeedTable.LoaderFrom(gf)
// 	loader.CreateDisposition = bigquery.CreateNever
// 	job, err = loader.Run(*ctx)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	u := myGrainFeedTable.Inserter()
// 	grainFeedSlice := []*GrainFeed{
// 		{115191, "SomeItem", 2.122, 1, "USA", 2022},
// 		{11919, "SomeOtherItem", 0.12, 1, "USA", 2023},
// 	}

// 	if err := u.Put(*ctx, grainFeedSlice); err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(err, job)

// }
