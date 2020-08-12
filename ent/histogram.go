// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/gobench-io/gobench/ent/histogram"
	"github.com/gobench-io/gobench/ent/metric"
)

// Histogram is the model entity for the Histogram schema.
type Histogram struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Time holds the value of the "time" field.
	Time int64 `json:"time"`
	// Count holds the value of the "count" field.
	Count int64 `json:"count"`
	// Min holds the value of the "min" field.
	Min int64 `json:"min"`
	// Max holds the value of the "max" field.
	Max int64 `json:"max"`
	// Mean holds the value of the "mean" field.
	Mean float64 `json:"mean"`
	// Stddev holds the value of the "stddev" field.
	Stddev float64 `json:"stddev"`
	// Median holds the value of the "median" field.
	Median float64 `json:"median"`
	// P75 holds the value of the "p75" field.
	P75 float64 `json:"p75"`
	// P95 holds the value of the "p95" field.
	P95 float64 `json:"p95"`
	// P99 holds the value of the "p99" field.
	P99 float64 `json:"p99"`
	// P999 holds the value of the "p999" field.
	P999 float64 `json:"p999"`
	// WID holds the value of the "wID" field.
	WID string `json:"wId"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HistogramQuery when eager-loading is set.
	Edges             HistogramEdges `json:"edges"`
	metric_histograms *int
}

// HistogramEdges holds the relations/edges for other nodes in the graph.
type HistogramEdges struct {
	// Metric holds the value of the metric edge.
	Metric *Metric
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MetricOrErr returns the Metric value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HistogramEdges) MetricOrErr() (*Metric, error) {
	if e.loadedTypes[0] {
		if e.Metric == nil {
			// The edge metric was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: metric.Label}
		}
		return e.Metric, nil
	}
	return nil, &NotLoadedError{edge: "metric"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Histogram) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullInt64{},   // time
		&sql.NullInt64{},   // count
		&sql.NullInt64{},   // min
		&sql.NullInt64{},   // max
		&sql.NullFloat64{}, // mean
		&sql.NullFloat64{}, // stddev
		&sql.NullFloat64{}, // median
		&sql.NullFloat64{}, // p75
		&sql.NullFloat64{}, // p95
		&sql.NullFloat64{}, // p99
		&sql.NullFloat64{}, // p999
		&sql.NullString{},  // wID
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Histogram) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // metric_histograms
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Histogram fields.
func (h *Histogram) assignValues(values ...interface{}) error {
	if m, n := len(values), len(histogram.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	h.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field time", values[0])
	} else if value.Valid {
		h.Time = value.Int64
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field count", values[1])
	} else if value.Valid {
		h.Count = value.Int64
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field min", values[2])
	} else if value.Valid {
		h.Min = value.Int64
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field max", values[3])
	} else if value.Valid {
		h.Max = value.Int64
	}
	if value, ok := values[4].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field mean", values[4])
	} else if value.Valid {
		h.Mean = value.Float64
	}
	if value, ok := values[5].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field stddev", values[5])
	} else if value.Valid {
		h.Stddev = value.Float64
	}
	if value, ok := values[6].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field median", values[6])
	} else if value.Valid {
		h.Median = value.Float64
	}
	if value, ok := values[7].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field p75", values[7])
	} else if value.Valid {
		h.P75 = value.Float64
	}
	if value, ok := values[8].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field p95", values[8])
	} else if value.Valid {
		h.P95 = value.Float64
	}
	if value, ok := values[9].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field p99", values[9])
	} else if value.Valid {
		h.P99 = value.Float64
	}
	if value, ok := values[10].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field p999", values[10])
	} else if value.Valid {
		h.P999 = value.Float64
	}
	if value, ok := values[11].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field wID", values[11])
	} else if value.Valid {
		h.WID = value.String
	}
	values = values[12:]
	if len(values) == len(histogram.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field metric_histograms", value)
		} else if value.Valid {
			h.metric_histograms = new(int)
			*h.metric_histograms = int(value.Int64)
		}
	}
	return nil
}

// QueryMetric queries the metric edge of the Histogram.
func (h *Histogram) QueryMetric() *MetricQuery {
	return (&HistogramClient{config: h.config}).QueryMetric(h)
}

// Update returns a builder for updating this Histogram.
// Note that, you need to call Histogram.Unwrap() before calling this method, if this Histogram
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Histogram) Update() *HistogramUpdateOne {
	return (&HistogramClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (h *Histogram) Unwrap() *Histogram {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Histogram is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Histogram) String() string {
	var builder strings.Builder
	builder.WriteString("Histogram(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteString(", time=")
	builder.WriteString(fmt.Sprintf("%v", h.Time))
	builder.WriteString(", count=")
	builder.WriteString(fmt.Sprintf("%v", h.Count))
	builder.WriteString(", min=")
	builder.WriteString(fmt.Sprintf("%v", h.Min))
	builder.WriteString(", max=")
	builder.WriteString(fmt.Sprintf("%v", h.Max))
	builder.WriteString(", mean=")
	builder.WriteString(fmt.Sprintf("%v", h.Mean))
	builder.WriteString(", stddev=")
	builder.WriteString(fmt.Sprintf("%v", h.Stddev))
	builder.WriteString(", median=")
	builder.WriteString(fmt.Sprintf("%v", h.Median))
	builder.WriteString(", p75=")
	builder.WriteString(fmt.Sprintf("%v", h.P75))
	builder.WriteString(", p95=")
	builder.WriteString(fmt.Sprintf("%v", h.P95))
	builder.WriteString(", p99=")
	builder.WriteString(fmt.Sprintf("%v", h.P99))
	builder.WriteString(", p999=")
	builder.WriteString(fmt.Sprintf("%v", h.P999))
	builder.WriteString(", wID=")
	builder.WriteString(h.WID)
	builder.WriteByte(')')
	return builder.String()
}

// Histograms is a parsable slice of Histogram.
type Histograms []*Histogram

func (h Histograms) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}