package quic

import "github.com/quic-go/quic-go/internal/protocol"

type BandwidthEstimator interface {
	BandwidthEstimate() protocol.ByteCount
}
