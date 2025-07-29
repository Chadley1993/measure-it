import math
import time
import requests
import json

theta = 0.1
while True:
    temp = math.sin(theta)
    response = requests.post("http://localhost:8080/data-bridge", json={"temperatureCelsius": temp}, headers={"sensor-name": "test-sensor-1"})
    print(response.status_code)
    time.sleep(0.1)
    theta += 0.1
