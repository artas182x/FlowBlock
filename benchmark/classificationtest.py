
# Start and completion date can be read from backend logs
# Testing classification speed. It's recommended to load data in the way that 1 day is 50 pictures
import time
import time
import datetime
from aiohttp import request
import requests

dateToStart = datetime.date(2020,1,1)
dateToStartTimestamp = int(time.mktime(dateToStart.timetuple()))

dateToEnd = datetime.date(2020,1,9)
dateToEndTimestamp = int(time.mktime(dateToEnd.timetuple()))+86399

base_url = "http://20.232.138.236/"

# Put your auth token here
bearer_token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjZXJ0aWZpY2F0ZSI6Ii0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLVxuTUlJQzFqQ0NBbnlnQXdJQkFnSVVFMnp1M01PSUJSaVp1TFRsb1BFM3h4YTBBSlV3Q2dZSUtvWkl6ajBFQXdJd1xuY0RFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1ROHdEUVlEVlFRSFxuRXdaRWRYSm9ZVzB4R1RBWEJnTlZCQW9URUc5eVp6RXVaWGhoYlhCc1pTNWpiMjB4SERBYUJnTlZCQU1URTJOaFxuTG05eVp6RXVaWGhoYlhCc1pTNWpiMjB3SGhjTk1qSXdNekk1TVRjME5UQXdXaGNOTWpNd016STVNVGMxTURBd1xuV2pCZ01Rc3dDUVlEVlFRR0V3SlZVekVYTUJVR0ExVUVDQk1PVG05eWRHZ2dRMkZ5YjJ4cGJtRXhGREFTQmdOVlxuQkFvVEMwaDVjR1Z5YkdWa1oyVnlNUTR3REFZRFZRUUxFd1ZoWkcxcGJqRVNNQkFHQTFVRUF4TUpiM0puTVdGa1xuYldsdU1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRUhuQklBYWVaKzBXaXRjQ0VkYjV3dGFjWVxuNGZjankwZWtIbGNQUzdBQnUyaW9GcnJGSXR2WStzbHQ5OWQvQ2FBZDFBMW14WjFEWnYxS3llMGRMSmZwN0tPQ1xuQVFJd2dmOHdEZ1lEVlIwUEFRSC9CQVFEQWdlQU1Bd0dBMVVkRXdFQi93UUNNQUF3SFFZRFZSME9CQllFRkErMlxubGJaU1gzMEFTRXVQaExLcFdqb3UvbTBVTUI4R0ExVWRJd1FZTUJhQUZBckxwYmtWN1QxZEIrL2oyOVk0OHFBOVxuNVdFek1CVUdBMVVkRVFRT01BeUNDazl5WnpGTllYTjBaWEl3Z1ljR0NDb0RCQVVHQndnQkJIdDdJbUYwZEhKelxuSWpwN0lsSmxZV1JQZEdobGNuTkVZWFJoSWpvaU1TSXNJbEpsY1hWbGMzUlViMnRsYmxKdmJHVWlPaUl4SWl3aVxuYUdZdVFXWm1hV3hwWVhScGIyNGlPaUlpTENKb1ppNUZibkp2Ykd4dFpXNTBTVVFpT2lKdmNtY3hZV1J0YVc0aVxuTENKb1ppNVVlWEJsSWpvaVlXUnRhVzRpZlgwd0NnWUlLb1pJemowRUF3SURTQUF3UlFJaEFQVUFHcnI4S0hyVFxub0RaYVMzQ0ZGOW5ySlJ4S2JQUWxpb1daVElqbjFYeUlBaUJFa3NWcjZZTTExNXRoTEZMaHgzWW4vVVVyYm1Ialxua1kvTVJ0ZXBndHlVOWc9PVxuLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLVxuIiwiZXhwIjoxNjQ4NTg0MzU5LCJtc3BJRCI6Ik9yZzFNU1AiLCJvcmlnX2lhdCI6MTY0ODU4MDc1OSwicHJpdmF0ZUtleSI6Ii0tLS0tQkVHSU4gUFJJVkFURSBLRVktLS0tLVxuTUlHSEFnRUFNQk1HQnlxR1NNNDlBZ0VHQ0NxR1NNNDlBd0VIQkcwd2F3SUJBUVFnTFdGWVNHekJwWUM2Nlcxd1xuMEc1Rkxxd1BRRlJPSE9ySFdDaUNRKzRTSUJLaFJBTkNBQVFlY0VnQnA1bjdSYUsxd0lSMXZuQzFweGpoOXlQTFxuUjZRZVZ3OUxzQUc3YUtnV3VzVWkyOWo2eVczMzEzOEpvQjNVRFdiRm5VTm0vVXJKN1Iwc2wrbnNcbi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS1cbiIsInJvbGVzIjpbIm1hbmFnZSBvdGhlcnMgZGF0YSIsImNvbXB1dGF0aW9uIiwiYWRtaW4iXSwidXNlck5hbWUiOiJvcmcxYWRtaW4ifQ._nwgcZbTqeOhu3GBDDjaNjJr_RwOdj5kxrZNmnONy-Y"

request_flow_json = {"nodes":[{"type":"TokenNode","id":"node_16477087948944","name":"XRayPneumoniaCases","chaincodeName":"examplealgorithm","tokenId":"","methodName":"ExampleAlgorithmSmartContract:XRayPneumoniaCases","options":[{"name":"Description","value":"Test"},{"name":"startDateTimestamp","value":f"{dateToStartTimestamp}"},{"name":"endDateTimestamp","value":f"{dateToEndTimestamp}"}],"state":{},"interfaces":[{"name":"Output","id":"ni_16477087948945"}],"position":{"x":155,"y":184},"width":200,"twoColumn":"false","customClasses":""}],"connections":[],"panning":{"x":0,"y":0},"scaling":1}


auth_header = {'content-type': 'application/json',
               'Authorization': bearer_token}

def flow():
    response = requests.post(base_url + 'api/v1/computation/requestflow', json=request_flow_json, headers=auth_header)
    json_token_resp = response.json()
    token_id = json_token_resp['nodes'][0]['tokenId']
    requests.post(f"{base_url}api/v1/computation/token/{token_id}/start", headers=auth_header)

if __name__ == "__main__":
    flow()