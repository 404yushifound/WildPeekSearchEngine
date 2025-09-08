# WildPeekSearchEngine
A simple Go-based search API for exotic animals. Users can search animals by name, color, height, weight, or species. The API returns matching animals in JSON format.

## Features

- REST API built with Go (`net/http`)  
- Case-insensitive search across multiple fields  
- Handles multiple matches gracefully  
- Returns results in JSON format  
- Minimal frontend for interactive searching  

## Tech Stack

- **Backend:** Go  
- **Data Format:** JSON  

### Available Animals

- Macaw (Bird) – Blue-Yellow, 85 cm, 1.2 kg
- Axolotl (Amphibian) – Pink, 23 cm, 250 g
- Fennec Fox (Mammal) – Cream, 20 cm, 1.5 kg
- Red Panda (Mammal) – Reddish-Brown, 60 cm, 6 kg
- Koi Fish (Fish) – Orange-White, 70 cm, 4 kg

##  How to Search
1.Run the server with:
  go run main.go

2.Then open your browser and visit:

http://localhost:10000/search?q=<term>

3.Replace <term> with any keyword (like macaw, pink, or mammal). The API will return matching animals in JSON format. Search is case-insensitive and works across name, color, height, weight, and species.
