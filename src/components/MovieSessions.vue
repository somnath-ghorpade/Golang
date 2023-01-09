<template>
  <div class="movie-sessions">
    <div
      v-for="session in filteredSession"
      :key="session.id"
      class="session-time-wrapper tooltip-wrapper"
      v-tooltip="session.seats"
    >
      <div class="session-time">
        {{ formatSessionTime(session.time) }}
      </div>
    </div>
  </div>
</template>

<script>
import dayjs from "dayjs";
import times from "@/util/times";
export default {
  props: {
    movie: Object,
    day: Object,
    times: Array
  },
  computed: {
    filteredSession() {
      return this.movie.sessions.filter(this.dayFilter).filter(this.timeFilter);
    }
  },
  methods: {
    formatSessionTime(raw) {
      return dayjs(raw).format("h:mm A");
    },
    dayFilter(session) {
      return dayjs(session.time).isSame(this.day, "day");
    },
    timeFilter(session) {
      if (!this.times.length || this.times.length === 2) {
        return true;
      } else if (this.times[0] === times.BEFORE_6PM) {
        return dayjs(session.time).hour() < 18;
      } else {
        return dayjs(session.time).hour() >= 18;
      }
    }
  }
};
</script>

<style scoped></style>
