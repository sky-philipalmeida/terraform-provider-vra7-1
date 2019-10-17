package sdk

import (
	"time"
)

// ResourceActionTemplate - is used to store information
// related to resource action template information.
type ResourceActionTemplate struct {
	Type        string                 `json:"type,omitempty"`
	ResourceID  string                 `json:"resourceId,omitempty"`
	ActionID    string                 `json:"actionId,omitempty"`
	Description string                 `json:"description,omitempty"`
	Reasons     string                 `json:"reasons,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
}

// RequestStatusView - used to store REST response of
// request triggered against any resource.
type RequestStatusView struct {
	RequestCompletion struct {
		RequestCompletionState string `json:"requestCompletionState"`
		CompletionDetails      string `json:"CompletionDetails"`
	} `json:"requestCompletion"`
	Phase string `json:"phase"`
}

// BusinessGroups - list of business groups
type BusinessGroups struct {
	Content []BusinessGroup `json:"content,omitempty"`
}

// BusinessGroup - detail view of a business group
type BusinessGroup struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

// RequestResourceView - resource view of a provisioned request
type RequestResourceView struct {
	Content []DeploymentResource `json:"content,omitempty"`
	Links   []interface{}        `json:"links,omitempty"`
}

// DeploymentResource - deployment level view of the provisionined request
type DeploymentResource struct {
	RequestState    string                 `json:"requestState,omitempty"`
	Description     string                 `json:"description,omitempty"`
	LastUpdated     string                 `json:"lastUpdated,omitempty"`
	TenantID        string                 `json:"tenantId,omitempty"`
	Name            string                 `json:"name,omitempty"`
	BusinessGroupID string                 `json:"businessGroupId,omitempty"`
	DateCreated     string                 `json:"dateCreated,omitempty"`
	Status          string                 `json:"status,omitempty"`
	RequestID       string                 `json:"requestId,omitempty"`
	ResourceID      string                 `json:"resourceId,omitempty"`
	ResourceType    string                 `json:"resourceType,omitempty"`
	ResourcesData   DeploymentResourceData `json:"data,omitempty"`
}

// DeploymentResourceData - view of the resources/machines in a deployment
type DeploymentResourceData struct {
	Memory                      int         `json:"MachineMemory,omitempty"`
	CPU                         int         `json:"MachineCPU,omitempty"`
	IPAddress                   string      `json:"ip_address,omitempty"`
	Storage                     int         `json:"MachineStorage,omitempty"`
	MachineInterfaceType        string      `json:"MachineInterfaceType,omitempty"`
	MachineName                 string      `json:"MachineName,omitempty"`
	MachineGuestOperatingSystem string      `json:"MachineGuestOperatingSystem,omitempty"`
	MachineDestructionDate      string      `json:"MachineDestructionDate,omitempty"`
	MachineGroupName            string      `json:"MachineGroupName,omitempty"`
	MachineBlueprintName        string      `json:"MachineBlueprintName,omitempty"`
	MachineReservationName      string      `json:"MachineReservationName,omitempty"`
	MachineType                 string      `json:"MachineType,omitempty"`
	MachineID                   string      `json:"machineId,omitempty"`
	MachineExpirationDate       string      `json:"MachineExpirationDate,omitempty"`
	Component                   string      `json:"Component,omitempty"`
	Expire                      bool        `json:"Expire,omitempty"`
	Reconfigure                 bool        `json:"Reconfigure,omitempty"`
	Reset                       bool        `json:"Reset,omitempty"`
	Reboot                      bool        `json:"Reboot,omitempty"`
	PowerOff                    bool        `json:"PowerOff,omitempty"`
	Destroy                     bool        `json:"Destroy,omitempty"`
	Shutdown                    bool        `json:"Shutdown,omitempty"`
	Suspend                     bool        `json:"Suspend,omitempty"`
	Reprovision                 bool        `json:"Reprovision,omitempty"`
	ChangeLease                 bool        `json:"ChangeLease,omitempty"`
	ChangeOwner                 bool        `json:"ChangeOwner,omitempty"`
	CreateSnapshot              bool        `json:"CreateSnapshot,omitempty"`
	Networks                    []NWDetails `json:"NETWORK_LIST,omitempty"`
	CollectedIPAddress          string      `json:"__datacollected_ipaddress,omitempty"`
}

//Networks Related structs
//
type NWDetails struct {
	ComponentTypeId    string      `json:"componentTypeId,omitempty"`
	NetworkAddressInfo NetworkInfo `json:"data,omitempty"`
}
type NetworkInfo struct {
	IPAddress  string `json:"NETWORK_ADDRESS,omitempty"`
	MACAddress string `json:"NETWORK_MAC_ADDRESS,omitempty"`
	Name       string `json:"NETWORK_NAME,omitempty"`
}

// ResourceActions - Retrieves the resources that were provisioned as a result of a given request.
// Also returns the actions allowed on the resources and their templates
type ResourceActions struct {
	Links   []interface{}           `json:"links,omitempty"`
	Content []ResourceActionContent `json:"content,omitempty"`
}

// ResourceActionContent - Detailed view of the resource provisioned and the operation allowed
type ResourceActionContent struct {
	ID              string          `json:"id,omitempty"`
	Name            string          `json:"name,omitempty"`
	ResourceTypeRef ResourceTypeRef `json:"resourceTypeRef,omitempty"`
	Status          string          `json:"status,omitempty"`
	RequestID       string          `json:"requestId,omitempty"`
	RequestState    string          `json:"requestState,omitempty"`
	Operations      []Operation     `json:"operations,omitempty"`
	ResourceData    ResourceDataMap `json:"resourceData,omitempty"`
}

// ResourceTypeRef - type of resource (deployment, or machine, etc)
type ResourceTypeRef struct {
	ID    string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
}

// Operation - detailed view of an operation allowed on a resource
type Operation struct {
	Name        string `json:"name,omitempty"`
	OperationID string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}

// ResourceDataMap - properties of a provisioned resource
type ResourceDataMap struct {
	Entries []ResourceDataEntry `json:"entries,omitempty"`
}

// ResourceDataEntry - the property key and value of a resource
type ResourceDataEntry struct {
	Key   string                 `json:"key,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
}

