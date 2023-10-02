import time
from paho.mqtt import client as mqtt_client

def on_connect(client, userdata, flags, rc):
    if rc == 0:
        print("Connected to MQTT Broker!")
    else:
        print("Failed to connect, return code %d\n", rc)
        
client = mqtt_client.Client()
client.on_connect = on_connect
client.connect("localhost")

topic = "blabla"

msg_count = 1
while True:
    time.sleep(1)
    msg = f"Messages: {msg_count}"
    result = client.publish(topic, msg)

    if result[0] == 0:
        print(f"Send `{msg}` to topic `{topic}`")
    else:
        print(f"Failed to send message to topic {topic}")

    msg_count += 1
    if msg_count > 5:
        break