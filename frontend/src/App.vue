<style lang="scss">
@import 'styles/custom.scss';
</style>
<template>
  <div id="app">

    <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarTogglerDemo01" aria-controls="navbarTogglerDemo01" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarTogglerDemo01">
        <a class="navbar-brand" href="#">Blockchain data processor</a>
        <ul class="navbar-nav me-auto mt-2 mt-lg-0">
          <li class="nav-item">
            <router-link to="/home" class="nav-link">
              <font-awesome-icon icon="home" /> Home
            </router-link>
          </li>
          <li v-if="showComputationBoard" class="nav-item">
            <router-link to="/computations" class="nav-link">
              <font-awesome-icon icon="microchip" /> Workflows
            </router-link>
          </li>
        </ul>
        <div v-if="!loginData" class="navbar-nav ml-auto">
          <li class="nav-item">
            <router-link to="/register" class="nav-link">
              <font-awesome-icon icon="user-plus" /> Sign Up
            </router-link>
          </li>
          <li class="nav-item">
            <router-link to="/login" class="nav-link">
              <font-awesome-icon icon="sign-in-alt" /> Login
            </router-link>
          </li>
        </div>

        <div v-if="loginData" class="navbar-nav ml-auto">
          <li class="nav-item">
            <router-link to="/profile" class="nav-link">
              <font-awesome-icon icon="user" />
              {{ loginData.user.UserName }}
            </router-link>
          </li>
          <li class="nav-item">
            <a class="nav-link" @click.prevent="logOut">
              <font-awesome-icon icon="sign-out-alt" /> LogOut
            </a>
          </li>
        </div>
      </div>
    </nav>
      <p></p>
    <div class="container">
      <router-view />
    </div>
  </div>
</template>

<script>
export default {
  computed: {
    loginData() {
      return this.$store.state.auth.user;
    },
    showComputationBoard() {
      if (this.loginData && this.loginData.user.Roles) {
        return this.loginData.user.Roles.includes('computation');
      }

      return false;
    },

  },
  methods: {
    logOut() {
      this.$store.dispatch('auth/logout');
      this.$router.push('/login');
    }
  }
};
</script>
