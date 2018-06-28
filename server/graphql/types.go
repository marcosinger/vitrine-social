package graphql

import (
	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/graphql-go/graphql"
)

var (
	dateInput = &graphql.InputObjectFieldConfig{
		Type: Date,
	}

	intListArgument = &graphql.ArgumentConfig{
		Type: graphql.NewList(graphql.Int),
	}

	intArgument = &graphql.ArgumentConfig{
		Type: graphql.Int,
	}

	stringField = &graphql.Field{
		Type: graphql.String,
	}

	intField = &graphql.Field{
		Type: graphql.Int,
	}

	nonNullIntArgument = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.Int),
	}

	nonNullIntInput = &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.Int),
	}

	intInput = &graphql.InputObjectFieldConfig{
		Type: graphql.Int,
	}

	nonNullIntField = &graphql.Field{
		Type: graphql.NewNonNull(graphql.Int),
	}

	stringArgument = &graphql.ArgumentConfig{
		Type: graphql.String,
	}

	nonNullStringInput = &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.String),
	}

	stringInput = &graphql.InputObjectFieldConfig{
		Type: graphql.String,
	}

	nonNullStringArgument = &graphql.ArgumentConfig{
		Type: graphql.NewNonNull(graphql.String),
	}

	nonNullStringField = &graphql.Field{
		Type: graphql.NewNonNull(graphql.String),
	}

	organizationType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Organization",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if o, ok := p.Source.(*model.Organization); ok && o != nil {
						return o.ID, nil
					}
					return nil, nil
				},
			},
			"name": nonNullStringField,
			"logo": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if o, ok := p.Source.(*model.Organization); ok && o != nil && o.Logo != nil {
						return o.Logo.URL, nil
					}
					return nil, nil
				},
			},
			"slug":  nonNullStringField,
			"phone": nonNullStringField,
			"about": nonNullStringField,
			"video": nonNullStringField,
			"email": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if o, ok := p.Source.(*model.Organization); ok && o != nil {
						return o.Email, nil
					}
					return nil, nil
				},
			},
			"address": &graphql.Field{
				Type: graphql.NewNonNull(addressType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if o, ok := p.Source.(*model.Organization); ok && o != nil {
						return o.Address, nil
					}
					return nil, nil
				},
			},
		},
	})

	addressType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Address",
		Fields: graphql.Fields{
			"street":        stringField,
			"number":        stringField,
			"complement":    stringField,
			"neighbordhood": stringField,
			"city":          stringField,
			"state":         stringField,
			"zipcode":       stringField,
		},
	})

	needStatusEnum = graphql.NewEnum(graphql.EnumConfig{
		Name:        "NeedStatus",
		Description: "Status of a Need",
		Values: graphql.EnumValueConfigMap{
			"ACTIVE": &graphql.EnumValueConfig{
				Value:       model.NeedStatusActive,
				Description: "A active Need",
			},
			"INACTIVE": &graphql.EnumValueConfig{
				Value:       model.NeedStatusInactive,
				Description: "A inactive Need",
			},
		},
	})

	needStatusInput = &graphql.InputObjectFieldConfig{
		Type: needStatusEnum,
	}

	needType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Need",
		Fields: graphql.Fields{
			"id":               nonNullIntField,
			"title":            nonNullStringField,
			"description":      nonNullStringField,
			"requiredQuantity": intField,
			"reachedQuantity":  intField,
			"unit":             stringField,
			"dueDate": &graphql.Field{
				Type: Date,
			},
			"status": &graphql.Field{
				Type: needStatusEnum,
			},
			"createdAt": &graphql.Field{
				Type: graphql.NewNonNull(DateTime),
			},
			"updatedAt": &graphql.Field{
				Type: DateTime,
			},
		},
	})

	categoryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id":   nonNullIntField,
			"name": nonNullStringField,
			"slug": nonNullStringField,
		},
	})
)