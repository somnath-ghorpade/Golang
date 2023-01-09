<template>
  <div id="day-select">
    <ul class="days">
      <li
        class="day"
        v-for="day in days"
        :key="day"
        :class="{ active: isActive(day) }"
        v-on:click="$emit('day-selected', day)"
      >
        {{ formatDay(day) }}
      </li>
    </ul>
  </div>
</template>

<script>
import dayjs from "dayjs";
export default {
  data: () => ({
    days: [0, 1, 2, 3, 4, 5, 6].map(num => dayjs().add(num, "day"))
  }),
  props: {
    day: Object
  },
  emits: ["day-selected"],
  methods: {
    formatDay(raw) {
      if (raw.isSame(dayjs(), "day")) {
        return "Today";
      } else {
        return raw.format("ddd DD/MM");
      }
    },
    isActive(day) {
      return day.isSame(this.day, "day");
    }
  }
};
</script>

<style scoped></style>
