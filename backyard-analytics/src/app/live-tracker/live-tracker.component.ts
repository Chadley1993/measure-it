import { Component, OnInit, ElementRef } from '@angular/core';
import * as d3 from 'd3';

@Component({
  selector: 'app-live-tracker',
  standalone: true,
  imports: [],
  templateUrl: './live-tracker.component.html',
  styleUrl: './live-tracker.component.scss'
})
export class LiveTrackerComponent implements OnInit {
  constructor(private elRef: ElementRef) {}
  private trackLineData: any[] = []
  ngOnInit(): void {
    this.trackLineData = require("./resource/killarney_trackline_v1.json")
    this.drawPolygon(this.trackLineData);
  }

  drawPolygon(trackLineData: any[]): void {
    let data: any[] = []
    trackLineData.forEach(datapoint => {
      if (datapoint.type == "trackline") {
        
        data.push([datapoint.x / 2, datapoint.y / 2])
      }
      return data
    })
    const points: Iterable<[number, number]> = data

    const svg = d3.select(this.elRef.nativeElement.querySelector('#polygon-container'));
    
    const lineGenerator = d3.line()
      .x(d => d[0])
      .y(d => d[1])
      .curve(d3.curveLinear)

    const pathData = lineGenerator(points);

    svg
      .append('path')
      .attr('d', pathData)
      .attr('fill', 'none')
      .attr('stroke', 'darkblue')
      .attr('stroke-width', 2);
  }
}
