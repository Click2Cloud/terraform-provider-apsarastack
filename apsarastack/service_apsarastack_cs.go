package apsarastack

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aliyun/terraform-provider-apsarastack/apsarastack/connectivity"
	"github.com/denverdino/aliyungo/cs"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"log"
	"regexp"
	"strings"
	"time"
)

type CsService struct {
	client *connectivity.ApsaraStackClient
}

const (
	DefaultECSTag         = "k8s.aliyun.com"
	UpgradeClusterTimeout = 30 * time.Minute
)

var (
	ATTACH_SCRIPT_WITH_VERSION = `#!/bin/sh
curl http://aliacs-k8s-%s.oss-%s.aliyuncs.com/public/pkg/run/attach/%s/attach_node.sh | bash -s -- --openapi-token %s --ess true `
	NETWORK_ADDON_NAMES = []string{"terway", "kube-flannel-ds", "terway-eni", "terway-eniip"}
)

func (s *CsService) DescribeCsKubernetes(id string) (cluster *cs.KubernetesClusterDetail, err error) {
	invoker := NewInvoker()
	var requestInfo *cs.Client
	var response interface{}

	if err := invoker.Run(func() error {
		raw, err := s.client.WithCsClient(func(csClient *cs.Client) (interface{}, error) {
			requestInfo = csClient
			return csClient.DescribeKubernetesClusterDetail(id)
		})
		response = raw
		return err
	}); err != nil {
		if IsExpectedErrors(err, []string{"ErrorClusterNotFound"}) {
			return cluster, WrapErrorf(err, NotFoundMsg, DenverdinoApsaraStackgo)
		}
		return cluster, WrapErrorf(err, DefaultErrorMsg, id, "DescribeKubernetesCluster", DenverdinoApsaraStackgo)
	}
	if debugOn() {
		requestMap := make(map[string]interface{})
		requestMap["ClusterId"] = id
		addDebug("DescribeKubernetesCluster", response, requestInfo, requestMap)
	}
	cluster, _ = response.(*cs.KubernetesClusterDetail)
	if cluster.ClusterId != id {
		return cluster, WrapErrorf(Error(GetNotFoundMessage("CsKubernetes", id)), NotFoundMsg, ProviderERROR)
	}
	return
}

func (s *CsService) CsKubernetesInstanceStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeCsKubernetes(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if string(object.State) == failState {
				return object, string(object.State), WrapError(Error(FailedToReachTargetStatus, string(object.State)))
			}
		}
		return object, string(object.State), nil
	}
}

func (s *CsService) tagsToMap(tags []cs.Tag) map[string]string {
	result := make(map[string]string)
	for _, t := range tags {
		if !s.ignoreTag(t) {
			result[t.Key] = t.Value
		}
	}
	return result
}

func (s *CsService) ignoreTag(t cs.Tag) bool {
	filter := []string{"^http://", "^https://"}
	for _, v := range filter {
		log.Printf("[DEBUG] Matching prefix %v with %v\n", v, t.Key)
		ok, _ := regexp.MatchString(v, t.Key)
		if ok {
			log.Printf("[DEBUG] Found Apasarastack Cloud specific t %s (val: %s), ignoring.\n", t.Key, t.Value)
			return true
		}
	}
	return false
}

func (s *CsService) GetPermanentToken(clusterId string) (string, error) {

	describeClusterTokensResponse, err := s.client.WithCsClient(func(csClient *cs.Client) (interface{}, error) {
		return csClient.DescribeClusterTokens(clusterId)
	})
	if err != nil {
		return "", WrapError(fmt.Errorf("failed to get permanent token,because of %v", err))
	}

	tokens, ok := describeClusterTokensResponse.([]*cs.ClusterTokenResponse)

	if ok != true {
		return "", WrapError(fmt.Errorf("failed to parse ClusterTokenResponse of cluster %s", clusterId))
	}

	permanentTokens := make([]string, 0)

	for _, token := range tokens {
		if token.Expired == 0 && token.IsActive == 1 {
			permanentTokens = append(permanentTokens, token.Token)
			break
		}
	}

	// create a new token
	if len(permanentTokens) == 0 {
		createClusterTokenResponse, err := s.client.WithCsClient(func(csClient *cs.Client) (interface{}, error) {
			clusterTokenReqeust := &cs.ClusterTokenReqeust{}
			clusterTokenReqeust.IsPermanently = true
			return csClient.CreateClusterToken(clusterId, clusterTokenReqeust)
		})
		if err != nil {
			return "", WrapError(fmt.Errorf("failed to create permanent token,because of %v", err))
		}

		token, ok := createClusterTokenResponse.(*cs.ClusterTokenResponse)
		if ok != true {
			return "", WrapError(fmt.Errorf("failed to parse token of %s", clusterId))
		}
		return token.Token, nil
	}

	return permanentTokens[0], nil
}

