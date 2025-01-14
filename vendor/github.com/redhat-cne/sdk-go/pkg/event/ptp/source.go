// Copyright 2021 The Cloud Native Events Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ptp

// EventSource ...
type EventSource string

const (
	// GnssSyncStatus notification is signalled from equipment at state change
	GnssSyncStatus EventSource = "/sync/gnss-status/gnss-sync-status"

	// OsClockSyncState State of node OS clock synchronization is notified at state change
	OsClockSyncState EventSource = "/sync/sync-status/os-clock-sync-state"

	// PtpClockClass notification is generated when the clock-class changes.
	PtpClockClass EventSource = "/sync/ptp-status/ptp-clock-class-change"

	// PtpLockState notification is signalled from equipment at state change
	PtpLockState EventSource = "/sync/ptp-status/lock-state"

	// SynceClockQuality notification is generated when the clock-quality changes.
	SynceClockQuality EventSource = "/sync/synce-status/clock-quality"

	// SynceLockState Notification used to inform about synce synchronization state change
	SynceLockState EventSource = "/sync/synce-status/lock-state"

	// SynceLockStateExtended notification is signalled from equipment at state change, enhanced information
	SynceLockStateExtended EventSource = "/sync/synce-status/lock-state-extended"

	// SyncStatusState State of equipment synchronization is notified at state change
	SyncStatusState EventSource = "/sync/sync-status/sync-state"
)
