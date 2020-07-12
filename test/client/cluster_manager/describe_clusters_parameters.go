// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeClustersParams creates a new DescribeClustersParams object
// with the default values initialized.
func NewDescribeClustersParams() *DescribeClustersParams {
	var ()
	return &DescribeClustersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeClustersParamsWithTimeout creates a new DescribeClustersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeClustersParamsWithTimeout(timeout time.Duration) *DescribeClustersParams {
	var ()
	return &DescribeClustersParams{

		timeout: timeout,
	}
}

// NewDescribeClustersParamsWithContext creates a new DescribeClustersParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeClustersParamsWithContext(ctx context.Context) *DescribeClustersParams {
	var ()
	return &DescribeClustersParams{

		Context: ctx,
	}
}

// NewDescribeClustersParamsWithHTTPClient creates a new DescribeClustersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeClustersParamsWithHTTPClient(client *http.Client) *DescribeClustersParams {
	var ()
	return &DescribeClustersParams{
		HTTPClient: client,
	}
}

/*DescribeClustersParams contains all the parameters to send to the API endpoint
for the describe clusters operation typically these are written to a http.Request
*/
type DescribeClustersParams struct {

	/*AppID
	  app ids.

	*/
	AppID []string
	/*ClusterID
	  cluster ids.

	*/
	ClusterID []string
	/*ClusterType
	  cluster type, frontgate or normal cluster.

	*/
	ClusterType *string
	/*CreatedDate
	  cluster created duration eg.[1 day].

	*/
	CreatedDate *int64
	/*DisplayColumns
	  select column to display.

	*/
	DisplayColumns []string
	/*ExternalClusterID
	  external cluster id.

	*/
	ExternalClusterID *string
	/*FrontgateID
	  frontgate ids.

	*/
	FrontgateID []string
	/*Limit
	  data limit per page, default value 20, max value 200.

	*/
	Limit *int64
	/*MinCreateTime*/
	MinCreateTime *strfmt.DateTime
	/*Offset
	  data offset, default 0.

	*/
	Offset *int64
	/*Owner
	  owner.

	*/
	Owner []string
	/*Reverse
	  value = 0 sort ASC, value = 1 sort DESC.

	*/
	Reverse *bool
	/*RuntimeID
	  runtime ids.

	*/
	RuntimeID []string
	/*SearchWord
	  query key, support these fields(cluster_id, app_id, version_id, status, runtime_id, frontgate_id, owner, cluster_type).

	*/
	SearchWord *string
	/*SortKey
	  sort key, order by sort_key, default create_time.

	*/
	SortKey *string
	/*Status
	  cluster status eg.[active|used|enabled|disabled|deleted|stopped|ceased].

	*/
	Status []string
	/*VersionID
	  version ids.

	*/
	VersionID []string
	/*WithDetail
	  get cluster detail info or not.

	*/
	WithDetail *bool
	/*Zone
	  namespace.

	*/
	Zone []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe clusters params
func (o *DescribeClustersParams) WithTimeout(timeout time.Duration) *DescribeClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe clusters params
func (o *DescribeClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe clusters params
func (o *DescribeClustersParams) WithContext(ctx context.Context) *DescribeClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe clusters params
func (o *DescribeClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe clusters params
func (o *DescribeClustersParams) WithHTTPClient(client *http.Client) *DescribeClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe clusters params
func (o *DescribeClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppID adds the appID to the describe clusters params
func (o *DescribeClustersParams) WithAppID(appID []string) *DescribeClustersParams {
	o.SetAppID(appID)
	return o
}

// SetAppID adds the appId to the describe clusters params
func (o *DescribeClustersParams) SetAppID(appID []string) {
	o.AppID = appID
}

// WithClusterID adds the clusterID to the describe clusters params
func (o *DescribeClustersParams) WithClusterID(clusterID []string) *DescribeClustersParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the describe clusters params
func (o *DescribeClustersParams) SetClusterID(clusterID []string) {
	o.ClusterID = clusterID
}

// WithClusterType adds the clusterType to the describe clusters params
func (o *DescribeClustersParams) WithClusterType(clusterType *string) *DescribeClustersParams {
	o.SetClusterType(clusterType)
	return o
}

// SetClusterType adds the clusterType to the describe clusters params
func (o *DescribeClustersParams) SetClusterType(clusterType *string) {
	o.ClusterType = clusterType
}

// WithCreatedDate adds the createdDate to the describe clusters params
func (o *DescribeClustersParams) WithCreatedDate(createdDate *int64) *DescribeClustersParams {
	o.SetCreatedDate(createdDate)
	return o
}

// SetCreatedDate adds the createdDate to the describe clusters params
func (o *DescribeClustersParams) SetCreatedDate(createdDate *int64) {
	o.CreatedDate = createdDate
}

// WithDisplayColumns adds the displayColumns to the describe clusters params
func (o *DescribeClustersParams) WithDisplayColumns(displayColumns []string) *DescribeClustersParams {
	o.SetDisplayColumns(displayColumns)
	return o
}

// SetDisplayColumns adds the displayColumns to the describe clusters params
func (o *DescribeClustersParams) SetDisplayColumns(displayColumns []string) {
	o.DisplayColumns = displayColumns
}

// WithExternalClusterID adds the externalClusterID to the describe clusters params
func (o *DescribeClustersParams) WithExternalClusterID(externalClusterID *string) *DescribeClustersParams {
	o.SetExternalClusterID(externalClusterID)
	return o
}

// SetExternalClusterID adds the externalClusterId to the describe clusters params
func (o *DescribeClustersParams) SetExternalClusterID(externalClusterID *string) {
	o.ExternalClusterID = externalClusterID
}

// WithFrontgateID adds the frontgateID to the describe clusters params
func (o *DescribeClustersParams) WithFrontgateID(frontgateID []string) *DescribeClustersParams {
	o.SetFrontgateID(frontgateID)
	return o
}

// SetFrontgateID adds the frontgateId to the describe clusters params
func (o *DescribeClustersParams) SetFrontgateID(frontgateID []string) {
	o.FrontgateID = frontgateID
}

// WithLimit adds the limit to the describe clusters params
func (o *DescribeClustersParams) WithLimit(limit *int64) *DescribeClustersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe clusters params
func (o *DescribeClustersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMinCreateTime adds the minCreateTime to the describe clusters params
func (o *DescribeClustersParams) WithMinCreateTime(minCreateTime *strfmt.DateTime) *DescribeClustersParams {
	o.SetMinCreateTime(minCreateTime)
	return o
}

// SetMinCreateTime adds the minCreateTime to the describe clusters params
func (o *DescribeClustersParams) SetMinCreateTime(minCreateTime *strfmt.DateTime) {
	o.MinCreateTime = minCreateTime
}

// WithOffset adds the offset to the describe clusters params
func (o *DescribeClustersParams) WithOffset(offset *int64) *DescribeClustersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe clusters params
func (o *DescribeClustersParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe clusters params
func (o *DescribeClustersParams) WithOwner(owner []string) *DescribeClustersParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe clusters params
func (o *DescribeClustersParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithReverse adds the reverse to the describe clusters params
func (o *DescribeClustersParams) WithReverse(reverse *bool) *DescribeClustersParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe clusters params
func (o *DescribeClustersParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithRuntimeID adds the runtimeID to the describe clusters params
func (o *DescribeClustersParams) WithRuntimeID(runtimeID []string) *DescribeClustersParams {
	o.SetRuntimeID(runtimeID)
	return o
}

// SetRuntimeID adds the runtimeId to the describe clusters params
func (o *DescribeClustersParams) SetRuntimeID(runtimeID []string) {
	o.RuntimeID = runtimeID
}

// WithSearchWord adds the searchWord to the describe clusters params
func (o *DescribeClustersParams) WithSearchWord(searchWord *string) *DescribeClustersParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe clusters params
func (o *DescribeClustersParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe clusters params
func (o *DescribeClustersParams) WithSortKey(sortKey *string) *DescribeClustersParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe clusters params
func (o *DescribeClustersParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe clusters params
func (o *DescribeClustersParams) WithStatus(status []string) *DescribeClustersParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe clusters params
func (o *DescribeClustersParams) SetStatus(status []string) {
	o.Status = status
}

// WithVersionID adds the versionID to the describe clusters params
func (o *DescribeClustersParams) WithVersionID(versionID []string) *DescribeClustersParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the describe clusters params
func (o *DescribeClustersParams) SetVersionID(versionID []string) {
	o.VersionID = versionID
}

// WithWithDetail adds the withDetail to the describe clusters params
func (o *DescribeClustersParams) WithWithDetail(withDetail *bool) *DescribeClustersParams {
	o.SetWithDetail(withDetail)
	return o
}

// SetWithDetail adds the withDetail to the describe clusters params
func (o *DescribeClustersParams) SetWithDetail(withDetail *bool) {
	o.WithDetail = withDetail
}

// WithZone adds the zone to the describe clusters params
func (o *DescribeClustersParams) WithZone(zone []string) *DescribeClustersParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the describe clusters params
func (o *DescribeClustersParams) SetZone(zone []string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAppID := o.AppID

	joinedAppID := swag.JoinByFormat(valuesAppID, "multi")
	// query array param app_id
	if err := r.SetQueryParam("app_id", joinedAppID...); err != nil {
		return err
	}

	valuesClusterID := o.ClusterID

	joinedClusterID := swag.JoinByFormat(valuesClusterID, "multi")
	// query array param cluster_id
	if err := r.SetQueryParam("cluster_id", joinedClusterID...); err != nil {
		return err
	}

	if o.ClusterType != nil {

		// query param cluster_type
		var qrClusterType string
		if o.ClusterType != nil {
			qrClusterType = *o.ClusterType
		}
		qClusterType := qrClusterType
		if qClusterType != "" {
			if err := r.SetQueryParam("cluster_type", qClusterType); err != nil {
				return err
			}
		}

	}

	if o.CreatedDate != nil {

		// query param created_date
		var qrCreatedDate int64
		if o.CreatedDate != nil {
			qrCreatedDate = *o.CreatedDate
		}
		qCreatedDate := swag.FormatInt64(qrCreatedDate)
		if qCreatedDate != "" {
			if err := r.SetQueryParam("created_date", qCreatedDate); err != nil {
				return err
			}
		}

	}

	valuesDisplayColumns := o.DisplayColumns

	joinedDisplayColumns := swag.JoinByFormat(valuesDisplayColumns, "multi")
	// query array param display_columns
	if err := r.SetQueryParam("display_columns", joinedDisplayColumns...); err != nil {
		return err
	}

	if o.ExternalClusterID != nil {

		// query param external_cluster_id
		var qrExternalClusterID string
		if o.ExternalClusterID != nil {
			qrExternalClusterID = *o.ExternalClusterID
		}
		qExternalClusterID := qrExternalClusterID
		if qExternalClusterID != "" {
			if err := r.SetQueryParam("external_cluster_id", qExternalClusterID); err != nil {
				return err
			}
		}

	}

	valuesFrontgateID := o.FrontgateID

	joinedFrontgateID := swag.JoinByFormat(valuesFrontgateID, "multi")
	// query array param frontgate_id
	if err := r.SetQueryParam("frontgate_id", joinedFrontgateID...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.MinCreateTime != nil {

		// query param min_create_time
		var qrMinCreateTime strfmt.DateTime
		if o.MinCreateTime != nil {
			qrMinCreateTime = *o.MinCreateTime
		}
		qMinCreateTime := qrMinCreateTime.String()
		if qMinCreateTime != "" {
			if err := r.SetQueryParam("min_create_time", qMinCreateTime); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "multi")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	valuesRuntimeID := o.RuntimeID

	joinedRuntimeID := swag.JoinByFormat(valuesRuntimeID, "multi")
	// query array param runtime_id
	if err := r.SetQueryParam("runtime_id", joinedRuntimeID...); err != nil {
		return err
	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesVersionID := o.VersionID

	joinedVersionID := swag.JoinByFormat(valuesVersionID, "multi")
	// query array param version_id
	if err := r.SetQueryParam("version_id", joinedVersionID...); err != nil {
		return err
	}

	if o.WithDetail != nil {

		// query param with_detail
		var qrWithDetail bool
		if o.WithDetail != nil {
			qrWithDetail = *o.WithDetail
		}
		qWithDetail := swag.FormatBool(qrWithDetail)
		if qWithDetail != "" {
			if err := r.SetQueryParam("with_detail", qWithDetail); err != nil {
				return err
			}
		}

	}

	valuesZone := o.Zone

	joinedZone := swag.JoinByFormat(valuesZone, "multi")
	// query array param zone
	if err := r.SetQueryParam("zone", joinedZone...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
