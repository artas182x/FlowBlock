<template>
  <div class="jumbotron">
    <h1 class="display-4">Token submit</h1>
  </div>

  <div v-for="error in errors" v-bind:key="error" class="alert alert-danger">
    {{ error }}
  </div>

  <div class="form-group form-row col-md-4">
    <label>Chaincode: </label>
    <select v-model="chainCodeNameSelected" class="form-select" @change="onChainCodeChange()" >
      <option v-for="chainCode in chainCodes" v-bind:value="chainCode.value" v-bind:key="chainCode.value">
        {{ chainCode.text }}
      </option>
    </select>
  </div>

  <div class="form-group form-row col-md-4">
    <label>Method: </label>
    <select v-model="methodSelected" class="form-select">
      <option v-for="method in methods" v-bind:value="method" v-bind:key="method.Name">
        {{ method.Name }}
      </option>
    </select>
    <p>{{ methodSelected.Description }}</p>
  </div>

  <div class="form-group form-row col-md-2" v-for="argument in methodSelected.Arguments" v-bind:key="argument">
    <label>{{ argument.Name }}</label>
    <input class="form-text" v-if="argument.Type === 'string'" v-model="argument.Value">
    <datepicker v-if="argument.Type === 'ts'" v-model="argument.Value" />
  </div>

  <button type="button" :disabled='isSubmitDisabled()' class="btn btn-primary" @click="submit()">Submit</button>

</template>


<script>
import userService from "@/services/user.service";
import Datepicker from 'vue3-datepicker'

export default {
  name: "TokenSubmit",
  components: { Datepicker },
  data: () => ({
    errors: [],
    chainCodeNameSelected: "",
    chainCodes: [
      { text: 'Example algorithm', value: 'examplealgorithm' },
    ],
    methodSelected: {"Name": "", "Description": "", "RetType": "", "Arguments": []},
    methods: [
    ],
  }),
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
  },
  methods: {
    isSubmitDisabled() {
      return this.methodSelected.Name === "" || this.chainCodeNameSelected === ""
    },
    onChainCodeChange() {
      this.methods = []
      userService.getAvailableMethods(this.chainCodeNameSelected).then(
          (response) => {
            let methods = response.data
            methods.forEach(method => {
              method.Arguments.forEach(arg => {
                if (arg.Type === "ts") {
                  arg.Value = new Date();
                } else {
                  arg.Value = ""
                }
              })
              this.methods.push(method)
            })
          },
          (error) => {
            if (error.response.status === 401) {
              this.logOut()
            }
          }
      );
    },
    submit() {
        let argumentsFlat = []
        this.errors = []
        this.methodSelected.Arguments.forEach(arg => {
          if (arg.Type === "ts") {
            argumentsFlat.push('' + Math.round(arg.Value/1000))
          } else {
            argumentsFlat.push(arg.Value)
          }

        })
        let request = {"arguments": argumentsFlat, "chaincodeName": this.chainCodeNameSelected, "method": this.methodSelected.Name}
        userService.requestToken(request).then(
            () => {
              this.$router.push('/computations');
            },
            (error) => {
              if (error.response.status === 401) {
                this.logOut()
              } else {
                this.errors.push("Error during submitting a request")
              }
            }
        )
    },
    logOut() {
      this.$store.dispatch('auth/logout');
      this.$router.push('/login');
    }
  },
  mounted() {
    if (!this.currentUser) {
      this.$router.push('/login');
    }
  },

}
</script>

<style scoped>
.form-row {
  margin: 20px 0;
  padding-bottom: 10px;
}
</style>