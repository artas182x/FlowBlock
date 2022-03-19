from datetime import datetime
import json
import re
import time
from locust import FastHttpUser, HttpUser, run_single_user, task
from pexpect import TIMEOUT

# Put your auth token here
bearer_token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjZXJ0aWZpY2F0ZSI6Ii0tLS0tQkVHSU4gQ0VSVElGSUNBVEUtLS0tLVxuTUlJQzFqQ0NBbnlnQXdJQkFnSVVNZzJ0WDlnTHBhSFhrdkpFU0xUTDZvQnFTdFF3Q2dZSUtvWkl6ajBFQXdJd1xuY0RFTE1Ba0dBMVVFQmhNQ1ZWTXhGekFWQmdOVkJBZ1REazV2Y25Sb0lFTmhjbTlzYVc1aE1ROHdEUVlEVlFRSFxuRXdaRWRYSm9ZVzB4R1RBWEJnTlZCQW9URUc5eVp6RXVaWGhoYlhCc1pTNWpiMjB4SERBYUJnTlZCQU1URTJOaFxuTG05eVp6RXVaWGhoYlhCc1pTNWpiMjB3SGhjTk1qSXdNekU1TWpFd01EQXdXaGNOTWpNd016RTVNakV3TlRBd1xuV2pCZ01Rc3dDUVlEVlFRR0V3SlZVekVYTUJVR0ExVUVDQk1PVG05eWRHZ2dRMkZ5YjJ4cGJtRXhGREFTQmdOVlxuQkFvVEMwaDVjR1Z5YkdWa1oyVnlNUTR3REFZRFZRUUxFd1ZoWkcxcGJqRVNNQkFHQTFVRUF4TUpiM0puTVdGa1xuYldsdU1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRWlmcWlGcVpyTkhMSWVBSGVuQlBLa1ZpTlxuVlI1Q3BBbWV3Slljd002QndtTEhXLzJoL09UQ1ZxaktsdGwxYTNmWVpYenVxcUREa3BMUUxVYWV5ZnlERUtPQ1xuQVFJd2dmOHdEZ1lEVlIwUEFRSC9CQVFEQWdlQU1Bd0dBMVVkRXdFQi93UUNNQUF3SFFZRFZSME9CQllFRkVoOVxuTGp3SW0xSHlwL05UamZTc2pGZVpIU2pvTUI4R0ExVWRJd1FZTUJhQUZENUF1OTlZMUNTZmtNcS81Y21QQldadlxuRjhLaE1CVUdBMVVkRVFRT01BeUNDbk4zWVhKdGJtOWtaVEV3Z1ljR0NDb0RCQVVHQndnQkJIdDdJbUYwZEhKelxuSWpwN0lsSmxZV1JQZEdobGNuTkVZWFJoSWpvaU1TSXNJbEpsY1hWbGMzUlViMnRsYmxKdmJHVWlPaUl4SWl3aVxuYUdZdVFXWm1hV3hwWVhScGIyNGlPaUlpTENKb1ppNUZibkp2Ykd4dFpXNTBTVVFpT2lKdmNtY3hZV1J0YVc0aVxuTENKb1ppNVVlWEJsSWpvaVlXUnRhVzRpZlgwd0NnWUlLb1pJemowRUF3SURTQUF3UlFJaEFPclJlbHFnUExHdlxuUlM4QWdQQVlIdS9QSmQ3dVFFMStJZFRaN2t3ZXp3WGVBaUJUa0d0eW8yazlEQm0vaUduejZ0ZUZzVEVOc21FVFxubEtDUXBhVE9OWWh6dmc9PVxuLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLVxuIiwiZXhwIjoxNjQ3NzI3OTY5LCJtc3BJRCI6Ik9yZzFNU1AiLCJvcmlnX2lhdCI6MTY0NzcyNDM2OSwicHJpdmF0ZUtleSI6Ii0tLS0tQkVHSU4gUFJJVkFURSBLRVktLS0tLVxuTUlHSEFnRUFNQk1HQnlxR1NNNDlBZ0VHQ0NxR1NNNDlBd0VIQkcwd2F3SUJBUVFncitCZmE2cFE5YWJVd284NFxudXFQOUVTZ01jdW0zRGl2Z2kxNmlGdDJ4MDZ5aFJBTkNBQVNKK3FJV3BtczBjc2g0QWQ2Y0U4cVJXSTFWSGtLa1xuQ1o3QWxoekF6b0hDWXNkYi9hSDg1TUpXcU1xVzJYVnJkOWhsZk82cW9NT1NrdEF0UnA3Si9JTVFcbi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS1cbiIsInJvbGVzIjpbIm1hbmFnZSBvdGhlcnMgZGF0YSIsImNvbXB1dGF0aW9uIiwiYWRtaW4iXSwidXNlck5hbWUiOiJvcmcxYWRtaW4ifQ.tEz28He4CbC3EEAt3gSjThkdS3bQCHySa-01539gx_4"

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

            token_resp = self.client.get(f"api/v1/computation/token/{token_id}", headers=auth_header, name="api/v1/computation/token/id")
            token_resp_json = token_resp.json()

            if token_resp_json['ret']['RetValue'] != '':
                break

            time_delta = datetime.now() - start_time
            print(time_delta)
            if time_delta.total_seconds() >= 30*60:
                raise TimeoutError()

            time.sleep(2.0)

        token_resp = self.client.get(f"api/v1/computation/token/{token_id}", headers=auth_header, name="Computation finished")



if __name__ == "__main__":
    run_single_user(Flow)