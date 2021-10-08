
export type GameInfo = {
  id: string,
  name: string,
  players: {
    joined: number,
    required: number,
  },
  status: string,
}
