package command

import (
	"fmt"
	"strings"

	"github.com/ryanuber/columnize"
)

type NodeStatusCommand struct {
	Meta
}

func (c *NodeStatusCommand) Help() string {
	helpText := `
Usage: nomad node-status [options] [node]

  Display status information about a given node. The list of nodes
  returned includes only nodes which jobs may be scheduled to, and
  includes status and other high-level information.

  If a node ID is passed, information for that specific node will
  be displayed. If no node ID's are passed, then a short-hand
  list of all nodes will be displayed.

General Options:

  ` + generalOptionsUsage()
	return strings.TrimSpace(helpText)
}

func (c *NodeStatusCommand) Synopsis() string {
	return "Display status information about nodes"
}

func (c *NodeStatusCommand) Run(args []string) int {
	flags := c.Meta.FlagSet("node-status", FlagSetClient)
	flags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := flags.Parse(args); err != nil {
		return 1
	}

	// Check that we got either a single node or none
	args = flags.Args()
	if len(args) > 1 {
		c.Ui.Error(c.Help())
		return 1
	}

	// Get the HTTP client
	client, err := c.Meta.Client()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error initializing client: %s", err))
		return 1
	}

	// Use list mode if no node name was provided
	if len(args) == 0 {
		// Query the node info
		nodes, _, err := client.Nodes().List(nil)
		if err != nil {
			c.Ui.Error(fmt.Sprintf("Error querying node status: %s", err))
			return 1
		}

		// Return nothing if no nodes found
		if len(nodes) == 0 {
			return 0
		}

		// Format the nodes list
		out := make([]string, len(nodes)+1)
		out[0] = "ID|DC|Name|Class|Drain|Status"
		for i, node := range nodes {
			out[i+1] = fmt.Sprintf("%s|%s|%s|%s|%v|%s",
				node.ID,
				node.Datacenter,
				node.Name,
				node.NodeClass,
				node.Drain,
				node.Status)
		}

		// Dump the output
		c.Ui.Output(columnize.SimpleFormat(out))
		return 0
	}

	// Query the specific node
	nodeID := args[0]
	node, _, err := client.Nodes().Info(nodeID, nil)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error querying node info: %s", err))
		return 1
	}

	// Format the output
	out := []string{
		fmt.Sprintf("ID | %s", node.ID),
		fmt.Sprintf("Name | %s", node.Name),
		fmt.Sprintf("Class | %s", node.NodeClass),
		fmt.Sprintf("Datacenter | %s", node.Datacenter),
		fmt.Sprintf("Drain | %v", node.Drain),
		fmt.Sprintf("Status | %s", node.Status),
	}

	// Make the column config so we can dump k = v pairs
	columnConf := columnize.DefaultConfig()
	columnConf.Glue = " = "

	// Dump the output
	c.Ui.Output(columnize.Format(out, columnConf))
	return 0
}
