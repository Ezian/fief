import { LOCAL_STORAGE_AUTH_TOKEN_KEY } from "./auth.constants";

export default function authHeader(){
  const storedAuthToken = localStorage.getItem(LOCAL_STORAGE_AUTH_TOKEN_KEY)
  if (storedAuthToken){
    const authToken = JSON.parse(storedAuthToken as string)
    return { Authorization: 'Bearer ' + authToken}
  } else {
    return {}
  }
}