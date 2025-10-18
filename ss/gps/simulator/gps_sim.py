import json
import requests
import time

f = open("test_points.json", "r")
data = json.load(f)
longitude_min, longitude_max = 1000, 0
latitude_min, latitude_max = 1000, 0
f.close()

for d in data:
    requests.post("http://localhost:8080/test", json={"sensorName": "gps-position-1", "y": d["y"], "x": d["x"]})
    time.sleep(1)
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