# SCDB Updater

Learn more about the project [here](https://scdb-landing-001e.twc1.net/) 🙂

[Документация на Русском языке](./README.md)

A system for updating the database of educational institutions in the Russian Federation.

## Description

This project is designed to collect, process, and update information about educational institutions in Russia. The system processes data from various sources and creates a structured database with information on:

- Educational institutions
- Regions
- Cities
- Federal districts
- Types of educational institutions

## Requirements

- Go 1.24.3 or higher
- Internet access for downloading data

## Installation

1. Clone the repository:

```bash
git clone https://github.com/MaksimAndreevich/scdb-updater.git
cd scdb-updater
```

2. Install dependencies:

```bash
go mod download
```

3. Create a `.env` file in the root directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=scdb
```

## Data Update Process

The system performs the following steps:

1. Import federal districts
2. Import regions
3. Import cities
4. Import types of educational institutions
5. Import educational institutions

## Types of Educational Institutions

The system supports the following types:

1. Preschool education
2. General education
3. Secondary vocational education
4. Higher education
5. Additional education
6. Vocational training
7. Special education
8. Military education
9. Religious education
10. International education

Detailed information about the types and their keywords is located in [data/org_types.json](data/org_types.json)

## Data Sources

- **Regions**: [russia-cities](https://github.com/arbaev/russia-cities)
- **Cities**: [city](https://github.com/hflabs/city)
- **Educational institutions**: [Organization Registry](https://obrnadzor.gov.ru/otkrytoe-pravitelstvo/opendata/7701537808-raoo/)
- **Types of educational institutions**: [Federal Law "On Education in the Russian Federation"](https://base.garant.ru/70291362/)

## Database Structure

### Tables

#### education_organizations

- id (text) - unique identifier
- full_name (text) - full name
- short_name (text) - short name
- post_address (text) - address
- phone (text) - phone
- email (text) - email
- website (text) - website
- inn (text) - INN (Taxpayer Identification Number)
- kpp (text) - KPP (Tax Registration Reason Code)
- ogrn (text) - OGRN (Primary State Registration Number)
- fk_city_id (text) - relation to city
- fk_region_id (int) - relation to region
- fk_federal_district_id (int) - relation to federal district
- fk_education_type_key (text) - relation to education type

#### education_types

- key (text) - unique type key
- title (text) - type name
- level (text) - education level
- ownership_forms (text\[]) - forms of ownership
- keywords (text\[]) - keywords for search

#### cities

- fias_id (text) - FIAS ID
- name (text) - name
- fk_region_id (int) - relation to region

#### regions

- id (int) - unique identifier
- name (text) - name
- fk_federal_district_id (int) - relation to federal district

#### federal_districts

- id (int) - unique identifier
- name (text) - name

## Logging

The system uses multi-level logging:

- INFO - informational messages
- WARNING - warnings
- ERROR - errors
- FATAL - critical errors

## Development

### Adding New Types of Educational Institutions

1. Open the file `data/org_types.json`
2. Add a new type to the array
3. Specify:

   - key - unique key
   - title - name
   - level - education level
   - ownership_forms - forms of ownership
   - keywords - keywords for search

### Adding New Data Sources

1. Create a new file in the `data/` directory
2. Add a parser in `internal/services/`
3. Update `main.go`
