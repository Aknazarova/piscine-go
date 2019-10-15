url=https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json
curl -s ${url} |jq '.[70]. name'