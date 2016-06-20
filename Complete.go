/* 
 * Titan API
 *
 * The ultimate, language agnostic, container based job processing framework.
 *
 * OpenAPI spec version: 0.4.0
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

type Complete struct {

	// Time when job was completed. Always in UTC.
	CompletedAt time.Time `json:"completed_at,omitempty"`

	// Machine readable reason failure, if status=error. Only used by the /error endpoint.
	Reason string `json:"reason,omitempty"`

	// Error message, if status=error. Only used by the /error endpoint.
	Error_ string `json:"error,omitempty"`
}
