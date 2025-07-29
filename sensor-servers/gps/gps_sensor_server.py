import serial
from datetime import datetime
import time
import requests

SERIAL_PORT = "/dev/ttyACM0"
running = True

START_TIME = datetime.now()
gps_payload = {'longitude': 0.0, 'latitude': 0.0, 'speedKPH': 0.0}
# In the NMEA message, the position gets transmitted as:
# DDMM.MMMMM, where DD denotes the degrees and MM.MMMMM denotes
# the minutes. However, I want to convert this format to the following:
# DD.MMMM. This method converts a transmitted string to the desired format
def formatDegreesMinutes(coordinates, digits):
    
    parts = coordinates.split(".")

    if (len(parts) != 2):
        print("something happened.")
        return coordinates

    if (digits > 3 or digits < 2):
        print("something weirder happened!")
        return coordinates
    
    left = parts[0]
    right = parts[1]
    degrees = float(left[:digits])
    minutes = float(coordinates[digits:]) / 60.0

    return degrees + minutes

def getPositionData(gps):
    data = gps.readline()
    message = data[0:6].decode("utf-8")
    gps_data = data[6:].decode("utf-8")
    
    if (message == "$GPRMC"):
        parts = gps_data.split(",")
        if parts[2] == 'V':
            print("GPS receiver warning")
        else:
            gps_payload["latitude"] = formatDegreesMinutes(parts[3], 2)
            gps_payload["longitude"] = formatDegreesMinutes(parts[5], 3)
            requests.post("http://localhost:8080/data-bridge", json=gps_payload, headers={'sensor-name': 'gps-1'})
    elif (message == "$GPVTG"):
        delta = datetime.now() - START_TIME
        parts = gps_data.split(",")
        gps_payload["speedKPH"] = int(float(parts[7]))
        requests.post("http://localhost:8080/data-bridge", json=gps_payload, headers={'sensor-name': 'gps-1'})
    else:
        pass

print("Application started!")
gps = serial.Serial(SERIAL_PORT, baudrate = 9600, timeout = 0.5)

while running:
    try:
        getPositionData(gps)
        START_TIME = datetime.now()
        time.sleep(0.1)
    except KeyboardInterrupt:
        running = False
        gps.close()
        print("Application closed!")
    except Exception as e:
        print ("Application error!", e)
