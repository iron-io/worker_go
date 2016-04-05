package titan

import (
    "time"
)

type Job struct {
    Image  string  `json:"image,omitempty"`
    Payload  string  `json:"payload,omitempty"`
    Delay  int32  `json:"delay,omitempty"`
    Timeout  int32  `json:"timeout,omitempty"`
    Priority  int32  `json:"priority,omitempty"`
    MaxRetries  int32  `json:"max_retries,omitempty"`
    RetriesDelay  int32  `json:"retries_delay,omitempty"`
    Id  string  `json:"id,omitempty"`
    Status  string  `json:"status,omitempty"`
    GroupName  string  `json:"group_name,omitempty"`
    Reason  Reason  `json:"reason,omitempty"`
    CreatedAt  time.Time  `json:"created_at,omitempty"`
    StartedAt  time.Time  `json:"started_at,omitempty"`
    CompletedAt  time.Time  `json:"completed_at,omitempty"`
    RetryOf  string  `json:"retry_of,omitempty"`
    RetryId  string  `json:"retry_id,omitempty"`
    
}
