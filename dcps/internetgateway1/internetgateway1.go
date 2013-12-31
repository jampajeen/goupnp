package internetgateway1

import (
	"time"

	"github.com/huin/goupnp/soap"
)

// Hack to avoid Go complaining if time isn't used.
var _ time.Time

// Device URNs:
const (
	URN_LANDevice_1           = "urn:schemas-upnp-org:device:LANDevice:1"
	URN_WANConnectionDevice_1 = "urn:schemas-upnp-org:device:WANConnectionDevice:1"
	URN_WANDevice_1           = "urn:schemas-upnp-org:device:WANDevice:1"
)

// Service URNs:
const (
	URN_LANHostConfigManagement_1  = "urn:schemas-upnp-org:service:LANHostConfigManagement:1"
	URN_Layer3Forwarding_1         = "urn:schemas-upnp-org:service:Layer3Forwarding:1"
	URN_WANCableLinkConfig_1       = "urn:schemas-upnp-org:service:WANCableLinkConfig:1"
	URN_WANCommonInterfaceConfig_1 = "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1"
	URN_WANDSLLinkConfig_1         = "urn:schemas-upnp-org:service:WANDSLLinkConfig:1"
	URN_WANEthernetLinkConfig_1    = "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1"
	URN_WANIPConnection_1          = "urn:schemas-upnp-org:service:WANIPConnection:1"
	URN_WANPOTSLinkConfig_1        = "urn:schemas-upnp-org:service:WANPOTSLinkConfig:1"
	URN_WANPPPConnection_1         = "urn:schemas-upnp-org:service:WANPPPConnection:1"
)

// LANHostConfigManagement1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:LANHostConfigManagement:1".
type LANHostConfigManagement1 struct {
	SOAPClient soap.SOAPClient
}

