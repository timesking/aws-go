// Package support provides a client for AWS Support.
package support

import (
	"fmt"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
)

// Support is a client for AWS Support.
type Support struct {
	client *aws.JSONClient
}

// New returns a new Support client.
func New(key, secret, region string, client *http.Client) *Support {
	if client == nil {
		client = http.DefaultClient
	}

	return &Support{
		client: &aws.JSONClient{
			Client:       client,
			Region:       region,
			Endpoint:     fmt.Sprintf("https://support.%s.amazonaws.com", region),
			Prefix:       "support",
			Key:          key,
			Secret:       secret,
			JSONVersion:  "1.1",
			TargetPrefix: "AWSSupport_20130415",
		},
	}
}

// AddAttachmentsToSet adds one or more attachments to an attachment set.
// If an AttachmentSetId is not specified, a new attachment set is created,
// and the ID of the set is returned in the response. If an AttachmentSetId
// is specified, the attachments are added to the specified set, if it
// exists. An attachment set is a temporary container for attachments that
// are to be added to a case or case communication. The set is available
// for one hour after it is created; the ExpiryTime returned in the
// response indicates when the set expires. The maximum number of
// attachments in a set is 3, and the maximum size of any attachment in the
// set is 5
func (c *Support) AddAttachmentsToSet(req AddAttachmentsToSetRequest) (resp *AddAttachmentsToSetResponse, err error) {
	resp = &AddAttachmentsToSetResponse{}
	err = c.client.Do("AddAttachmentsToSet", "POST", "/", req, resp)
	return
}

// AddCommunicationToCase adds additional customer communication to an AWS
// Support case. You use the CaseId value to identify the case to add
// communication to. You can list a set of email addresses to copy on the
// communication using the CcEmailAddresses value. The CommunicationBody
// value contains the text of the communication. The response indicates the
// success or failure of the request. This operation implements a subset of
// the features of the AWS Support Center.
func (c *Support) AddCommunicationToCase(req AddCommunicationToCaseRequest) (resp *AddCommunicationToCaseResponse, err error) {
	resp = &AddCommunicationToCaseResponse{}
	err = c.client.Do("AddCommunicationToCase", "POST", "/", req, resp)
	return
}

// CreateCase creates a new case in the AWS Support Center. This operation
// is modeled on the behavior of the AWS Support Center Create Case page.
// Its parameters require you to specify the following information:
// IssueType. The type of issue for the case. You can specify either
// "customer-service" or "technical." If you do not indicate a value, the
// default is "technical." ServiceCode. The code for an AWS service. You
// obtain the ServiceCode by calling DescribeServices . CategoryCode. The
// category for the service defined for the ServiceCode value. You also
// obtain the category code for a service by calling DescribeServices .
// Each AWS service defines its own set of category codes. SeverityCode. A
// value that indicates the urgency of the case, which in turn determines
// the response time according to your service level agreement with AWS
// Support. You obtain the SeverityCode by calling DescribeSeverityLevels
// Subject. The Subject field on the AWS Support Center Create Case page.
// CommunicationBody. The Description field on the AWS Support Center
// Create Case page. AttachmentSetId. The ID of a set of attachments that
// has been created by using AddAttachmentsToSet Language. The human
// language in which AWS Support handles the case. English and Japanese are
// currently supported. CcEmailAddresses. The AWS Support Center field on
// the Create Case page. You can list email addresses to be copied on any
// correspondence about the case. The account that opens the case is
// already identified by passing the AWS Credentials in the method or in a
// method or function call from one of the programming languages supported
// by an AWS . To add additional communication or attachments to an
// existing case, use AddCommunicationToCase A successful CreateCase
// request returns an AWS Support case number. Case numbers are used by the
// DescribeCases operation to retrieve existing AWS Support cases.
func (c *Support) CreateCase(req CreateCaseRequest) (resp *CreateCaseResponse, err error) {
	resp = &CreateCaseResponse{}
	err = c.client.Do("CreateCase", "POST", "/", req, resp)
	return
}

