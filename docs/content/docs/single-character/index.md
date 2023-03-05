---
weight: 3
title: Get Character by ID
summary: Retrieve character info by their ID.
---
If you know the id of the character you want to get, you can request just their info.

Endpoint:
```
/api/characters/{id}
```

Example Request: 
```
/api/characters/2
```

Example Response:
```json
{
  "id": 2,
  "name": "Aang",
  "url": "https://avatar.fandom.com/wiki/Aang",
  "image": "https://static.wikia.nocookie.net/avatar/images/a/ae/Aang_at_Jasmine_Dragon.png/revision/latest?cb=20130612174003",
  "nicknames": [
    "Aangy (by Koko)",
    "Sweetie (by Katara)",
    "Twinkle Toes (by Toph)"
  ],
  "aliases": [
    "Bonzu Pippinpaddleopsicopolis the Third",
    "Kuzon (while at the Fire Nation school)"
  ],
  "nationality": [
    "Southern Air Temple"
  ],
  "ethnicity": [
    "Air Nomad"
  ],
  "fightingStyle": [
    "Airbending",
    "waterbending (Northern and Southern style)",
    "earthbending (Chu Gar Praying Mantis Kung Fu)",
    "firebending (Dancing Dragon)",
    "energybending"
  ],
  "died": [
    "Spring 100 AG (revived by Katara using spirit water)",
    "153 AG"
  ],
  "hair_color": null,
  "gender": "Man",
  "weapons": [
    "The elements",
    "glider staff"
  ],
  "profession": [
    "Air Nomad culture teacher",
    "Airbending instructor",
    "Avatar",
    "Monk"
  ],
  "affiliation": [
    "Air Acolytes",
    "Air Nomads",
    "Team Avatar"
  ],
  "voiced_by": null
}
```