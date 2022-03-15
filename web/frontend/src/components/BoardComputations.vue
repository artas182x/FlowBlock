<template>
  <div class="jumbotron">
    <h1 class="display-4">
      My workflows
    </h1>
  </div>
  <div>
    <div
      v-if="submitSuccess"
      class="alert alert-success"
    >
      You successfully scheduled a task
    </div>

    <div
      v-if="submitError"
      class="alert alert-danger"
    >
      Error during scheduling a task
    </div>

    <button
      type="button"
      class="btn btn-primary float-sm-end"
      @click="goToTokenSubmit()"
    >
      Workflow designer
    </button>

    <h1>Queue</h1>

    <table-lite
      :is-slot-mode="true"
      :is-loading="tableQueue.isLoading"
      :columns="tableQueue.columns"
      :rows="tableQueue.rows"
      :total="tableQueue.totalRecordCount"
      @do-search="doSearchQueue"
    >
      <template #name="data">
        {{ data.value.name }}
      </template>
    </table-lite>

    <p />
    <h1>My tokens</h1>

    <table-lite
      :is-slot-mode="true"
      :is-loading="tableTokens.isLoading"
      :columns="tableTokens.columns"
      :rows="tableTokens.rows"
      :total="tableTokens.totalRecordCount"
      @do-search="doSearchTokens"
      @is-finished="tableLoadingFinish"
    >
      <template #name="data">
        {{ data.value.name }}
      </template>
    </table-lite>
  </div>
</template>

<script>
import { reactive } from "vue";
import TableLite from 'vue3-table-lite'
import UserService from "@/services/user.service";
import userService from "@/services/user.service";
import moment from "moment";
import AuthService from "@/services/auth.service";

