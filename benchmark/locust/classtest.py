
import json
import time
import time
import datetime
from aiohttp import request
import requests

dateToStart = datetime.date(2020,1,1)
dateToStartTimestamp = int(time.mktime(dateToStart.timetuple()))

dateToEnd = datetime.date(2020,1,2)
dateToEndTimestamp = int(time.mktime(dateToEnd.timetuple()))+86399

base_url = "http://20.232.138.236/"

# Put your auth token here
bearer_token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjZXJ0aWZpY2F0ZSI6Ii0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLVxuTUlJQzFUQ0NBbnlnQXdJQkFnSVVRNmh2eDBZbFN6UXlMSXNJUmEwU0tDaFYydHd3Q2dZSUtvWkl6ajBFQXdJd1xuY0RFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1ROHdEUVlEVlFRSFxuRXdaRWRYSm9ZVzB4R1RBWEJnTlZCQW9URUc5eVp6RXVaWGhoYlhCc1pTNWpiMjB4SERBYUJnTlZCQU1URTJOaFxuTG05eVp6RXVaWGhoYlhCc1pTNWpiMjB3SGhjTk1qSXdNekk1TVRZek56QXdXaGNOTWpNd016STVNVFkwTWpBd1xuV2pCZ01Rc3dDUVlEVlFRR0V3SlZVekVYTUJVR0ExVUVDQk1PVG05eWRHZ2dRMkZ5YjJ4cGJtRXhGREFTQmdOVlxuQkFvVEMwaDVjR1Z5YkdWa1oyVnlNUTR3REFZRFZRUUxFd1ZoWkcxcGJqRVNNQkFHQTFVRUF4TUpiM0puTVdGa1xuYldsdU1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRXU5dlA4c3MyR0JINzdma0l2WHNOalUvblxuUjBYQUg1amV5LzhFQUxCbWtvZWQ4Z2NwWjFmdEhEMnI0Wkl1TVZNbHlMckdXMG54eXpwY0MzUUtBMjRLNEtPQ1xuQVFJd2dmOHdEZ1lEVlIwUEFRSC9CQVFEQWdlQU1Bd0dBMVVkRXdFQi93UUNNQUF3SFFZRFZSME9CQllFRk43RVxuNUJuVk9may90SjhqV2tsUTZnOTBzTTRKTUI4R0ExVWRJd1FZTUJhQUZQME9PSmhXT2xBa0pnTjZLR0FBekdORFxuQ2ZxeU1CVUdBMVVkRVFRT01BeUNDazl5WnpGTllYTjBaWEl3Z1ljR0NDb0RCQVVHQndnQkJIdDdJbUYwZEhKelxuSWpwN0lsSmxZV1JQZEdobGNuTkVZWFJoSWpvaU1TSXNJbEpsY1hWbGMzUlViMnRsYmxKdmJHVWlPaUl4SWl3aVxuYUdZdVFXWm1hV3hwWVhScGIyNGlPaUlpTENKb1ppNUZibkp2Ykd4dFpXNTBTVVFpT2lKdmNtY3hZV1J0YVc0aVxuTENKb1ppNVVlWEJsSWpvaVlXUnRhVzRpZlgwd0NnWUlLb1pJemowRUF3SURSd0F3UkFJZ05ZVTNPZHJyWWRhOFxuK3pKV3czVlhaS1A1SDgwTnRhTVM0ZDUzVUdTWkFEa0NJSFg3emNtdGV1dWo3UVhZMU02TU00Vm1lU3BXeFl6UFxuTG9KSFE1UHlhRUdyXG4tLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tXG4iLCJleHAiOjE2NDg1Nzc3MzUsIm1zcElEIjoiT3JnMU1TUCIsIm9yaWdfaWF0IjoxNjQ4NTc0MTM1LCJwcml2YXRlS2V5IjoiLS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tXG5NSUdIQWdFQU1CTUdCeXFHU000OUFnRUdDQ3FHU000OUF3RUhCRzB3YXdJQkFRUWczL3ZTbzlwNklpRWJzRUpVXG5QZkU1cnhHc1NJYlNWaFVOemE3amNrMGllT3FoUkFOQ0FBUzcyOC95eXpZWUVmdnQrUWk5ZXcyTlQrZEhSY0FmXG5tTjdML3dRQXNHYVNoNTN5QnlsblYrMGNQYXZoa2k0eFV5WEl1c1piU2ZITE9sd0xkQW9EYmdyZ1xuLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLVxuIiwicm9sZXMiOlsibWFuYWdlIG90aGVycyBkYXRhIiwiY29tcHV0YXRpb24iLCJhZG1pbiJdLCJ1c2VyTmFtZSI6Im9yZzFhZG1pbiJ9.9XE2cPJutQjWSP29W-f8qE5YmCgUrqkJ7vcHGLZ2wRA"

request_flow_json = {"nodes":[{"type":"TokenNode","id":"node_16477087948944","name":"XRayPneumoniaCases","chaincodeName":"examplealgorithm","tokenId":"","methodName":"ExampleAlgorithmSmartContract:XRayPneumoniaCases","options":[{"name":"Description","value":"Test"},{"name":"startDateTimestamp","value":f"{dateToStartTimestamp}"},{"name":"endDateTimestamp","value":f"{dateToEndTimestamp}"}],"state":{},"interfaces":[{"name":"Output","id":"ni_16477087948945"}],"position":{"x":155,"y":184},"width":200,"twoColumn":"false","customClasses":""}],"connections":[],"panning":{"x":0,"y":0},"scaling":1}


auth_header = {'content-type': 'application/json',
               'Authorization': bearer_token}

def flow():
    response = requests.post(base_url + 'api/v1/computation/requestflow', json=request_flow_json, headers=auth_header)
    json_token_resp = response.json()
    token_id = json_token_resp['nodes'][0]['tokenId']
    requests.post(f"{base_url}api/v1/computation/token/{token_id}/start", headers=auth_header)

    start_time = datetime.datetime.now()
    print(f"{start_time}")
    while True:

        token_resp = requests.get(f"{base_url}api/v1/computation/token/{token_id}", headers=auth_header)
        token_resp_json = token_resp.json()

        if token_resp_json['ret']['RetValue'] != '':
            break

        time_delta = datetime.datetime.now() - start_time
        if time_delta.total_seconds() >= 30*60:
            raise TimeoutError()


    print(f"{datetime.datetime.now()}")



if __name__ == "__main__":
    flow()