# Uptime Monitoring Service
- A system which checks if the requested URL is up or not.
- Data corresponding to the request, is stored and managed in the database.
- Installation and usage has been explained in the later sections.

## Tech Stack Used:
- Golang
- Mysql
  - Gorm as orm library
- Docker

## Installation

There are 3 ways to install this service:


#### 1. Using Docker Compose (Suggested Way)

This method of installation, doesn't require for you to setup database on your local machine.

Prerequisites | Following should be installed on the local machine:
```
Docker desktop
docker-compose
```
Run following commands in the terminal:
```
git clone https://github.com/sammarth-kapse/UptimeMonitoringService.git
cd UptimeMonitoringService
docker-compose up --build
```
________________

##### For the other 2 methods of installation, you would have to setup database on local machine
```
mysql database named as 'uptime_monitoring_system'
Username = sammarth
Password = mysql
```
OR
```
You can clone the repository.
Edit the environment variables present in .env file.
Update the username, database-name and password according to your use.
```
---------------

#### 2. On local machine 
Prerequisites | Following should be installed on the local machine:
```
Golang
```
Run following commands in the terminal:
```
git clone https://github.com/sammarth-kapse/UptimeMonitoringService.git
cd UptimeMonitoringService
go mod download
```

Build
```
go build .
```
Run
```
DOCKER=false ./UptimeMonitoringSevice
```
--------------
#### 3. Using Docker Image

Run following commands in the terminal:
```
docker image pull sammarth/uptime-monitoring-service
docker run -p 8080:8080 sammarth/uptime-monitoring-service
```


## API Endpoints:
#### Base URL
```
http://localhost:8080
```

#### Add a URL to monitor:
Use `POST /urls/` to add a URL to the service.
- crawl_timeout is the time for which system wait before giving up on URL.
- URL is pinged after every given set of time that is equal to the frequency.
- Failure Threshold is the failure limit, once failure count reaches the threshold, url is marked as inactive and is no longer monitored.

Request: 
```
{
   "url":                        ”abc.com”,
   "crawl_timeout":              20,
   “frequency”:                  30, 
   “failure_threshold” :         50  
}
```

Response:
```
{
  "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”abc.com”,
  "crawl_timeout":              20,
  “frequency”:                  30,
  “failure_threshold” :         50,
  “status”:                     “active”,
  “failure_count”:               0
}

```
---------------
#### GET URL Information
Use `GET /urls/:id` to get URL Info for the corresponding ID

Response:
```
{
  "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”abc.com”,
  "crawl_timeout":              20,
  “frequency”:                  30,
  “failure_threshold” :         50,
  “status”:                     “active”,
  “failure_count”:              5
}
```
---------------   
#### Update URL Parameters
Use `PATCH /urls/:id` to update parameters.

Parameters that can be updated: Frequency, Crawl-Timeout and Failure-Threshold.

Request:
```
{
  "frequency": 10,
  "failure_threshold": 5,
  "crawl_timeout": 7
}
```

Response:
```
{
  "id":"                        b7f32a21-b863-4dd1-bd86-e99e8961ffc6",
  "url":                        ”abc.com”,
  "crawl_timeout":              7,
  “frequency”:                  10,
  “failure_threshold” :         5,
  “status”:                     “active”,
  “failure_count”:              0
}
```
---------------
#### Activate URL status
Use `POST /urls/:id/activate` to activate the monitoring of the URL.

---------------
#### Deactivate URL status
Use `POST /urls/:id/deactivate` to deactivate the monitoring of the URL.

---------------
#### Delete URL from System
Use `DELETE /urls/:id` to remove the URL from system.


## Testing
To run test functions:

#### 1. With Docker Compose (Suggested Way)

Run following commands in the terminal:
```
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
```



#### 2. On local machine

Run following commands in the terminal:
```
cd monitor
go test
```

## HTTP status code returned
 Status Code | Explanation |
| ------- | --- |
| 200 | POST ,GET, PATCH Request successful |
| 204 | DELETE request successful |
| 400 | Request is not in proper format |
| 404 | ID DNE|