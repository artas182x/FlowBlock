# Stress testing reading medical data and running classification. Recommended commands:
# ulimit -S -n 400000
# locust -H http://xxx.xxx.xxx.xxx/ -u 15 -r 0.1 -t 5m --autostart --csv csv1name MedicalData
# locust -H http://xxx.xxx.xxx.xxx/ -u 6 -r 0.1 -t 5m --autostart --csv csv2name Flow

from datetime import datetime
import json
import re
import time
from locust import FastHttpUser, HttpUser, run_single_user, task
from pexpect import TIMEOUT

# Put your auth token here
bearer_token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjZXJ0aWZpY2F0ZSI6Ii0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLVxuTUlJQzFUQ0NBbnlnQXdJQkFnSVVhMFRKRXlXTkQzclpWRXIwUFpkWDFaMXpZRll3Q2dZSUtvWkl6ajBFQXdJd1xuY0RFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1ROHdEUVlEVlFRSFxuRXdaRWRYSm9ZVzB4R1RBWEJnTlZCQW9URUc5eVp6RXVaWGhoYlhCc1pTNWpiMjB4SERBYUJnTlZCQU1URTJOaFxuTG05eVp6RXVaWGhoYlhCc1pTNWpiMjB3SGhjTk1qSXdNekk1TWpNeE16QXdXaGNOTWpNd016STVNak14T0RBd1xuV2pCZ01Rc3dDUVlEVlFRR0V3SlZVekVYTUJVR0ExVUVDQk1PVG05eWRHZ2dRMkZ5YjJ4cGJtRXhGREFTQmdOVlxuQkFvVEMwaDVjR1Z5YkdWa1oyVnlNUTR3REFZRFZRUUxFd1ZoWkcxcGJqRVNNQkFHQTFVRUF4TUpiM0puTVdGa1xuYldsdU1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRXRFYjdtR3h4VnVUWVF0NkpGY04xd0NXQlxuaGRVZ3p3SnJHTzlkTFRlaDF6WjZTK3hvbXZxS1FsbklTMjlFeENkZTRYNXNLQmJhZWxxa3E2Q1ArQzhTSzZPQ1xuQVFJd2dmOHdEZ1lEVlIwUEFRSC9CQVFEQWdlQU1Bd0dBMVVkRXdFQi93UUNNQUF3SFFZRFZSME9CQllFRkpBd1xuVytScmV5OGJjSHlqRVBvQkJXNFQwSW0zTUI4R0ExVWRJd1FZTUJhQUZONjFEYitFbFNSYW9VKytUNDVxejk1N1xuS1pRMk1CVUdBMVVkRVFRT01BeUNDazl5WnpGTllYTjBaWEl3Z1ljR0NDb0RCQVVHQndnQkJIdDdJbUYwZEhKelxuSWpwN0lsSmxZV1JQZEdobGNuTkVZWFJoSWpvaU1TSXNJbEpsY1hWbGMzUlViMnRsYmxKdmJHVWlPaUl4SWl3aVxuYUdZdVFXWm1hV3hwWVhScGIyNGlPaUlpTENKb1ppNUZibkp2Ykd4dFpXNTBTVVFpT2lKdmNtY3hZV1J0YVc0aVxuTENKb1ppNVVlWEJsSWpvaVlXUnRhVzRpZlgwd0NnWUlLb1pJemowRUF3SURSd0F3UkFJZ2MrVTRvM3JLTGVzZ1xuMHhQUjZJUlFYVkRXRXFGdi9oZ0U0S0xxRzJDSmxPRUNJRGlmN3IxdEZlOXl3SDQydXI5RFZEK0pHUVZ1WE5wMFxuYzRCbG5jSkpIRzdNXG4tLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tXG4iLCJleHAiOjE2NDg2MDA0NDQsIm1zcElEIjoiT3JnMU1TUCIsIm9yaWdfaWF0IjoxNjQ4NTk2ODQ0LCJwcml2YXRlS2V5IjoiLS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tXG5NSUdIQWdFQU1CTUdCeXFHU000OUFnRUdDQ3FHU000OUF3RUhCRzB3YXdJQkFRUWdnTlFKRHFGd1lwQXJWOXlOXG5YbjJNN3laaFlCai9WTGxneHJaT2RHVXp2dStoUkFOQ0FBUzBSdnVZYkhGVzVOaEMzb2tWdzNYQUpZR0YxU0RQXG5BbXNZNzEwdE42SFhObnBMN0dpYStvcENXY2hMYjBURUoxN2hmbXdvRnRwNldxU3JvSS80THhJclxuLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLVxuIiwicm9sZXMiOlsibWFuYWdlIG90aGVycyBkYXRhIiwiY29tcHV0YXRpb24iLCJhZG1pbiJdLCJ1c2VyTmFtZSI6Im9yZzFhZG1pbiJ9.d9CDTPnqvZ8o0Q8vPq_y7RKjKXI3ORgiiLEz9Bs3R-c"

request_flow_json = {"nodes":[{"type":"TokenNode","id":"node_16477087948944","name":"XRayPneumoniaCases","chaincodeName":"examplealgorithm","tokenId":"","methodName":"ExampleAlgorithmSmartContract:XRayPneumoniaCases","options":[{"name":"Description","value":"Test"},{"name":"startDateTimestamp","value":"1577897580"},{"name":"endDateTimestamp","value":"1647708780"}],"state":{},"interfaces":[{"name":"Output","id":"ni_16477087948945"}],"position":{"x":155,"y":184},"width":200,"twoColumn":"false","customClasses":""}],"connections":[],"panning":{"x":0,"y":0},"scaling":1}

auth_header = {'content-type': 'application/json',
               'Authorization': bearer_token}


class MedicalData(FastHttpUser):
    @task
    def getData(self):
        getDataJson = {
            "medicalEntryName": "",
            "dateStartTimestamp": "0",
            "dateEndTimestamp": "1647711421"
            }
        self.client.post('api/v1/medicaldata/request', json=getDataJson, headers=auth_header)

class Flow(FastHttpUser):
    @task
    def flow(self):
        response = self.client.post('api/v1/computation/requestflow', json=request_flow_json, headers=auth_header)
        json_token_resp = response.json()
        token_id = json_token_resp['nodes'][0]['tokenId']
        self.client.post(f"api/v1/computation/token/{token_id}/start", headers=auth_header, name="api/v1/computation/token/id/start")

        start_time = datetime.now()
        while True:

            with self.client.get(f"api/v1/computation/token/{token_id}", headers=auth_header, name="api/v1/computation/token/id", catch_response=True) as token_resp:
                if token_resp.status_code == 400:
                    token_resp.success()
                    continue
                else:
                    token_resp_json = token_resp.json()

                    token_resp.success()

                    if 'ret' not in token_resp_json:
                        continue

                    if token_resp_json['ret']['RetValue'] != '':
                        break
                    
            time_delta = datetime.now() - start_time
            #print(time_delta)
            if time_delta.total_seconds() >= 30*60:
                raise TimeoutError()

            time.sleep(2.0)

        token_resp = self.client.get(f"api/v1/computation/token/{token_id}", headers=auth_header, name="Computation finished")



if __name__ == "__main__":
    run_single_user(Flow)