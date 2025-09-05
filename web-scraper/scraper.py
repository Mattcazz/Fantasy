import requests
import json

url = "https://app.analiticafantasy.com/api/fantasy-players/mercado"

headers = {
    "Accept": "application/json, text/plain, */*",
    "Content-Type": "application/json",
    "Origin": "https://www.analiticafantasy.com",
    "Referer": "https://www.analiticafantasy.com/",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 "
                  "(KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36"
}

payload = {
    "last":1,"league":2
}

response = requests.post(url, headers=headers, json=payload)

# Comprobar si la respuesta es JSON
if response.headers.get("Content-Type", "").startswith("application/json"):
    data = response.json()
    # Guardar el archivo completo
    with open("mercado.json", "w", encoding="utf-8") as f:
        json.dump(data, f, ensure_ascii=False, indent=2)
    print("Archivo 'mercado.json' guardado correctamente")
else:
    print("La respuesta no es JSON")
    print("Status:", response.status_code)
    print(response.text[:500])
