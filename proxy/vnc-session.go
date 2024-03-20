package proxy

type SessionStatus int
type SessionType int

const (
	SessionStatusInit SessionStatus = iota
	SessionStatusActive
	SessionStatusError
)

const (
	SessionTypeRecordingProxy SessionType = iota
	SessionTypeReplayServer
	SessionTypeProxyPass
)

type VncSession struct {
	Target         string
	TargetPassword string
	ID             string
	Status         SessionStatus
	Type           SessionType
}
