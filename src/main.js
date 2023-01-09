
/* eslint-disable */
import { createApp } from "vue";
import "@/assets/style.scss";
import router from "@/router";
import store from "@/store";
import tooltip from "@/plugins/tooltip";
import App from "./App.vue";
import i18n from "./i18n";
const app = createApp(App);
app.use(i18n);
app.use(router);
app.use(store);
app.use(tooltip);
app.mount("#app");
