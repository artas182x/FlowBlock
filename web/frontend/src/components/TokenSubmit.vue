<template>
  <div class="jumbotron">
    <h1 class="display-4">
      Token submit
    </h1>
  </div>

  <div
    v-for="error in errors"
    :key="error"
    class="alert alert-danger"
  >
    {{ error }}
  </div>

  <button
    type="button"
    class="btn btn-primary"
    @click="save()"
  >
    Save
  </button>

  <div style="height: 100vh; width: 100vw">
    <hint-overlay />
    <baklava-editor :plugin="viewPlugin" />
  </div>

  <div class="form-group form-row col-md-4">
    <label>Chaincode: </label>
    <select
      v-model="chainCodeNameSelected"
      class="form-select"
      @change="onChainCodeChange()"
    >
      <option
        v-for="chainCode in chainCodes"
        :key="chainCode.value"
        :value="chainCode.value"
      >
        {{ chainCode.text }}
      </option>
    </select>
  </div>

  <div class="form-group form-row col-md-4">
    <label>Method: </label>
    <select
      v-model="methodSelected"
      class="form-select"
    >
      <option
        v-for="method in methods"
        :key="method.Name"
        :value="method"
      >
        {{ method.Name }}
      </option>
    </select>
    <p>{{ methodSelected.Description }}</p>
  </div>

  <div
    v-for="argument in methodSelected.Arguments"
    :key="argument"
    class="form-group form-row col-md-2"
  >
    <label>{{ argument.Name }}</label>
    <input
      v-if="argument.Type === 'string'"
      v-model="argument.Value"
      class="form-text"
    >
    <datepicker
      v-if="argument.Type === 'ts'"
      v-model="argument.Value"
    />
  </div>

  <button
    type="button"
    :disabled="isSubmitDisabled()"
    class="btn btn-primary"
    @click="submit()"
  >
    Submit
  </button>
</template>


<script>
import userService from "@/services/user.service";
import Datepicker from 'vue3-datepicker'
import {Editor} from "@baklavajs/core";
import {ViewPlugin} from "@baklavajs/plugin-renderer-vue3";
import {InputOption, OptionPlugin} from "@baklavajs/plugin-options-vue3";
import {ComputeBlockBuilder} from "@/components/ComputeBlockBuilder";
import DateOption from "@/components/DateOption";
import MetadataOption from "@/components/MetadataOption";
import TokenNode from "@/components/TokenNode.ts";
import AddOption from "@/components/AddOption.vue";

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
    editor: new Editor(),
    viewPlugin: new ViewPlugin()
  }),
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
  },
  created() {
    this.editor.use(this.viewPlugin);
    this.editor.use(new OptionPlugin());

    this.viewPlugin.registerOption("DateOption", DateOption);
    this.viewPlugin.registerOption("InputOption", InputOption);
    this.viewPlugin.registerOption("MetadataOption", MetadataOption);
    this.viewPlugin.registerOption("AddOption", AddOption);

    this.chainCodes.forEach(chaincode => {

      userService.getAvailableMethods(chaincode.value).then(
          (response) => {
            let methods = response.data
            methods.forEach(method => {

              const Block1 = ComputeBlockBuilder({
                MethodData: method,
                ChaincodeName: chaincode.value
              });
              this.editor.registerNodeType(method.Name, Block1);

            })
          },
          (error) => {
            if (error.response.status === 401) {
              this.logOut()
            }
          }
      );
    });



  },
  mounted() {
    if (!this.currentUser) {
      this.$router.push('/login');
    }
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
    save() {
      const state = JSON.stringify(this.editor.save());
      console.log("onSaveButtonClick:", state);
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

}
</script>

<style scoped>
.form-row {
  margin: 20px 0;
  padding-bottom: 10px;
}
</style>