```graphql
mutation Seed(
  $airlines: [AirlinesInsertInput!]!
  $manufacturers: [ManufacturersInsertInput!]!
  $aircrafts: [AircraftsInsertInput!]!
  $airports: [AirportsInsertInput!]!
  $employees: [EmployeesInsertInput!]!
  $flights: [FlightsInsertInput!]!
  $staff: [FlightStaffInsertInput!]!
) {
  airlines: insertAirlines(
    objects: $airlines
    onConflict: { constraint: airlinesPkey, updateColumns: [id] }
  ) {
    __typename
  }

  manufacturers: insertManufacturers(
    objects: $manufacturers
    onConflict: { constraint: manufacturersPkey, updateColumns: [id] }
  ) {
    __typename
  }

  aircrafts: insertAircrafts(
    objects: $aircrafts
    onConflict: { constraint: aircraftsPkey, updateColumns: [id] }
  ) {
    __typename
  }

  airports: insertAirports(
    objects: $airports
    onConflict: { constraint: airportsPkey, updateColumns: [id] }
  ) {
    __typename
  }

  employees: insertEmployees(
    objects: $employees
    onConflict: { constraint: employeesPkey, updateColumns: [id] }
  ) {
    __typename
  }

  flights: insertFlights(
    objects: $flights
    onConflict: { constraint: flightsPkey, updateColumns: [id] }
  ) {
    __typename
  }

  staff: insertFlightStaff(
    objects: $staff
    onConflict: { constraint: flightStaffPkey, updateColumns: [employeeId] }
  ) {
    __typename
  }
}
```

### JSON Variables

