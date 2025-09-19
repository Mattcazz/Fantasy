import requests
import json

player_url = "https://app.analiticafantasy.com/api/fantasy-stats/get-fantasy-stats"
market_url = "https://app.analiticafantasy.com/api/fantasy-players/mercado"

HEADERS = {
    "Accept": "application/json, text/plain, */*",
    "Content-Type": "application/json",
    "Origin": "https://www.analiticafantasy.com",
    "Referer": "https://www.analiticafantasy.com/",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "
                  "(KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36"
}

player_payload = {
    "league":1, 
    "position" : 0, 
    "season" : 2025,
    "teams": None, 
    "week": -1 
}

market_payload = {
    "last":1,"league":2
}


def scrape(url, headers, payload, file_name):        
    response = requests.post(url, headers=headers, json=payload)

    # Comprobar si la respuesta es JSON
    if response.headers.get("Content-Type", "").startswith("application/json"):
        data = response.json()
        # Guardar el archivo completo
        with open(file_name, "w", encoding="utf-8") as f:
            json.dump(data, f, ensure_ascii=False, indent=2)
        print(f"Archivo '{file_name}' guardado correctamente")
    else:
        print("La respuesta no es JSON")
        print("Status:", response.status_code)
        print(response.text[:500])

scrape(market_url, HEADERS, market_payload, "market.json")
scrape(player_url, HEADERS, player_payload, "player.json")
