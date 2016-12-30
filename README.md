#Tank in Dungeon API #

Small and fast API for game [Tank in Dungeon](https://github.com/MrJaqbq/LudumDare36)

Endpoints: 
* http://domain:/results/N - GET, where N is a number of the best results
```json
[
  {
    "PlayerName": "John",
    "Score": 605,
    "Time": 31
  },
  {
    "PlayerName": "Sylwia",
    "Score": 100,
    "Time": 756
  },
  {
    "PlayerName": "Lukasz",
    "Score": 5,
    "Time": 51
  }
]
```

* http://domain:/results - POST, with data:

```json
{
    "PlayerName": "Tommy",
    "Score": 100,
    "Time": 71
}
```