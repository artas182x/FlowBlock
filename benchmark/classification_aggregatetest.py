
# Start and completion date can be read from backend logs
# Testing running 2 classification alghoritm + aggregating results to chart that is uploaded to minio.
import time
import time
import datetime
from aiohttp import request
import requests

dateToStart = datetime.date(2020,1,1)
dateToStartTimestamp = int(time.mktime(dateToStart.timetuple()))

dateToEnd = datetime.date(2020,1,1)
dateToEndTimestamp = int(time.mktime(dateToEnd.timetuple()))+86399

base_url = "http://20.232.138.236/"

# Put your auth token here
bearer_token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjZXJ0aWZpY2F0ZSI6Ii0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLVxuTUlJQzFUQ0NBbnlnQXdJQkFnSVVUZWo4UzNvM0NZMlhjWmg0dnNyNlp1RlphNVF3Q2dZSUtvWkl6ajBFQXdJd1xuY0RFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1ROHdEUVlEVlFRSFxuRXdaRWRYSm9ZVzB4R1RBWEJnTlZCQW9URUc5eVp6RXVaWGhoYlhCc1pTNWpiMjB4SERBYUJnTlZCQU1URTJOaFxuTG05eVp6RXVaWGhoYlhCc1pTNWpiMjB3SGhjTk1qSXdNekk1TWpJek56QXdXaGNOTWpNd016STVNakkwTWpBd1xuV2pCZ01Rc3dDUVlEVlFRR0V3SlZVekVYTUJVR0ExVUVDQk1PVG05eWRHZ2dRMkZ5YjJ4cGJtRXhGREFTQmdOVlxuQkFvVEMwaDVjR1Z5YkdWa1oyVnlNUTR3REFZRFZRUUxFd1ZoWkcxcGJqRVNNQkFHQTFVRUF4TUpiM0puTVdGa1xuYldsdU1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRVpBKzdXbjZSNy9wckYwOFQxZEw4NlY5MFxuVXRXRksxVmRHOVEwbjJtSDIwRVRlWndhbjhvNWRROUhEdVJjUVVTOXZpYUc3SDZ5Y1pXRTFuT3ErMWlNMXFPQ1xuQVFJd2dmOHdEZ1lEVlIwUEFRSC9CQVFEQWdlQU1Bd0dBMVVkRXdFQi93UUNNQUF3SFFZRFZSME9CQllFRkJjUFxuakVnTnJkL3AyV3V5elVWTFF2UHVBblFqTUI4R0ExVWRJd1FZTUJhQUZBK09FUkZvNjVscksyNGh1TzJIM2ZuZlxuQUxxaU1CVUdBMVVkRVFRT01BeUNDazl5WnpGTllYTjBaWEl3Z1ljR0NDb0RCQVVHQndnQkJIdDdJbUYwZEhKelxuSWpwN0lsSmxZV1JQZEdobGNuTkVZWFJoSWpvaU1TSXNJbEpsY1hWbGMzUlViMnRsYmxKdmJHVWlPaUl4SWl3aVxuYUdZdVFXWm1hV3hwWVhScGIyNGlPaUlpTENKb1ppNUZibkp2Ykd4dFpXNTBTVVFpT2lKdmNtY3hZV1J0YVc0aVxuTENKb1ppNVVlWEJsSWpvaVlXUnRhVzRpZlgwd0NnWUlLb1pJemowRUF3SURSd0F3UkFJZ0syWUwyVFRKbGxzelxuOVlHTzNkaVZ5NGRxdFBXK0VyTnBnaWxhRXA0NUgyVUNJQzlSVFZRcDJvZ0RJWXpzSm42cE9qL29IMkRnT0dRZVxudTNIc1F0YXZJWGczXG4tLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tXG4iLCJleHAiOjE2NDg1OTc4MzEsIm1zcElEIjoiT3JnMU1TUCIsIm9yaWdfaWF0IjoxNjQ4NTk0MjMxLCJwcml2YXRlS2V5IjoiLS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tXG5NSUdIQWdFQU1CTUdCeXFHU000OUFnRUdDQ3FHU000OUF3RUhCRzB3YXdJQkFRUWdjY2ZIVE4xZzk3amk0cHJLXG42UkhmZ2s0WWJsdTQvWktxakRiR0hzQkh5SVNoUkFOQ0FBUmtEN3RhZnBIdittc1hUeFBWMHZ6cFgzUlMxWVVyXG5WVjBiMURTZmFZZmJRUk41bkJxZnlqbDFEMGNPNUZ4QlJMMitKb2JzZnJKeGxZVFdjNnI3V0l6V1xuLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLVxuIiwicm9sZXMiOlsibWFuYWdlIG90aGVycyBkYXRhIiwiY29tcHV0YXRpb24iLCJhZG1pbiJdLCJ1c2VyTmFtZSI6Im9yZzFhZG1pbiJ9.wYcDd6J1cpYURRAsp5DxmZFq5lTNSJKnOPAD_4qkJm8"

