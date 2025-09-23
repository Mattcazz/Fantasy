import requests
import json
from bs4 import BeautifulSoup

player_url = "https://app.analiticafantasy.com/api/oraculo-fantasy"
market_url = "https://app.analiticafantasy.com/api/fantasy-players/mercado"

HEADERS = {
    "Accept": "application/json, text/plain, */*",
    "Content-Type": "application/json",
    "Origin": "https://www.analiticafantasy.com",
    "Referer": "https://www.analiticafantasy.com/",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "
                  "(KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36"
}


market_payload = {
    "last":1,"league":2
}


def scrape_get(url, headers, file_name):        
    response = requests.get(url, headers=headers)

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


def scrape_post(url, headers, payload,file_name):        
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


def team_scrape (url, headers, file_name):
    response  = requests.get(url, headers=headers)  
    soup = BeautifulSoup(response.text, "html.parser")

    script_tag = soup.find("script", {"id": "__NEXT_DATA__", "type": "application/json"})

    if script_tag:
        data = json.loads(script_tag.string)
        with open(file_name, "w", encoding="utf-8") as f:
            json.dump(data, f, ensure_ascii=False, indent=2)
        print(f"Archivo '{file_name}' guardado correctamente")
    else:
        print("Element not found")

def web_scrape():
    scrape_get(player_url, HEADERS, "player.json")
    team_scrape("https://analiticafantasy.com/clasificacion", HEADERS, "team.json")

web_scrape()
scrape_post(market_url, HEADERS, market_payload, "market.json")
