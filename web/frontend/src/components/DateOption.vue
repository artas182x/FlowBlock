
<template>
  {{ name }}
  <Datepicker
    v-model="date"
    class="dark-input"
    :format="format"
    @update:modelValue="listeners"
  />
</template>

<script>

import {defineComponent} from "vue";
import Datepicker from 'vue3-date-time-picker';
import 'vue3-date-time-picker/dist/main.css'

export default defineComponent({
  components: { Datepicker },
  props: ["value", "name"],
  emits: ["input"],
  setup() {
    const format = (date) => {
      const day = date.getDate();
      const month = date.getMonth() + 1;
      const year = date.getFullYear();

      return `${day}/${month}/${year}`;
    }

    return {
      format,
    }
  },
  data() {
    return {
      date: this.value,
      nodeName: this.name
    }
  },
  methods: {
     listeners(event) {
       return this.$emit("input", '' + Math.round(event/1000));
    }
  }


});


</script>