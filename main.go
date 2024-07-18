package main

import (
	ex_graphql "go-example/graphql"
)

// func main() {
// 	h := handler.New(&handler.Config{
// 		Schema:     &testutil.StarWarsSchema,
// 		Pretty:     true,
// 		GraphiQL:   true,
// 		Playground: false,
// 	})

// 	http.Handle("/graphql", h)

// 	log.Println("GraphQL Server running on [POST]: localhost:8080/graphql")
// 	log.Println("GraphQL Playground running on [GET]: localhost:8080/graphql")

// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	ex_graphql.RunningCustomScalarType()
	// ex_graphql.RunningHelloWorld()

	// schema, err := graphql.NewSchema(graphql.SchemaConfig{
	// 	Query: ex_graphql.QueryType,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// http.Handle("/foo", handler.New(
	// 	&handler.Config{
	// 		Schema:     &schema,
	// 		Pretty:     true,
	// 		GraphiQL:   true,
	// 		Playground: false,
	// 	},
	// ))

	// http.Handle("/product", ex_graphql.HandlerProduct())

	// // http.HandleFunc("/product", ex_graphql.HandlerProduct)
	// http.ListenAndServe(":8080", nil)

	// query := `
	// 	query {
	// 		concurrentFieldFoo {
	// 			name
	// 		}
	// 		concurrentFieldBar {
	// 			name
	// 		}
	// 	}
	// `
	//
	// result := graphql.Do(graphql.Params{
	// 	RequestString: query,
	// 	Schema:        schema,
	// })
	// b, err := json.Marshal(result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", b)
	/*
		{
			 "data": {
				"concurrentFieldBar": {
					"name": "Bar's name"
				},
				"concurrentFieldFoo": {
					"name": "Foo's name"
				}
			}
		}
	*/
}
