<template>
  <div class="jumbotron">
    <h1 class="display-4">My workflows</h1>
  </div>
  <div>

    <div v-if="submitSuccess" class="alert alert-success">
      You successfully scheduled a task
    </div>

    <div v-if="submitError" class="alert alert-danger">
      Error during scheduling a task
    </div>

    <button type="button" class="btn btn-primary float-sm-end" @click="goToTokenSubmit()">Submit a workflow</button>

    <h1>Queue</h1>

    <table-lite
        :is-slot-mode="true"
        :is-loading="tableQueue.isLoading"
        :columns="tableQueue.columns"
        :rows="tableQueue.rows"
        :total="tableQueue.totalRecordCount"
        @do-search="doSearchQueue"
    >
      <template v-slot:name="data">
        {{ data.value.name }}
      </template>
    </table-lite>

    <p></p>
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
      <template v-slot:name="data">
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
          },
          {
            label: "Time requested",
            field: "timeRequested",
          },
          {
            label: "Token expiration time",
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
                return("Already running")
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
                  this.doSearchTokens(0, 20);
                },
                () => {
                  this.submitSuccess = false;
                  this.submitError = true;
                  this.doSearchQueue(0, 20);
                  this.doSearchTokens(0, 20);
                }
            );
          });
        }
        if (element.classList.contains("details-btn")) {
          element.addEventListener("click", function () {
            console.log(this.dataset.id + " details-btn click!!");
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
              const max = Math.min(limit, tokens.length);
              for (let i = offset; i < max ; i++) {
                data.push({
                  id: tokens[i]["Result"]["ID"],
                  chaincodeName: tokens[i]["Result"]["ChaincodeName"],
                  method: tokens[i]["Result"]["Method"].split(":")[1],
                  arguments: tokens[i]["Result"]["Arguments"],
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
              const max = Math.min(limit, tokens.length);
              for (let i = offset; i < max ; i++) {
                data.push({
                  id: tokens[i]["ID"],
                  chaincodeName: tokens[i]["ChaincodeName"],
                  method: tokens[i]["Method"].split(":")[1],
                  arguments: tokens[i]["Arguments"].split(";"),
                  retval: tokens[i]["ret"]["RetValue"],
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
  mounted() {
    if (!this.currentUser) {
      this.$router.push('/login');
    }
    this.doSearchQueue(0, 20)
    this.doSearchTokens(0, 20)
  },

}
</script>

<style scoped>
::v-deep(.vtl-table .vtl-thead .vtl-thead-th) {
  color: #000000;
  background-color: #e9ecef;
  border-color: #e9ecef;
}
</style>