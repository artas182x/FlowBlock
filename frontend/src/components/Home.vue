<template>
  <div class="jumbotron">
    <h1 class="display-4">{{ content }}</h1>
    <p class="lead">Welcome to Blockchain processing service</p>

  </div>
</template>

<script>

export default {
  name: "Home",
  data() {
    return {
      content: "Main page",
    };
  },
};
</script>
