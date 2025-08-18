package service

import (
	"context"
	"encoding/json"
	"fmt"
	"morpher-controller/common"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
)

var regionToLocation = map[string]string{
	"ap-northeast-2": "Asia Pacific (Seoul)",
	"ap-northeast-1": "Asia Pacific (Tokyo)",
	"us-east-1":      "US East (N. Virginia)",
	"us-west-2":      "US West (Oregon)",
	"eu-central-1":   "EU (Frankfurt)",
}

type GetEc2AwsPriceService struct {
}

func NewGetEc2AwsPriceService() *GetEc2AwsPriceService {
	return &GetEc2AwsPriceService{}
}

func (service *GetEc2AwsPriceService) GetEc2Price(
	query GetEc2PriceQuery,
) ([]EC2Instance, error) {
	ctx := context.Background()

	location, ok := regionToLocation[query.Region]
	if !ok {
		return nil, fmt.Errorf("Location mapping not found for region: %s", query.Region)
	}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		return nil, err
	}

	svc := pricing.NewFromConfig(cfg)
	filters := createFilter(location)

	out, err := svc.GetProducts(ctx, &pricing.GetProductsInput{
		ServiceCode:   common.AwsStr("AmazonEC2"),
		Filters:       filters,
		FormatVersion: common.AwsStr("aws_v1"),
		MaxResults:    common.AwsI32(10),
	})
	if err != nil {
		return nil, err
	}
	var instances []EC2Instance
	for _, priceData := range out.PriceList {
		var item priceItem
		if err := json.Unmarshal([]byte(priceData), &item); err != nil {
			continue
		}

		attrs := item.Product.Attributes

		vcpuStr := attrs["vcpu"]
		memoryStr := attrs["memory"]

		if !checkSpecsMatch(vcpuStr, memoryStr, query.MinVCpu, query.MinMemoryGB) {
			continue
		}

		pricePerHour := ""
		description := ""
		for _, term := range item.Terms.OnDemand {
			for _, dim := range term.PriceDimensions {
				pricePerHour = dim.PricePerUnit["USD"]
				description = dim.Description
				break
			}
			if pricePerHour != "" {
				break
			}
		}

		instance := EC2Instance{
			InstanceType:    attrs["instanceType"],
			VCpu:            vcpuStr,
			Memory:          memoryStr,
			Storage:         attrs["storage"],
			NetworkPerf:     attrs["networkPerformance"],
			PricePerHour:    pricePerHour,
			Description:     description,
			Location:        attrs["location"],
			OperatingSystem: attrs["operatingSystem"],
		}

		instances = append(instances, instance)
	}
	return instances, nil
}

func createFilter(location string) []types.Filter {
	return []types.Filter{
		{Type: types.FilterTypeTermMatch, Field: common.AwsStr("ServiceCode"), Value: common.AwsStr("AmazonEC2")},
		{Type: types.FilterTypeTermMatch, Field: common.AwsStr("location"), Value: common.AwsStr(location)},
		{Type: types.FilterTypeTermMatch, Field: common.AwsStr("operatingSystem"), Value: common.AwsStr("linux")},
		{Type: types.FilterTypeTermMatch, Field: common.AwsStr("tenancy"), Value: common.AwsStr("Shared")},
		{Type: types.FilterTypeTermMatch, Field: common.AwsStr("capacitystatus"), Value: common.AwsStr("Used")},
		{Type: types.FilterTypeTermMatch, Field: common.AwsStr("preInstalledSw"), Value: common.AwsStr("NA")},
		{Type: types.FilterTypeTermMatch, Field: common.AwsStr("termType"), Value: common.AwsStr("OnDemand")},
	}
}

func checkSpecsMatch(vcpuStr, memoryStr string, minVCpu int, minMemoryGB float64) bool {
	vcpu, err := strconv.Atoi(vcpuStr)
	if err != nil {
		return false
	}
	if vcpu < minVCpu {
		return false
	}

	memoryGB, err := parseMemoryGB(memoryStr)
	if err != nil {
		return false
	}
	if memoryGB < minMemoryGB {
		return false
	}

	return true
}

func parseMemoryGB(memoryStr string) (float64, error) {
	parts := strings.Fields(memoryStr)
	if len(parts) < 2 {
		return 0, fmt.Errorf("invalid memory format: %s", memoryStr)
	}

	return strconv.ParseFloat(parts[0], 64)
}

type priceItem struct {
	Product struct {
		Attributes map[string]string `json:"attributes"`
	} `json:"product"`
	Terms struct {
		OnDemand map[string]struct {
			PriceDimensions map[string]struct {
				Unit         string            `json:"unit"`
				PricePerUnit map[string]string `json:"pricePerUnit"`
				Description  string            `json:"description"`
			} `json:"priceDimensions"`
		} `json:"OnDemand"`
		Reserved map[string]any `json:"Reserved"`
	} `json:"terms"`
}
