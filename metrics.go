package metrics

import (
	"expvar"
)

type counter struct {
	v expvar.Float
}

type gauge struct {
	v expvar.Float
	t expvar.Var
}

type histogram struct {
}

type Metrics struct {
	counters   map[string]*counter
	gaugs      map[string]*gauge
	histograms map[string]histogram
}

func (m *Metrics) Start() {

}

func (m *Metrics) Stop() {

}

func (m *Metrics) Counter(name string) *counter {
	if c, ok := m.counters[name]; ok {
		return c
	}
	var newCounter counter
	m.counters[name] = &newCounter
	return &newCounter
}

func (m *Metrics) Gauge(name string) *gauge {
	if g, ok := m.gaugs[name]; ok {
		return g
	}
	var newGauge gauge
	m.gaugs[name] = &newGauge
	return &newGauge
}


// m.Counter("DevOps.Test").Add(2)
// m.Counter("DevOps.Test").Reset()

func (c *counter) Add(delta float64) {
	c.v.Add(delta)
}

func (g *gauge) Set(value float64) {
	g.v.Set(value)
}

func (g *gauge) Get() float64 {

}

func (g *gauge) GetAndReset() float64 {
	
}

func (g *gauge) Add(delta float64) {
	g.v.Add(delta)
}
