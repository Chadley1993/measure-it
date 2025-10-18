import { Component, OnInit, PLATFORM_ID, Inject, signal, ElementRef, effect } from '@angular/core';
import { MyService } from '../my-service.service';
import { isPlatformBrowser } from '@angular/common';
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
  svg: any = null
  isBrowser = signal(false);

  constructor(private elRef: ElementRef, private myService: MyService, @Inject(PLATFORM_ID) platformId: object) {
    this.isBrowser.set(isPlatformBrowser(platformId));
    
    effect(() => {
      const latitude = this.myService.sensorData().latitude / 2;
      const longitude = this.myService.sensorData().longitude / 2;
      console.log(longitude, latitude)
      if (latitude != 0 && longitude != 0) {
        this.circle.transition()
          .duration(500)
          .attr("cx", latitude)
          .attr("cy", longitude)
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

    this.svg = d3.select(this.elRef.nativeElement.querySelector('#polygon-container'));
    
    const lineGenerator = d3.line()
      .x(d => d[0])
      .y(d => d[1])
      .curve(d3.curveLinear)

    const pathData = lineGenerator(points);

    this.svg
      .append('path')
      .attr('d', pathData)
      .attr('fill', 'none')
      .attr('stroke', 'rgba(147, 156, 20, 1)')
      .attr('stroke-width', 2);

    let sector1Points: number[] = createSectorPoints(309.33693982357454, 317.86306017669745, 259.34655658270805, 262.2284434167221, 2)
    this.svg
      .append("line")
      .attr("x1", sector1Points[0] / 2)
      .attr("y1", sector1Points[1] / 2)
      .attr("x2", sector1Points[2] / 2)
      .attr("y2", sector1Points[3] / 2)
      .attr('stroke', 'white')
      .attr('stroke-width', 2);
    
    let sector2Points: number[] = createSectorPoints(30.150715243115588, 37.482618089990254, 252.07774902337758, 257.2972509770118, 1)
    this.svg
      .append("line")
      .attr("x1", sector2Points[0] / 2)
      .attr("y1", sector2Points[1] / 2)
      .attr("x2", sector2Points[2] / 2)
      .attr("y2", sector2Points[3] / 2)
      .attr('stroke', 'rgba(248, 71, 248, 1)')
      .attr('stroke-width', 2);
    
    let sector3Points: number[] = createSectorPoints(424.3013224077671, 417.4320109256272, 280.2824284066782, 274.4675715931924, 2)
    this.svg
      .append("line")
      .attr("x1", sector3Points[0] / 2)
      .attr("y1", sector3Points[1] / 2)
      .attr("x2", sector3Points[2] / 2)
      .attr("y2", sector3Points[3] / 2)
      .attr('stroke', 'rgba(248, 71, 248, 1)')
      .attr('stroke-width', 2);

    if (this.isBrowser()) {
      console.log("init run!")
      this.circle = this.svg.append("circle")
        .attr("cx", 200)
        .attr("cy", 110)
        .attr("r", 5)
        .attr("fill", "rgba(40, 230, 255, 1)");
    }
  }
}
function createSectorPoints(x0: number, x1: number, y0: number, y1: number, factor: number) {
  let dx = Math.abs(x0 - x1)
  let dy = Math.abs(y0 - y1)
  
  // p1 = np.array([data[i]["x"][0], data[i]["y"][0]])
  // p2 = np.array([data[i]["x"][1], data[i]["y"][1]])
  // direction = p1 - p2

  let newX0 = x0 - dx * factor
  let newY0 = y0 - dy * factor
  let newX1 = x1 + dx * factor
  let newY1 = y1 + dy * factor
  return [newX0, newY0, newX1, newY1]  
}