request_flow_json = {"nodes":[{"type":"TokenNode","id":"node_16485881807690","name":"XRayPneumoniaCases","chaincodeName":"examplealgorithm","tokenId":"","methodName":"ExampleAlgorithmSmartContract:XRayPneumoniaCases","options":[{"name":"Description","value":"Test1"},{"name":"startDateTimestamp","value":"1577916540"},{"name":"endDateTimestamp","value":"1578003000"}],"state":{},"interfaces":[{"name":"Output","id":"ni_16485881807691"}],"position":{"x":143.934847388879,"y":90.22803413892346},"width":200,"twoColumn":"false","customClasses":""},{"type":"TokenNode","id":"node_16485882211562","name":"XRayPneumoniaCases","chaincodeName":"examplealgorithm","tokenId":"","methodName":"ExampleAlgorithmSmartContract:XRayPneumoniaCases","options":[{"name":"Description","value":"Test2"},{"name":"startDateTimestamp","value":"1577916600"},{"name":"endDateTimestamp","value":"1577916600"}],"state":{},"interfaces":[{"name":"Output","id":"ni_16485882211563"}],"position":{"x":143.0196077940895,"y":439.6654315470086},"width":200,"twoColumn":"false","customClasses":""},{"type":"TokenNode","id":"node_16485882411304","name":"CreateDonutChart","chaincodeName":"examplealgorithm","tokenId":"","methodName":"ExampleAlgorithmSmartContract:CreateDonutChart","options":[{"name":"Description","value":"Chart"},{"name":"Add Input","value":"null"},{"name":"Remove Input","value":"null"},{"name":"title","value":"Chart"}],"state":{},"interfaces":[{"name":"Output","id":"ni_16485882411305"},{"name":"Input 1","id":"ni_16485882427876"},{"name":"Input 2","id":"ni_16485882483638"}],"position":{"x":630.8812576890722,"y":253.5237047790975},"width":200,"twoColumn":"false","customClasses":""}],"connections":[{"id":"164858825426711","from":"ni_16485881807691","to":"ni_16485882427876"},{"id":"164858825661916","from":"ni_16485882211563","to":"ni_16485882483638"}],"panning":{"x":88.1367849242315,"y":23.863966482887676},"scaling":0.8219543390760873}

auth_header = {'content-type': 'application/json',
               'Authorization': bearer_token}

def flow():
    response = requests.post(base_url + 'api/v1/computation/requestflow', json=request_flow_json, headers=auth_header)
    json_token_resp = response.json()
    token_id = json_token_resp['nodes'][0]['tokenId']
    # You need to manually trigger this job from UI
    #requests.post(f"{base_url}api/v1/computation/token/{token_id}/start", headers=auth_header)

if __name__ == "__main__":
    flow()