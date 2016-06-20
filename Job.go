/* 
 * Titan API
 *
 * The ultimate, language agnostic, container based job processing framework.
 *
 * OpenAPI spec version: 0.4.6
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package titan

import (
	"time"
)

type Job struct {

	// Name of Docker image to use. This is optional and can be used to override the image defined at the group level.
	Image string `json:"image,omitempty"`

	// Payload for the job. This is what you pass into each job to make it do something.
	Payload string `json:"payload,omitempty"`

	// Number of seconds to wait before queueing the job for consumption for the first time. Must be a positive integer. Jobs with a delay start in state \"delayed\" and transition to \"running\" after delay seconds.
	Delay int32 `json:"delay,omitempty"`

	// Maximum runtime in seconds. If a consumer retrieves the job, but does not change it's status within timeout seconds, the job is considered failed, with reason timeout (Titan may allow a small grace period). The consumer should also kill the job after timeout seconds. If a consumer tries to change status after Titan has already timed out the job, the consumer will be ignored. 
	Timeout int32 `json:"timeout,omitempty"`

	// Priority of the job. Higher has more priority. 3 levels from 0-2. Jobs at same priority are processed in FIFO order.
	Priority int32 `json:"priority,omitempty"`

	// \"Number of automatic retries this job is allowed. A retry will be attempted if a task fails. Max 25. Automatic retries are performed by titan when a task reaches a failed state and has `max_retries` > 0. A retry is performed by queueing a new job with the same image id and payload. The new job's max_retries is one less than the original. The new job's `retry_of` field is set to the original Job ID.  Titan will delay the new job for retries_delay seconds before queueing it. Cancelled or successful tasks are never automatically retried.\" 
	MaxRetries int32 `json:"max_retries,omitempty"`

	// Time in seconds to wait before retrying the job. Must be a non-negative integer.
	RetriesDelay int32 `json:"retries_delay,omitempty"`

	// Unique identifier representing a specific job.
	Id string `json:"id,omitempty"`

	// States and valid transitions.                   +---------+        +---------> delayed <----------------+                  +----+----+                |                       |                     |                       |                     |                  +----v----+                |        +---------> queued  <----------------+                  +----+----+                *                       |                     *                       |               retry * creates new job                  +----v----+                *                  | running |                *                  +--+-+-+--+                |           +---------|-|-|-----+-------------+       +---|---------+ | +-----|---------+   |       |   |           |       |         |   | +-----v---^-+      +--v-------^+     +--v---^-+ | success   |      | cancelled |     |  error | +-----------+      +-----------+     +--------+  * delayed - has a delay. * queued - Ready to be consumed when it's turn comes. * running - Currently consumed by a runner which will attempt to process it. * success - (or complete? success/error is common javascript terminology) * error - Something went wrong. In this case more information can be obtained   by inspecting the \"reason\" field.   - timeout   - killed - forcibly killed by worker due to resource restrictions or access     violations.   - bad_exit - exited with non-zero status due to program termination/crash. * cancelled - cancelled via API. More information in the reason field.   - client_request - Request was cancelled by a client. 
	Status string `json:"status,omitempty"`

	// Group this job belongs to.
	GroupName string `json:"group_name,omitempty"`

	// The error message, if status is 'error'. This is errors due to things outside the job itself. Errors from user code will be found in the log.
	Error_ string `json:"error,omitempty"`

	// Machine usable reason for job being in this state. Valid values for error status are `timeout | killed | bad_exit`. Valid values for cancelled status are `client_request`. For everything else, this is undefined. 
	Reason string `json:"reason,omitempty"`

	// Time when job was submitted. Always in UTC.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Time when job started execution. Always in UTC.
	StartedAt time.Time `json:"started_at,omitempty"`

	// Time when job completed, whether it was successul or failed. Always in UTC.
	CompletedAt time.Time `json:"completed_at,omitempty"`

	// If this field is set, then this job is a retry of the ID in this field.
	RetryOf string `json:"retry_of,omitempty"`

	// If this field is set, then this job was retried by the job referenced in this field.
	RetryAt string `json:"retry_at,omitempty"`

	// Env vars for the task. Comes from the ones set on the Group.
	EnvVars map[string]string `json:"env_vars,omitempty"`
}
