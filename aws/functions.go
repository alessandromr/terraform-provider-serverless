package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	homedir "github.com/mitchellh/go-homedir"
	"io/ioutil"
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

// Takes the result of flatmap.Expand for an array of strings
// and returns a []*string
func expandStringList(configured []interface{}) []*string {
	vs := make([]*string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, aws.String(v.(string)))
		}
	}
	return vs
}

// Takes the result of schema.Set of strings and returns a []*string
func expandStringSet(configured *schema.Set) []*string {
	return expandStringList(configured.List())
}

func readEnvironmentVariables(ev map[string]interface{}) map[string]string {
	variables := make(map[string]string)
	for k, v := range ev {
		variables[k] = v.(string)
	}

	return variables
}

// loadFileContent returns contents of a file in a given path
func loadFileContent(v string) ([]byte, error) {
	filename, err := homedir.Expand(v)
	if err != nil {
		return nil, err
	}
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return fileContent, nil
}
