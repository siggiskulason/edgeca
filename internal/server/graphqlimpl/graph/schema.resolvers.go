package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/edgesec-org/edgeca/internal/issuer"
	"github.com/edgesec-org/edgeca/internal/server/graphqlimpl/graph/generated"
	"github.com/edgesec-org/edgeca/internal/server/graphqlimpl/graph/model"
	"github.com/edgesec-org/edgeca/internal/state"
)

func (r *mutationResolver) CreateCertificate(ctx context.Context, input model.NewCertificate) (*model.Certificate, error) {
	certificate, key, expiryStr, err := issuer.GenerateCertificateUsingX509SubjectOptionalValues(input.CommonName,
		input.Organization, input.OrganizationalUnit, input.Locality, input.Province, input.Country,
		state.GetSubCACert(), state.GetSubCAKey())

	var cert model.Certificate
	cert.Certificate = string(certificate)
	cert.Expiry = expiryStr
	cert.Key = string(key)
	return &cert, err
}

func (r *queryResolver) Certificate(ctx context.Context) ([]*model.Certificate, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
