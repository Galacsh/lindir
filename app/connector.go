package app

import (
	"lindir/common/constants"
	"lindir/common/types"
)

type connector struct {
	file           types.Path
	connectedPaths types.PathSet
}

func createConnectorFile(dir types.Path) error {
	defaultConnections := types.PathSet{}
	defaultConnections.AddPath(dir)

	return connectorFileOf(dir).Write(defaultConnections)
}

func connectorFileOf(dir types.Path) types.Path {
	return dir.Join(constants.CONNECTOR)
}

func newConnector(dir types.Path) (*connector, error) {
	connectorFile := connectorFileOf(dir)
	connections, err := connectorFile.Read()
	if err != nil {
		return nil, err
	}

	return &connector{connectorFile, connections}, nil
}

func (c connector) connections() types.PathSet {
	return c.connectedPaths
}

func (c connector) hasConnection(to types.Path) bool {
	return c.connectedPaths.ContainsPath(to)
}

func (c *connector) connect(to types.Path) {
	c.connectedPaths.AddPath(to)
}
