<template>
  <div id="select">
    <select v-model="$i18n.locale">
      <option v-for="(local, i) in locals" :key="i" :value="local">
        {{ local }}
      </option>
    </select>
  </div>
</template>
<script>
export default {
  data() {
    return {
      locals: ["en", "mr", "hi"]
    };
  }
};
</script>
