# Design
## Layers
- common - helper functions e.g. logger
- data - main business logic e.g. extract data from db
- env - env variable and initilization e.g. start db connection
- handlers - handlers for api endpoints
- models - data models e.g. items 
- routes - api endpoint routes
- vendor - imported packages
 
## Tools
- docker
- go dep - dependancy management


## Start & Stop 
1. build: `make build` (There might be a little warning message about file out of sync, can be ignored)
2. run: `make run`
3. stop: `make stop`

## Env
Listen Port 9060

## Endpoints Specs
- health/readiness check `/_ping`
- search endpoint `/search?searchTerm&lat&lng` Query params all needed
  - response: 
  ```[
	{
		"item": {
			"item_name": "Nord Stage 2 HA76 piano",
			"lat": 51.4217796,
			"lng": -0.294957995,
			"item_url": "london/hire-nord-stage-ex-73-piano-11680228",
			"img_urls": "[\"nord-stage-2-ha76-piano-74980366.png\"]"
		}
	},
	{
		"item": {
			"item_name": "Yamaha electric upright piano ",
			"lat": 51.3388901,
			"lng": -0.434143096,
			"item_url": "london/hire-yamaha-electric-upright-piano--68948282",
			"img_urls": "[\"yamaha-electric-upright-piano--83066402.jpeg\"]"
		}
	}
]

  - response: 
	```{
	"code": 400,
	"description": "Bad Request",
	"error": "missing_lng"}```
