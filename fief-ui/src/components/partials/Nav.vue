<template>
<nav class="navbar container" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <a class="navbar-item" href="/">
      <strong class="is-size-4">Fief</strong>
    </a>
    <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>
  <div id="navbar" class="navbar-menu">
    <div class="navbar-start">
      <router-link to="/" class="navbar-item">Games</router-link>
      <router-link to="/about" class="navbar-item">About</router-link>
    </div>
    <div v-if="!isLogged()" class="navbar-end">
      <div class="navbar-item">
        <div class="buttons">
          <router-link to="/signin" class="button is-dark">
            <strong>Sign In</strong>
          </router-link>
        </div>
      </div>
      <div class="navbar-item">
        <div class="buttons">
          <router-link to="/signup" class="button is-dark">
            <strong>Sign up</strong>
          </router-link>
        </div>
      </div>
    </div>
    <div v-if="isLogged()" class="navbar-end">
      <div class="navbar-item">
        Logged as {{  userName() }}
      </div>
      <div class="navbar-item">
        <div class="buttons">
          <button class="button is-danger" @click="logout()">
            <font-awesome-icon icon="sign-out-alt" />
          </button>
        </div>
      </div>
    </div>
  </div>
</nav>
</template>
<script lang="ts">
import authService from "@/services/auth.service"
import { defineComponent } from "@vue/runtime-core"

export default  defineComponent({
    name: 'Nav',
    methods: {
      isLogged(): boolean{
        return authService.isLogged();
      },
      userName(): string{
        return authService.loggedUserInfo().login;
      },
      logout(): void{
        authService.logout()
      }
    }
});
</script>

<style scoped>
  nav {
    margin-top: 25px;
    margin-bottom: 30px;
    a {
      font-weight: bold;
      color: #2c3e50;
      &.router-link-exact-active {
        color: #d88d00;
      }
    }  
  } 
</style>