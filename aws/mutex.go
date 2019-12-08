package aws

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
)

// This is a global MutexKV for use within this plugin.
var awsMutexKV = mutexkv.NewMutexKV()
