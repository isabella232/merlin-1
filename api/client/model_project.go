/*
 * Merlin
 *
 * API Guide for accessing Merlin's model management, deployment, and serving functionalities
 *
 * API version: 0.14.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

type Project struct {
	Id                int32     `json:"id,omitempty"`
	Name              string    `json:"name"`
	MlflowTrackingUrl string    `json:"mlflow_tracking_url,omitempty"`
	Administrators    []string  `json:"administrators,omitempty"`
	Readers           []string  `json:"readers,omitempty"`
	Team              string    `json:"team,omitempty"`
	Stream            string    `json:"stream,omitempty"`
	Labels            []Label   `json:"labels,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}
