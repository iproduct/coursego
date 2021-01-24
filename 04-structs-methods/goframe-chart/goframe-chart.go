package main

import (
	"context"
	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/plot"
	wc "github.com/rocketlaunchr/dataframe-go/plot/wcharczuk/go-chart"
	chart "github.com/wcharczuk/go-chart"
)

func main() {
	sales := dataframe.NewSeriesFloat64("sales", nil, 50.3, nil, 23.4, 56.2, 89, 32, 84.2, 72, 89)
	ctx := context.Background()
	cs, _ := wc.S(ctx, sales, nil, nil)

	graph := chart.Chart{Series: []chart.Series{cs}}

	plt, _ := plot.Open("Monthly sales", 450, 300)
	graph.Render(chart.SVG, plt)
	plt.Display(plot.None)
	<-plt.Closed

}
