package main

func main() {

	//mappingTpl := `{
	//"mappings":{
	// "properties":{
	//  "Uid":     { "type": "integer" },
	//  "Username":   { "type": "keyword" },
	//  "Status":   { "type": "keyword" },
	//  "Email":   { "type": "keyword" },
	//  "Password":    { "type": "keyword" },
	//  "City":  { "type": "keyword" },
	//  "State":         { "type": "keyword" },
	//  "Country":  { "type": "keyword" },
	//  "Profile":  { "type": "keyword" }
	//  }
	// }
	//}`
	//err := CreateIndex(mappingTpl)
	//if err != nil {
	//	fmt.Println(err)
	//}
	client := NewEsClient()
	AddDocument(client)
}
