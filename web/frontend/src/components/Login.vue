<template>
  <div class="col-md-12">
    <div class="card card-container">
      <img
        id="profile-img"
        src="//ssl.gstatic.com/accounts/ui/avatar_2x.png"
        class="profile-img-card"
      >
      <Form :validation-schema="schema">
        <div class="form-group">
          <label for="wallet">Wallet</label>
          <input
            type="file"
            class="inputFile"
            @change="onFileChange"
          >
          <ErrorMessage
            name="certificate"
            class="error-feedback"
          />
        </div>

        <div class="form-group">
          <div
            v-if="message"
            class="alert alert-danger"
            role="alert"
          >
            {{ message }}
          </div>
        </div>
      </Form>
    </div>
  </div>
</template>

<script>
import { Form, ErrorMessage } from "vee-validate";
import * as yup from "yup";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Login",
  components: {
    Form,
    ErrorMessage,
  },
  data() {
    const schema = yup.object().shape({
      certificate: yup.string().required("Certificate is required!"),
      privateKey: yup.string().required("Private Key is required!"),
      mspid: yup.string().required("MSP ID"),
    });

    return {
      loading: false,
      message: "",
      schema,
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    },
  },
  created() {
    if (this.loggedIn) {
      this.$router.push("/profile");
    }
  },
  methods: {
    onFileChange(e) {
     let files = e.target.files || e.dataTransfer.files;
     if (!files.length) return;
     this.readFile(files[0]);
   },
   readFile(file) {
     let reader = new FileReader();
     reader.onload = e => {
       let json = JSON.parse(e.target.result);
       const user = {};
       user.certificate = json.credentials.certificate;
       user.privateKey = json.credentials.privateKey;
       user.mspid = json.mspId;

      this.$store.dispatch("auth/login", user).then(
        () => {
          this.$router.push("/profile");
        },
        (error) => {
          this.loading = false;
          this.message =
            (error.response &&
              error.response.data &&
              error.response.data.message) ||
            error.message ||
            error.toString();
        }
      );
     };
     reader.readAsText(file);
    },
  },
};
</script>

<style scoped>
label {
  display: block;
  margin-top: 10px;
}

.card-container.card {
  max-width: 350px !important;
  padding: 40px 40px;
}

.inputFile {
  max-width: 300px !important;
}

.card {
  background-color: #f7f7f7;
  padding: 20px 25px 30px;
  margin: 0 auto 25px;
  margin-top: 50px;
  -moz-border-radius: 2px;
  -webkit-border-radius: 2px;
  border-radius: 2px;
  -moz-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  -webkit-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
}

.profile-img-card {
  width: 96px;
  height: 96px;
  margin: 0 auto 10px;
  display: block;
  -moz-border-radius: 50%;
  -webkit-border-radius: 50%;
  border-radius: 50%;
}

.error-feedback {
  color: red;
}
</style>
