import { Component, effect } from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import { MyService } from '../my-service.service';

@Component({
  selector: 'app-gps-speed',
  standalone: true,
  imports: [MatCardModule],
  templateUrl: './gps-speed.component.html',
  styleUrl: './gps-speed.component.scss',
})
export class GpsSpeedComponent {
  errorMessage: string | null = null;
  speed: any = "x";

  constructor(private myService: MyService) {
    effect(() => {
      const speed = this.myService.sensorData().speedKPH;
      this.speed = speed;
      console.log("speed -> " + speed);
    });
  }
}
