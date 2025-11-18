import json
import requests
import time

f = open("trackday_session_2.json", "r")
data = json.load(f)
# longitude_min, longitude_max = 1000, 0
# latitude_min, latitude_max = 1000, 0
f.close()

for d in data:
    if 'latitude' in d:
        requests.post("http://localhost:8080/sensorData", json={"sensorName": "gps-position-1", "longitude": d["longitude"], "latitude": d["latitude"]})
        time.sleep(0.1)
#     if d["x"] < longitude_min:
#         longitude_min = d['x']
#     if d["x"] > longitude_max:
#         longitude_max = d['x']

#     if d["y"] < latitude_min:
#         latitude_min = d['y']
#     if d["y"] > latitude_max:
#         latitude_max = d['y']

# print(longitude_min, longitude_max)
# print(latitude_min, latitude_max)