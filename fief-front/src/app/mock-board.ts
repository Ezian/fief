import { Board, City, Road } from './model'

export const BOARD: Board = {
 cities: [
   {id: 10, name: "Blaye"},
   {id: 11, name: "La Salle"},
   {id: 12, name: "Beaujeu"},
   {id: 13, name: "Charolles"},
   {id: 20, name: "Sigy"},
   {id: 21, name: "Saint Clers D'Abzac"},
   {id: 22, name: "Marcamps"},
   {id: 30, name: "Saint Paul"},
   {id: 31, name: "Les Essarts"},
   {id: 40, name: "Bourg"},
   {id: 41, name: "Saint Vivien"},
   {id: 42, name: "Villeneuve"},
   {id: 43, name: "La Bussière"},
   {id: 50, name: "Saint Gérôme"},
   {id: 51, name: "Saint Médard"},
   {id: 52, name: "Sant André"},
   {id: 60, name: "Tournus"},
   {id: 61, name: "Belleville"},
   {id: 62, name: "Pugnac"},
   {id: 70, name: "Mazilles"},
   {id: 71, name: "Saint Andrôme"},
   {id: 80, name: "Sennecy"},
   {id: 81, name: "L'Epervier"},
   {id: 82, name: "Châteauneuf"},
 ],
 roads: [
   {cityId1: 10,  cityId2:13},
   {cityId1: 10,  cityId2:11},
   {cityId1: 11,  cityId2:12},
   {cityId1: 12,  cityId2:21},
   {cityId1: 20,  cityId2:21},
   {cityId1: 21,  cityId2:22},
   {cityId1: 20,  cityId2:30},
   {cityId1: 22,  cityId2:31},
   {cityId1: 31,  cityId2:42},
   {cityId1: 30,  cityId2:31},
   {cityId1: 30,  cityId2:41},
   {cityId1: 40,  cityId2:41},
   {cityId1: 40,  cityId2:42},
   {cityId1: 40,  cityId2:43},
   {cityId1: 13,  cityId2:42},
   {cityId1: 42,  cityId2:51},
   {cityId1: 43,  cityId2:52},
   {cityId1: 50,  cityId2:51},
   {cityId1: 50,  cityId2:52},
   {cityId1: 50,  cityId2:62},
   {cityId1: 51,  cityId2:71},
   {cityId1: 60,  cityId2:62},
   {cityId1: 60,  cityId2:71},
   {cityId1: 61,  cityId2:62},
   {cityId1: 61,  cityId2:70},
   {cityId1: 70,  cityId2:71},
   {cityId1: 70,  cityId2:82},
   {cityId1: 13,  cityId2:81},
   {cityId1: 80,  cityId2:81},
   {cityId1: 81,  cityId2:82},
   {cityId1: 11,  cityId2:80},
 ],
}
