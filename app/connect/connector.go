package connect

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

func (c connector) HasConnection(to types.Path) bool {
	return c.connectedPaths.ContainsPath(to)
}

func (c connector) ErrIfConnected(to types.Path) error {
	if c.HasConnection(to) {
		return alreadyConnectedError{c.base, to}
	} else {
		return connectedToOtherDirectoriesError{to}
	}
}

func (c *connector) Connect(to types.Path) {
	c.connectedPaths.AddPath(to)
}
