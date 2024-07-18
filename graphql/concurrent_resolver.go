package ex_graphql

import "github.com/graphql-go/graphql"

type Foo struct{ Name string }

var FieldFooType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Foo",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

type Bar struct{ Name string }

var FieldBarType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Bar",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"concurrentFieldFoo": &graphql.Field{
				Type: FieldFooType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					foo := Foo{Name: "Foo's name"}
					return func() (interface{}, error) {
						return &foo, nil
					}, nil
				},
			},
			"concurrentFieldBar": &graphql.Field{
				Type: FieldBarType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					type result struct {
						data interface{}
						err  error
					}
					ch := make(chan *result, 1)
					go func() {
						defer close(ch)
						bar := &Bar{Name: "Bar's name"}
						ch <- &result{data: bar, err: nil}
					}()
					return func() (interface{}, error) {
						r := <-ch
						return r.data, r.err
					}, nil
				},
			},
		},
	})
