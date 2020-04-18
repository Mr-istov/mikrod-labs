import hug
import socket


@hug.get("/sum")
def complex_algorithm(number: hug.types.number):
    res = sum(num for num in range(1, number + 1))
    return {
        "message": f"The sum from 1 to {number} is {res}. From {socket.gethostname()}"
    }