// DescribeAttachment returns the attachment that has the specified ID.
// Attachment IDs are generated by the case management system when you add
// an attachment to a case or case communication. Attachment IDs are
// returned in the AttachmentDetails objects that are returned by the
// DescribeCommunications operation.
func (c *Support) DescribeAttachment(req DescribeAttachmentRequest) (resp *DescribeAttachmentResponse, err error) {
	resp = &DescribeAttachmentResponse{}
	err = c.client.Do("DescribeAttachment", "POST", "/", req, resp)
	return
}

// DescribeCases returns a list of cases that you specify by passing one or
// more case IDs. In addition, you can filter the cases by date by setting
// values for the AfterTime and BeforeTime request parameters. You can set
// values for the IncludeResolvedCases and IncludeCommunications request
// parameters to control how much information is returned. Case data is
// available for 12 months after creation. If a case was created more than
// 12 months ago, a request for data might cause an error. The response
// returns the following in format: One or more CaseDetails data types. One
// or more NextToken values, which specify where to paginate the returned
// records represented by the CaseDetails objects.
func (c *Support) DescribeCases(req DescribeCasesRequest) (resp *DescribeCasesResponse, err error) {
	resp = &DescribeCasesResponse{}
	err = c.client.Do("DescribeCases", "POST", "/", req, resp)
	return
}

// DescribeCommunications returns communications (and attachments) for one
// or more support cases. You can use the AfterTime and BeforeTime
// parameters to filter by date. You can use the CaseId parameter to
// restrict the results to a particular case. Case data is available for 12
// months after creation. If a case was created more than 12 months ago, a
// request for data might cause an error. You can use the MaxResults and
// NextToken parameters to control the pagination of the result set. Set
// MaxResults to the number of cases you want displayed on each page, and
// use NextToken to specify the resumption of pagination.
func (c *Support) DescribeCommunications(req DescribeCommunicationsRequest) (resp *DescribeCommunicationsResponse, err error) {
	resp = &DescribeCommunicationsResponse{}
	err = c.client.Do("DescribeCommunications", "POST", "/", req, resp)
	return
}

// DescribeServices returns the current list of AWS services and a list of
// service categories that applies to each one. You then use service names
// and categories in your CreateCase requests. Each AWS service has its own
// set of categories. The service codes and category codes correspond to
// the values that are displayed in the Service and Category drop-down
// lists on the AWS Support Center Create Case page. The values in those
// fields, however, do not necessarily match the service codes and
// categories returned by the DescribeServices request. Always use the
// service codes and categories obtained programmatically. This practice
// ensures that you always have the most recent set of service and category
// codes.
func (c *Support) DescribeServices(req DescribeServicesRequest) (resp *DescribeServicesResponse, err error) {
	resp = &DescribeServicesResponse{}
	err = c.client.Do("DescribeServices", "POST", "/", req, resp)
	return
}

// DescribeSeverityLevels returns the list of severity levels that you can
// assign to an AWS Support case. The severity level for a case is also a
// field in the CaseDetails data type included in any CreateCase request.
func (c *Support) DescribeSeverityLevels(req DescribeSeverityLevelsRequest) (resp *DescribeSeverityLevelsResponse, err error) {
	resp = &DescribeSeverityLevelsResponse{}
	err = c.client.Do("DescribeSeverityLevels", "POST", "/", req, resp)
	return
}

// DescribeTrustedAdvisorCheckRefreshStatuses returns the refresh status of
// the Trusted Advisor checks that have the specified check IDs. Check IDs
// can be obtained by calling DescribeTrustedAdvisorChecks
func (c *Support) DescribeTrustedAdvisorCheckRefreshStatuses(req DescribeTrustedAdvisorCheckRefreshStatusesRequest) (resp *DescribeTrustedAdvisorCheckRefreshStatusesResponse, err error) {
	resp = &DescribeTrustedAdvisorCheckRefreshStatusesResponse{}
	err = c.client.Do("DescribeTrustedAdvisorCheckRefreshStatuses", "POST", "/", req, resp)
	return
}

