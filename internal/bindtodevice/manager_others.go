//go:build !linux

package bindtodevice

import (
	"context"
	"net/netip"

	"github.com/AdguardTeam/AdGuardDNS/internal/agd"
	"github.com/AdguardTeam/AdGuardDNS/internal/dnsserver/netext"
	"github.com/AdguardTeam/golibs/errors"
)

// Manager creates individual listeners and dispatches connections to them.
//
// It is only suported on Linux.
type Manager struct{}

// NewManager returns a new manager of interface listeners.
//
// It is only suported on Linux.
func NewManager(c *ManagerConfig) (m *Manager) {
	return &Manager{}
}

// errUnsupported is returned from all [Manager] methods on OSs other than
// Linux.
const errUnsupported errors.Error = "bindtodevice is only supported on linux"

// Add creates a new interface-listener record in m.
//
// It is only suported on Linux.
func (m *Manager) Add(id ID, ifaceName string, port uint16) (err error) { return errUnsupported }

// ListenConfig returns a new netext.ListenConfig that receives connections from
// the interface listener with the given id and the destination addresses of
// which fall within subnet.  subnet should be masked.
//
// It is only suported on Linux.
func (m *Manager) ListenConfig(id ID, subnet netip.Prefix) (c netext.ListenConfig, err error) {
	return nil, errUnsupported
}

// type check
var _ agd.Service = (*Manager)(nil)

// Start implements the [agd.Service] interface for *Manager.
//
// It is only suported on Linux.
func (m *Manager) Start() (err error) { return errUnsupported }

// Shutdown implements the [agd.Service] interface for *Manager.
//
// It is only suported on Linux.
func (m *Manager) Shutdown(_ context.Context) (err error) { return errUnsupported }
