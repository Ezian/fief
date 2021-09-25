<template>

<section class="hero  is-dark">
  <div class="hero-body">
    <div class="container">
      <div class="columns is-centered">
        <div class="column is-5-tablet is-4-desktop is-3-widescreen  box ">
          <Form @submit="handleRegister" :validation-schema="schema">
            <div v-if="!successful">
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
                <label for="email" class="label">Email</label>
                <div class="control has-icons-left">                
                  <Field name="email" type="email" class="input" required />
                  <ErrorMessage name="email" class="is-warning" />
                  <span class="icon is-small is-left">
                    <font-awesome-icon icon="envelope" />
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
                <span>Sign Up</span>
                </button>
              </div>
            </div>
          </Form>
              
          <div
            v-if="message"
            class="alert"
            :class="successful ? 'is-success' : 'is-danger'"
          >
            {{ message }}
          </div>
        </div>
      </div>
    </div>
  </div>
</section>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import authService, { RegisterUser } from "@/services/auth.service";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default defineComponent({
  name: "Register",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data() {
    const schema = yup.object().shape({
      username: yup
        .string()
        .required("Username is required!")
        .min(3, "Must be at least 3 characters!")
        .max(20, "Must be maximum 20 characters!"),
      email: yup
        .string()
        .required("Email is required!")
        .email("Email is invalid!")
        .max(50, "Must be maximum 50 characters!"),
      password: yup
        .string()
        .required("Password is required!")
        .min(6, "Must be at least 6 characters!")
        .max(40, "Must be maximum 40 characters!"),
    });

    return {
      successful: false,
      loading: false,
      message: "",
      schema,
    };
  },
  computed: {
    loggedIn() {
      return authService.isLogged();
    },
  },
  mounted() {
    if (this.loggedIn) {
      this.$router.push("/");
    }
  },
  methods: {
    handleRegister(user : RegisterUser) {
      this.message = "";
      this.successful = false;
      this.loading = true;

      authService.register(user).then(
        (data) => {
          this.message = data.message;
          this.successful = true;
          this.loading = false;
        },
        (error) => {
          this.message =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();
          this.successful = false;
          this.loading = false;
        }
      );
    },
  },
});
</script>