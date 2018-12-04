package configcommands

import (
	"errors"

	"github.com/pivotalservices/cf-mgmt/config"
)

type OrgNamedQuotaConfigurationCommand struct {
	ConfigManager config.Manager
	BaseConfigCommand
	Name  string        `long:"name" description:"Name of quota" required:"true"`
	Quota NamedOrgQuota `group:"quota"`
}

//Execute - updates named quota
func (c *OrgNamedQuotaConfigurationCommand) Execute(args []string) error {
	c.initConfig()
	orgQuota, err := c.ConfigManager.GetOrgQuota(c.Name)
	if err != nil {
		return err
	}
	if orgQuota == nil {
		orgQuota = &config.OrgQuota{
			Name:                    c.Name,
			TotalPrivateDomains:     "unlimited",
			TotalReservedRoutePorts: "unlimited",
			TotalServiceKeys:        "unlimited",
			AppInstanceLimit:        "unlimited",
			AppTaskLimit:            "unlimited",
			InstanceMemoryLimit:     "unlimited",
			TotalRoutes:             "unlimited",
			TotalServices:           "unlimited",
			PaidServicePlansAllowed: true,
		}
	}
	errorString := ""
	updateOrgNamedQuotaConfig(orgQuota, c.Quota, &errorString)

	if errorString != "" {
		return errors.New(errorString)
	}
	return c.ConfigManager.SaveOrgQuota(orgQuota)
}

func updateOrgNamedQuotaConfig(namedOrgQuota *config.OrgQuota, orgQuota NamedOrgQuota, errorString *string) {
	convertToGB("memory-limit", &namedOrgQuota.MemoryLimit, orgQuota.MemoryLimit, errorString)
	convertToGB("instance-memory-limit", &namedOrgQuota.InstanceMemoryLimit, orgQuota.InstanceMemoryLimit, errorString)
	convertToFormattedInt("total-routes", &namedOrgQuota.TotalRoutes, orgQuota.TotalRoutes, errorString)
	convertToFormattedInt("total-services", &namedOrgQuota.TotalServices, orgQuota.TotalServices, errorString)
	convertToBool("paid-service-plans-allowed", &namedOrgQuota.PaidServicePlansAllowed, orgQuota.PaidServicesAllowed, errorString)
	convertToFormattedInt("total-private-domains", &namedOrgQuota.TotalPrivateDomains, orgQuota.TotalPrivateDomains, errorString)
	convertToFormattedInt("total-reserved-route-ports", &namedOrgQuota.TotalReservedRoutePorts, orgQuota.TotalReservedRoutePorts, errorString)
	convertToFormattedInt("total-service-keys", &namedOrgQuota.TotalServiceKeys, orgQuota.TotalServiceKeys, errorString)
	convertToFormattedInt("app-instance-limit", &namedOrgQuota.AppInstanceLimit, orgQuota.AppInstanceLimit, errorString)
	convertToFormattedInt("app-task-limit", &namedOrgQuota.AppTaskLimit, orgQuota.AppTaskLimit, errorString)
}

func (c *OrgNamedQuotaConfigurationCommand) initConfig() {
	if c.ConfigManager == nil {
		c.ConfigManager = config.NewManager(c.ConfigDirectory)
	}
}
