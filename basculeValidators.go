package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/xmidt-org/bascule"
	"github.com/xmidt-org/webpa-common/v2/basculechecks"
)

var requirePartnersJWTClaim bascule.ValidatorFunc = func(_ context.Context, token bascule.Token) error {
	partnerVal, ok := bascule.GetNestedAttribute(token.Attributes(), basculechecks.PartnerKeys()...)
	if !ok {
		return fmt.Errorf("Partner IDs not found at keys %v", basculechecks.PartnerKeys())
	}
	ids, ok := partnerVal.([]string)
	if !ok || len(ids) < 1 {
		return errors.New("Partner ID JWT claim should be a non-empty list of strings")
	}
	return nil
}
