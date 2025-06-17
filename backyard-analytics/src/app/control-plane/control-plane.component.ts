import { Component, OnInit, ElementRef } from '@angular/core';
import { MatButtonModule, } from '@angular/material/button';
import { MatSlideToggleModule, MatSlideToggleChange } from '@angular/material/slide-toggle';
import { MyService } from '../my-service.service';
import * as d3 from 'd3';

@Component({
  selector: 'app-control-plane',
  standalone: true,
  imports: [MatSlideToggleModule, MatButtonModule],
  templateUrl: './control-plane.component.html',
  styleUrl: './control-plane.component.scss'
})
export class ControlPlaneComponent implements OnInit {
  // private theService: MyService;
  connectionBtnColor: string = '#424242';
  recordBtnColor: string = '#424242';
  isRecordDisabled: boolean = true;
  isConnectionDisabled: boolean = false;

  constructor(private myService: MyService, private elRef: ElementRef) {
    // this.theService = myService;
  }

  ngOnInit(): void {
    // const svg = d3.select(this.elRef.nativeElement.querySelector('#delta-container'));

    // const points = "20,25 15,35 25,35";

    // // Append the triangle to the SVG
    // svg.append("polygon")
    //     .attr("points", points)
    //     .attr("fill", "red"); // Triangle color

    const svg = d3.select(this.elRef.nativeElement.querySelector('#delta-container'));

    const points = "20,35 15,25 25,25";

    // Append the triangle to the SVG
    svg.append("polygon")
        .attr("points", points)
        .attr("fill", "#00FF00");
  }

  toggleState: boolean = false;

  onToggleConnect() {
    console.log(this.connectionBtnColor)
    if (this.connectionBtnColor == '#424242') {
      this.connectionBtnColor = 'red';
      this.isRecordDisabled = false;
    } else {
      this.connectionBtnColor = '#424242';
      this.isRecordDisabled = true;
    }
  }

  onToggleRecord() {
    if (this.recordBtnColor == '#424242') {
      this.recordBtnColor = 'red';
      this.isConnectionDisabled = true;
    } else {
      this.recordBtnColor = '#424242';
      this.isConnectionDisabled = false;
    }
  }

  onToggleGPSPrimary(event: MatSlideToggleChange) {
    console.log('Toggle state:', event.checked);

    if (event.checked) {
      this.startStreaming();
    } else {
      this.endStreaming();
    }
  }

  onToggleGPSSecondary(event: MatSlideToggleChange) {
    console.log('Toggle state:', event.checked);

    if (event.checked) {
      this.startStreaming();
    } else {
      this.endStreaming();
    }
  }

  startStreaming() {
    console.log("starting...")
  }

  endStreaming() {
    console.log("...ending")
  }
}
