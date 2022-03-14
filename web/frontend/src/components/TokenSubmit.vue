<template>
  <div class="jumbotron">
    <h1 class="display-4">
      Workflow designer
    </h1>
    <p>Click right button on board to start designing</p>
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
    class="btn btn-primary "
    @click="save()"
  >
    Save
  </button>

  <div style="height: 70vh; width: 100%" class="form-group form-row col-md-4">
    <hint-overlay />
    <baklava-editor :plugin="viewPlugin" />
  </div>

</template>

<style>
.node-editor .background {
  background-color: #fafafa;
}

.node {
  background: #212121;
}

.dark-input {
  background-color: #ffffff;
  color: #484848;

}

.dark-input:hover {
  background-color: #d6d6d7;
}

.dark-context-menu {
  background: #212121;
}

.connection {
  stroke: #593196;
}

</style>


<script>
import userService from "@/services/user.service";
import {Editor} from "@baklavajs/core";
import {ViewPlugin} from "@baklavajs/plugin-renderer-vue3";
import {InputOption, OptionPlugin} from "@baklavajs/plugin-options-vue3";
import {ComputeBlockBuilder} from "@/components/ComputeBlockBuilder";
import DateOption from "@/components/DateOption";
import MetadataOption from "@/components/MetadataOption";
import AddOption from "@/components/AddOption.vue";

export default {
  name: "TokenSubmit",
  components: {  },
  data: () => ({
    errors: [],
    chainCodeNameSelected: "",
    chainCodes: [
      { text: 'Example algorithm', value: 'examplealgorithm' },
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
    save() {
      const state = JSON.stringify(this.editor.save());
      console.log(state);
      userService.requestFlow(state).then(
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