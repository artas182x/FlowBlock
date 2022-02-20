<template>
  <div class="container">
    <header class="jumbotron">
      <h3>
        Profile
      </h3>
    </header>
    <p>
      <strong>Token:</strong>
      {{ loginData.token.substring(0, 20) }} ... {{ loginData.token.substr(loginData.token.length - 20) }}
    </p>
    <p>
      <strong>Certificate:</strong>
      {{ loginData.user.Login.certificate.substring(40, 60)  }} ...
    </p>
    <p>
      <strong>Id:</strong>
      {{ loginData.user.UserName }}
    </p>
    <p>
      <strong>Organization ID:</strong>
      {{ loginData.user.Login.mspid }}
    </p>
    <strong>Roles:</strong>
    <ul>
      <li v-for="role in loginData.user.Roles" :key="role">{{role}}</li>
    </ul>
  </div>
</template>

<script>
import UserService from "@/services/user.service";

export default {
  name: 'Profile',
  computed: {
    loginData() {
      return this.$store.state.auth.user;
    }
  },
  mounted() {
    if (!this.loginData) {
      this.$router.push('/login');
    }
    UserService.refreshToken().then(
        () => {},
        (error) => {
          if (error.response.status === 401) {
            this.logOut()
          }
        }
    )
  }
};
</script>