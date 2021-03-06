/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type ClustersAwsAttributes struct {
	FirstOnDemand int32 `json:"first_on_demand,omitempty"`

	Availability *ClustersAwsAvailability `json:"availability,omitempty"`

	ZoneId string `json:"zone_id,omitempty"`

	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`

	SpotBidPricePercent int32 `json:"spot_bid_price_percent,omitempty"`

	EbsVolumeType *ClustersEbsVolumeType `json:"ebs_volume_type,omitempty"`

	EbsVolumeCount int32 `json:"ebs_volume_count,omitempty"`

	EbsVolumeSize int32 `json:"ebs_volume_size,omitempty"`
}