// GetUserData of cluster
func (s *CsService) GetUserData(clusterId string, labels string, taints string) (string, error) {

	token, err := s.GetPermanentToken(clusterId)

	if err != nil {
		return "", err
	}

	if labels == "" {
		labels = fmt.Sprintf("%s=true", DefaultECSTag)
	} else {
		labels = fmt.Sprintf("%s,%s=true", labels, DefaultECSTag)
	}

	cluster, err := s.DescribeCsKubernetes(clusterId)

	if err != nil {
		return "", WrapError(fmt.Errorf("failed to describe cs kuberentes cluster,because of %v", err))
	}

	extra_options := make([]string, 0)

	if len(labels) > 0 || len(taints) > 0 {

		if len(labels) != 0 {
			extra_options = append(extra_options, fmt.Sprintf("--labels %s", labels))
		}

		if len(taints) != 0 {
			extra_options = append(extra_options, fmt.Sprintf("--taints %s", taints))
		}
	}

	if network, err := GetKubernetesNetworkName(cluster); err == nil && network != "" {
		extra_options = append(extra_options, fmt.Sprintf("--network %s", network))
	}

	extra_options_in_line := strings.Join(extra_options, " ")

	version := cluster.CurrentVersion
	region := cluster.RegionId

	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(ATTACH_SCRIPT_WITH_VERSION+extra_options_in_line, region, region, version, token))), nil
}

func (s *CsService) UpgradeCluster(clusterId string, args *cs.UpgradeClusterArgs) error {
	invoker := NewInvoker()
	err := invoker.Run(func() error {
		_, e := s.client.WithCsClient(func(csClient *cs.Client) (interface{}, error) {
			return nil, csClient.UpgradeCluster(clusterId, args)
		})
		if e != nil {
			return e
		}
		return nil
	})

	if err != nil {
		return WrapError(err)
	}

	state, upgradeError := s.WaitForUpgradeCluster(clusterId, "Upgrade")
	if state == cs.Task_Status_Success && upgradeError == nil {
		return nil
	}

	// if upgrade failed cancel the task
	err = invoker.Run(func() error {
		_, e := s.client.WithCsClient(func(csClient *cs.Client) (interface{}, error) {
			return nil, csClient.CancelUpgradeCluster(clusterId)
		})
		if e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		return WrapError(upgradeError)
	}

	if state, err := s.WaitForUpgradeCluster(clusterId, "CancelUpgrade"); err != nil || state != cs.Task_Status_Success {
		log.Printf("[WARN] %s ACK Cluster cancel upgrade error: %#v", clusterId, err)
	}

	return WrapError(upgradeError)
}

func (s *CsService) CsServerlessKubernetesInstanceStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeCsServerlessKubernetes(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if string(object.State) == failState {
				return object, string(object.State), WrapError(Error(FailedToReachTargetStatus, string(object.State)))
			}
		}
		return object, string(object.State), nil
	}
}

func (s *CsService) DescribeCsServerlessKubernetes(id string) (*cs.ServerlessClusterResponse, error) {
	cluster := &cs.ServerlessClusterResponse{}
	var requestInfo *cs.Client
	invoker := NewInvoker()
	var response interface{}

	if err := invoker.Run(func() error {
		raw, err := s.client.WithCsClient(func(csClient *cs.Client) (interface{}, error) {
			requestInfo = csClient
			return csClient.DescribeServerlessKubernetesCluster(id)
		})
		response = raw
		return err
	}); err != nil {
		if IsExpectedErrors(err, []string{"ErrorClusterNotFound"}) {
			return cluster, WrapErrorf(err, NotFoundMsg, DenverdinoAliyungo)
		}
		return cluster, WrapErrorf(err, DefaultErrorMsg, id, "DescribeServerlessKubernetesCluster", DenverdinoAliyungo)
	}
	if debugOn() {
		requestMap := make(map[string]interface{})
		requestMap["Id"] = id
		addDebug("DescribeServerlessKubernetesCluster", response, requestInfo, requestMap, map[string]interface{}{"Id": id})
	}
	cluster, _ = response.(*cs.ServerlessClusterResponse)
	if cluster != nil && cluster.ClusterId != id {
		return cluster, WrapErrorf(Error(GetNotFoundMessage("CSServerlessKubernetes", id)), NotFoundMsg, ProviderERROR)
	}
	return cluster, nil

}

func GetKubernetesNetworkName(cluster *cs.KubernetesClusterDetail) (network string, err error) {

	metadata := make(map[string]interface{})
	if err := json.Unmarshal([]byte(cluster.MetaData), &metadata); err != nil {
		return "", fmt.Errorf("unmarshal metaData failed. error: %s", err)
	}

	for _, name := range NETWORK_ADDON_NAMES {
		if _, ok := metadata[fmt.Sprintf("%s%s", name, "Version")]; ok {
			return name, nil
		}
	}
	return "", fmt.Errorf("no network addon found")
}

func (s *CsService) WaitForUpgradeCluster(clusterId string, action string) (string, error) {
	err := resource.Retry(UpgradeClusterTimeout, func() *resource.RetryError {
		resp, err := s.client.WithCsClient(func(csClient *cs.Client) (interface{}, error) {
			return csClient.QueryUpgradeClusterResult(clusterId)
		})
		if err != nil || resp == nil {
			return resource.RetryableError(err)
		}

		upgradeResult := resp.(*cs.UpgradeClusterResult)
		if upgradeResult.UpgradeStep == cs.UpgradeStep_Success {
			return nil
		}

		if upgradeResult.UpgradeStep == cs.UpgradeStep_Pause && upgradeResult.UpgradeStatus.Failed == "true" {
			msg := ""
			events := upgradeResult.UpgradeStatus.Events
			if len(events) > 0 {
				msg = events[len(events)-1].Message
			}
			return resource.NonRetryableError(fmt.Errorf("faild to %s cluster, error: %s", action, msg))
		}
		return resource.RetryableError(fmt.Errorf("%s cluster state not matched", action))
	})

	if err == nil {
		log.Printf("[INFO] %s ACK Cluster %s successed", action, clusterId)
		return cs.Task_Status_Success, nil
	}

	return cs.Task_Status_Failed, WrapError(err)
}
