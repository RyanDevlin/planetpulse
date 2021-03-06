[![Go Report Card](https://goreportcard.com/badge/github.com/RyanDevlin/planetpulse)](https://goreportcard.com/report/github.com/RyanDevlin/planetpulse)
![example branch parameter](https://github.com/RyanDevlin/planetpulse/actions/workflows/release-apiserver.yml/badge.svg?branch=release-0.1.0)



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/RyanDevlin/planetpulse">
    <img src="https://github.com/RyanDevlin/planetpulse/blob/main/api/apiserver/images/planetpulse.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Planet Pulse API</h3>

  <p align="center">
    A REST API that serves NOAA climate data
    <br />
    <a href="https://github.com/RyanDevlin/planetpulse/blob/main/docs/README.md"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://planetpulse.io">View Demo</a>
    ·
    <a href="https://github.com/RyanDevlin/planetpulse/issues">Report Bug</a>
    ·
    <a href="https://github.com/RyanDevlin/planetpulse/issues">Request Feature</a>
  </p>
</div>

# What is this? 🤔
Planet Pulse is an API service that was designed to make climate data widely available for free. Planet Pulse serves climate data obtained from the [National Oceanic and Atmospheric Administration (NOAA)](https://www.noaa.gov/) via a REST API reachable at [api.planetpulse.io](https://api.planetpulse.io).

# Why NOAA Data? 🌡
Some of NOAA's most up-to-date climate data is still served as text/csv files over an [AFTP server](https://gml.noaa.gov/aftp/). This causes headaches for developers because a significant amount of parsing and cleaning code is needed to download and use it. To avoid this, Planet Pulse does the heavy lifting and serves the data over a simple REST API. This allows developers to request only the data they need and provides a predictable and fast endpoint to do so.

# How can I use this? 👨‍💻
Head over to [planetpulse.io](https://planetpulse.io) to see the API wrapped in a nice frontend! If you wish to use the API diretly, simply hit the [api.planetpulse.io](https://api.planetpulse.io) endpoint. If you wish to run this service yourself, read the build instructions below.

# Features (v1.0.0) 🌈
- Endpoints to request Carbon Dioxide (Co2) and Methane (Ch4) atmospheric data dating back to 1974!
- Compression (gzip) by default for all responses
- Query parameters for each endpoint that can be combined to filter results serverside
- Build pipeline to build/test/validate source code and upload artifacts as a container image
- Deployment automation to set everything up in AWS with a few commands!

# Limitations (v1.0.0) 🚧
This functionality is a work in progress and subject to change:
- API Keys are not yet in place, but are under development
- Ratelimiting requires API Keys as a prerequisite and therefore is not yet implemented
- Realistically this service should be deployed on a Kubernetes cluster. Because of financial limitations, it's currently deployed as two containers behind an NGINX server running on a single machine in AWS.

# Building 👷‍♂️
These build instructions assume you desire to run this service in AWS. If you wish to run it locally, skip the AWS deployment steps.

1.) Provsion a new AWS EC2 machine running RHEL8 (most of these steps will work on other linux distros, except things relating to SELinux).
2.) Install git and clone this repo
```
[ec2-user@ip-0-0-0-0 ~]$ sudo yum install git
[ec2-user@ip-0-0-0-0 ~]$ git clone https://github.com/RyanDevlin/planetpulse.git
```
3.) Move into the `planetpulse` directory and run the provisioning script. This will setup the AWS machine with everything needed to run the service.
```
[ec2-user@ip-0-0-0-0 ~]$ cd planetpulse/
[ec2-user@ip-0-0-0-0 ~]$ ./api/apiserver/preflight
```
4.) Run the startup script to download and run the API server container
```
[ec2-user@ip-0-0-0-0 ~]$ ./start.sh
```
