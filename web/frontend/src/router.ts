import { createWebHistory, createRouter } from "vue-router";
import Login from "./components/Login.vue";

// lazy-loaded
const Profile = () => import("./components/Profile.vue")
const BoardComputations = () => import("./components/BoardComputations.vue")
const TokenSubmit = () =>  import("@/components/TokenSubmit.vue");
const BoardMedicalData = () => import("./components/BoardMedicalData.vue")

const routes = [
  {
    path: "/",
    name: "computations",
    component: BoardComputations,
  },
  {
    path: "/login",
    component: Login,
  },
  {
    path: "/profile",
    name: "profile",
    // lazy-loaded
    component: Profile,
  },
  {
    path: "/computations",
    name: "computations",
    // lazy-loaded
    component: BoardComputations,
  },
  {
    path: "/tokensubmit",
    name: "tokensubmit",
    // lazy-loaded
    component: TokenSubmit,
  },
  {
    path: "/medicaldata",
    name: "medicaldata",
    // lazy-loaded
    component: BoardMedicalData,
  },

];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// router.beforeEach((to, from, next) => {
//   const publicPages = ['/login', '/register', '/home'];
//   const authRequired = !publicPages.includes(to.path);
//   const loggedIn = localStorage.getItem('user');

//   // trying to access a restricted page + not logged in
//   // redirect to login page
//   if (authRequired && !loggedIn) {
//     next('/login');
//   } else {
//     next();
//   }
// });

export default router;