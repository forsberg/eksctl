package cloudformation

import (
	"encoding/json"
)

// AWSRDSDBSecurityGroup_Ingress AWS CloudFormation Resource (AWS::RDS::DBSecurityGroup.Ingress)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html
type AWSRDSDBSecurityGroup_Ingress struct {

	// CIDRIP AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-cidrip
	CIDRIP *Value `json:"CIDRIP,omitempty"`

	// EC2SecurityGroupId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupid
	EC2SecurityGroupId *Value `json:"EC2SecurityGroupId,omitempty"`

	// EC2SecurityGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupname
	EC2SecurityGroupName *Value `json:"EC2SecurityGroupName,omitempty"`

	// EC2SecurityGroupOwnerId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-security-group-rule.html#cfn-rds-securitygroup-ec2securitygroupownerid
	EC2SecurityGroupOwnerId *Value `json:"EC2SecurityGroupOwnerId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSDBSecurityGroup_Ingress) AWSCloudFormationType() string {
	return "AWS::RDS::DBSecurityGroup.Ingress"
}

func (r *AWSRDSDBSecurityGroup_Ingress) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
