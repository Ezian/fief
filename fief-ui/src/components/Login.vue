<template>
<section class="hero  is-dark">
  <div class="hero-body">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-5-tablet is-4-desktop is-3-widescreen  box ">
          <Form @submit="handleLogin" :validation-schema="schema">
            <div class="field">
              <label for="username" class="label">Login</label>
              <div class="control has-icons-left">                
                <Field name="username" type="text" class="input" required />
                <ErrorMessage name="username" class="is-warning" />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="user" />
                </span>
              </div>
            </div>
            <div class="field">
              <label for="password" class="label">Password</label>                  
              <div class="control has-icons-left">
              <Field name="password" placeholder="*******" type="password" class="input" required />
              <ErrorMessage name="password" class="is-warning" />
                <span class="icon is-small is-left">
                  <font-awesome-icon icon="lock" />
                </span>
              </div>
            </div>
            <div class="field">
              <button class="button is-success" :disabled="loading">
              <span
                v-show="loading"
                class="spinner-border spinner-border-sm"
              ></span>
              <span>Login</span>
              </button>
              <div v-if="message" class="notification is-warning" role="alert">
                {{ message }}
              </div>
            </div>
          </Form>
        </div>
      </div>
    </div>
  </div>
</section>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import {Form, Field, ErrorMessage} from 'vee-validate'
import * as yup from 'yup'
import authService, { LoginInfo } from "@/services/auth.service";

export default defineComponent({
  name: "Login",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data(){
    const schema = yup.object().shape({
      username: yup.string().required("Username is required!"),
      password: yup.string().required("Password is required!"),
    });

    return {
      loading: false,
      message: "",
      schema,
    }
  },
  computed: {
    loggedIn(){
      return authService.isLogged()
    },
  },
  created(){
    if(this.loggedIn){
      this.$router.push("/")
    }
  },
  methods: {
    handleLogin(user: LoginInfo){
      this.loading = true
      authService.login(user).then(() => {
        this.$router.push("/")
      },
      error => {
        this.loading = false;
        this.message =
          (error.response &&
            error.response.data &&
            error.response.data.message) ||
          error.message ||
          error.toString();
      })
    }
  }
});
</script>