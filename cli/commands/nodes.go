package commands

import (
	"fmt"
	"strings"

	"github.com/kontena/kontena-client-go/api"
	"github.com/kontena/kontena-client-go/cli"
	"github.com/kontena/kontena-client-go/client"
)

func nodeStatus(node api.Node) string {
	if node.Connected {
		return "online"
	} else {
		return "offline"
	}
}

func nodeInitial(node api.Node) string {
	if node.InitialMember {
		return fmt.Sprintf("%d / %d", node.NodeNumber, node.Grid.InitialSize)
	} else {
		return "-"
	}
}

func nodeLabels(node api.Node) string {
	return strings.Join(node.Labels, ",")
}

func printNodes(nodes []api.Node) {
	tbl := cli.Table("Name", "Version", "Status", "Initial", "Labels")

	for _, node := range nodes {
		tbl.AddRow(node.Name,
			node.AgentVersion,
			nodeStatus(node),
			nodeInitial(node),
		)
	}

	tbl.Print()
}

type NodesCommand struct {
	*cli.CLI
	Grid string
}

func (cmd NodesCommand) List() error {
	if nodes, err := cmd.Client.Nodes.List(cmd.Grid); err != nil {
		return err
	} else {
		printNodes(nodes)
	}

	return nil
}

func (cmd NodesCommand) Show(name string) error {
	if node, err := cmd.Client.Nodes.Get(client.NodeID{cmd.Grid, name}); err != nil {
		return err
	} else {
		return cli.Print(node)
	}
}
