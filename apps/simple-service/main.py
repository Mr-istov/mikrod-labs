import hug
import socket
import os
import requests
import json

COMPLEX_SERVICE_URL = os.getenv("COMPLEX_SERVICE_URL")


@hug.get("/calculate")
def calculate(num: hug.types.number):
    response = requests.get(f"{COMPLEX_SERVICE_URL}/sum?number={num}")
    result = json.loads(response.content)

    return {
        "message": f"{socket.gethostname()} sent the number to the algorithm and got: {result['message']}"
    }
