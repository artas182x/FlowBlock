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
      {{ loginData.user.Login.certificate.substring(40, 60) }} ...
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
      <li
        v-for="role in loginData.user.Roles"
        :key="role"
      >
        {{ role }}
      </li>
    </ul>
  </div>
</template>

<script>
import AuthService from "@/services/auth.service";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
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
    AuthService.refreshToken()
  }
};
</script>