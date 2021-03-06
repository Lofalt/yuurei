import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";
import HomeView from "../views/HomeView.vue";
import PicPage from "@/views/ArticlePage.vue";
import BlogPage from "../views/MessageBoard.vue"
import VlogPage from "../views/VlogPage.vue"
import LovePage from "../views/LovePage.vue"
import AnliPage from "@/views/PicPage.vue"

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        redirect: "/home"
    },
    {
        path: "/home",
        name: "home",
        component: () => import("../components/RightView.vue"),
        children: [
            {
                path: "/archive/article/:id",
                component: () => import("@/views/Archive.vue"),
                props: true

            }
        ]
    },
    {
        path: "/admin",
        name: "admin",
        component: () => import("../views/admin/Board.vue"),
        props: true,
        meta: {
            needAuth: true
        }
    },
    {
        path: "/login",
        name: "login",
        component: () => import("../views/user/LogSign.vue"),
        props: true
    },
    // {
    //   path: "/home",
    //   name: "home",
    //   component: () => import("../components/RightBar.vue"),
    //   children: [
    //     {
    //       path: "pages",
    //       components: {
    //         default: PicPage,
    //         BlogPage,
    //         VlogPage,
    //         LovePage

    //       },
    //       children: [
    //         {
    //           path: "anli/:pageNum",
    //           component: AnliPage,
    //           props: true,
    //         }
    //       ]
    //     }
    //   ]
    // },
    {
        path: "/About",
        component: () => import("@/views/About.vue")
    },


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