//CatalogRequest - A structure that captures a vRA catalog request.
type CatalogRequest struct {
	ID           string      `json:"id"`
	IconID       string      `json:"iconId"`
	Version      int         `json:"version"`
	State        string      `json:"state"`
	Description  string      `json:"description"`
	Reasons      interface{} `json:"reasons"`
	RequestedFor string      `json:"requestedFor"`
	RequestedBy  string      `json:"requestedBy"`
	Organization struct {
		TenantRef      string `json:"tenantRef"`
		TenantLabel    string `json:"tenantLabel"`
		SubtenantRef   string `json:"subtenantRef"`
		SubtenantLabel string `json:"subtenantLabel"`
	} `json:"organization"`

	RequestorEntitlementID   string                 `json:"requestorEntitlementId"`
	PreApprovalID            string                 `json:"preApprovalId"`
	PostApprovalID           string                 `json:"postApprovalId"`
	DateCreated              time.Time              `json:"dateCreated"`
	LastUpdated              time.Time              `json:"lastUpdated"`
	DateSubmitted            time.Time              `json:"dateSubmitted"`
	DateApproved             time.Time              `json:"dateApproved"`
	DateCompleted            time.Time              `json:"dateCompleted"`
	Quote                    interface{}            `json:"quote"`
	RequestData              map[string]interface{} `json:"requestData"`
	RequestCompletion        string                 `json:"requestCompletion"`
	RetriesRemaining         int                    `json:"retriesRemaining"`
	RequestedItemName        string                 `json:"requestedItemName"`
	RequestedItemDescription string                 `json:"requestedItemDescription"`
	Components               string                 `json:"components"`
	StateName                string                 `json:"stateName"`

	CatalogItemProviderBinding struct {
		BindingID   string `json:"bindingId"`
		ProviderRef struct {
			ID    string `json:"id"`
			Label string `json:"label"`
		} `json:"providerRef"`
	} `json:"catalogItemProviderBinding"`

	Phase           string `json:"phase"`
	ApprovalStatus  string `json:"approvalStatus"`
	ExecutionStatus string `json:"executionStatus"`
	WaitingStatus   string `json:"waitingStatus"`
	CatalogItemRef  struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	} `json:"catalogItemRef"`
}

//CatalogItemRequestTemplate - A structure that captures a catalog request template, to be filled in and POSTED.
type CatalogItemRequestTemplate struct {
	Type            string                 `json:"type"`
	CatalogItemID   string                 `json:"catalogItemId"`
	RequestedFor    string                 `json:"requestedFor"`
	BusinessGroupID string                 `json:"businessGroupId"`
	Description     string                 `json:"description"`
	Reasons         string                 `json:"reasons"`
	Data            map[string]interface{} `json:"data"`
}

//catalogName - This struct holds catalog name from json response.
type catalogName struct {
	Name string `json:"name"`
	ID   string `json:"catalogItemId"`
}

//CatalogItem - This struct holds the value of response of catalog item list
type CatalogItem struct {
	CatalogItem catalogName `json:"catalogItem"`
}

// EntitledCatalogItemViews represents catalog items in an active state, the current user
// is entitled to consume
type EntitledCatalogItemViews struct {
	Links    interface{} `json:"links"`
	Content  interface{} `json:"content"`
	Metadata Metadata    `json:"metadata"`
}

// Metadata - Metadata  used to store metadata of resource list response
type Metadata struct {
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	Number        int `json:"number"`
}
