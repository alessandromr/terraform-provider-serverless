package aws

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

var validLambdaRuntimes = []string{
	lambda.RuntimeDotnetcore10,
	lambda.RuntimeDotnetcore20,
	lambda.RuntimeDotnetcore21,
	lambda.RuntimeGo1X,
	lambda.RuntimeJava8,
	lambda.RuntimeJava11,
	// lambda.RuntimeNodejs43, EOL
	// lambda.RuntimeNodejs43Edge, EOL
	// lambda.RuntimeNodejs610, EOL
	lambda.RuntimeNodejs810,
	lambda.RuntimeNodejs10X,
	lambda.RuntimeNodejs12X,
	lambda.RuntimeProvided,
	lambda.RuntimePython27,
	lambda.RuntimePython36,
	lambda.RuntimePython37,
	lambda.RuntimePython38,
	lambda.RuntimeRuby25,
}

var validS3Events = []string{
	"s3:ObjectCreated:*",
	"s3:ObjectCreated:Put",
	"s3:ObjectCreated:Post",
	"s3:ObjectCreated:Copy",
	"s3:ObjectCreated:CompleteMultipartUpload",
	"s3:ObjectRemoved:*",
	"s3:ObjectRemoved:Delete",
	"s3:ObjectRemoved:DeleteMarkerCreated",
	"s3:ObjectRestore:Post",
	"s3:ObjectRestore:Completed",
	"s3:ReducedRedundancyLostObject",
	"s3:Replication:OperationFailedReplication",
	"s3:Replication:OperationMissedThreshold",
	"s3:Replication:OperationReplicatedAfterThreshold",
	"s3:Replication:OperationNotTracked",
}

var validHTTPMethod = []string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"OPTION",
}

func ResourceFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceFunctionCreate,
		Read:   resourceFunctionRead,
		Update: resourceFunctionUpdate,
		Delete: resourceFunctionDelete,

		Schema: map[string]*schema.Schema{
			"function": {
				Type:     schema.TypeList,
				Optional: false,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filename": {
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"s3_bucket", "s3_key", "s3_object_version"},
						},
						"s3_bucket": {
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"filename"},
						},
						"s3_key": {
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"filename"},
						},
						"s3_object_version": {
							Type:          schema.TypeString,
							Optional:      true,
							ConflictsWith: []string{"filename"},
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"memory_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  128,
						},
						"runtime": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(validLambdaRuntimes, false),
						},
						"environment": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"variables": {
										Type:     schema.TypeMap,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"timeout": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  3,
						},
						"vpc_config": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subnet_ids": {
										Type:     schema.TypeSet,
										Required: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"security_group_ids": {
										Type:     schema.TypeSet,
										Required: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
										Set:      schema.HashString,
									},
									"vpc_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"function_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"handler": {
							Type:     schema.TypeString,
							Required: true,
						},
						"arn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_code_hash": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_code_size": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"qualified_arn": {}, //ToDo
						"invoke_arn":    {}, //ToDo
						// "role": {
						// },
						// "layers": {
						// 	Type:     schema.TypeList,
						// 	Optional: true,
						// 	MaxItems: 5,
						// 	Elem: &schema.Schema{
						// 		Type:         schema.TypeString,
						// 		ValidateFunc: validateArn,
						// 	},
						// },
					},
				},
			},
			"event": {
				Type:     schema.TypeList,
				Optional: false,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"s3": {
							Type:     schema.TypeList,
							Optional: false,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bucket": {
										Type:     schema.TypeString,
										Required: true,
									},
									"event_types": {
										Type:     schema.TypeList,
										Required: true,
										MinItems: 1,
										MaxItems: 10,
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringInSlice(validS3Events, false),
										},
									},
									"event_key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"object_prefix": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"object_suffix": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"bucket_domain_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"bucket_regional_domain_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"arn": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"http": {
							Type:     schema.TypeList,
							Optional: false,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path": {
										Type:     schema.TypeString,
										Required: true,
									},
									"http_method": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(validHTTPMethod, false),
									},
									"already_existing": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"apiId": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"apiName": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"executionRole": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"arn": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"root_resource_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"created_date": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"role": {
				Type:     schema.TypeList,
				Optional: false,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_policy": {},
					},
				},
			},
		},
	}
}

func resourceFunctionCreate(d *schema.ResourceData, m interface{}) error {
	return resourceFunctionRead(d, m)
}

func resourceFunctionRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFunctionUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceFunctionRead(d, m)
}

func resourceFunctionDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
