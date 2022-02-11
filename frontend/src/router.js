import { createWebHistory, createRouter } from "vue-router";
import Home from "./components/Home.vue";
import Login from "./components/Login.vue";

// lazy-loaded
const Profile = () => import("./components/Profile.vue")
const BoardModerator = () => import("./components/BoardComputations.vue")
const TokenSubmit = () =>  import("@/components/TokenSubmit");

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
  },
  {
    path: "/home",
    component: Home,
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
    component: BoardModerator,
  },
  {
    path: "/tokensubmit",
    name: "tokensubmit",
    // lazy-loaded
    component: TokenSubmit,
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