export default {
  components: { TableLite },
  data() {
    return {
      submitSuccess: false,
      submitError: false
    }
  },
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
    tableTokens() {
      return reactive({
        isLoading: true,
        columns: [
          {
            label: "Chaincode Name",
            field: "chaincodeName",
          },
          {
            label: "Method",
            field: "method",
          },
          {
            label: "Arguments",
            field: "arguments",
          },
          {
            label: "Returned value",
            field: "retval",
            display:  (row) => {

              if (row.retType !== "s3img") {
                return ('<div style="margin-left: 10px;">' + row.retval + "")
              } else {
                return (
                    '<button type="button" title="Download" data-id="' +
                    row.retval.split("?")[0] +
                    '" class="is-rows-el download-btn btn btn-light">&#128229;</button>' +

                    '<button title="Copy SHA256 checksum to clipboard" type="button" data-id="' +
                    row.retval.split("?")[1] +
                    '" class="is-rows-el copy-btn btn btn-light">&#x1f4cb;</button>'

                );
              }



            },
          },
          {
            label: "Time requested",
            field: "timeRequested",
          },
          {
            label: "Token expiration",
            field: "expireTime",
          },
          {
            label: "Control",
            field: "control",
            display:  (row) => {
              let alreadyRunning = false;
              this.tableQueue.rows.forEach(element => {
                if(element.id === row.id) {
                  alreadyRunning = true;
                }
              });

              if (alreadyRunning) {
                return ("Already running")
              } else if (!row.directlyExecutable) {
                return ("Dependent on token above")
              } else if(moment(row.expireTime, 'MMMM Do YYYY, h:mm:ss a').isBefore(moment(Date.now())) && row.retval === "") {
                return("Token has expired and workflow wasn't started")
              } else if(moment(row.expireTime, 'MMMM Do YYYY, h:mm:ss a').isAfter(moment(Date.now())) && row.retval === "") {
                return (
                    '<button type="button" data-id="' +
                    row.id +
                    '" class="is-rows-el run-btn btn btn-light">&#x23F5</button>'
                );
              } else {
                return('');
              }
            },
          },
        ],
        rows: [],
        totalRecordCount: 0,
      });
    },
    tableQueue() {
      return reactive({
        isLoading: true,
        columns: [
          {
            label: "Chaincode Name",
            field: "chaincodeName",
          },
          {
            label: "Method",
            field: "method",
          },
          {
            label: "Arguments",
            field: "arguments",
          },
          {
            label: "Time requested",
            field: "timeRequested",
          },
        ],
        rows: [],
        totalRecordCount: 0,
      });
    }
  },
  mounted() {
    if (!this.currentUser) {
      this.$router.push('/login');
    }
    AuthService.refreshToken()
    this.doSearchQueue(0, 10)
    this.doSearchTokens(0, 10)
  },
  methods: {
    goToTokenSubmit(){
      this.$router.push('/tokensubmit');
    },
    tableLoadingFinish(elements) {
      this.tableTokens.isLoading = false;
      Array.prototype.forEach.call(elements, (element) => {
        if (element.classList.contains("run-btn")) {
          element.addEventListener("click", () => {
            userService.startComputation(element.getAttribute("data-id")).then(
                () => {
                  this.submitSuccess = true;
                  this.submitError = false;
                  this.doSearchQueue(0, 20);
                },
                () => {
                  this.submitSuccess = false;
                  this.submitError = true;
                  this.doSearchQueue(0, 20);
                }
            );
          });
        }
        if (element.classList.contains("details-btn")) {
          element.addEventListener("click", function () {
            console.log(this.dataset.id + " details-btn click!!");
          });
        }
        if (element.classList.contains("download-btn")) {
          element.addEventListener("click", () => {

            UserService.downloadFile(element.getAttribute("data-id")).then((response) => {
              const fileURL = window.URL.createObjectURL(new Blob([response.data]));
              const fileLink = document.createElement('a');

              fileLink.href = fileURL;
              fileLink.setAttribute('download', element.getAttribute("data-id"));
              document.body.appendChild(fileLink);

              fileLink.click();
            })
          });
        }
        if (element.classList.contains("copy-btn")) {
          element.addEventListener("click", () => {
            navigator.clipboard.writeText(element.getAttribute("data-id"))
          });
        }
      });
    },
    doSearchQueue(offset, limit) {
      this.tableQueue.isLoading = true;
      setTimeout(() => {
        this.tableQueue.isReSearch = offset === undefined;
        if (offset >= 10 || limit >= 20) {
          limit = 20;
        }

        UserService.getUserQueue().then(
            (response) => {
              let tokens = response.data
              let data = [];
              const max = Math.min(offset+limit, tokens.length);
              for (let i = offset; i < max ; i++) {

                let argsStrings = tokens[i]["Result"]["Arguments"]

                let args = [];

                argsStrings.forEach((value) => {
                  let argVal = value["Value"]
                  const argName = value["Name"]
                  const argType = value["Type"]

                  if (argType === "ts") {
                    argVal = moment(new Date(parseInt(argVal))).format('MMMM Do YYYY, h:mm:ss a')
                  } else if (argType === "tokenInputs") {
                    argVal = "..." + argVal.substr(argVal.length-25, 25)
                  }

                  args.push(argName + ": " + argVal)
                });

                data.push({
                  id: tokens[i]["Result"]["ID"],
                  chaincodeName: tokens[i]["Result"]["ChaincodeName"],
                  method: tokens[i]["Result"]["Method"].split(":")[1],
                  arguments: args,
                  timeRequested: moment(new Date(tokens[i]["Result"]["TimeRequested"])).format('MMMM Do YYYY, h:mm:ss a'),
                  finished: tokens[i]["Finished"] ? "Yes" : "No",
                });
              }
              this.tableQueue.rows = data;
              this.tableQueue.totalRecordCount = tokens.length;
              this.tableQueue.isLoading = false;
            },
            (error) => {
              if (error.response.status === 401) {
                this.logOut()
              }
            }
        );

      }, 600);
    },
    doSearchTokens(offset, limit) {
      this.tableTokens.isLoading = true;
      setTimeout(() => {
        this.tableTokens.isReSearch = offset === undefined;
        if (offset >= 10 || limit >= 20) {
          limit = 20;
        }

        UserService.getUserTokens().then(
            (response) => {
              let tokens = response.data
              let data = [];
              const max = Math.min(offset+limit, tokens.length);
              for (let i = offset; i < max ; i++) {

                let argsStrings = tokens[i]["Arguments"]

                let args = [];

                argsStrings.forEach((value) => {
                  let argVal = value["Value"]
                  const argName = value["Name"]
                  const argType = value["Type"]

                  if (argType === "ts") {
                     argVal = moment(new Date(parseInt(argVal)*1000)).format('MMMM Do YYYY, h:mm:ss a')
                  } else if (argType === "s3img") {
                    const fileName = argVal.split("?")[0]
                    const fileSum = argVal.split("?")[1]
                    argVal = fileName + " (" + fileSum.substr(7) + "...)"
                  } else if (argType === "tokenInputs") {
                    argVal = "..." + argVal.substr(argVal.length-25, 25)
                  }

                  args.push(argName + ": " + argVal)

                });

                data.push({
                  id: tokens[i]["ID"],
                  chaincodeName: tokens[i]["ChaincodeName"],
                  directlyExecutable: tokens[i]["DirectlyExecutable"],
                  method: tokens[i]["Method"].split(":")[1],
                  arguments: args,
                  retval: tokens[i]["ret"]["RetValue"],
                  retType: tokens[i]["ret"]["RetType"],
                  timeRequested: moment(new Date(tokens[i]["TimeRequested"])).format('MMMM Do YYYY, h:mm:ss a'),
                  expireTime: moment(new Date(tokens[i]["ExpirationTime"])).format('MMMM Do YYYY, h:mm:ss a'),
                });
              }
              data.sort((a,b) => (a.timeRequested < b.timeRequested) ? 1 : ((b.timeRequested < a.timeRequested) ? -1 : 0))
              this.tableTokens.rows = data;
              this.tableTokens.totalRecordCount = tokens.length;
              this.tableTokens.isLoading = false;
            },
            (error) => {
              if (error.response.status === 401) {
                this.logOut()
              }
            }
        );

      }, 600);
    },
    logOut() {
      this.$store.dispatch('auth/logout');
      this.$router.push('/login');
    }
  },

}
</script>

<style>
::v-deep(.vtl-table .vtl-thead .vtl-thead-th)  {
  color: #000000;
  background-color: #e9ecef;
  border-color: #e9ecef;
}

table {
  display: table;
  table-layout: fixed;
  width: auto;
  word-break: break-all;
}

</style>