```json
{
  "airlines": [
    {
      "id": "1d86c3a5-eaab-48e0-aae3-109a6a04daf5",
      "name": "Southwest Airlines"
    },
    {
      "id": "4f1f7ee5-b3c1-4b91-bc2b-c754275daa96",
      "name": "Spirit Airlines"
    }
  ],
  "manufacturers": [
    {
      "id": "37bf29cd-3cbd-4431-9403-7e9f2e580308",
      "name": "Airbus"
    },
    {
      "id": "6aeda167-0dc2-44c2-97c1-c70ec614fe82",
      "name": "Boeing"
    }
  ],
  "aircrafts": [
    {
      "id": "91c55517-6f49-4aad-a009-f4649d7b8bcf",
      "modelNumber": "A21N-32Q",
      "serialNumber": "2128859141",
      "airlineId": "1d86c3a5-eaab-48e0-aae3-109a6a04daf5",
      "manufacturerId": "37bf29cd-3cbd-4431-9403-7e9f2e580308"
    },
    {
      "id": "703c0014-896a-4ef4-a53b-a4ffa859a8c6",
      "modelNumber": "B37M-7M7",
      "serialNumber": "2252734816",
      "airlineId": "4f1f7ee5-b3c1-4b91-bc2b-c754275daa96",
      "manufacturerId": "6aeda167-0dc2-44c2-97c1-c70ec614fe82"
    }
  ],
  "airports": [
    {
      "id": "1a6ece44-a5ac-402f-96ca-6f8487017c43",
      "name": "Los Angeles International Airport",
      "code": "LAX"
    },
    {
      "id": "b8dffeb7-f2a7-45fb-b2b1-7db283277baa",
      "name": "Hartsfield-Jackson Atlanta International Airport",
      "code": "ATL"
    },
    {
      "id": "f9c3aa9a-637c-4d1c-bab5-f3e216229b65",
      "name": "Denver International Airport",
      "code": "DEN"
    }
  ],
  "employees": [
    {
      "id": "6db66e81-aaf4-4c0d-beaa-020ec8da720d",
      "fullName": "Alice"
    },
    {
      "id": "6e9b95df-502e-4d93-8d3b-111c6f66b8f5",
      "fullName": "Bob"
    },
    {
      "id": "df919e6b-ef07-4f1c-8cbe-614ed4d3566e",
      "fullName": "Casey"
    },
    {
      "id": "737f52d7-0167-4a73-be95-0bee2c44b1c3",
      "fullName": "Dylan"
    },
    {
      "id": "74bd8afa-8917-4113-a5df-0aa4a41c419e",
      "fullName": "Emerson"
    },
    {
      "id": "1c5bcc1a-94fb-4521-b371-51e74c254a40",
      "fullName": "Frankie"
    },
    {
      "id": "923a30f6-af6a-4940-940e-e9d57837c8e1",
      "fullName": "Hayden"
    },
    {
      "id": "325a2d8c-0bd3-4590-a370-80e7b98bf92d",
      "fullName": "Indigo"
    }
  ],
  "flights": [
    {
      "id": "3a590547-865e-4c8e-844c-357ad3006492",
      "departingAirportId": "1a6ece44-a5ac-402f-96ca-6f8487017c43",
      "arrivingAirportId": "b8dffeb7-f2a7-45fb-b2b1-7db283277baa",
      "airlineId": "1d86c3a5-eaab-48e0-aae3-109a6a04daf5",
      "aircraftId": "91c55517-6f49-4aad-a009-f4649d7b8bcf",
      "expectedDurationMins": 293,
      "expectedDepartureTime": "2023-09-05T16:32:00Z",
      "expectedArrivalTime": "2023-09-05T21:25:00Z"
    },
    {
      "id": "49f4cc96-5c07-462b-ac5b-b32ba89efe3d",
      "departingAirportId": "b8dffeb7-f2a7-45fb-b2b1-7db283277baa",
      "arrivingAirportId": "1a6ece44-a5ac-402f-96ca-6f8487017c43",
      "airlineId": "4f1f7ee5-b3c1-4b91-bc2b-c754275daa96",
      "aircraftId": "703c0014-896a-4ef4-a53b-a4ffa859a8c6",
      "expectedDurationMins": 293,
      "expectedDepartureTime": "2023-09-05T16:32:00Z",
      "expectedArrivalTime": "2023-09-05T21:25:00Z"
    }
  ],
  "staff": [
    {
      "employeeId": "6db66e81-aaf4-4c0d-beaa-020ec8da720d",
      "flightId": "3a590547-865e-4c8e-844c-357ad3006492",
      "role": "FLIGHT_ATTENDANT"
    },
    {
      "employeeId": "6e9b95df-502e-4d93-8d3b-111c6f66b8f5",
      "flightId": "3a590547-865e-4c8e-844c-357ad3006492",
      "role": "FLIGHT_ATTENDANT"
    },
    {
      "employeeId": "df919e6b-ef07-4f1c-8cbe-614ed4d3566e",
      "flightId": "3a590547-865e-4c8e-844c-357ad3006492",
      "role": "PILOT"
    },
    {
      "employeeId": "737f52d7-0167-4a73-be95-0bee2c44b1c3",
      "flightId": "3a590547-865e-4c8e-844c-357ad3006492",
      "role": "CO_PILOT"
    },
    {
      "employeeId": "74bd8afa-8917-4113-a5df-0aa4a41c419e",
      "flightId": "49f4cc96-5c07-462b-ac5b-b32ba89efe3d",
      "role": "FLIGHT_ATTENDANT"
    },
    {
      "employeeId": "1c5bcc1a-94fb-4521-b371-51e74c254a40",
      "flightId": "49f4cc96-5c07-462b-ac5b-b32ba89efe3d",
      "role": "FLIGHT_ATTENDANT"
    },
    {
      "employeeId": "923a30f6-af6a-4940-940e-e9d57837c8e1",
      "flightId": "49f4cc96-5c07-462b-ac5b-b32ba89efe3d",
      "role": "PILOT"
    },
    {
      "employeeId": "325a2d8c-0bd3-4590-a370-80e7b98bf92d",
      "flightId": "49f4cc96-5c07-462b-ac5b-b32ba89efe3d",
      "role": "CO_PILOT"
    }
  ]
}
```
