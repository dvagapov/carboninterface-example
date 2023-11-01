# carboninterface-example
**Table of Contents**

- [Introduction](#introduction)
- [How use](#how-use)
- [Cleanup](#cleanup)
---

## Introduction
This golang cli, that allows us to estimate the carbon footprint using [carboninterface Estimates API](https://docs.carboninterface.com/#/?id=estimates-api)
The Estimates API makes it easy to generate accurate emissions estimates from flights, vehicles, shipping, electricity consumption and fuel

Please use the free [Carbon Interface API](https://www.carboninterface.com/dashboard) for retrieving the carbon usage of flights - you can get an API key from, after signing up for a free account.

Note that there is a limit of 200 requests a month, so be careful with how many requests you make whilst testing your solution.

## How use
    # build binary using make
    make

    # run cli tool
    ./carboninterface-example -apikey "YOUR_API_KEY" -body 'YOUR_REQUEST_BODY'

## Example
    # Flight Estimate Request
    ./carboninterface-example -apikey "YOUR_API_KEY" -body '{"type": "flight","passengers":2,"legs":[{"departure_airport":"sfo","destination_airport":"yyz"},{"departure_airport":"yyz","destination_airport":"sfo"}]}'
    
    # Flight Estimate Response JSON:
    {"data":{"id":"501dc17e-8578-453b-bac6-9f777ec5b93c","type":"estimate","attributes":{"passengers":2,"legs":[{"departure_airport":"SFO","destination_airport":"YYZ"},{"departure_airport":"YYZ","destination_airport":"SFO"}],"distance_value":7454.15,"distance_unit":"km","estimated_at":"2023-10-31T22:44:12.693Z","carbon_g":859824,"carbon_lb":1895.59,"carbon_kg":859.82,"carbon_mt":0.86}}}

## Cleanup
    # cleanup binary using make
    make clean

## Possible improvements
1. Store responses on persistence storage -> this can allow us to avoid using requests with the same body and reach monthly API usage that fast. (!) Need to double-check with the support team about the data refreshing
2. Add test of input API key befu using it
3. Add go unit-tests using lib "net/http/httptest" [example on Medium](https://bismobaruno.medium.com/unit-test-http-request-in-golang-a96d146406e6)