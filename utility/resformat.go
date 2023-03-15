package utility

func CreateNullSafeResponse(incoming interface{}, outgoing interface{}) {
	// fields := reflect.VisibleFields(reflect.TypeOf(incoming))
	// values := reflect.ValueOf(incoming)
	// // metaValue := reflect.ValueOf(output).
	// for _, field := range fields {

	// 	switch field.Type.String() {
	// 	case "sql.NullInt16":
	// 		if(values.FieldByName(field.Name).IsValid()) {
	// 			values.FieldByName(field.Name).Int16
	// 		} else {

	// 		}
	// 	case "sql.NullInt32":

	// 	case  "sql.NullInt64":

	// 	case "sql.NullFloat64":
	// 		fmt.Println("Something else")
	// 	case "sql.NullBool":
	// 		fmt.Println("Null Bool")
	// 	case "sql.NullTime":
	// 		fmt.Println("Null time")
	// 	}
	// 	fmt.Println(field.Name, field.Type, values.FieldByName(field.Name))

	// }
}
