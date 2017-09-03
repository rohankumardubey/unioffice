// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheet

import (
	"baliance.com/gooxml/chart"
	"baliance.com/gooxml/drawing"
	dml "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml"
	crt "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/chart"
)

type LineChartSeries struct {
	x *crt.CT_LineSer
}

// X returns the inner wrapped XML type.
func (c LineChartSeries) X() *crt.CT_LineSer {
	return c.x
}

// Index returns the index of the series
func (c LineChartSeries) Index() uint32 {
	return c.x.Idx.ValAttr
}

// SetIndex sets the index of the series
func (c LineChartSeries) SetIndex(idx uint32) {
	c.x.Idx.ValAttr = idx
}

// Order returns the order of the series
func (c LineChartSeries) Order() uint32 {
	return c.x.Order.ValAttr
}

// SetOrder sets the order of the series
func (c LineChartSeries) SetOrder(idx uint32) {
	c.x.Order.ValAttr = idx
}

// SetText sets the series text
func (c LineChartSeries) SetText(s string) {
	c.x.Tx = crt.NewCT_SerTx()
	c.x.Tx.Choice.V = &s
}

// Properties returns the line chart series shape properties.
func (c LineChartSeries) Properties() drawing.ShapeProperties {
	if c.x.SpPr == nil {
		c.x.SpPr = dml.NewCT_ShapeProperties()
	}
	return drawing.MakeShapeProperties(c.x.SpPr)
}

// Marker returns the marker properties.
func (c LineChartSeries) Marker() chart.Marker {
	if c.x.Marker == nil {
		c.x.Marker = crt.NewCT_Marker()
	}
	return chart.MakeMarker(c.x.Marker)
}

// Labels returns the data label properties.
func (c LineChartSeries) Labels() chart.DataLabels {
	if c.x.DLbls == nil {
		c.x.DLbls = crt.NewCT_DLbls()
	}
	return chart.MakeDataLabels(c.x.DLbls)
}

func (c LineChartSeries) Axis() chart.AxisDataSource {
	if c.x.Cat == nil {
		c.x.Cat = crt.NewCT_AxDataSource()
	}
	return chart.MakeAxisDataSource(c.x.Cat)
}