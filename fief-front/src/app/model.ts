export interface Board{
  cities: City[],
  roads: Road[]
}

export interface City{
  id: number,
  name: string,
}

export interface Road{
  cityId1: number,
  cityId2: number,
}
