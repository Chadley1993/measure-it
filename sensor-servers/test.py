import requests
from datetime import datetime


start_time = datetime.now()
# for i in range(1000):
data = ['test-sensor-1', 'world', 'gps-1']
# response = requests.get('https://webhook-test.mautic.com/1frcobw1', json=data)
response = requests.get('http://localhost:8080/pitWall-bridge', json=data)

print(response.status_code)
print(response.content)
delta = datetime.now() - start_time
print(delta.total_seconds())