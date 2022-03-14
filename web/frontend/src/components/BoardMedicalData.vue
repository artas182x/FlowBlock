<template>
  <div class="jumbotron">
    <h1 class="display-4">
      Medical data
    </h1>
  </div>
  <div>
    <h3>Show only data that you have rights to read</h3>
    <h4>
      If you are academic institution you are probably looking for <router-link to="/computations">
        <font-awesome-icon icon="microchip" /> Workflows
      </router-link>
    </h4>

    <table-lite
      :is-slot-mode="true"
      :is-loading="tableMedicalData.isLoading"
      :columns="tableMedicalData.columns"
      :rows="tableMedicalData.rows"
      :total="tableMedicalData.totalRecordCount"
      @do-search="doSearchMedicalData"
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
    tableMedicalData() {
      return reactive({
        isLoading: true,
        columns: [
          {
            label: "ID",
            field: "ID",
            display:  (row) => {

              return (
                  '<button type="button" title="Copy ID to clipboard" data-id="' +
                  row.ID +
                  '" class="is-rows-el copy-btn btn btn-light">&#x1f4cb;</button>'
              );

            },
          },
          {
            label: "Patient ID",
            field: "PatientID",
          },
          {
            label: "Name",
            field: "MedicalEntryName",
          },
          {
            label: "Value",
            field: "MedicalEntryValue",
            display:  (row) => {

              if (row.MedicalEntryType !== "s3img") {
                return ('<div style="margin-left: 10px;">' + row.MedicalEntryValue + "")
              } else {
                return (
                    '<button type="button" title="Download asset" data-id="' +
                    row.MedicalEntryValue +
                    '" class="is-rows-el download-btn btn btn-light">&#128229;</button>'
                );
              }



            },
          },
          {
            label: "Date added",
            field: "DateAdded",
          },
        ],
        rows: [],
        totalRecordCount: 0,
      });
    },

  },
  mounted() {
    if (!this.currentUser) {
      this.$router.push('/login');
    }
    UserService.refreshToken().then(
        // eslint-disable-next-line @typescript-eslint/no-empty-function
        () => {},
        (error) => {
          if (error.response.status === 401) {
            this.logOut()
          }
        }
    )
    this.doSearchMedicalData(0, 10)
  },
  methods: {

    tableLoadingFinish(elements) {
      this.tableMedicalData.isLoading = false;
      Array.prototype.forEach.call(elements, (element) => {
        if (element.classList.contains("copy-btn")) {
          element.addEventListener("click", () => {
            navigator.clipboard.writeText(element.getAttribute("data-id"))
          });
        }
        else if (element.classList.contains("download-btn")) {
          element.addEventListener("click", () => {

            UserService.downloadFile(element.getAttribute("data-id")).then((response) => {
              const fileURL = window.URL.createObjectURL(new Blob([response.data]));
              const fileLink = document.createElement('a');

              fileLink.href = fileURL;
              fileLink.setAttribute('download', element.getAttribute("data-id").split("?")[0]);
              document.body.appendChild(fileLink);

              fileLink.click();
            })
          });
        }
      });
    },

    doSearchMedicalData(offset, limit) {
      this.tableMedicalData.isLoading = true;
      setTimeout(() => {
        this.tableMedicalData.isReSearch = offset === undefined;

        let request = {"medicalEntryName": "", "dateStartTimestamp": "0", "dateEndTimestamp": Math.floor(Date.now() / 1000).toString()}

        UserService.requestMedicalData(request).then(
            (response) => {
              let dataResponse = response.data
              let data = [];
              const max = Math.min(offset+limit, dataResponse.length);
              for (let i = offset; i < max ; i++) {
                data.push({
                  ID: dataResponse[i]["ID"],
                  PatientID: dataResponse[i]["PatientID"],
                  MedicalEntryName: dataResponse[i]["MedicalEntryName"],
                  MedicalEntryValue: dataResponse[i]["MedicalEntryValue"],
                  MedicalEntryType: dataResponse[i]["MedicalEntryType"],
                  DateAdded: moment(new Date(dataResponse[i]["DateAdded"])).format('MMMM Do YYYY, h:mm:ss a'),
                });
              }
              this.tableMedicalData.rows = data;
              this.tableMedicalData.totalRecordCount = dataResponse.length;
              this.tableMedicalData.isLoading = false;
            },
            (error) => {
              this.tableMedicalData.isLoading = false;
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

<style scoped>
::v-deep(.vtl-table .vtl-thead .vtl-thead-th) {
  color: #000000;
  background-color: #e9ecef;
  border-color: #e9ecef;
}
</style>