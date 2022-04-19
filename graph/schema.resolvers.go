package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/costa-neto/go-api.git/graph/generated"
	"github.com/costa-neto/go-api.git/graph/model"
)

func (r *mutationResolver) Search(ctx context.Context, text string) (*model.Response, error) {
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var number []rune
	var auxNumber []rune
	totalValue, aux := 0, 0
	for _, char := range text {
		if romanValues[char] > 0 {
			aux += romanValues[char]
			auxNumber = append(auxNumber, char)
		} else {
			if aux > totalValue {
				totalValue = aux
				number = auxNumber
			}
			aux = 0
			auxNumber = nil
		}
	}
	if aux > totalValue {
		totalValue = aux
		number = auxNumber
	}

	result := &model.Response{
		Number: string(number),
		Value:  totalValue,
	}
	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