// DescribeTrustedAdvisorCheckResult returns the results of the Trusted
// Advisor check that has the specified check ID. Check IDs can be obtained
// by calling DescribeTrustedAdvisorChecks The response contains a
// TrustedAdvisorCheckResult object, which contains these three objects: In
// addition, the response contains these fields: Status. The alert status
// of the check: "ok" (green), "warning" (yellow), "error" (red), or
// "not_available". Timestamp. The time of the last refresh of the check.
// CheckId. The unique identifier for the check.
func (c *Support) DescribeTrustedAdvisorCheckResult(req DescribeTrustedAdvisorCheckResultRequest) (resp *DescribeTrustedAdvisorCheckResultResponse, err error) {
	resp = &DescribeTrustedAdvisorCheckResultResponse{}
	err = c.client.Do("DescribeTrustedAdvisorCheckResult", "POST", "/", req, resp)
	return
}

// DescribeTrustedAdvisorCheckSummaries returns the summaries of the
// results of the Trusted Advisor checks that have the specified check IDs.
// Check IDs can be obtained by calling DescribeTrustedAdvisorChecks The
// response contains an array of TrustedAdvisorCheckSummary objects.
func (c *Support) DescribeTrustedAdvisorCheckSummaries(req DescribeTrustedAdvisorCheckSummariesRequest) (resp *DescribeTrustedAdvisorCheckSummariesResponse, err error) {
	resp = &DescribeTrustedAdvisorCheckSummariesResponse{}
	err = c.client.Do("DescribeTrustedAdvisorCheckSummaries", "POST", "/", req, resp)
	return
}

// DescribeTrustedAdvisorChecks returns information about all available
// Trusted Advisor checks, including name, ID, category, description, and
// metadata. You must specify a language code; English ("en") and Japanese
// ("ja") are currently supported. The response contains a
// TrustedAdvisorCheckDescription for each check.
func (c *Support) DescribeTrustedAdvisorChecks(req DescribeTrustedAdvisorChecksRequest) (resp *DescribeTrustedAdvisorChecksResponse, err error) {
	resp = &DescribeTrustedAdvisorChecksResponse{}
	err = c.client.Do("DescribeTrustedAdvisorChecks", "POST", "/", req, resp)
	return
}

// RefreshTrustedAdvisorCheck requests a refresh of the Trusted Advisor
// check that has the specified check ID. Check IDs can be obtained by
// calling DescribeTrustedAdvisorChecks The response contains a
// TrustedAdvisorCheckRefreshStatus object, which contains these fields:
// Status. The refresh status of the check: "none", "enqueued",
// "processing", "success", or "abandoned". MillisUntilNextRefreshable. The
// amount of time, in milliseconds, until the check is eligible for
// refresh. CheckId. The unique identifier for the check.
func (c *Support) RefreshTrustedAdvisorCheck(req RefreshTrustedAdvisorCheckRequest) (resp *RefreshTrustedAdvisorCheckResponse, err error) {
	resp = &RefreshTrustedAdvisorCheckResponse{}
	err = c.client.Do("RefreshTrustedAdvisorCheck", "POST", "/", req, resp)
	return
}

// ResolveCase takes a CaseId and returns the initial state of the case
// along with the state of the case after the call to ResolveCase
// completed.
func (c *Support) ResolveCase(req ResolveCaseRequest) (resp *ResolveCaseResponse, err error) {
	resp = &ResolveCaseResponse{}
	err = c.client.Do("ResolveCase", "POST", "/", req, resp)
	return
}

type AddAttachmentsToSetRequest struct {
	AttachmentSetID string       `json:"attachmentSetId,omitempty"`
	Attachments     []Attachment `json:"attachments"`
}

type AddAttachmentsToSetResponse struct {
	AttachmentSetID string `json:"attachmentSetId,omitempty"`
	ExpiryTime      string `json:"expiryTime,omitempty"`
}

type AddCommunicationToCaseRequest struct {
	AttachmentSetID   string   `json:"attachmentSetId,omitempty"`
	CaseID            string   `json:"caseId,omitempty"`
	CcEmailAddresses  []string `json:"ccEmailAddresses,omitempty"`
	CommunicationBody string   `json:"communicationBody"`
}

type AddCommunicationToCaseResponse struct {
	Result bool `json:"result,omitempty"`
}

type Attachment struct {
	Data     []byte `json:"data,omitempty"`
	FileName string `json:"fileName,omitempty"`
}

type AttachmentDetails struct {
	AttachmentID string `json:"attachmentId,omitempty"`
	FileName     string `json:"fileName,omitempty"`
}

