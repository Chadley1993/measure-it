import { Component, OnInit, ElementRef, effect } from '@angular/core';
import { MyService } from '../my-service.service';
import * as d3 from 'd3';

@Component({
  selector: 'app-live-tracker',
  standalone: true,
  imports: [],
  templateUrl: './live-tracker.component.html',
  styleUrl: './live-tracker.component.scss'
})
export class LiveTrackerComponent implements OnInit {
  circle: any = null
  constructor(private elRef: ElementRef, private myService: MyService) {
    effect(() => {
      const latitude = this.myService.sensorData().latitude / 2;
      const longitude = this.myService.sensorData().longitude / 2;
      if (latitude != 0 && longitude != 0) {
        this.circle.transition()
          .duration(2000)
          .attr("cx", latitude)
          .attr("cy", latitude)
      }
    });
  }

  private trackLineData: any[] = []
  ngOnInit(): void {
    this.trackLineData = require("./resource/killarney_trackline_v1.json")
    this.drawPolygon(this.trackLineData);
  }

  drawPolygon(trackLineData: any[]): void {
    console.log("drawPoly")
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
      .attr('stroke', 'rgba(147, 156, 20, 1)')
      .attr('stroke-width', 2);

    this.circle = svg.append("circle")
      .attr("cx", 301.56666666663057 / 2)
      .attr("cy", 320.63749999995395 / 2)
      .attr("r", 5)
      .attr("fill", "rgba(68, 196, 185, 1)");
  }
}
