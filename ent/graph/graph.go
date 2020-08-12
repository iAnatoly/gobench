// Code generated by entc, DO NOT EDIT.

package graph

const (
	// Label holds the string label denoting the graph type in the database.
	Label = "graph"
	// FieldID holds the string denoting the id field in the database.
	FieldID    = "id"    // FieldTitle holds the string denoting the title vertex property in the database.
	FieldTitle = "title" // FieldUnit holds the string denoting the unit vertex property in the database.
	FieldUnit  = "unit"

	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeMetrics holds the string denoting the metrics edge name in mutations.
	EdgeMetrics = "metrics"

	// Table holds the table name of the graph in the database.
	Table = "graphs"
	// GroupTable is the table the holds the group relation/edge.
	GroupTable = "graphs"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_graphs"
	// MetricsTable is the table the holds the metrics relation/edge.
	MetricsTable = "metrics"
	// MetricsInverseTable is the table name for the Metric entity.
	// It exists in this package in order to avoid circular dependency with the "metric" package.
	MetricsInverseTable = "metrics"
	// MetricsColumn is the table column denoting the metrics relation/edge.
	MetricsColumn = "graph_metrics"
)

// Columns holds all SQL columns for graph fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldUnit,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Graph type.
var ForeignKeys = []string{
	"group_graphs",
}