type CaseDetails struct {
	CaseID               string                   `json:"caseId,omitempty"`
	CategoryCode         string                   `json:"categoryCode,omitempty"`
	CcEmailAddresses     []string                 `json:"ccEmailAddresses,omitempty"`
	DisplayID            string                   `json:"displayId,omitempty"`
	Language             string                   `json:"language,omitempty"`
	RecentCommunications RecentCaseCommunications `json:"recentCommunications,omitempty"`
	ServiceCode          string                   `json:"serviceCode,omitempty"`
	SeverityCode         string                   `json:"severityCode,omitempty"`
	Status               string                   `json:"status,omitempty"`
	Subject              string                   `json:"subject,omitempty"`
	SubmittedBy          string                   `json:"submittedBy,omitempty"`
	TimeCreated          string                   `json:"timeCreated,omitempty"`
}

type Category struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type Communication struct {
	AttachmentSet []AttachmentDetails `json:"attachmentSet,omitempty"`
	Body          string              `json:"body,omitempty"`
	CaseID        string              `json:"caseId,omitempty"`
	SubmittedBy   string              `json:"submittedBy,omitempty"`
	TimeCreated   string              `json:"timeCreated,omitempty"`
}

type CreateCaseRequest struct {
	AttachmentSetID   string   `json:"attachmentSetId,omitempty"`
	CategoryCode      string   `json:"categoryCode,omitempty"`
	CcEmailAddresses  []string `json:"ccEmailAddresses,omitempty"`
	CommunicationBody string   `json:"communicationBody"`
	IssueType         string   `json:"issueType,omitempty"`
	Language          string   `json:"language,omitempty"`
	ServiceCode       string   `json:"serviceCode,omitempty"`
	SeverityCode      string   `json:"severityCode,omitempty"`
	Subject           string   `json:"subject"`
}

type CreateCaseResponse struct {
	CaseID string `json:"caseId,omitempty"`
}

type DescribeAttachmentRequest struct {
	AttachmentID string `json:"attachmentId"`
}

type DescribeAttachmentResponse struct {
	Attachment Attachment `json:"attachment,omitempty"`
}

type DescribeCasesRequest struct {
	AfterTime             string   `json:"afterTime,omitempty"`
	BeforeTime            string   `json:"beforeTime,omitempty"`
	CaseIDList            []string `json:"caseIdList,omitempty"`
	DisplayID             string   `json:"displayId,omitempty"`
	IncludeCommunications bool     `json:"includeCommunications,omitempty"`
	IncludeResolvedCases  bool     `json:"includeResolvedCases,omitempty"`
	Language              string   `json:"language,omitempty"`
	MaxResults            int      `json:"maxResults,omitempty"`
	NextToken             string   `json:"nextToken,omitempty"`
}

type DescribeCasesResponse struct {
	Cases     []CaseDetails `json:"cases,omitempty"`
	NextToken string        `json:"nextToken,omitempty"`
}

type DescribeCommunicationsRequest struct {
	AfterTime  string `json:"afterTime,omitempty"`
	BeforeTime string `json:"beforeTime,omitempty"`
	CaseID     string `json:"caseId"`
	MaxResults int    `json:"maxResults,omitempty"`
	NextToken  string `json:"nextToken,omitempty"`
}

type DescribeCommunicationsResponse struct {
	Communications []Communication `json:"communications,omitempty"`
	NextToken      string          `json:"nextToken,omitempty"`
}

type DescribeServicesRequest struct {
	Language        string   `json:"language,omitempty"`
	ServiceCodeList []string `json:"serviceCodeList,omitempty"`
}

type DescribeServicesResponse struct {
	Services []Service `json:"services,omitempty"`
}

type DescribeSeverityLevelsRequest struct {
	Language string `json:"language,omitempty"`
}

type DescribeSeverityLevelsResponse struct {
	SeverityLevels []SeverityLevel `json:"severityLevels,omitempty"`
}

type DescribeTrustedAdvisorCheckRefreshStatusesRequest struct {
	CheckIds []string `json:"checkIds"`
}

type DescribeTrustedAdvisorCheckRefreshStatusesResponse struct {
	Statuses []TrustedAdvisorCheckRefreshStatus `json:"statuses"`
}

type DescribeTrustedAdvisorCheckResultRequest struct {
	CheckID  string `json:"checkId"`
	Language string `json:"language,omitempty"`
}

