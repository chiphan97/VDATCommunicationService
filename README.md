# Chat
A service to help people communication in VDAT ecosystem

## Plans
### Version 0.1
Basic functionality, including:
1. Basic UI interface
2. View messages and send message
3. Searching people

## Schedule
![](docs/schedule.png)

###### to update schedule modify `./docs/schedule.puml` and save the result in `./docs/schedule.png`

## Development

### Monitoring tools
use following command to run Prometheus, Grafana to monitor performance
```shell script
docker-compose -f dev/docker-compose.yml up
```
* access Grafana's Dashboards at http://localhost:3000/dashboards

### Project layout guideline

`/cmd` contains binary package
`/pkg` contains shared modules

see https://github.com/golang-standards/project-layout
