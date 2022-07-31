import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import HomeView from "../views/HomeView.vue";
import PicPage from "../../../src/views/ArticlePage.vue";
import BlogPage from "../views/BlogPage.vue"
import VlogPage from "../views/VlogPage.vue"
import LovePage from "../views/LovePage.vue"
import AnliPage from "../../../src/views/PicPage.vue"

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    redirect: "/home"
  },
  {
    path: "/home",
    name: "home",
    component: () => import("../components/RightBar.vue"),
    children: [
      {
        path: "pages",
        components: {
          default: PicPage,
          BlogPage,
          VlogPage,
          LovePage

        },
        children: [

        ]
      }
    ]
  },
  {
    path: "/anli/:pageNum",
    // name: "anli",
    component: AnliPage,
    props: true,
  }

  // {
  //   path: "/about",
  //   name: "about",
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: () =>
  //     import(/* webpackChunkName: "about" */ "../views/AboutView.vue"),
  // },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