type DescribeTrustedAdvisorCheckResultResponse struct {
	Result TrustedAdvisorCheckResult `json:"result,omitempty"`
}

type DescribeTrustedAdvisorCheckSummariesRequest struct {
	CheckIds []string `json:"checkIds"`
}

type DescribeTrustedAdvisorCheckSummariesResponse struct {
	Summaries []TrustedAdvisorCheckSummary `json:"summaries"`
}

type DescribeTrustedAdvisorChecksRequest struct {
	Language string `json:"language"`
}

type DescribeTrustedAdvisorChecksResponse struct {
	Checks []TrustedAdvisorCheckDescription `json:"checks"`
}

type RecentCaseCommunications struct {
	Communications []Communication `json:"communications,omitempty"`
	NextToken      string          `json:"nextToken,omitempty"`
}

type RefreshTrustedAdvisorCheckRequest struct {
	CheckID string `json:"checkId"`
}

type RefreshTrustedAdvisorCheckResponse struct {
	Status TrustedAdvisorCheckRefreshStatus `json:"status"`
}

type ResolveCaseRequest struct {
	CaseID string `json:"caseId,omitempty"`
}

type ResolveCaseResponse struct {
	FinalCaseStatus   string `json:"finalCaseStatus,omitempty"`
	InitialCaseStatus string `json:"initialCaseStatus,omitempty"`
}

type Service struct {
	Categories []Category `json:"categories,omitempty"`
	Code       string     `json:"code,omitempty"`
	Name       string     `json:"name,omitempty"`
}

type SeverityLevel struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type TrustedAdvisorCategorySpecificSummary struct {
	CostOptimizing TrustedAdvisorCostOptimizingSummary `json:"costOptimizing,omitempty"`
}

type TrustedAdvisorCheckDescription struct {
	Category    string   `json:"category"`
	Description string   `json:"description"`
	ID          string   `json:"id"`
	Metadata    []string `json:"metadata"`
	Name        string   `json:"name"`
}

type TrustedAdvisorCheckRefreshStatus struct {
	CheckID                    string `json:"checkId"`
	MillisUntilNextRefreshable int    `json:"millisUntilNextRefreshable"`
	Status                     string `json:"status"`
}

type TrustedAdvisorCheckResult struct {
	CategorySpecificSummary TrustedAdvisorCategorySpecificSummary `json:"categorySpecificSummary"`
	CheckID                 string                                `json:"checkId"`
	FlaggedResources        []TrustedAdvisorResourceDetail        `json:"flaggedResources"`
	ResourcesSummary        TrustedAdvisorResourcesSummary        `json:"resourcesSummary"`
	Status                  string                                `json:"status"`
	Timestamp               string                                `json:"timestamp"`
}

type TrustedAdvisorCheckSummary struct {
	CategorySpecificSummary TrustedAdvisorCategorySpecificSummary `json:"categorySpecificSummary"`
	CheckID                 string                                `json:"checkId"`
	HasFlaggedResources     bool                                  `json:"hasFlaggedResources,omitempty"`
	ResourcesSummary        TrustedAdvisorResourcesSummary        `json:"resourcesSummary"`
	Status                  string                                `json:"status"`
	Timestamp               string                                `json:"timestamp"`
}

type TrustedAdvisorCostOptimizingSummary struct {
	EstimatedMonthlySavings        float64 `json:"estimatedMonthlySavings"`
	EstimatedPercentMonthlySavings float64 `json:"estimatedPercentMonthlySavings"`
}

type TrustedAdvisorResourceDetail struct {
	IsSuppressed bool     `json:"isSuppressed,omitempty"`
	Metadata     []string `json:"metadata"`
	Region       string   `json:"region"`
	ResourceID   string   `json:"resourceId"`
	Status       string   `json:"status"`
}

type TrustedAdvisorResourcesSummary struct {
	ResourcesFlagged    int `json:"resourcesFlagged"`
	ResourcesIgnored    int `json:"resourcesIgnored"`
	ResourcesProcessed  int `json:"resourcesProcessed"`
	ResourcesSuppressed int `json:"resourcesSuppressed"`
}

var _ time.Time // to avoid errors if the time package isn't referenced