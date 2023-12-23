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

// Return the path of the connector file
func ConnectorFileOf(dir types.Path) types.Path {
	return dir.Join(constants.CONNECTOR)
}

// Create a connector file in the directory
func CreateConnectorFile(dir types.Path) error {
	defaultConnections := types.PathSet{}
	defaultConnections.AddPath(dir)

	return ConnectorFileOf(dir).Write(defaultConnections)
}

// Return a new connector object.
// The connector object contains the paths of the connected directories.
func NewConnector(dir types.Path) (*connector, error) {
	connectorFile := ConnectorFileOf(dir)
	connections, err := connectorFile.Read()
	if err != nil {
		return nil, err
	}

	return &connector{dir, connectorFile, connections}, nil
}

// Return the paths of the connected directories
func (c connector) Connections() types.PathSet {
	return c.connectedPaths
}

// Check if the directory is connected to the given directory
func (c connector) isConnected(to types.Path) bool {
	return c.connectedPaths.ContainsPath(to)
}

// Return error if the directory is connected to the given directory
func (c connector) ErrIfConnected(to types.Path) error {
	if c.isConnected(to) {
		return alreadyConnectedError{c.base, to}
	} else {
		return connectedToOtherDirectoriesError{to}
	}
}

// Check if the directory has any connections
func (c connector) hasConnection() bool {
	return c.connectedPaths.Len() > 0
}

// Return error if the directory has no connections
func (c connector) ErrIfNoConnections() error {
	if c.hasConnection() {
		return nil
	} else {
		return noConnectionsError{c.base}
	}
}

// Connect the directory to the given directory.
// Notice that this function doesn't save the changes.
func (c *connector) Connect(to types.Path) {
	c.connectedPaths.AddPath(to)
}

// Disconnect the directory from the given directory.
// Notice that this function doesn't save the changes.
func (c *connector) Disconnect(to types.Path) {
	c.connectedPaths.RemovePath(to)
}

// Save the changes to the connector file
func (c connector) Save() error {
	return c.file.Write(c.connectedPaths)
}
