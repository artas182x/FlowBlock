import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";


import { FontAwesomeIcon } from './plugins/font-awesome'
import { BaklavaVuePlugin } from "@baklavajs/plugin-renderer-vue3";
import "@baklavajs/plugin-renderer-vue3/dist/styles.css";

import { library } from "@fortawesome/fontawesome-svg-core";
import {
    faCoffee,
    faCocktail,
    faGlassMartini,
    faBeer,
    faMicrochip,
    faDatabase
} from "@fortawesome/free-solid-svg-icons";

library.add(
    faCoffee,
    faCocktail,
    faGlassMartini,
    faBeer,
    faMicrochip,
    faDatabase
);


createApp(App)
    .use(router)
    .use(store)
    .use(BaklavaVuePlugin)
    .component("font-awesome-icon", FontAwesomeIcon)
    .mount("#app");