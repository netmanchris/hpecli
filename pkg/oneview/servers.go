// (C) Copyright 2019 Hewlett Packard Enterprise Development LP.

package oneview

import (
	"encoding/json"
	"fmt"

	"github.com/HewlettPackard/hpecli/pkg/logger"
	"github.com/spf13/cobra"
)

var ovServersData struct {
	name string
}

// login represents the oneview login command
var serversCmd = &cobra.Command{
	Use:   "servers",
	Short: "Get servers from OneView: hpecli oneview get servers",
	RunE:  getServers,
}

func init() {
	ovGetCmd.AddCommand(serversCmd)
	serversCmd.Flags().StringVar(&ovServersData.name, "name", "", "name of the server to retrieve")
}

func getServers(_ *cobra.Command, _ []string) error {
	return getServerHardware()
}

func getServerHardware() error {
	d, err := getContext()
	if err != nil {
		logger.Debug("unable to retrieve apiKey because of: %#v", err)
		return fmt.Errorf("unable to retrieve the last login for OneView." +
			"Please login to OneView using: hpecli oneview login")
	}

	ovc := NewOVClientFromAPIKey(d.Host, d.APIKey)

	logger.Always("Retrieving data from: %s", d.Host)

	var sh interface{}
	if ovServersData.name != "" {
		sh, err = ovc.GetServerHardwareByName(ovServersData.name)
	} else {
		sh, err = ovc.GetServerHardwareList(nil, "", "", "", "")
	}

	if err != nil {
		logger.Warning("Unable to login with supplied credentials to OneView at: %s", d.Host)
		return err
	}

	out, err := json.MarshalIndent(sh, "", "  ")
	if err != nil {
		logger.Warning("Unable to output data as JSON.  Please try the command again.")
	}

	logger.Always("%s", out)

	return nil
}
