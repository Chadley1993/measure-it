import requests
import math
import time

gps_payload = {'longitude': 0.0, 'latitude': 0.0, 'speedKPH': 0.0}

theta = 0.0
while True:
    speed = math.sin(theta % 3) * 200
    gps_payload['speedKPH'] = int(speed)
    requests.post("http://localhost:8080/data-bridge", json=gps_payload, headers={'sensor-name': 'gps-1'})
    theta += 0.1
    time.sleep(1)
