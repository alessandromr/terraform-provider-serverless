package aws

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func isResourceNotFoundError(err error) bool {
	_, ok := err.(*resource.NotFoundError)
	return ok
}

func isResourceTimeoutError(err error) bool {
	timeoutErr, ok := err.(*resource.TimeoutError)
	return ok && timeoutErr.LastError == nil
}
