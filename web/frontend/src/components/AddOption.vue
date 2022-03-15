
<template>
  <button
    class="dark-input"
    @click="onClick"
  >
    {{ name }}
  </button>
  <p />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import TokenNode from "./TokenNode";
export default defineComponent({
  props: {
    name: {
      type: String,
      required: true
    },
    node: {
      type:  Object as () => TokenNode,
      required: true
    }
  },
  methods: {
    onClick() {
      this.node.action(this.name);
    },
  }
});
</script>