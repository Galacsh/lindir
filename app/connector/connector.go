package connector

import (
	"lindir/common/constants"
	"lindir/common/types"
)

type connector struct {
	base           types.Path
	file           types.Path
	connectedPaths types.PathSet
}

func ConnectorFileOf(dir types.Path) types.Path {
	return dir.Join(constants.CONNECTOR)
}

func CreateConnectorFile(dir types.Path) error {
	defaultConnections := types.PathSet{}
	defaultConnections.AddPath(dir)

	return ConnectorFileOf(dir).Write(defaultConnections)
}

func NewConnector(dir types.Path) (*connector, error) {
	connectorFile := ConnectorFileOf(dir)
	connections, err := connectorFile.Read()
	if err != nil {
		return nil, err
	}

	return &connector{dir, connectorFile, connections}, nil
}

func (c connector) Connections() types.PathSet {
	return c.connectedPaths
}

func (c connector) IsConnected(to types.Path) bool {
	return c.connectedPaths.ContainsPath(to)
}

func (c connector) ErrIfConnected(to types.Path) error {
	if c.IsConnected(to) {
		return alreadyConnectedError{c.base, to}
	} else {
		return connectedToOtherDirectoriesError{to}
	}
}

func (c connector) HasConnection() bool {
	return len(c.connectedPaths) > 0
}

func (c connector) ErrIfNoConnections() error {
	if c.HasConnection() {
		return nil
	} else {
		return noConnectionsError{c.base}
	}
}

func (c *connector) Connect(to types.Path) {
	c.connectedPaths.AddPath(to)
}

func (c *connector) Disconnect(to types.Path) {
	c.connectedPaths.RemovePath(to)
}

func (c connector) Save() error {
	return c.file.Write(c.connectedPaths)
}
