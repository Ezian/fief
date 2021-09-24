import { LOCAL_STORAGE_AUTH_TOKEN_KEY } from "./auth.constants";

export default function authHeader(){
  let storedAuthToken = localStorage.getItem(LOCAL_STORAGE_AUTH_TOKEN_KEY)
  if (storedAuthToken){
    let authToken = JSON.parse(storedAuthToken as string)
    return { Authorization: 'Bearer ' + authToken}
  } else {
    return {}
  }
}