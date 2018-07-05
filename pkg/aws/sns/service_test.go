package sns

import (
	pkgAws "github.com/andream16/aws-sdk-go-bindings/pkg/aws"
	"github.com/andream16/aws-sdk-go-bindings/testdata"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSession_SnsPublish(t *testing.T) {

	cfg := testdata.MockConfiguration(t)

	svcIn, svcInErr := pkgAws.NewSessionInput(cfg.Region)

	assert.NoError(t, svcInErr)
	awsSvc, awsSvcErr := pkgAws.New(svcIn)

	assert.NoError(t, awsSvcErr)
	assert.NotEmpty(t, awsSvc)

	snsSvc, snsSvcErr := New(awsSvc)

	assert.NoError(t, snsSvcErr)
	assert.NotEmpty(t, snsSvc)

	in := &PublishInput{
		PublishInput: &sns.PublishInput{
			Message:          aws.String(`{"default":"{\"par1\":\"pr1\",\"par2\":\"pr2\"}"}`),
			TargetArn:        aws.String(cfg.SNS.TargetArn),
			MessageStructure: aws.String(`json`),
		},
	}

	err := snsSvc.SnsPublish(in)

	assert.NoError(t, err)

}
