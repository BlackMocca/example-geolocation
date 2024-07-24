# Example-GEOLOCATION

### This repository is example to save lat lon from mobile 

### Save Tracking
```shell
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"plate_number": "กข-1234", "lat": 13.7465389, "lon": 100.5342779}' \
  http://127.0.0.1:3000/trackings
```