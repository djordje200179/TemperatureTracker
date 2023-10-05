from atexit import register
import time
import requests

device_name = "Pi Zero"
device_key = "test123"
sensor_name = "DS1820"

base_url = "http://localhost:8081"

def register_device():
	body = {
		"name": device_name,
		"key": device_key
	}
	
	response = requests.post(base_url + "/devices", json=body)
	if response.status_code == 409:
		print("Device already registered")
	elif response.status_code == 201:
		print("Device registered")
	else:
		response.raise_for_status()

jwt: str
def login():
	body = {
		"name": device_name,
		"key": device_key
	}
	
	response = requests.post(base_url + "/devices/auth", json=body)
	response.raise_for_status()
	
	print("Logged in successfully")
	
	global jwt
	jwt = response.text
	
def register_sensor():
	body = {
		"name": sensor_name
	}
	
	while True:
		headers = {
			"Authorization": "Bearer " + jwt
		}
		
		response = requests.post(base_url + "/sensors", json=body, headers=headers)
		match response.status_code:
			case 409:
				print("Sensor already registered")
				break
			case 201:
				print("Sensor registered")
				break
			case 401:
				print("Login expired, logging in again...")
				login()
			case _:
				response.raise_for_status()
	
def send_data(temperature: float):
	body = {
		"sensor": sensor_name,
		"temperature": temperature
	}
	
	while True:
		headers = {
			"Authorization": "Bearer " + jwt
		}
	
		response = requests.post(base_url + "/readings", json=body, headers=headers)
		match response.status_code:
			case 201:
				print("Data sent")
				break
			case 401:
				print("Login expired, logging in again...")
				login()
			case _:
				response.raise_for_status()
			
def read_temp():
	return 39

def main():
	register_device()
	login()
	
	register_sensor()
	
	while True:
		temperature = read_temp()
		print(f"Temperature: {temperature}")

		send_data(temperature)
		
		print("Sleeping...")
		time.sleep(5)
		
if __name__ == "__main__":
	main()