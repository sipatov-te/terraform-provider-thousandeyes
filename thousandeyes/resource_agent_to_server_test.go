package thousandeyes

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/thousandeyes/thousandeyes-sdk-go/v3/tests"	
)

func TestAccThousandEyesAgentToServer(t *testing.T) {
	var resourceName = "thousandeyes_agent_to_server.test"
	var testCases = []struct {
		name                 string
		resourceFile         string
		resourceName         string
		checkDestroyFunction func(*terraform.State) error
		checkFunc            []resource.TestCheckFunc
	}{
		{
			name:                 "basic",
			resourceFile:         "acceptance_resources/agent_to_server/basic.tf",
			resourceName:         resourceName,
			checkDestroyFunction: testAccCheckAgentToServerResourceDestroy,
			checkFunc: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr(resourceName, "test_name", "User Acceptance Test - Agent To Server"),
				resource.TestCheckResourceAttr(resourceName, "server", "api.stg.thousandeyes.com"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "tcp"),
				resource.TestCheckResourceAttr(resourceName, "port", "443"),
				resource.TestCheckResourceAttr(resourceName, "interval", "120"),
				resource.TestCheckResourceAttr(resourceName, "probe_mode", "sack"),
				resource.TestCheckResourceAttr(resourceName, "alerts_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "alert_rules.#", "2"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:          func() { testAccPreCheck(t) },
				ProviderFactories: providerFactories,
				CheckDestroy:      tc.checkDestroyFunction,
				Steps: []resource.TestStep{
					{
						Config: testAccThousandEyesAgentToServerConfig(tc.resourceFile),
						Check:  resource.ComposeTestCheckFunc(tc.checkFunc...),
					},
				},
			})
		})
	}
}

func testAccCheckAgentToServerResourceDestroy(s *terraform.State) error {
	resourceList := []ResourceType{
		{
			ResourceName: "thousandeyes_agent_to_server",
			GetResource: func(id string) (interface{}, error) {
				return getAgentToServer(id)
			}},
	}
	return testAccCheckResourceDestroy(resourceList, s)
}

func testAccThousandEyesAgentToServerConfig(testResource string) string {
	content, err := os.ReadFile(testResource)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func getAgentToServer(id string) (interface{}, error) {
	api := (*tests.AgentToServerTestsAPIService)(&testClient.Common)
	req := api.GetAgentToServerTest(id).Expand(tests.AllowedExpandTestOptionsEnumValues)
	req = SetAidFromContext(testClient.GetConfig().Context, req)
	resp, _, err := req.Execute()
	return resp, err
}