// Arguments:
//
// * NewDHCPServerConfigurable:
//
//
func (client *LANHostConfigManagement1) SetDHCPServerConfigurable(NewDHCPServerConfigurable bool) (err error) {
	// Request structure.
	var request struct {
		NewDHCPServerConfigurable string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDHCPServerConfigurable, err = soap.MarshalBoolean(NewDHCPServerConfigurable); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPServerConfigurable", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDHCPServerConfigurable:
func (client *LANHostConfigManagement1) GetDHCPServerConfigurable() (NewDHCPServerConfigurable bool, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDHCPServerConfigurable string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDHCPServerConfigurable", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDHCPServerConfigurable, err = soap.UnmarshalBoolean(response.NewDHCPServerConfigurable); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewDHCPRelay:
//
//
func (client *LANHostConfigManagement1) SetDHCPRelay(NewDHCPRelay bool) (err error) {
	// Request structure.
	var request struct {
		NewDHCPRelay string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDHCPRelay, err = soap.MarshalBoolean(NewDHCPRelay); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDHCPRelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDHCPRelay:
func (client *LANHostConfigManagement1) GetDHCPRelay() (NewDHCPRelay bool, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDHCPRelay string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDHCPRelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDHCPRelay, err = soap.UnmarshalBoolean(response.NewDHCPRelay); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewSubnetMask:
//
//
func (client *LANHostConfigManagement1) SetSubnetMask(NewSubnetMask string) (err error) {
	// Request structure.
	var request struct {
		NewSubnetMask string
	}
	// BEGIN Marshal arguments into request.

	if request.NewSubnetMask, err = soap.MarshalString(NewSubnetMask); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetSubnetMask", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewSubnetMask:
func (client *LANHostConfigManagement1) GetSubnetMask() (NewSubnetMask string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewSubnetMask string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetSubnetMask", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewSubnetMask, err = soap.UnmarshalString(response.NewSubnetMask); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewIPRouters:
//
//
func (client *LANHostConfigManagement1) SetIPRouter(NewIPRouters string) (err error) {
	// Request structure.
	var request struct {
		NewIPRouters string
	}
	// BEGIN Marshal arguments into request.

	if request.NewIPRouters, err = soap.MarshalString(NewIPRouters); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetIPRouter", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewIPRouters:
//
//
func (client *LANHostConfigManagement1) DeleteIPRouter(NewIPRouters string) (err error) {
	// Request structure.
	var request struct {
		NewIPRouters string
	}
	// BEGIN Marshal arguments into request.

	if request.NewIPRouters, err = soap.MarshalString(NewIPRouters); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteIPRouter", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewIPRouters:
func (client *LANHostConfigManagement1) GetIPRoutersList() (NewIPRouters string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewIPRouters string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetIPRoutersList", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewIPRouters, err = soap.UnmarshalString(response.NewIPRouters); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewDomainName:
//
//
func (client *LANHostConfigManagement1) SetDomainName(NewDomainName string) (err error) {
	// Request structure.
	var request struct {
		NewDomainName string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDomainName, err = soap.MarshalString(NewDomainName); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDomainName", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDomainName:
func (client *LANHostConfigManagement1) GetDomainName() (NewDomainName string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDomainName string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDomainName", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDomainName, err = soap.UnmarshalString(response.NewDomainName); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewMinAddress:
//
// * NewMaxAddress:
//
//
func (client *LANHostConfigManagement1) SetAddressRange(NewMinAddress string, NewMaxAddress string) (err error) {
	// Request structure.
	var request struct {
		NewMinAddress string
		NewMaxAddress string
	}
	// BEGIN Marshal arguments into request.

	if request.NewMinAddress, err = soap.MarshalString(NewMinAddress); err != nil {
		return
	}
	if request.NewMaxAddress, err = soap.MarshalString(NewMaxAddress); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetAddressRange", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewMinAddress:
//
// * NewMaxAddress:
func (client *LANHostConfigManagement1) GetAddressRange() (NewMinAddress string, NewMaxAddress string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewMinAddress string
		NewMaxAddress string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetAddressRange", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewMinAddress, err = soap.UnmarshalString(response.NewMinAddress); err != nil {
		return
	}
	if NewMaxAddress, err = soap.UnmarshalString(response.NewMaxAddress); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewReservedAddresses:
//
//
func (client *LANHostConfigManagement1) SetReservedAddress(NewReservedAddresses string) (err error) {
	// Request structure.
	var request struct {
		NewReservedAddresses string
	}
	// BEGIN Marshal arguments into request.

	if request.NewReservedAddresses, err = soap.MarshalString(NewReservedAddresses); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetReservedAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewReservedAddresses:
//
//
func (client *LANHostConfigManagement1) DeleteReservedAddress(NewReservedAddresses string) (err error) {
	// Request structure.
	var request struct {
		NewReservedAddresses string
	}
	// BEGIN Marshal arguments into request.

	if request.NewReservedAddresses, err = soap.MarshalString(NewReservedAddresses); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteReservedAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewReservedAddresses:
func (client *LANHostConfigManagement1) GetReservedAddresses() (NewReservedAddresses string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewReservedAddresses string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetReservedAddresses", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewReservedAddresses, err = soap.UnmarshalString(response.NewReservedAddresses); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewDNSServers:
//
//
func (client *LANHostConfigManagement1) SetDNSServer(NewDNSServers string) (err error) {
	// Request structure.
	var request struct {
		NewDNSServers string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDNSServers, err = soap.MarshalString(NewDNSServers); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "SetDNSServer", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewDNSServers:
//
//
func (client *LANHostConfigManagement1) DeleteDNSServer(NewDNSServers string) (err error) {
	// Request structure.
	var request struct {
		NewDNSServers string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDNSServers, err = soap.MarshalString(NewDNSServers); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "DeleteDNSServer", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDNSServers:
func (client *LANHostConfigManagement1) GetDNSServers() (NewDNSServers string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDNSServers string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_LANHostConfigManagement_1, "GetDNSServers", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDNSServers, err = soap.UnmarshalString(response.NewDNSServers); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Layer3Forwarding1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:Layer3Forwarding:1".
type Layer3Forwarding1 struct {
	SOAPClient soap.SOAPClient
}

// Arguments:
//
// * NewDefaultConnectionService:
//
//
func (client *Layer3Forwarding1) SetDefaultConnectionService(NewDefaultConnectionService string) (err error) {
	// Request structure.
	var request struct {
		NewDefaultConnectionService string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDefaultConnectionService, err = soap.MarshalString(NewDefaultConnectionService); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "SetDefaultConnectionService", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDefaultConnectionService:
func (client *Layer3Forwarding1) GetDefaultConnectionService() (NewDefaultConnectionService string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDefaultConnectionService string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_Layer3Forwarding_1, "GetDefaultConnectionService", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDefaultConnectionService, err = soap.UnmarshalString(response.NewDefaultConnectionService); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// WANCableLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANCableLinkConfig:1".
type WANCableLinkConfig1 struct {
	SOAPClient soap.SOAPClient
}

//
//
// Return values:
//
// * NewCableLinkConfigState: allowed values: notReady, dsSyncComplete, usParamAcquired, rangingComplete, ipComplete, todEstablished, paramTransferComplete, registrationComplete, operational, accessDenied
//
// * NewLinkType: allowed values: Ethernet
func (client *WANCableLinkConfig1) GetCableLinkConfigInfo() (NewCableLinkConfigState string, NewLinkType string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewCableLinkConfigState string
		NewLinkType             string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetCableLinkConfigInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewCableLinkConfigState, err = soap.UnmarshalString(response.NewCableLinkConfigState); err != nil {
		return
	}
	if NewLinkType, err = soap.UnmarshalString(response.NewLinkType); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDownstreamFrequency:
func (client *WANCableLinkConfig1) GetDownstreamFrequency() (NewDownstreamFrequency uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDownstreamFrequency string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetDownstreamFrequency", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDownstreamFrequency, err = soap.UnmarshalUi4(response.NewDownstreamFrequency); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDownstreamModulation: allowed values: 64QAM, 256QAM
func (client *WANCableLinkConfig1) GetDownstreamModulation() (NewDownstreamModulation string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDownstreamModulation string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetDownstreamModulation", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDownstreamModulation, err = soap.UnmarshalString(response.NewDownstreamModulation); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewUpstreamFrequency:
func (client *WANCableLinkConfig1) GetUpstreamFrequency() (NewUpstreamFrequency uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUpstreamFrequency string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamFrequency", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUpstreamFrequency, err = soap.UnmarshalUi4(response.NewUpstreamFrequency); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewUpstreamModulation: allowed values: QPSK, 16QAM
func (client *WANCableLinkConfig1) GetUpstreamModulation() (NewUpstreamModulation string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUpstreamModulation string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamModulation", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUpstreamModulation, err = soap.UnmarshalString(response.NewUpstreamModulation); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewUpstreamChannelID:
func (client *WANCableLinkConfig1) GetUpstreamChannelID() (NewUpstreamChannelID uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUpstreamChannelID string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamChannelID", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUpstreamChannelID, err = soap.UnmarshalUi4(response.NewUpstreamChannelID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewUpstreamPowerLevel:
func (client *WANCableLinkConfig1) GetUpstreamPowerLevel() (NewUpstreamPowerLevel uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUpstreamPowerLevel string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetUpstreamPowerLevel", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUpstreamPowerLevel, err = soap.UnmarshalUi4(response.NewUpstreamPowerLevel); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewBPIEncryptionEnabled:
func (client *WANCableLinkConfig1) GetBPIEncryptionEnabled() (NewBPIEncryptionEnabled bool, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewBPIEncryptionEnabled string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetBPIEncryptionEnabled", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewBPIEncryptionEnabled, err = soap.UnmarshalBoolean(response.NewBPIEncryptionEnabled); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewConfigFile:
func (client *WANCableLinkConfig1) GetConfigFile() (NewConfigFile string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConfigFile string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetConfigFile", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConfigFile, err = soap.UnmarshalString(response.NewConfigFile); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewTFTPServer:
func (client *WANCableLinkConfig1) GetTFTPServer() (NewTFTPServer string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewTFTPServer string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCableLinkConfig_1, "GetTFTPServer", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTFTPServer, err = soap.UnmarshalString(response.NewTFTPServer); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// WANCommonInterfaceConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANCommonInterfaceConfig:1".
type WANCommonInterfaceConfig1 struct {
	SOAPClient soap.SOAPClient
}

// Arguments:
//
// * NewEnabledForInternet:
//
//
func (client *WANCommonInterfaceConfig1) SetEnabledForInternet(NewEnabledForInternet bool) (err error) {
	// Request structure.
	var request struct {
		NewEnabledForInternet string
	}
	// BEGIN Marshal arguments into request.

	if request.NewEnabledForInternet, err = soap.MarshalBoolean(NewEnabledForInternet); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "SetEnabledForInternet", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewEnabledForInternet:
func (client *WANCommonInterfaceConfig1) GetEnabledForInternet() (NewEnabledForInternet bool, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewEnabledForInternet string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetEnabledForInternet", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewEnabledForInternet, err = soap.UnmarshalBoolean(response.NewEnabledForInternet); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewWANAccessType: allowed values: DSL, POTS, Cable, Ethernet
//
// * NewLayer1UpstreamMaxBitRate:
//
// * NewLayer1DownstreamMaxBitRate:
//
// * NewPhysicalLinkStatus: allowed values: Up, Down
func (client *WANCommonInterfaceConfig1) GetCommonLinkProperties() (NewWANAccessType string, NewLayer1UpstreamMaxBitRate uint32, NewLayer1DownstreamMaxBitRate uint32, NewPhysicalLinkStatus string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewWANAccessType              string
		NewLayer1UpstreamMaxBitRate   string
		NewLayer1DownstreamMaxBitRate string
		NewPhysicalLinkStatus         string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetCommonLinkProperties", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewWANAccessType, err = soap.UnmarshalString(response.NewWANAccessType); err != nil {
		return
	}
	if NewLayer1UpstreamMaxBitRate, err = soap.UnmarshalUi4(response.NewLayer1UpstreamMaxBitRate); err != nil {
		return
	}
	if NewLayer1DownstreamMaxBitRate, err = soap.UnmarshalUi4(response.NewLayer1DownstreamMaxBitRate); err != nil {
		return
	}
	if NewPhysicalLinkStatus, err = soap.UnmarshalString(response.NewPhysicalLinkStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewWANAccessProvider:
func (client *WANCommonInterfaceConfig1) GetWANAccessProvider() (NewWANAccessProvider string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewWANAccessProvider string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetWANAccessProvider", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewWANAccessProvider, err = soap.UnmarshalString(response.NewWANAccessProvider); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewMaximumActiveConnections: allowed value range: minimum=1, step=1
func (client *WANCommonInterfaceConfig1) GetMaximumActiveConnections() (NewMaximumActiveConnections uint16, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewMaximumActiveConnections string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetMaximumActiveConnections", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewMaximumActiveConnections, err = soap.UnmarshalUi2(response.NewMaximumActiveConnections); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewTotalBytesSent:
func (client *WANCommonInterfaceConfig1) GetTotalBytesSent() (NewTotalBytesSent uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewTotalBytesSent string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalBytesSent", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTotalBytesSent, err = soap.UnmarshalUi4(response.NewTotalBytesSent); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewTotalBytesReceived:
func (client *WANCommonInterfaceConfig1) GetTotalBytesReceived() (NewTotalBytesReceived uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewTotalBytesReceived string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalBytesReceived", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTotalBytesReceived, err = soap.UnmarshalUi4(response.NewTotalBytesReceived); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewTotalPacketsSent:
func (client *WANCommonInterfaceConfig1) GetTotalPacketsSent() (NewTotalPacketsSent uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewTotalPacketsSent string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalPacketsSent", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTotalPacketsSent, err = soap.UnmarshalUi4(response.NewTotalPacketsSent); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewTotalPacketsReceived:
func (client *WANCommonInterfaceConfig1) GetTotalPacketsReceived() (NewTotalPacketsReceived uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewTotalPacketsReceived string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetTotalPacketsReceived", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewTotalPacketsReceived, err = soap.UnmarshalUi4(response.NewTotalPacketsReceived); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewActiveConnectionIndex:
//
// Return values:
//
// * NewActiveConnDeviceContainer:
//
// * NewActiveConnectionServiceID:
func (client *WANCommonInterfaceConfig1) GetActiveConnection(NewActiveConnectionIndex uint16) (NewActiveConnDeviceContainer string, NewActiveConnectionServiceID string, err error) {
	// Request structure.
	var request struct {
		NewActiveConnectionIndex string
	}
	// BEGIN Marshal arguments into request.

	if request.NewActiveConnectionIndex, err = soap.MarshalUi2(NewActiveConnectionIndex); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewActiveConnDeviceContainer string
		NewActiveConnectionServiceID string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANCommonInterfaceConfig_1, "GetActiveConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewActiveConnDeviceContainer, err = soap.UnmarshalString(response.NewActiveConnDeviceContainer); err != nil {
		return
	}
	if NewActiveConnectionServiceID, err = soap.UnmarshalString(response.NewActiveConnectionServiceID); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// WANDSLLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANDSLLinkConfig:1".
type WANDSLLinkConfig1 struct {
	SOAPClient soap.SOAPClient
}

// Arguments:
//
// * NewLinkType:
//
//
func (client *WANDSLLinkConfig1) SetDSLLinkType(NewLinkType string) (err error) {
	// Request structure.
	var request struct {
		NewLinkType string
	}
	// BEGIN Marshal arguments into request.

	if request.NewLinkType, err = soap.MarshalString(NewLinkType); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDSLLinkType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewLinkType:
//
// * NewLinkStatus: allowed values: Up, Down
func (client *WANDSLLinkConfig1) GetDSLLinkInfo() (NewLinkType string, NewLinkStatus string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewLinkType   string
		NewLinkStatus string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetDSLLinkInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewLinkType, err = soap.UnmarshalString(response.NewLinkType); err != nil {
		return
	}
	if NewLinkStatus, err = soap.UnmarshalString(response.NewLinkStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewAutoConfig:
func (client *WANDSLLinkConfig1) GetAutoConfig() (NewAutoConfig bool, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewAutoConfig string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetAutoConfig", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewAutoConfig, err = soap.UnmarshalBoolean(response.NewAutoConfig); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewModulationType:
func (client *WANDSLLinkConfig1) GetModulationType() (NewModulationType string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewModulationType string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetModulationType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewModulationType, err = soap.UnmarshalString(response.NewModulationType); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewDestinationAddress:
//
//
func (client *WANDSLLinkConfig1) SetDestinationAddress(NewDestinationAddress string) (err error) {
	// Request structure.
	var request struct {
		NewDestinationAddress string
	}
	// BEGIN Marshal arguments into request.

	if request.NewDestinationAddress, err = soap.MarshalString(NewDestinationAddress); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetDestinationAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDestinationAddress:
func (client *WANDSLLinkConfig1) GetDestinationAddress() (NewDestinationAddress string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDestinationAddress string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetDestinationAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDestinationAddress, err = soap.UnmarshalString(response.NewDestinationAddress); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewATMEncapsulation:
//
//
func (client *WANDSLLinkConfig1) SetATMEncapsulation(NewATMEncapsulation string) (err error) {
	// Request structure.
	var request struct {
		NewATMEncapsulation string
	}
	// BEGIN Marshal arguments into request.

	if request.NewATMEncapsulation, err = soap.MarshalString(NewATMEncapsulation); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetATMEncapsulation", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewATMEncapsulation:
func (client *WANDSLLinkConfig1) GetATMEncapsulation() (NewATMEncapsulation string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewATMEncapsulation string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetATMEncapsulation", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewATMEncapsulation, err = soap.UnmarshalString(response.NewATMEncapsulation); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewFCSPreserved:
//
//
func (client *WANDSLLinkConfig1) SetFCSPreserved(NewFCSPreserved bool) (err error) {
	// Request structure.
	var request struct {
		NewFCSPreserved string
	}
	// BEGIN Marshal arguments into request.

	if request.NewFCSPreserved, err = soap.MarshalBoolean(NewFCSPreserved); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "SetFCSPreserved", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewFCSPreserved:
func (client *WANDSLLinkConfig1) GetFCSPreserved() (NewFCSPreserved bool, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewFCSPreserved string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANDSLLinkConfig_1, "GetFCSPreserved", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewFCSPreserved, err = soap.UnmarshalBoolean(response.NewFCSPreserved); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// WANEthernetLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANEthernetLinkConfig:1".
type WANEthernetLinkConfig1 struct {
	SOAPClient soap.SOAPClient
}

//
//
// Return values:
//
// * NewEthernetLinkStatus: allowed values: Up, Down
func (client *WANEthernetLinkConfig1) GetEthernetLinkStatus() (NewEthernetLinkStatus string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewEthernetLinkStatus string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANEthernetLinkConfig_1, "GetEthernetLinkStatus", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewEthernetLinkStatus, err = soap.UnmarshalString(response.NewEthernetLinkStatus); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// WANIPConnection1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANIPConnection:1".
type WANIPConnection1 struct {
	SOAPClient soap.SOAPClient
}

// Arguments:
//
// * NewConnectionType:
//
//
func (client *WANIPConnection1) SetConnectionType(NewConnectionType string) (err error) {
	// Request structure.
	var request struct {
		NewConnectionType string
	}
	// BEGIN Marshal arguments into request.

	if request.NewConnectionType, err = soap.MarshalString(NewConnectionType); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetConnectionType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewConnectionType:
//
// * NewPossibleConnectionTypes: allowed values: Unconfigured, IP_Routed, IP_Bridged
func (client *WANIPConnection1) GetConnectionTypeInfo() (NewConnectionType string, NewPossibleConnectionTypes string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConnectionType          string
		NewPossibleConnectionTypes string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetConnectionTypeInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConnectionType, err = soap.UnmarshalString(response.NewConnectionType); err != nil {
		return
	}
	if NewPossibleConnectionTypes, err = soap.UnmarshalString(response.NewPossibleConnectionTypes); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
//
func (client *WANIPConnection1) RequestConnection() (err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
//
func (client *WANIPConnection1) RequestTermination() (err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "RequestTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
//
func (client *WANIPConnection1) ForceTermination() (err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "ForceTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewAutoDisconnectTime:
//
//
func (client *WANIPConnection1) SetAutoDisconnectTime(NewAutoDisconnectTime uint32) (err error) {
	// Request structure.
	var request struct {
		NewAutoDisconnectTime string
	}
	// BEGIN Marshal arguments into request.

	if request.NewAutoDisconnectTime, err = soap.MarshalUi4(NewAutoDisconnectTime); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewIdleDisconnectTime:
//
//
func (client *WANIPConnection1) SetIdleDisconnectTime(NewIdleDisconnectTime uint32) (err error) {
	// Request structure.
	var request struct {
		NewIdleDisconnectTime string
	}
	// BEGIN Marshal arguments into request.

	if request.NewIdleDisconnectTime, err = soap.MarshalUi4(NewIdleDisconnectTime); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewWarnDisconnectDelay:
//
//
func (client *WANIPConnection1) SetWarnDisconnectDelay(NewWarnDisconnectDelay uint32) (err error) {
	// Request structure.
	var request struct {
		NewWarnDisconnectDelay string
	}
	// BEGIN Marshal arguments into request.

	if request.NewWarnDisconnectDelay, err = soap.MarshalUi4(NewWarnDisconnectDelay); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "SetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewConnectionStatus: allowed values: Unconfigured, Connected, Disconnected
//
// * NewLastConnectionError: allowed values: ERROR_NONE
//
// * NewUptime:
func (client *WANIPConnection1) GetStatusInfo() (NewConnectionStatus string, NewLastConnectionError string, NewUptime uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConnectionStatus    string
		NewLastConnectionError string
		NewUptime              string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetStatusInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConnectionStatus, err = soap.UnmarshalString(response.NewConnectionStatus); err != nil {
		return
	}
	if NewLastConnectionError, err = soap.UnmarshalString(response.NewLastConnectionError); err != nil {
		return
	}
	if NewUptime, err = soap.UnmarshalUi4(response.NewUptime); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewAutoDisconnectTime:
func (client *WANIPConnection1) GetAutoDisconnectTime() (NewAutoDisconnectTime uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewAutoDisconnectTime string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewAutoDisconnectTime, err = soap.UnmarshalUi4(response.NewAutoDisconnectTime); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewIdleDisconnectTime:
func (client *WANIPConnection1) GetIdleDisconnectTime() (NewIdleDisconnectTime uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewIdleDisconnectTime string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewIdleDisconnectTime, err = soap.UnmarshalUi4(response.NewIdleDisconnectTime); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewWarnDisconnectDelay:
func (client *WANIPConnection1) GetWarnDisconnectDelay() (NewWarnDisconnectDelay uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewWarnDisconnectDelay string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewWarnDisconnectDelay, err = soap.UnmarshalUi4(response.NewWarnDisconnectDelay); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewRSIPAvailable:
//
// * NewNATEnabled:
func (client *WANIPConnection1) GetNATRSIPStatus() (NewRSIPAvailable bool, NewNATEnabled bool, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewRSIPAvailable string
		NewNATEnabled    string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetNATRSIPStatus", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewRSIPAvailable, err = soap.UnmarshalBoolean(response.NewRSIPAvailable); err != nil {
		return
	}
	if NewNATEnabled, err = soap.UnmarshalBoolean(response.NewNATEnabled); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewPortMappingIndex:
//
// Return values:
//
// * NewRemoteHost:
//
// * NewExternalPort:
//
// * NewProtocol: allowed values: TCP, UDP
//
// * NewInternalPort:
//
// * NewInternalClient:
//
// * NewEnabled:
//
// * NewPortMappingDescription:
//
// * NewLeaseDuration:
func (client *WANIPConnection1) GetGenericPortMappingEntry(NewPortMappingIndex uint16) (NewRemoteHost string, NewExternalPort uint16, NewProtocol string, NewInternalPort uint16, NewInternalClient string, NewEnabled bool, NewPortMappingDescription string, NewLeaseDuration uint32, err error) {
	// Request structure.
	var request struct {
		NewPortMappingIndex string
	}
	// BEGIN Marshal arguments into request.

	if request.NewPortMappingIndex, err = soap.MarshalUi2(NewPortMappingIndex); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetGenericPortMappingEntry", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewRemoteHost, err = soap.UnmarshalString(response.NewRemoteHost); err != nil {
		return
	}
	if NewExternalPort, err = soap.UnmarshalUi2(response.NewExternalPort); err != nil {
		return
	}
	if NewProtocol, err = soap.UnmarshalString(response.NewProtocol); err != nil {
		return
	}
	if NewInternalPort, err = soap.UnmarshalUi2(response.NewInternalPort); err != nil {
		return
	}
	if NewInternalClient, err = soap.UnmarshalString(response.NewInternalClient); err != nil {
		return
	}
	if NewEnabled, err = soap.UnmarshalBoolean(response.NewEnabled); err != nil {
		return
	}
	if NewPortMappingDescription, err = soap.UnmarshalString(response.NewPortMappingDescription); err != nil {
		return
	}
	if NewLeaseDuration, err = soap.UnmarshalUi4(response.NewLeaseDuration); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewRemoteHost:
//
// * NewExternalPort:
//
// * NewProtocol: allowed values: TCP, UDP
//
// Return values:
//
// * NewInternalPort:
//
// * NewInternalClient:
//
// * NewEnabled:
//
// * NewPortMappingDescription:
//
// * NewLeaseDuration:
func (client *WANIPConnection1) GetSpecificPortMappingEntry(NewRemoteHost string, NewExternalPort uint16, NewProtocol string) (NewInternalPort uint16, NewInternalClient string, NewEnabled bool, NewPortMappingDescription string, NewLeaseDuration uint32, err error) {
	// Request structure.
	var request struct {
		NewRemoteHost   string
		NewExternalPort string
		NewProtocol     string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetSpecificPortMappingEntry", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewInternalPort, err = soap.UnmarshalUi2(response.NewInternalPort); err != nil {
		return
	}
	if NewInternalClient, err = soap.UnmarshalString(response.NewInternalClient); err != nil {
		return
	}
	if NewEnabled, err = soap.UnmarshalBoolean(response.NewEnabled); err != nil {
		return
	}
	if NewPortMappingDescription, err = soap.UnmarshalString(response.NewPortMappingDescription); err != nil {
		return
	}
	if NewLeaseDuration, err = soap.UnmarshalUi4(response.NewLeaseDuration); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewRemoteHost:
//
// * NewExternalPort:
//
// * NewProtocol: allowed values: TCP, UDP
//
// * NewInternalPort:
//
// * NewInternalClient:
//
// * NewEnabled:
//
// * NewPortMappingDescription:
//
// * NewLeaseDuration:
//
//
func (client *WANIPConnection1) AddPortMapping(NewRemoteHost string, NewExternalPort uint16, NewProtocol string, NewInternalPort uint16, NewInternalClient string, NewEnabled bool, NewPortMappingDescription string, NewLeaseDuration uint32) (err error) {
	// Request structure.
	var request struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}
	if request.NewInternalPort, err = soap.MarshalUi2(NewInternalPort); err != nil {
		return
	}
	if request.NewInternalClient, err = soap.MarshalString(NewInternalClient); err != nil {
		return
	}
	if request.NewEnabled, err = soap.MarshalBoolean(NewEnabled); err != nil {
		return
	}
	if request.NewPortMappingDescription, err = soap.MarshalString(NewPortMappingDescription); err != nil {
		return
	}
	if request.NewLeaseDuration, err = soap.MarshalUi4(NewLeaseDuration); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "AddPortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewRemoteHost:
//
// * NewExternalPort:
//
// * NewProtocol: allowed values: TCP, UDP
//
//
func (client *WANIPConnection1) DeletePortMapping(NewRemoteHost string, NewExternalPort uint16, NewProtocol string) (err error) {
	// Request structure.
	var request struct {
		NewRemoteHost   string
		NewExternalPort string
		NewProtocol     string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "DeletePortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewExternalIPAddress:
func (client *WANIPConnection1) GetExternalIPAddress() (NewExternalIPAddress string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewExternalIPAddress string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANIPConnection_1, "GetExternalIPAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewExternalIPAddress, err = soap.UnmarshalString(response.NewExternalIPAddress); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// WANPOTSLinkConfig1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANPOTSLinkConfig:1".
type WANPOTSLinkConfig1 struct {
	SOAPClient soap.SOAPClient
}

// Arguments:
//
// * NewISPPhoneNumber:
//
// * NewISPInfo:
//
// * NewLinkType: allowed values: PPP_Dialup
//
//
func (client *WANPOTSLinkConfig1) SetISPInfo(NewISPPhoneNumber string, NewISPInfo string, NewLinkType string) (err error) {
	// Request structure.
	var request struct {
		NewISPPhoneNumber string
		NewISPInfo        string
		NewLinkType       string
	}
	// BEGIN Marshal arguments into request.

	if request.NewISPPhoneNumber, err = soap.MarshalString(NewISPPhoneNumber); err != nil {
		return
	}
	if request.NewISPInfo, err = soap.MarshalString(NewISPInfo); err != nil {
		return
	}
	if request.NewLinkType, err = soap.MarshalString(NewLinkType); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetISPInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewNumberOfRetries:
//
// * NewDelayBetweenRetries:
//
//
func (client *WANPOTSLinkConfig1) SetCallRetryInfo(NewNumberOfRetries uint32, NewDelayBetweenRetries uint32) (err error) {
	// Request structure.
	var request struct {
		NewNumberOfRetries     string
		NewDelayBetweenRetries string
	}
	// BEGIN Marshal arguments into request.

	if request.NewNumberOfRetries, err = soap.MarshalUi4(NewNumberOfRetries); err != nil {
		return
	}
	if request.NewDelayBetweenRetries, err = soap.MarshalUi4(NewDelayBetweenRetries); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "SetCallRetryInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewISPPhoneNumber:
//
// * NewISPInfo:
//
// * NewLinkType: allowed values: PPP_Dialup
func (client *WANPOTSLinkConfig1) GetISPInfo() (NewISPPhoneNumber string, NewISPInfo string, NewLinkType string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewISPPhoneNumber string
		NewISPInfo        string
		NewLinkType       string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetISPInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewISPPhoneNumber, err = soap.UnmarshalString(response.NewISPPhoneNumber); err != nil {
		return
	}
	if NewISPInfo, err = soap.UnmarshalString(response.NewISPInfo); err != nil {
		return
	}
	if NewLinkType, err = soap.UnmarshalString(response.NewLinkType); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewNumberOfRetries:
//
// * NewDelayBetweenRetries:
func (client *WANPOTSLinkConfig1) GetCallRetryInfo() (NewNumberOfRetries uint32, NewDelayBetweenRetries uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewNumberOfRetries     string
		NewDelayBetweenRetries string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetCallRetryInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewNumberOfRetries, err = soap.UnmarshalUi4(response.NewNumberOfRetries); err != nil {
		return
	}
	if NewDelayBetweenRetries, err = soap.UnmarshalUi4(response.NewDelayBetweenRetries); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewFclass:
func (client *WANPOTSLinkConfig1) GetFclass() (NewFclass string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewFclass string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetFclass", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewFclass, err = soap.UnmarshalString(response.NewFclass); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDataModulationSupported:
func (client *WANPOTSLinkConfig1) GetDataModulationSupported() (NewDataModulationSupported string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDataModulationSupported string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataModulationSupported", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDataModulationSupported, err = soap.UnmarshalString(response.NewDataModulationSupported); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDataProtocol:
func (client *WANPOTSLinkConfig1) GetDataProtocol() (NewDataProtocol string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDataProtocol string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataProtocol", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDataProtocol, err = soap.UnmarshalString(response.NewDataProtocol); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewDataCompression:
func (client *WANPOTSLinkConfig1) GetDataCompression() (NewDataCompression string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewDataCompression string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetDataCompression", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewDataCompression, err = soap.UnmarshalString(response.NewDataCompression); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewPlusVTRCommandSupported:
func (client *WANPOTSLinkConfig1) GetPlusVTRCommandSupported() (NewPlusVTRCommandSupported bool, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPlusVTRCommandSupported string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPOTSLinkConfig_1, "GetPlusVTRCommandSupported", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPlusVTRCommandSupported, err = soap.UnmarshalBoolean(response.NewPlusVTRCommandSupported); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// WANPPPConnection1 is a client for UPnP SOAP service with URN "urn:schemas-upnp-org:service:WANPPPConnection:1".
type WANPPPConnection1 struct {
	SOAPClient soap.SOAPClient
}

// Arguments:
//
// * NewConnectionType:
//
//
func (client *WANPPPConnection1) SetConnectionType(NewConnectionType string) (err error) {
	// Request structure.
	var request struct {
		NewConnectionType string
	}
	// BEGIN Marshal arguments into request.

	if request.NewConnectionType, err = soap.MarshalString(NewConnectionType); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetConnectionType", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewConnectionType:
//
// * NewPossibleConnectionTypes: allowed values: Unconfigured, IP_Routed, DHCP_Spoofed, PPPoE_Bridged, PPTP_Relay, L2TP_Relay, PPPoE_Relay
func (client *WANPPPConnection1) GetConnectionTypeInfo() (NewConnectionType string, NewPossibleConnectionTypes string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConnectionType          string
		NewPossibleConnectionTypes string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetConnectionTypeInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConnectionType, err = soap.UnmarshalString(response.NewConnectionType); err != nil {
		return
	}
	if NewPossibleConnectionTypes, err = soap.UnmarshalString(response.NewPossibleConnectionTypes); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewUserName:
//
// * NewPassword:
//
//
func (client *WANPPPConnection1) ConfigureConnection(NewUserName string, NewPassword string) (err error) {
	// Request structure.
	var request struct {
		NewUserName string
		NewPassword string
	}
	// BEGIN Marshal arguments into request.

	if request.NewUserName, err = soap.MarshalString(NewUserName); err != nil {
		return
	}
	if request.NewPassword, err = soap.MarshalString(NewPassword); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ConfigureConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
//
func (client *WANPPPConnection1) RequestConnection() (err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestConnection", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
//
func (client *WANPPPConnection1) RequestTermination() (err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "RequestTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
//
func (client *WANPPPConnection1) ForceTermination() (err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "ForceTermination", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewAutoDisconnectTime:
//
//
func (client *WANPPPConnection1) SetAutoDisconnectTime(NewAutoDisconnectTime uint32) (err error) {
	// Request structure.
	var request struct {
		NewAutoDisconnectTime string
	}
	// BEGIN Marshal arguments into request.

	if request.NewAutoDisconnectTime, err = soap.MarshalUi4(NewAutoDisconnectTime); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewIdleDisconnectTime:
//
//
func (client *WANPPPConnection1) SetIdleDisconnectTime(NewIdleDisconnectTime uint32) (err error) {
	// Request structure.
	var request struct {
		NewIdleDisconnectTime string
	}
	// BEGIN Marshal arguments into request.

	if request.NewIdleDisconnectTime, err = soap.MarshalUi4(NewIdleDisconnectTime); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewWarnDisconnectDelay:
//
//
func (client *WANPPPConnection1) SetWarnDisconnectDelay(NewWarnDisconnectDelay uint32) (err error) {
	// Request structure.
	var request struct {
		NewWarnDisconnectDelay string
	}
	// BEGIN Marshal arguments into request.

	if request.NewWarnDisconnectDelay, err = soap.MarshalUi4(NewWarnDisconnectDelay); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "SetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewConnectionStatus: allowed values: Unconfigured, Connected, Disconnected
//
// * NewLastConnectionError: allowed values: ERROR_NONE
//
// * NewUptime:
func (client *WANPPPConnection1) GetStatusInfo() (NewConnectionStatus string, NewLastConnectionError string, NewUptime uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewConnectionStatus    string
		NewLastConnectionError string
		NewUptime              string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetStatusInfo", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewConnectionStatus, err = soap.UnmarshalString(response.NewConnectionStatus); err != nil {
		return
	}
	if NewLastConnectionError, err = soap.UnmarshalString(response.NewLastConnectionError); err != nil {
		return
	}
	if NewUptime, err = soap.UnmarshalUi4(response.NewUptime); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewUpstreamMaxBitRate:
//
// * NewDownstreamMaxBitRate:
func (client *WANPPPConnection1) GetLinkLayerMaxBitRates() (NewUpstreamMaxBitRate uint32, NewDownstreamMaxBitRate uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUpstreamMaxBitRate   string
		NewDownstreamMaxBitRate string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetLinkLayerMaxBitRates", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUpstreamMaxBitRate, err = soap.UnmarshalUi4(response.NewUpstreamMaxBitRate); err != nil {
		return
	}
	if NewDownstreamMaxBitRate, err = soap.UnmarshalUi4(response.NewDownstreamMaxBitRate); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewPPPEncryptionProtocol:
func (client *WANPPPConnection1) GetPPPEncryptionProtocol() (NewPPPEncryptionProtocol string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPPPEncryptionProtocol string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPEncryptionProtocol", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPPPEncryptionProtocol, err = soap.UnmarshalString(response.NewPPPEncryptionProtocol); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewPPPCompressionProtocol:
func (client *WANPPPConnection1) GetPPPCompressionProtocol() (NewPPPCompressionProtocol string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPPPCompressionProtocol string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPCompressionProtocol", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPPPCompressionProtocol, err = soap.UnmarshalString(response.NewPPPCompressionProtocol); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewPPPAuthenticationProtocol:
func (client *WANPPPConnection1) GetPPPAuthenticationProtocol() (NewPPPAuthenticationProtocol string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPPPAuthenticationProtocol string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPPPAuthenticationProtocol", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPPPAuthenticationProtocol, err = soap.UnmarshalString(response.NewPPPAuthenticationProtocol); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewUserName:
func (client *WANPPPConnection1) GetUserName() (NewUserName string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewUserName string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetUserName", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewUserName, err = soap.UnmarshalString(response.NewUserName); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewPassword:
func (client *WANPPPConnection1) GetPassword() (NewPassword string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewPassword string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetPassword", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewPassword, err = soap.UnmarshalString(response.NewPassword); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewAutoDisconnectTime:
func (client *WANPPPConnection1) GetAutoDisconnectTime() (NewAutoDisconnectTime uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewAutoDisconnectTime string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetAutoDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewAutoDisconnectTime, err = soap.UnmarshalUi4(response.NewAutoDisconnectTime); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewIdleDisconnectTime:
func (client *WANPPPConnection1) GetIdleDisconnectTime() (NewIdleDisconnectTime uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewIdleDisconnectTime string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetIdleDisconnectTime", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewIdleDisconnectTime, err = soap.UnmarshalUi4(response.NewIdleDisconnectTime); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewWarnDisconnectDelay:
func (client *WANPPPConnection1) GetWarnDisconnectDelay() (NewWarnDisconnectDelay uint32, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewWarnDisconnectDelay string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetWarnDisconnectDelay", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewWarnDisconnectDelay, err = soap.UnmarshalUi4(response.NewWarnDisconnectDelay); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewRSIPAvailable:
//
// * NewNATEnabled:
func (client *WANPPPConnection1) GetNATRSIPStatus() (NewRSIPAvailable bool, NewNATEnabled bool, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewRSIPAvailable string
		NewNATEnabled    string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetNATRSIPStatus", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewRSIPAvailable, err = soap.UnmarshalBoolean(response.NewRSIPAvailable); err != nil {
		return
	}
	if NewNATEnabled, err = soap.UnmarshalBoolean(response.NewNATEnabled); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewPortMappingIndex:
//
// Return values:
//
// * NewRemoteHost:
//
// * NewExternalPort:
//
// * NewProtocol: allowed values: TCP, UDP
//
// * NewInternalPort:
//
// * NewInternalClient:
//
// * NewEnabled:
//
// * NewPortMappingDescription:
//
// * NewLeaseDuration:
func (client *WANPPPConnection1) GetGenericPortMappingEntry(NewPortMappingIndex uint16) (NewRemoteHost string, NewExternalPort uint16, NewProtocol string, NewInternalPort uint16, NewInternalClient string, NewEnabled bool, NewPortMappingDescription string, NewLeaseDuration uint32, err error) {
	// Request structure.
	var request struct {
		NewPortMappingIndex string
	}
	// BEGIN Marshal arguments into request.

	if request.NewPortMappingIndex, err = soap.MarshalUi2(NewPortMappingIndex); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetGenericPortMappingEntry", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewRemoteHost, err = soap.UnmarshalString(response.NewRemoteHost); err != nil {
		return
	}
	if NewExternalPort, err = soap.UnmarshalUi2(response.NewExternalPort); err != nil {
		return
	}
	if NewProtocol, err = soap.UnmarshalString(response.NewProtocol); err != nil {
		return
	}
	if NewInternalPort, err = soap.UnmarshalUi2(response.NewInternalPort); err != nil {
		return
	}
	if NewInternalClient, err = soap.UnmarshalString(response.NewInternalClient); err != nil {
		return
	}
	if NewEnabled, err = soap.UnmarshalBoolean(response.NewEnabled); err != nil {
		return
	}
	if NewPortMappingDescription, err = soap.UnmarshalString(response.NewPortMappingDescription); err != nil {
		return
	}
	if NewLeaseDuration, err = soap.UnmarshalUi4(response.NewLeaseDuration); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewRemoteHost:
//
// * NewExternalPort:
//
// * NewProtocol: allowed values: TCP, UDP
//
// Return values:
//
// * NewInternalPort:
//
// * NewInternalClient:
//
// * NewEnabled:
//
// * NewPortMappingDescription:
//
// * NewLeaseDuration:
func (client *WANPPPConnection1) GetSpecificPortMappingEntry(NewRemoteHost string, NewExternalPort uint16, NewProtocol string) (NewInternalPort uint16, NewInternalClient string, NewEnabled bool, NewPortMappingDescription string, NewLeaseDuration uint32, err error) {
	// Request structure.
	var request struct {
		NewRemoteHost   string
		NewExternalPort string
		NewProtocol     string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetSpecificPortMappingEntry", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewInternalPort, err = soap.UnmarshalUi2(response.NewInternalPort); err != nil {
		return
	}
	if NewInternalClient, err = soap.UnmarshalString(response.NewInternalClient); err != nil {
		return
	}
	if NewEnabled, err = soap.UnmarshalBoolean(response.NewEnabled); err != nil {
		return
	}
	if NewPortMappingDescription, err = soap.UnmarshalString(response.NewPortMappingDescription); err != nil {
		return
	}
	if NewLeaseDuration, err = soap.UnmarshalUi4(response.NewLeaseDuration); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewRemoteHost:
//
// * NewExternalPort:
//
// * NewProtocol: allowed values: TCP, UDP
//
// * NewInternalPort:
//
// * NewInternalClient:
//
// * NewEnabled:
//
// * NewPortMappingDescription:
//
// * NewLeaseDuration:
//
//
func (client *WANPPPConnection1) AddPortMapping(NewRemoteHost string, NewExternalPort uint16, NewProtocol string, NewInternalPort uint16, NewInternalClient string, NewEnabled bool, NewPortMappingDescription string, NewLeaseDuration uint32) (err error) {
	// Request structure.
	var request struct {
		NewRemoteHost             string
		NewExternalPort           string
		NewProtocol               string
		NewInternalPort           string
		NewInternalClient         string
		NewEnabled                string
		NewPortMappingDescription string
		NewLeaseDuration          string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}
	if request.NewInternalPort, err = soap.MarshalUi2(NewInternalPort); err != nil {
		return
	}
	if request.NewInternalClient, err = soap.MarshalString(NewInternalClient); err != nil {
		return
	}
	if request.NewEnabled, err = soap.MarshalBoolean(NewEnabled); err != nil {
		return
	}
	if request.NewPortMappingDescription, err = soap.MarshalString(NewPortMappingDescription); err != nil {
		return
	}
	if request.NewLeaseDuration, err = soap.MarshalUi4(NewLeaseDuration); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "AddPortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

// Arguments:
//
// * NewRemoteHost:
//
// * NewExternalPort:
//
// * NewProtocol: allowed values: TCP, UDP
//
//
func (client *WANPPPConnection1) DeletePortMapping(NewRemoteHost string, NewExternalPort uint16, NewProtocol string) (err error) {
	// Request structure.
	var request struct {
		NewRemoteHost   string
		NewExternalPort string
		NewProtocol     string
	}
	// BEGIN Marshal arguments into request.

	if request.NewRemoteHost, err = soap.MarshalString(NewRemoteHost); err != nil {
		return
	}
	if request.NewExternalPort, err = soap.MarshalUi2(NewExternalPort); err != nil {
		return
	}
	if request.NewProtocol, err = soap.MarshalString(NewProtocol); err != nil {
		return
	}
	// END Marshal arguments into request.

	// Response structure.
	var response struct{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "DeletePortMapping", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	// END Unmarshal arguments from response.
	return
}

//
//
// Return values:
//
// * NewExternalIPAddress:
func (client *WANPPPConnection1) GetExternalIPAddress() (NewExternalIPAddress string, err error) {
	// Request structure.
	var request struct{}
	// BEGIN Marshal arguments into request.

	// END Marshal arguments into request.

	// Response structure.
	var response struct {
		NewExternalIPAddress string
	}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(URN_WANPPPConnection_1, "GetExternalIPAddress", &request, &response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.

	if NewExternalIPAddress, err = soap.UnmarshalString(response.NewExternalIPAddress); err != nil {
		return
	}
	// END Unmarshal arguments from response.
	return
}
