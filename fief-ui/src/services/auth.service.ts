import axios from "axios"
import { API_URL, LOCAL_STORAGE_AUTH_TOKEN_KEY } from "./auth.constants"

const SIGNIN_PATH = 'auth/signin'
const SIGNUP_PATH = 'auth/signup'

/**
 * Signin informations
 */
export type LoginInfo = {
  login: string,
  password: string, 
}

/**
 * Response from Auth server after signin
 */
export type LoginSuccess = {
  success: boolean,
  token: string,
}

/**
 * Information required to register an user
 */
export type RegisterUser = {
  email: string,
  login: string,
  password: string
}

/**
 * Response from Auth server after signup
 */
export type SuccessResponse = {
  success: boolean,
  message: string,
}

/**
 * Handle user registration and authentication
 */
class AuthService {

  /**
   * Authentication through AWT Token
   * 
   * Send user/password to auth server then retrieve and store in local storage the accessToken
   * 
   * @param loginInfo - The login and password
   * @returns A promise containing a boolean (indicating if the login has been successful or not) and the accessToken
   */
  async login(loginInfo: LoginInfo): Promise<LoginSuccess>{
    const anyResponse = await axios.post(API_URL + SIGNIN_PATH, {
      login: loginInfo.login,
      password: loginInfo.password
    })
    let response = anyResponse.data as LoginSuccess
    if (response.success && response.token) {
      localStorage.setItem(LOCAL_STORAGE_AUTH_TOKEN_KEY, JSON.stringify(response.token))
    }
    return response
  }

  /**
   * Delete the access token from local storage
   */
  logout(): void {
    localStorage.removeItem(LOCAL_STORAGE_AUTH_TOKEN_KEY)
  }

/**
 * Register a user to the Auth server
 *
 * @param user - User information to register
 * @returns a boolean for success and a message indicating cause of failure.
 */
  register(user: RegisterUser): Promise<SuccessResponse>{
    return axios.post(API_URL + SIGNUP_PATH, {
      email: user.email,
      password: user.password,
      login: user.login,
    })
  }

}

export default new AuthService()